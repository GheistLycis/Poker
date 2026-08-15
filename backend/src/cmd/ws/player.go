package ws

import (
	"backend/src/app"
	"encoding/json"

	"github.com/google/uuid"
)

type Player struct {
	Id        uuid.UUID     `json:"id"`
	Name      string        `json:"name"`
	Score     int           `json:"score"`
	Cards     Cards         `json:"cards"`
	SeatIndex app.SeatIndex `json:"seatIndex"`
}

type Cards [2]app.Card

func (h Cards) MarshalJSON() ([]byte, error) {
	if h == (Cards{}) {
		return []byte("[]"), nil
	}

	return json.Marshal([2]app.Card(h))
}

func (h *Cards) UnmarshalJSON(data []byte) error {
	var raw []app.Card

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*h = Cards{}
	copy(h[:], raw)

	return nil
}
