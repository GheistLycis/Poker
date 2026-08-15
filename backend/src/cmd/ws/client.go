package ws

import (
	"backend/src/app"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	addr   net.Addr
	conn   *websocket.Conn
	player *app.Player
	hub    *Hub
}

func newClient(c *websocket.Conn, h *Hub) *Client {
	return &Client{
		addr: c.RemoteAddr(),
		conn: c,
		hub:  h,
	}
}

func (c *Client) handleMessages() {
	for {
		msg := &Message[any]{}
		if err := c.conn.ReadJSON(msg); err != nil {
			log.Printf("read error (%T): %v", err, err)
			break
		}
		log.Printf("RECEIVED: type=%s payload=%+v", msg.Type, msg.Payload)

		switch msg.Type {
		case USER_LOGIN:
			if err := c.handleLogin(msg); err != nil {
				c.sendMessage(
					msg.RequestId,
					msg.Type,
					nil,
					newError(fmt.Sprintf("failed to handle login: %v", err), nil),
				)
			}

		case USER_ACTION:
			if err := c.handleAction(msg); err != nil {
				c.sendMessage(
					msg.RequestId,
					msg.Type,
					nil,
					newError(fmt.Sprintf("failed to handle action: %v", err), nil),
				)
			}

		default:
			log.Printf("TODO: handle message type '%s'", msg.Type)
		}
	}
}

func (c *Client) sendMessage(rId *uuid.UUID, t MessageType, p any, err *Error) error {
	msg := newOutMessage(rId, t, p, err)

	if err := c.conn.WriteJSON(msg); err != nil {
		log.Printf("write error (%T): %v", err, err)
		return err
	}

	return nil
}

func (c *Client) handleLogin(m *Message[any]) error {
	payload, ok := m.Payload.(struct {
		userName string
	})
	if !ok {
		return errors.New("malformed payload")
	}

	var availableSeat *app.Seat
	for _, s := range c.hub.match.Seats {
		if s.Player == nil {
			availableSeat = s
			break
		}
	}
	if availableSeat == nil {
		return errors.New("no available seats for user in this match")
	}

	match := c.hub.match
	player := app.NewPlayer(payload.userName, availableSeat.Index)

	availableSeat.Player = player
	c.player = player
	c.hub.sendPlayersInfo()
	if len(match.RoundSeats) == 0 {
		playersCount := 0

		for _, s := range match.Seats {
			if s.Player != nil {
				playersCount++
			}
			if playersCount == 2 {
				match.InitRound()
				break
			}
		}
	}

	return nil
}

func (c *Client) handleAction(m *Message[any]) error {
	payload, ok := m.Payload.(struct {
		action app.PlayerAction
		amount *int
	})
	if !ok {
		return errors.New("malformed payload")
	}

	match := c.hub.match

	switch payload.action {
	case app.BET:
		if payload.amount == nil {
			return fmt.Errorf("no amount provided for bet")
		}

		value := *payload.amount

		if value <= match.LastBet {
			return errors.New("bets/raises are only allowed if greater than the last bet")
		}
		match.DoPotTransaction(value, c.player)
		match.LastBet = value

	case app.CALL:
		match.DoPotTransaction(match.LastBet, c.player)

	case app.FOLD:
		playerSeatIdx := c.player.SeatIndex
		var newRoundSeats [8]*app.Seat

		for i, s := range match.RoundSeats {
			if s.Index != playerSeatIdx {
				newRoundSeats[i] = s
			}
		}
		match.RoundSeats = newRoundSeats
	}
	if payload.action == app.BET || payload.action == app.CALL {
		potAmountMsg := newOutMessage(
			nil,
			MATCH_POT_AMOUNT,
			map[string]int{
				"amount": match.Pot,
			},
			nil,
		)
		c.hub.broadcast <- potAmountMsg.asAny()
	}
	c.hub.sendPlayersInfo()
	c.hub.endTurn <- struct{}{}

	return nil
}
