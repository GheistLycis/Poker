package ws

import (
	"backend/src/app"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/gorilla/websocket"
)

type Hub struct {
	mailbox    chan any
	turnTicker *time.Ticker
	match      *app.Match
	clients    map[net.Addr]*Client
}

const TurnDuration = 15 * time.Second

// Messages sent by Client goroutines. Every field either the caller or the
// Hub needs back travels through a reply/done channel — nothing is ever
// read back out of Hub state directly.
type registerMsg struct {
	conn  *websocket.Conn
	reply chan *Client
}

type unregisterMsg struct {
	addr net.Addr
	done chan struct{}
}

type loginMsg struct {
	client   *Client
	userName string
	reply    chan error
}

type actionMsg struct {
	client *Client
	action app.PlayerAction
	amount *int
	reply  chan error
}

func newHub() *Hub {
	return &Hub{
		mailbox: make(chan any),
		match:   app.NewMatch(),
		clients: map[net.Addr]*Client{},
	}
}

// run is the actor loop. h.match, h.clients and h.turnTicker must only ever
// be touched from inside this goroutine — that's the entire invariant the
// rest of the package has to respect.
func (h *Hub) run() {
	h.turnTicker = time.NewTicker(TurnDuration)
	defer h.turnTicker.Stop()

	for {
		select {
		case <-h.turnTicker.C:
			h.endTurn()
		case raw := <-h.mailbox:
			h.dispatch(raw)
		}
	}
}

func (h *Hub) dispatch(raw any) {
	switch msg := raw.(type) {
	case registerMsg:
		msg.reply <- h.registerClient(msg.conn)
	case unregisterMsg:
		h.unregisterClient(msg.addr)
		close(msg.done)
	case loginMsg:
		msg.reply <- h.handleLogin(msg.client, msg.userName)
	case actionMsg:
		msg.reply <- h.handleAction(msg.client, msg.action, msg.amount)
	default:
		log.Printf("hub: unhandled message type %T", raw)
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
	client := newClient(c, h.mailbox)

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
	seatTurnMsg := newServerMessage(ServerMessageArgs[any]{
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

// --- moved from Client: these mutate match/seat state, so they belong to
// the Hub actor, not to whichever goroutine happened to read the socket.

func (h *Hub) handleLogin(c *Client, userName string) error {
	var availableSeat *app.Seat
	for _, s := range h.match.Seats {
		if s.Player == nil {
			availableSeat = s
			break
		}
	}
	if availableSeat == nil {
		return errors.New("no available seats for user in this match")
	}

	player := app.NewPlayer(userName, availableSeat.Index)
	availableSeat.Player = player
	c.player = player

	h.sendPlayersInfo()

	if h.match.RoundSeats == [8]*app.Seat{} {
		playersCount := 0
		for _, s := range h.match.Seats {
			if s.Player != nil {
				playersCount++
			}
			if playersCount == 2 {
				h.match.InitRound()
				break
			}
		}
	}

	return nil
}

func (h *Hub) handleAction(c *Client, action app.PlayerAction, amount *int) error {
	switch action {
	case app.BET:
		if amount == nil {
			return fmt.Errorf("no amount provided for bet")
		}
		value := *amount
		if value <= h.match.LastBet {
			return errors.New("bets/raises are only allowed if greater than the last bet")
		}
		h.match.DoPotTransaction(value, c.player)
		h.match.LastBet = value

	case app.CALL:
		h.match.DoPotTransaction(h.match.LastBet, c.player)

	case app.FOLD:
		playerSeatIdx := c.player.SeatIndex
		var newRoundSeats [8]*app.Seat
		for i, s := range h.match.RoundSeats {
			if s != nil && s.Index != playerSeatIdx {
				newRoundSeats[i] = s
			}
		}
		h.match.RoundSeats = newRoundSeats
	}

	if action == app.BET || action == app.CALL {
		potAmountMsg := newServerMessage(ServerMessageArgs[any]{
			Type: MATCH_POT_AMOUNT,
			Payload: map[string]int{
				"amount": h.match.Pot,
			},
		})
		h.broadcast(potAmountMsg)
	}
	h.endTurn()

	return nil
}
