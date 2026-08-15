package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func InitServer(r string, p int) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // TODO
		},
	}
	hub := newHub()
	port := ":" + fmt.Sprintf("%d", p)

	go hub.handleTicks()
	http.HandleFunc(r, func(w http.ResponseWriter, r *http.Request) {
		handleNewConn(hub, upgrader, w, r)
	})
	log.Println("Server started on " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleNewConn(h *Hub, u websocket.Upgrader, w http.ResponseWriter, r *http.Request) {
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
	defer conn.Close()
	log.Println("new conn stablished:", conn.RemoteAddr())
	client := h.registerClient(conn)
	defer h.unregisterClient(client.addr)
	client.handleMessages()
}
