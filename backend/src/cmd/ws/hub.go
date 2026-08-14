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
	endTurn    chan struct{}
}

func newHub() *Hub {
	return &Hub{
		match:      app.NewMatch(),
		clients:    map[net.Addr]*Client{},
		register:   make(chan *Client),
		unregister: make(chan net.Addr),
		broadcast:  make(chan *Message[any]),
		direct:     make(chan *DirectMessage[any]),
		endTurn:    make(chan struct{}),
	}
}

var TurnDuration = 15 * time.Second

func (h *Hub) handleTicks() {
	turnTimer := time.NewTicker(TurnDuration)
	defer turnTimer.Stop()

	h.match.InitRound()

	for {
		select {
		case <-turnTimer.C:
			h.endTurn <- struct{}{}

		case <-h.endTurn:
			h.handleNextTurn()
			turnTimer.Reset(TurnDuration)

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
		MATCH_SEAT_TURN,
		map[string]app.SeatIndex{
			"seatIndex": nextSeatTurn.Index,
		},
		nil,
	)
	h.broadcast <- seatTurnMsg.asAny()

	for _, c := range h.match.TableCards {
		if c == app.BACK {
			tableCardsMsg := newOutMessage(
				nil,
				MATCH_TABLE_CARDS,
				h.match.TableCards,
				nil,
			)

			h.broadcast <- tableCardsMsg.asAny()
			break
		}
	}
	potAmountMsg := newOutMessage(
		nil,
		MATCH_POT_AMOUNT,
		map[string]int{
			"amount": c.hub.match.Pot,
		},
		nil,
	)
	c.hub.broadcast <- potAmountMsg.asAny()
	h.sendPlayersInfo(true)
}

func (h *Hub) handleRegisterClient(c *Client) {
	h.clients[c.addr] = c
	h.sendPlayersInfo(true)
}

func (h *Hub) handleUnregisterClient(addr net.Addr) {
	playerSeatIdx := h.clients[addr].player.SeatIndex
	playerTurnIdx := h.match.SeatTurn.Index

	if playerSeatIdx == playerTurnIdx {
		h.endTurn <- struct{}{}
	}
	h.match.Seats[playerSeatIdx].Player = nil
	delete(h.clients, addr)
	h.sendPlayersInfo(true)
}

func (h *Hub) handleBroadcastMessage(m *Message[any]) {
	for _, c := range h.clients {
		c.sendMessage(m.RequestId, m.Type, m.Payload, nil)
	}
}

func (h *Hub) handleDirectMessage(dm *DirectMessage[any]) {
	client := h.clients[dm.conn]

	client.sendMessage(
		dm.message.RequestId,
		dm.message.Type,
		dm.message.Payload,
		nil,
	)
}

func (h *Hub) sendPlayersInfo(sendUserInfo bool) {
	for _, c := range h.clients {
		opponents := []*app.Player{}

		for _, o := range h.clients {
			if o != c {
				opponent := &app.Player{
					Id:        o.player.Id,
					Name:      o.player.Name,
					Score:     o.player.Score,
					Cards:     [2]app.Card{app.BACK, app.BACK},
					SeatIndex: o.player.SeatIndex,
				}
				opponents = append(opponents, opponent)
			} else if sendUserInfo {
				c.sendMessage(nil, USER_INFO, o.player, nil)
			}
		}

		c.sendMessage(nil, OPPONENTS_INFO, opponents, nil)
	}
}
