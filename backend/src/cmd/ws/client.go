package ws

import (
	"backend/src/app"
	"encoding/json"
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
		msg := &Message[json.RawMessage]{}
		if err := c.conn.ReadJSON(msg); err != nil {
			log.Printf("read error (%T): %v", err, err)
			break
		}
		clientId := c.addr.String()
		if c.player != nil {
			clientId = c.player.Name
		}
		log.Printf("[CLIENT %s] received msg: type=%s payload=%s", clientId, msg.Type, msg.Payload)

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

func (c *Client) handleLogin(m *Message[json.RawMessage]) error {
	var payload struct {
		UserName string `json:"userName"`
	}
	if err := json.Unmarshal(m.Payload, &payload); err != nil {
		return err
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
	player := app.NewPlayer(payload.UserName, availableSeat.Index)

	availableSeat.Player = player
	c.player = player
	c.hub.sendPlayersInfo()
	if match.RoundSeats == [8]*app.Seat{} {
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

func (c *Client) handleAction(m *Message[json.RawMessage]) error {
	var payload struct {
		Action app.PlayerAction `json:"action"`
		Amount *int             `json:"amount"`
	}
	if err := json.Unmarshal(m.Payload, &payload); err != nil {
		return err
	}

	match := c.hub.match

	switch payload.Action {
	case app.BET:
		if payload.Amount == nil {
			return fmt.Errorf("no amount provided for bet")
		}

		value := *payload.Amount

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
	if payload.Action == app.BET || payload.Action == app.CALL {
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
	c.hub.handleEndTurn()

	return nil
}
