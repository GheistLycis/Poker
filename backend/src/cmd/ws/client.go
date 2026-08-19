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
	addr       net.Addr
	conn       *websocket.Conn
	hubMailbox chan HubMsg
	// owned by the Hub goroutine. Readonly within Client via `hubMailbox <- getPlayerMsg{}`
	player   *app.Player
	sendChan chan ServerMessageArgs[any]
}

func newClient(c *websocket.Conn, mailbox chan HubMsg) *Client {
	client := &Client{
		addr:       c.RemoteAddr(),
		conn:       c,
		hubMailbox: mailbox,
		sendChan:   make(chan ServerMessageArgs[any]), // ? buffer
	}

	go client.writePump()

	return client
}

func (c *Client) writePump() {
	for m := range c.sendChan {
		msg := newServerMessage(m)
		if err := c.conn.WriteJSON(msg); err != nil {
			log.Printf("write error (%T): %v", err, err)
		}
	}
}

func (c *Client) handleMessages() {
	loggedInAs := c.addr.String()

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
				c.sendChan <- ServerMessageArgs[any]{
					RequestId:  msg.RequestId,
					Type:       msg.Type,
					ErrMessage: fmt.Sprintf("failed to handle login: %v", err),
				}
				continue
			}

			reply := make(chan error)
			c.hubMailbox <- loginMsg{
				client:   c,
				userName: payload.UserName,
				reply:    reply,
			}
			if err := <-reply; err != nil {
				c.sendChan <- ServerMessageArgs[any]{
					RequestId:  msg.RequestId,
					Type:       msg.Type,
					ErrMessage: fmt.Sprintf("failed to handle login: %v", err),
				}
			} else {
				loggedInAs = payload.UserName
			}

		case USER_ACTION:
			var payload struct {
				Action app.PlayerAction `json:"action"`
				Amount *int             `json:"amount"`
			}
			if err := json.Unmarshal(msg.Payload, &payload); err != nil {
				c.sendChan <- ServerMessageArgs[any]{
					RequestId:  msg.RequestId,
					Type:       msg.Type,
					ErrMessage: fmt.Sprintf("failed to handle action: %v", err),
				}
				continue
			}

			reply := make(chan error)
			c.hubMailbox <- actionMsg{
				client: c,
				action: payload.Action,
				amount: payload.Amount,
				reply:  reply,
			}
			if err := <-reply; err != nil {
				c.sendChan <- ServerMessageArgs[any]{
					RequestId:  msg.RequestId,
					Type:       msg.Type,
					ErrMessage: fmt.Sprintf("failed to handle action: %v", err),
				}
			}

		default:
			log.Printf("TODO: handle message type '%s'", msg.Type)
		}
	}
}
