package ws

import (
	"backend/src/app"
	"errors"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	addr   net.Addr
	conn   *websocket.Conn
	player *app.Player
}

func newClient(c *websocket.Conn) *Client {
	return &Client{
		addr: c.RemoteAddr(),
		conn: c,
	}
}

func (c *Client) receiveMessages() {
	for {
		msg := &Message[any]{}
		if err := c.conn.ReadJSON(msg); err != nil {
			log.Printf("read error (%T): %v", err, err)
			hub.unregister <- c.addr
			return
		}
		log.Printf("RECEIVED: type=%s payload=%+v", msg.Type, msg.Payload)
		if err := c.handleReceivedMessage(msg); err != nil {
			log.Printf("failed to handle message: %v", err)
		}
	}
}

// TODO: pass use cases handling down to app pkg
func (c *Client) handleReceivedMessage(m *Message[any]) error {
	if m.Type == "user.login" {
		payload, ok := m.Payload.(map[string]any)
		if !ok {
			return errors.New("malformed payload")
		}
		userName, ok := payload["userName"].(string)
		if !ok {
			return errors.New("userName must be a valid string")
		}

		var availableSeat *app.Seat
		for _, s := range hub.match.Seats {
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

		if err := c.sendMessage(
			m.RequestId,
			"user.info",
			map[string]any{
				"id":        player.Id,
				"name":      player.Name,
				"score":     player.Score,
				"seatIndex": player.SeatIndex,
				"cards":     player.Cards,
			},
		); err != nil {
			return err
		}

		if err := c.sendMessage(
			nil,
			"match.seat-turn",
			map[string]app.SeatIndex{"seatIndex": player.SeatIndex},
		); err != nil {
			return err
		}

		return nil
	}

	log.Printf("// TODO: handle message type '%s'", m.Type)

	return nil
}

func (c *Client) sendMessage(rId *uuid.UUID, t string, p any) error {
	msg := newOutMessage(rId, t, p)

	if err := c.conn.WriteJSON(msg); err != nil {
		log.Printf("write error (%T): %v", err, err)
		return err
	}

	return nil
}
