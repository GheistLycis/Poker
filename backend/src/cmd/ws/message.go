package ws

import "github.com/google/uuid"

type Message[T any] struct {
	RequestId *uuid.UUID `json:"requestId"`
	Origin    Origin     `json:"origin"`
	Type      string     `json:"type"`
	Payload   T          `json:"payload"`
}

func newOutMessage[T any](rId *uuid.UUID, t string, p T) *Message[T] {
	return &Message[T]{
		RequestId: rId,
		Origin:    SERVER,
		Type:      t,
		Payload:   p,
	}
}

func (m *Message[T]) asAny() *Message[any] {
	return newOutMessage[any](m.RequestId, m.Type, m.Payload)
}
