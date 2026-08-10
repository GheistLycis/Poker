package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func InitServer(r string, p int) {
	hub := newHub()
	port := ":" + fmt.Sprintf("%d", p)

	go hub.handleTicks()
	http.HandleFunc(r, func(w http.ResponseWriter, r *http.Request) {
		handleNewConn(hub, w, r)
	})
	log.Println("Server started on " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleNewConn(h *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
	connAddr := conn.RemoteAddr()
	log.Println("client connected:", connAddr)
	defer log.Println("client disconnected:", connAddr)
	defer conn.Close()

	client := newClient(conn, h)

	h.register <- client
	client.handleMessages()
}
