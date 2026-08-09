package app

import (
	"backend/cmd/ws"

	"github.com/gorilla/websocket"
)

type Player struct {
	name      string
	score     float32
	cards     []Card
	seatIndex int
	conn      *websocket.Conn
	send      chan *ws.Message[any]
}
