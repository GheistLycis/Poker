package ws

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

/*
const (

	SERVER   Origin = "SERVER"
	CLIENT Origin = "CLIENT"

)
*/
type Origin string

const (
	SERVER Origin = "SERVER"
	CLIENT Origin = "CLIENT"
)

type Message[T any] struct {
	RequestId *uuid.UUID `json:"requestId"`
	Origin    Origin     `json:"origin"`
	Type      string     `json:"type"`
	Payload   T          `json:"payload"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocketConn(w http.ResponseWriter, r *http.Request) {
	conn, err := stablishConn(w, r)
	if err != nil {
		return
	}

	for {
		msg, err := readNextMessage(conn)
		if err != nil {
			break
		}
		if msg.Type == "user.login" {
			payload, ok := msg.Payload.(map[string]any)
			if !ok {
				log.Printf("400: bad payload")
			}
			userName, ok := payload["userName"].(string)
			if !ok {
				log.Printf("400: userName must be a valid string")
			}

			loginRes := &Message[map[string]any]{}

			loginRes.RequestId = msg.RequestId
			loginRes.Type = "user.info"
			loginRes.Payload = map[string]any{
				"name":  userName,
				"score": 0,
				"cards": []any{},
			}
			if err = sendMessage(conn, loginRes); err != nil {
				break
			}

			playerTurnRes := &Message[map[string]int]{}

			playerTurnRes.Type = "match.seat-turn"
			playerTurnRes.Payload = map[string]int{"seatIndex": 0}
			if err = sendMessage(conn, playerTurnRes); err != nil {
				break
			}
		}
	}
}

func stablishConn(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return conn, err
	}

	log.Println("client connected:", conn.RemoteAddr())
	defer log.Println("handler exited:", conn.RemoteAddr())
	defer conn.Close()

	return conn, nil
}

func readNextMessage(c *websocket.Conn) (*Message[any], error) {
	msg := &Message[any]{}

	if err := c.ReadJSON(msg); err != nil {
		log.Printf("read error (%T): %v", err, err)
		return msg, err
	}
	log.Printf("RECEIVED: type=%s payload=%+v", msg.Type, msg.Payload)

	return msg, nil
}

func sendMessage[T any](c *websocket.Conn, m *Message[T]) error {
	m.Origin = "SERVER"
	if err := c.WriteJSON(m); err != nil {
		log.Printf("write error (%T): %v", err, err)
		return err
	}

	return nil
}
