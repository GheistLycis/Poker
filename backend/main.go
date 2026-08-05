package main

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

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
	log.Println("client connected:", conn.RemoteAddr())
	defer log.Println("handler exited:", conn.RemoteAddr())
	defer conn.Close()

	for {
		var msg Message[any]

		if err := conn.ReadJSON(&msg); err != nil {
			log.Printf("read error (%T): %v", err, err)
			break
		}
		log.Printf("RECEIVED: type=%s payload=%+v", msg.Type, msg.Payload)

		if msg.Type == "user.login" {
			payload, ok := msg.Payload.(map[string]any)
			if !ok {
				log.Printf("400: bad payload")
			}
			userName, ok := payload["userName"].(string)
			if !ok {
				log.Printf("400: userName must be a valid string")
			}
			var loginRes Message[map[string]any]

			loginRes.RequestId = msg.RequestId
			loginRes.Origin = "SERVER"
			loginRes.Type = "user.info"
			loginRes.Payload = map[string]any{
				"name":  userName,
				"score": 0,
				"cards": []any{},
			}
			if err := conn.WriteJSON(&loginRes); err != nil {
				log.Printf("write error (%T): %v", err, err)
			}

			var playerTurnRes Message[map[string]int]

			playerTurnRes.Origin = "SERVER"
			playerTurnRes.Type = "match.seat-turn"
			playerTurnRes.Payload = map[string]int{"seatIndex": 0}
			if err := conn.WriteJSON(&playerTurnRes); err != nil {
				log.Printf("write error (%T): %v", err, err)
			}
		}
	}
}

func main() {
	http.HandleFunc("/", handleWebSocket)
	log.Println("Server started on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
