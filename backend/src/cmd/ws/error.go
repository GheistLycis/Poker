package ws

type Error struct {
	message string
	details any
}

func newError(m string, d any) *Error {
	return &Error{
		message: m,
		details: d,
	}
}
