package ws

import (
	"backend/src/app"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/gorilla/websocket"
)

type Client struct {
	addr     net.Addr
	conn     *websocket.Conn
	player   *app.Player // owned exclusively by Hub.run()'s goroutine
	hub      *Hub
	sendChan chan ServerMessageArgs[any]
}

func newClient(c *websocket.Conn, h *Hub) *Client {
	client := &Client{
		addr:     c.RemoteAddr(),
		conn:     c,
		hub:      h,
		sendChan: make(chan ServerMessageArgs[any], 16),
	}
	go client.writePump()
	return client
}

// writePump is the only goroutine that calls conn.WriteJSON, so broadcasts
// from Hub and error replies from handleMessages never race on the socket.
func (c *Client) writePump() {
	for m := range c.sendChan {
		msg := newServerMessage(m)
		if err := c.conn.WriteJSON(msg); err != nil {
			log.Printf("write error (%T): %v", err, err)
		}
	}
}

func (c *Client) sendMessage(m ServerMessageArgs[any]) {
	c.sendChan <- m
}

func (c *Client) handleMessages() {
	loggedInAs := c.addr.String() // local to this goroutine only — no shared read of c.player

	for {
		msg := &Message[json.RawMessage]{}
		if err := c.conn.ReadJSON(msg); err != nil {
			log.Printf("read error (%T): %v", err, err)
			break
		}
		log.Printf("[CLIENT %s] received msg: type=%s payload=%s", loggedInAs, msg.Type, msg.Payload)

		switch msg.Type {
		case USER_LOGIN:
			var payload struct {
				UserName string `json:"userName"`
			}
			if err := json.Unmarshal(msg.Payload, &payload); err != nil {
				c.sendMessage(ServerMessageArgs[any]{
					RequestId:  msg.RequestId,
					Type:       msg.Type,
					ErrMessage: fmt.Sprintf("failed to handle login: %v", err),
				})
				continue
			}

			reply := make(chan error)
			c.hub.mailbox <- loginMsg{client: c, userName: payload.UserName, reply: reply}
			if err := <-reply; err != nil {
				c.sendMessage(ServerMessageArgs[any]{
					RequestId:  msg.RequestId,
					Type:       msg.Type,
					ErrMessage: fmt.Sprintf("failed to handle login: %v", err),
				})
			} else {
				loggedInAs = payload.UserName
			}

		case USER_ACTION:
			var payload struct {
				Action app.PlayerAction `json:"action"`
				Amount *int             `json:"amount"`
			}
			if err := json.Unmarshal(msg.Payload, &payload); err != nil {
				c.sendMessage(ServerMessageArgs[any]{
					RequestId:  msg.RequestId,
					Type:       msg.Type,
					ErrMessage: fmt.Sprintf("failed to handle action: %v", err),
				})
				continue
			}

			reply := make(chan error)
			c.hub.mailbox <- actionMsg{client: c, action: payload.Action, amount: payload.Amount, reply: reply}
			if err := <-reply; err != nil {
				c.sendMessage(ServerMessageArgs[any]{
					RequestId:  msg.RequestId,
					Type:       msg.Type,
					ErrMessage: fmt.Sprintf("failed to handle action: %v", err),
				})
			}

		default:
			log.Printf("TODO: handle message type '%s'", msg.Type)
		}
	}
}
