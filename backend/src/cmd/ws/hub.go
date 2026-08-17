package ws

import (
	"backend/src/app"
	"log"
	"net"
	"time"

	"github.com/gorilla/websocket"
)

// TODO: implement actor model pattern to prevent race condition from clients goroutines
type Hub struct {
	turnTicker *time.Ticker
	match      *app.Match
	clients    map[net.Addr]*Client
}

const TurnDuration = 15 * time.Second

func newHub() *Hub {
	return &Hub{
		match:   app.NewMatch(),
		clients: map[net.Addr]*Client{},
	}
}

func (h *Hub) handleTicks() {
	turnTicker := time.NewTicker(TurnDuration)
	defer turnTicker.Stop()

	h.turnTicker = turnTicker
	for range h.turnTicker.C {
		h.endTurn()
	}
}

func (h *Hub) endTurn() {
	h.turnTicker.Reset(TurnDuration)

	if h.match.RoundSeats == [8]*app.Seat{} {
		return
	}

	var lastRoundSeat *app.Seat
	for i := len(h.match.RoundSeats) - 1; i >= 0; i-- {
		roundSeat := h.match.RoundSeats[i]
		if roundSeat != nil && roundSeat.Player != nil {
			lastRoundSeat = roundSeat
			break
		}
	}
	isBettingRoundOver := lastRoundSeat == h.match.SeatTurn

	if isBettingRoundOver {
		if h.match.AllTableCardsAreRevealed() {
			h.showdown()
			h.initRound()
		} else {
			h.revealNextTableCard()
			h.passTurn()
		}
	} else {
		h.passTurn()
	}
	h.sendPlayersInfo()
	h.turnTicker.Reset(TurnDuration)
}

func (h *Hub) registerClient(c *websocket.Conn) *Client {
	client := newClient(c, h)

	h.clients[client.addr] = client
	log.Printf("client registered: %s (current clients = %d)", client.addr, len(h.clients))

	return client
}

func (h *Hub) unregisterClient(addr net.Addr) {
	client := h.clients[addr]
	if client == nil {
		return
	}
	player := client.player
	if player != nil {
		playerSeatIdx := player.SeatIndex

		h.match.Seats[playerSeatIdx].Player = nil
		for i, s := range h.match.RoundSeats {
			if s != nil && s.Index == playerSeatIdx {
				h.match.RoundSeats[i] = nil
				break
			}
		}
		if h.match.SeatTurn != nil && playerSeatIdx == h.match.SeatTurn.Index {
			h.endTurn()
		}
	}
	delete(h.clients, addr)
	h.sendPlayersInfo()
	playerName := "N/A"
	if player != nil {
		playerName = player.Name
	}
	log.Printf("client '%s (%s)' unregistered (current clients = %d)", addr, playerName, len(h.clients))
}

func (h *Hub) broadcast(m *Message[any]) {
	for _, c := range h.clients {
		c.sendMessage(ServerMessageArgs[any]{
			RequestId:  m.RequestId,
			Type:       m.Type,
			Payload:    m.Payload,
			ErrMessage: m.Error.Message,
			ErrDetails: m.Error.Details,
		})
	}
}

// TODO: active players in the round are never informed to clients
func (h *Hub) sendPlayersInfo() {
	var loggedInClients []*Client
	for _, c := range h.clients {
		if c.player != nil {
			loggedInClients = append(loggedInClients, c)
		}
	}

	for _, c1 := range loggedInClients {
		user := newPlayer(c1.player, false)
		c1.sendMessage(ServerMessageArgs[any]{
			Type:    USER_INFO,
			Payload: user,
		})

		opponents := make([]*Player, len(loggedInClients))
		for _, c2 := range loggedInClients {
			if c1 != c2 {
				opponent := newPlayer(c2.player, true)
				opponents = append(opponents, opponent)
			}
		}
		c1.sendMessage(ServerMessageArgs[any]{
			Type:    OPPONENTS_INFO,
			Payload: opponents,
		})
	}
}

func (h *Hub) initRound() {
	h.match.InitRound()
	tableCardsMsg := newServerMessage(ServerMessageArgs[any]{
		Type:    MATCH_TABLE_CARDS,
		Payload: h.match.TableCards,
	})

	h.broadcast(tableCardsMsg)
}

func (h *Hub) passTurn() {
	nextSeatToPlay := h.match.PassTurn()
	seatTurnMsg := newServerMessage(
		ServerMessageArgs[any]{
			Type: MATCH_SEAT_TURN,
			Payload: map[string]app.SeatIndex{
				"seatIndex": nextSeatToPlay.Index,
			},
		})
	h.broadcast(seatTurnMsg)
}

func (h *Hub) revealNextTableCard() {
	if err := h.match.RevealNextTableCard(); err != nil {
		log.Println(err)
		return
	}

	tableCardsMsg := newServerMessage(ServerMessageArgs[any]{
		Type:    MATCH_TABLE_CARDS,
		Payload: h.match.TableCards,
	})

	h.broadcast(tableCardsMsg)
}

func (h *Hub) showdown() {
	h.match.Showdown()
	// TODO: communicate winners
}
