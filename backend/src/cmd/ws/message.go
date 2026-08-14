package ws

import "github.com/google/uuid"

type Message[T any] struct {
	RequestId *uuid.UUID  `json:"requestId"`
	Origin    Origin      `json:"origin"`
	Type      MessageType `json:"type"`
	Payload   T           `json:"payload"`
	Error     *Error      `json:"error"`
}

func newOutMessage[T any](rId *uuid.UUID, t MessageType, p T, err *Error) *Message[T] {
	return &Message[T]{
		RequestId: rId,
		Origin:    SERVER,
		Type:      t,
		Payload:   p,
		Error:     err,
	}
}

func (m *Message[T]) asAny() *Message[any] {
	return newOutMessage[any](m.RequestId, m.Type, m.Payload, m.Error)
}
