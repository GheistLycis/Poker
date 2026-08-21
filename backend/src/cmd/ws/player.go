package ws

import (
	"backend/src/app"

	"github.com/google/uuid"
)

type Player struct {
	Id        uuid.UUID     `json:"id"`
	Name      string        `json:"name"`
	Score     int           `json:"score"`
	Cards     []app.Card    `json:"cards"`
	SeatIndex app.SeatIndex `json:"seatIndex"`
}

func newPlayer(p *app.Player, hideCards bool) *Player {
	cards := []app.Card{}

	if hideCards {
		cards = []app.Card{app.BACK, app.BACK}
	} else if p.Cards != [2]app.Card{} {
		cards = p.Cards[:]
	}

	return &Player{
		Id:        p.Id,
		Name:      p.Name,
		Score:     p.Score,
		Cards:     cards,
		SeatIndex: p.SeatIndex,
	}
}
