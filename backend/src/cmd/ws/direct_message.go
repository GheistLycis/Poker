package ws

import (
	"net"
)

type DirectMessage[T any] struct {
	conn    net.Addr
	message *Message[T]
}
