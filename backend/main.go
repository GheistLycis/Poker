package main

import (
	"log"
	"net/http"

	"backend/cmd/ws"
)

func main() {
	http.HandleFunc("/", ws.HandleWebSocketConn)
	log.Println("Server started on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
