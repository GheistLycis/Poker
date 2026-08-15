package ws

type Error struct {
	Message string `json:"message"`
	Details any    `json:"details"`
}

func newError(m string, d any) *Error {
	return &Error{
		Message: m,
		Details: d,
	}
}
