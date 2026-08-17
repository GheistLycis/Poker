package ws

import "github.com/google/uuid"

type Message[T any] struct {
	RequestId *uuid.UUID    `json:"requestId"`
	Origin    Origin        `json:"origin"`
	Type      MessageType   `json:"type"`
	Payload   T             `json:"payload"`
	Error     *MessageError `json:"error"`
}

type MessageError struct {
	Message string `json:"message"`
	Details any    `json:"details"`
}

type ServerMessageArgs[T any] struct {
	RequestId  *uuid.UUID
	Type       MessageType
	Payload    T
	ErrMessage string
	ErrDetails any
}

func newServerMessage[T any](p ServerMessageArgs[T]) *Message[T] {
	var msgErr *MessageError
	if p.ErrMessage != "" {
		msgErr = &MessageError{
			Message: p.ErrMessage,
			Details: p.ErrDetails,
		}
	}

	return &Message[T]{
		RequestId: p.RequestId,
		Origin:    SERVER,
		Type:      p.Type,
		Payload:   p.Payload,
		Error:     msgErr,
	}
}
