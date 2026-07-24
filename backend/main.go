package main

import (
	"log"
	"net/http"

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

// Define your request and response structs
type Message struct {
	Origin Origin `json:"origin"`
	Type  string `json:"type"`
	Payload map[int]*string `json:"payload"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for testing
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// Receive and send loop
	for {
		var msg Message

		// ReadJSON automatically decodes the incoming JSON payload into the struct
		err := conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		log.Printf("Received message: %s\n", msg.Payload)
		payload := map[int]*string{}

		user := "user"
		opponent := "opponent"

		payload[0] = &user
		payload[1] = &opponent


		// Prepare a response struct
		reply := Message{
			Origin:  "SERVER",
			Type: "match.seats",
			Payload: payload,
		}

		// WriteJSON automatically encodes the struct to JSON and sends it
		err = conn.WriteJSON(reply)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("Server started on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
