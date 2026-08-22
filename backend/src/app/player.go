package app

import (
	"github.com/google/uuid"
)

type Player struct {
	Id        uuid.UUID
	Name      string
	Score     int
	Cards     [2]Card
	SeatIndex SeatIndex
}

func NewPlayer(n string, s SeatIndex) *Player {
	return &Player{
		Id:        uuid.New(),
		Name:      n,
		SeatIndex: s,
		Score:     10_000,
	}
}
