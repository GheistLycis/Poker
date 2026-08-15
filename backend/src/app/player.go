package app

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Player struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Score     int       `json:"score"`
	Cards     CardHand  `json:"cards"`
	SeatIndex SeatIndex `json:"seatIndex"`
}

func NewPlayer(n string, s SeatIndex) *Player {
	return &Player{
		Id:        uuid.New(),
		Name:      n,
		SeatIndex: s,
	}
}

type CardHand [2]Card

func (h CardHand) MarshalJSON() ([]byte, error) {
	if h == (CardHand{}) {
		return []byte("[]"), nil
	}

	return json.Marshal([2]Card(h))
}

func (h *CardHand) UnmarshalJSON(data []byte) error {
	var raw []Card

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*h = CardHand{}
	copy(h[:], raw)

	return nil
}
