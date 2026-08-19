package ws

import (
	"backend/src/app"
	"net"

	"github.com/gorilla/websocket"
)

type HubMsg interface {
	isHubMsg()
}

type registerClientMsg struct {
	conn  *websocket.Conn
	reply chan *Client
}

func (registerClientMsg) isHubMsg() {}

type unregisterClientMsg struct {
	addr net.Addr
	done chan struct{}
}

func (unregisterClientMsg) isHubMsg() {}

type loginMsg struct {
	client   *Client
	userName string
	reply    chan error
}

func (loginMsg) isHubMsg() {}

type actionMsg struct {
	client *Client
	action app.PlayerAction
	amount *int
	reply  chan error
}

func (actionMsg) isHubMsg() {}

type getPlayerMsg struct {
	client *Client
	reply  chan *app.Player
}

func (getPlayerMsg) isHubMsg() {}
