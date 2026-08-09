package ws

import (
	"backend/src/app"
	"net"
)

type Hub struct {
	match      *app.Match
	clients    map[net.Addr]*Client
	register   chan *Client
	unregister chan net.Addr
	broadcast  chan *Message[any]
	direct     chan *DirectMessage[any]
}

func newHub() *Hub {
	return &Hub{
		match:      app.NewMatch(),
		clients:    map[net.Addr]*Client{},
		register:   make(chan *Client),
		unregister: make(chan net.Addr),
		broadcast:  make(chan *Message[any]),
		direct:     make(chan *DirectMessage[any]),
	}
}

func (h *Hub) handleTicks() {
	for {
		select {
		case c := <-h.register:
			h.clients[c.addr] = c

		case addr := <-h.unregister:
			delete(h.clients, addr)

		case m := <-h.broadcast:
			for _, c := range h.clients {
				c.sendMessage(m.RequestId, m.Type, m.Payload)
			}

		case dm := <-h.direct:
			client := h.clients[dm.conn]

			client.sendMessage(
				dm.message.RequestId,
				dm.message.Type,
				dm.message.Payload,
			)
		}
	}
}
