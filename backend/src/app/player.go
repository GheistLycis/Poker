package app

import (
	"github.com/google/uuid"
)

type Player struct {
	Id        uuid.UUID
	Name      string
	Score     float32
	Cards     [2]Card
	SeatIndex SeatIndex
}

func NewPlayer(n string, s SeatIndex) *Player {
	return &Player{
		Id:        uuid.New(),
		Name:      n,
		Cards:     [2]Card{0: BACK, 1: BACK},
		SeatIndex: s,
	}
}
