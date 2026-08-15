package ws

import (
	"backend/src/app"
	"log"
	"net"
	"time"

	"github.com/gorilla/websocket"
)

type Hub struct {
	match     *app.Match
	clients   map[net.Addr]*Client
	broadcast chan *Message[any]
	direct    chan *DirectMessage[any]
	endTurn   chan struct{}
}

func newHub() *Hub {
	return &Hub{
		match:     app.NewMatch(),
		clients:   map[net.Addr]*Client{},
		broadcast: make(chan *Message[any]),
		direct:    make(chan *DirectMessage[any]),
		endTurn:   make(chan struct{}),
	}
}

func (h *Hub) handleTicks() {
	const turnDuration = 15 * time.Second
	turnTimer := time.NewTicker(turnDuration)
	defer turnTimer.Stop()

	for {
		select {
		case <-turnTimer.C:
			h.endTurn <- struct{}{}

		case <-h.endTurn:
			h.handleNextTurn()
			turnTimer.Reset(turnDuration)

		case m := <-h.broadcast:
			h.handleBroadcastMessage(m)

		case dm := <-h.direct:
			h.handleDirectMessage(dm)
		}
	}
}

func (h *Hub) handleNextTurn() {
	noSeatsInThisRound := true
	for _, s := range h.match.RoundSeats {
		if s != nil {
			noSeatsInThisRound = false
			break
		}
	}
	if noSeatsInThisRound {
		return
	}

	var lastRoundSeat *app.Seat
	for i := len(h.match.RoundSeats) - 1; i >= 0; i-- {
		if roundSeat := h.match.RoundSeats[i]; roundSeat.Player != nil {
			lastRoundSeat = roundSeat
			break
		}
	}
	isBettingRoundOver := lastRoundSeat == h.match.SeatTurn

	if isBettingRoundOver {
		if h.match.AllTableCardsAreRevealed() {
			h.handleShowdown()
			return
		}
		h.handleRevealNextTableCard()
		h.match.InitRound()
	}
	h.handlePassTurn()
}

func (h *Hub) handleRegisterClient(c *websocket.Conn) *Client {
	client := newClient(c, h)

	h.clients[client.addr] = client
	log.Println("new client joined:", client.addr)
	log.Println("current clients:", len(h.clients))

	return client
}

func (h *Hub) handleUnregisterClient(addr net.Addr) {
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
		if playerSeatIdx == h.match.SeatTurn.Index {
			h.endTurn <- struct{}{}
		}
	}
	delete(h.clients, addr)
	h.sendPlayersInfo()
	log.Println("client left:", addr)
	log.Println("current clients:", len(h.clients))
}

func (h *Hub) handleBroadcastMessage(m *Message[any]) {
	for _, c := range h.clients {
		c.sendMessage(m.RequestId, m.Type, m.Payload, nil)
	}
}

func (h *Hub) handleDirectMessage(dm *DirectMessage[any]) {
	client := h.clients[dm.conn]

	client.sendMessage(
		dm.message.RequestId,
		dm.message.Type,
		dm.message.Payload,
		nil,
	)
}

// TODO: active players in the round are never informed to clients
func (h *Hub) sendPlayersInfo() {
	loggedInClients := []*Client{}
	for _, c := range h.clients {
		if c.player != nil {
			loggedInClients = append(loggedInClients, c)
		}
	}

	for _, c1 := range loggedInClients {
		c1.sendMessage(nil, USER_INFO, c1.player, nil)

		opponents := []*app.Player{}
		for _, c2 := range loggedInClients {
			if c1 != c2 {
				opponent := &app.Player{
					Id:        c2.player.Id,
					Name:      c2.player.Name,
					Score:     c2.player.Score,
					Cards:     [2]app.Card{app.BACK, app.BACK},
					SeatIndex: c2.player.SeatIndex,
				}
				opponents = append(opponents, opponent)
			}
		}
		c1.sendMessage(nil, OPPONENTS_INFO, opponents, nil)
	}
}

func (h *Hub) handlePassTurn() {
	nextSeatToPlay := h.match.PassTurn()
	seatTurnMsg := newOutMessage(
		nil,
		MATCH_SEAT_TURN,
		map[string]app.SeatIndex{
			"seatIndex": nextSeatToPlay.Index,
		},
		nil,
	)
	h.broadcast <- seatTurnMsg.asAny()
}

func (h *Hub) handleRevealNextTableCard() {
	if err := h.match.RevealNextTableCard(); err != nil {
		log.Println(err)
		return
	}
	for _, c := range h.match.TableCards {
		if c == app.BACK {
			tableCardsMsg := newOutMessage(
				nil,
				MATCH_TABLE_CARDS,
				h.match.TableCards,
				nil,
			)

			h.broadcast <- tableCardsMsg.asAny()
			break
		}
	}
}

func (h *Hub) handleShowdown() {
	h.match.Showdown()
	// TODO: communicate winners
}
