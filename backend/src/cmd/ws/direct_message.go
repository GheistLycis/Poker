package ws

import (
	"net"

	"github.com/google/uuid"
)

type DirectMessage[T any] struct {
	conn    net.Addr
	message *Message[T]
}

func newDirectMessage[T any](c net.Addr, rId *uuid.UUID, t string, p T) *DirectMessage[T] {
	return &DirectMessage[T]{
		conn:    c,
		message: newOutMessage(rId, t, p),
	}
}
