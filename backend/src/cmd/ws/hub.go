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
	clients    map[net.Addr]*Client
	mailbox    chan HubMsg
	match      *app.Match
	turnTicker *time.Ticker
}

const TurnDuration = 15 * time.Second

func newHub() *Hub {
	return &Hub{
		clients: map[net.Addr]*Client{},
		mailbox: make(chan HubMsg), // ? buffer
		match:   app.NewMatch(),
	}
}

func (h *Hub) run() {
	h.turnTicker = time.NewTicker(TurnDuration)
	defer h.turnTicker.Stop()

	for {
		select {
		case <-h.turnTicker.C:
			h.endTurn()
		case m := <-h.mailbox:
			h.dispatch(m)
		}
	}
}

func (h *Hub) dispatch(m HubMsg) {
	switch msg := m.(type) {
	case registerClientMsg:
		msg.reply <- h.registerClient(msg.conn)
	case unregisterClientMsg:
		h.unregisterClient(msg.addr)
		close(msg.done)
	case loginMsg:
		msg.reply <- h.handleLogin(msg.client, msg.userName)
	case actionMsg:
		msg.reply <- h.handleAction(msg.client, msg.action, msg.amount)
	case getPlayerMsg:
		msg.reply <- msg.client.player
	default:
		log.Printf("hub: unhandled message type %T", m)
	}
}

// TODO: a new round needs to init whenever someone raises
func (h *Hub) endTurn() {
	h.turnTicker.Reset(1 * time.Hour)
	defer h.turnTicker.Reset(TurnDuration)

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
	h.sendPlayersInfo(true)
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
		h.sendSeatsInfo()
		if h.match.SeatTurn != nil && playerSeatIdx == h.match.SeatTurn.Index {
			h.endTurn()
		} else {
			h.sendPlayersInfo(true)
		}
	}
	delete(h.clients, addr)
	playerName := "N/A"
	if player != nil {
		playerName = player.Name
	}
	log.Printf("client '%s (%s)' unregistered (current clients = %d)", addr, playerName, len(h.clients))
}

func (h *Hub) broadcast(m *Message[any]) {
	var errMsg string
	var errDetails any
	if m.Error != nil {
		errMsg = m.Error.Message
		errDetails = m.Error.Details
	}

	for _, c := range h.clients {
		c.sendChan <- ServerMessageArgs[any]{
			RequestId:  m.RequestId,
			Type:       m.Type,
			Payload:    m.Payload,
			ErrMessage: errMsg,
			ErrDetails: errDetails,
		}
	}
}

// TODO: active players in the round are never informed to clients
func (h *Hub) sendPlayersInfo(hideOpponentsHands bool) {
	var loggedInClients []*Client
	for _, c := range h.clients {
		if c.player != nil {
			loggedInClients = append(loggedInClients, c)
		}
	}

	for _, c1 := range loggedInClients {
		user := newPlayer(c1.player, false)
		c1.sendChan <- ServerMessageArgs[any]{
			Type:    USER_INFO,
			Payload: user,
		}

		opponents := []*Player{}
		for _, c2 := range loggedInClients {
			if c1 != c2 {
				opponent := newPlayer(c2.player, hideOpponentsHands)
				opponents = append(opponents, opponent)
			}
		}
		c1.sendChan <- ServerMessageArgs[any]{
			Type:    OPPONENTS_INFO,
			Payload: opponents,
		}
	}
}

func (h *Hub) sendSeatsInfo() {
	seatPlayerMap := map[app.SeatIndex]*string{}
	for _, s := range h.match.Seats {
		if s.Player != nil {
			id := s.Player.Id.String()
			seatPlayerMap[s.Index] = &id
		} else {
			seatPlayerMap[s.Index] = nil
		}
	}

	seatsMsg := newServerMessage(ServerMessageArgs[any]{
		Type:    MATCH_SEATS,
		Payload: seatPlayerMap,
	})
	h.broadcast(seatsMsg)
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
	winners := h.match.Showdown()
	winnersIds := make([]string, len(winners))
	for i, w := range winners {
		winnersIds[i] = w.Id.String()
	}
	winnersMsg := newServerMessage(ServerMessageArgs[any]{
		Type:    MATCH_WINNERS,
		Payload: winnersIds,
	})

	h.broadcast(winnersMsg)
	h.sendPlayersInfo(false)
}

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

	h.sendPlayersInfo(true)
	h.sendSeatsInfo()

	if h.match.RoundSeats == [8]*app.Seat{} {
		playersCount := 0
		for _, s := range h.match.Seats {
			if s.Player != nil {
				playersCount++
			}
			if playersCount == 2 {
				h.initRound()
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
