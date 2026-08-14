package ws

import (
	"backend/src/app"
	"errors"
	"fmt"
	"log"
	"net"
	"slices"

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
			c.hub.unregister <- c.addr // ? unregister only if read error == conn shut
			return
		}
		log.Printf("RECEIVED: type=%s payload=%+v", msg.Type, msg.Payload)

		switch msg.Type {
		case "user.login":
			if err := c.handleLogin(msg); err != nil {
				c.sendMessage(
					msg.RequestId,
					msg.Type,
					nil,
					newError(fmt.Sprintf("failed to handle login: %v", err), nil),
				)
				c.hub.unregister <- c.conn.RemoteAddr()
			}

		case "user.action":
			if err := c.handleAction(msg); err != nil {
				c.sendMessage(
					msg.RequestId,
					msg.Type,
					nil,
					newError(fmt.Sprintf("failed to handle action: %v", err), nil),
				)
			}

		default:
			log.Printf("// TODO: handle message type '%s'", msg.Type)
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

	player := app.NewPlayer(payload.userName, availableSeat.Index)
	availableSeat.Player = player
	c.player = player
	c.hub.sendPlayersInfo(false)

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
	if slices.Contains(app.ActionsWithAmount, payload.action) && payload.amount == nil {
		return fmt.Errorf("no amount provided for action %s", payload.action)
	}

	switch payload.action {
	case app.CALL:
		lastBet := c.hub.match.LastBet

		if err := c.player.Call(lastBet); err != nil {
			return err
		}
		c.hub.match.Pot += lastBet

	case app.FOLD:
		playerSeatIdx := c.player.SeatIndex
		var newRoundSeats []*app.Seat

		for _, s := range c.hub.match.RoundSeats {
			if s.Index != playerSeatIdx {
				newRoundSeats = append(newRoundSeats, s)
			}
		}

	case app.BET:
		c.player.Bet()

	case app.RAISE:
		c.player.Raise()
	}

	c.hub.endTurn <- struct{}{}

	return nil
}
