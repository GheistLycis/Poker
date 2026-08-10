package ws

import (
	"backend/src/app"
	"net"
	"time"
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
	turnTimer := time.NewTicker(20 * time.Second)

	defer turnTimer.Stop()
	for {
		select {
		case <-turnTimer.C:
			h.handleNextTurn()

		case c := <-h.register:
			h.handleRegisterClient(c)

		case addr := <-h.unregister:
			h.handleUnregisterClient(addr)

		case m := <-h.broadcast:
			h.handleBroadcastMessage(m)

		case dm := <-h.direct:
			h.handleDirectMessage(dm)
		}
	}
}

func (h *Hub) handleNextTurn() {
	nextSeatTurn := h.match.PassTurn()
	seatTurnMsg := newOutMessage(
		nil,
		"match.seat-turn",
		map[string]app.SeatIndex{
			"seatIndex": nextSeatTurn.Index,
		},
	)

	h.broadcast <- seatTurnMsg.asAny()
}

func (h *Hub) handleRegisterClient(c *Client) {
	h.clients[c.addr] = c
}

func (h *Hub) handleUnregisterClient(addr net.Addr) {
	delete(h.clients, addr)
}

func (h *Hub) handleBroadcastMessage(m *Message[any]) {
	for _, c := range h.clients {
		c.sendMessage(m.RequestId, m.Type, m.Payload)
	}
}

func (h *Hub) handleDirectMessage(dm *DirectMessage[any]) {
	client := h.clients[dm.conn]

	client.sendMessage(
		dm.message.RequestId,
		dm.message.Type,
		dm.message.Payload,
	)
}
