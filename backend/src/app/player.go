package app

import (
	"errors"

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
		Cards:     [2]Card{0: BACK, 1: BACK},
		SeatIndex: s,
	}
}

func (p *Player) Call(v int) error {
	if v > p.Score {
		return errors.New("player has insufficient score to call the last bet")
	}
	p.Score -= v

	return nil
}

func (p *Player) Bet() error {

	return nil
}

func (p *Player) Raise() error {

	return nil
}
