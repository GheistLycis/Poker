package app

import (
	"errors"
	"math/rand"
)

type Match struct {
	Seats      [8]*Seat
	Pot        int
	TableCards [5]Card
	Deck       *map[Card]bool
	SeatTurn   *Seat
	LastBet    int
	RoundSeats [8]*Seat
}

func NewMatch() *Match {
	deck := &map[Card]bool{
		CLUB_1:     false,
		CLUB_2:     false,
		CLUB_3:     false,
		CLUB_4:     false,
		CLUB_5:     false,
		CLUB_6:     false,
		CLUB_7:     false,
		CLUB_8:     false,
		CLUB_9:     false,
		CLUB_10:    false,
		CLUB_11:    false,
		CLUB_12:    false,
		CLUB_13:    false,
		DIAMOND_1:  false,
		DIAMOND_2:  false,
		DIAMOND_3:  false,
		DIAMOND_4:  false,
		DIAMOND_5:  false,
		DIAMOND_6:  false,
		DIAMOND_7:  false,
		DIAMOND_8:  false,
		DIAMOND_9:  false,
		DIAMOND_10: false,
		DIAMOND_11: false,
		DIAMOND_12: false,
		DIAMOND_13: false,
		HEART_1:    false,
		HEART_2:    false,
		HEART_3:    false,
		HEART_4:    false,
		HEART_5:    false,
		HEART_6:    false,
		HEART_7:    false,
		HEART_8:    false,
		HEART_9:    false,
		HEART_10:   false,
		HEART_11:   false,
		HEART_12:   false,
		HEART_13:   false,
		SPADE_1:    false,
		SPADE_2:    false,
		SPADE_3:    false,
		SPADE_4:    false,
		SPADE_5:    false,
		SPADE_6:    false,
		SPADE_7:    false,
		SPADE_8:    false,
		SPADE_9:    false,
		SPADE_10:   false,
		SPADE_11:   false,
		SPADE_12:   false,
		SPADE_13:   false,
	}
	seats := [8]*Seat{}
	for i := range seats {
		seats[i] = &Seat{
			Index: SeatIndex(i),
		}
	}

	return &Match{
		Seats:      seats,
		TableCards: [5]Card{0: BACK, 1: BACK, 2: BACK, 3: BACK, 4: BACK},
		Deck:       deck,
		SeatTurn:   seats[ZERO],
		RoundSeats: [8]*Seat{},
	}
}

func (m *Match) InitRound() {
	for i := range m.RoundSeats {
		m.RoundSeats[i] = nil
	}
	for i, s := range m.Seats {
		if s.Player != nil {
			m.RoundSeats[i] = s
		}
	}
	for _, s := range m.RoundSeats {
		var hand [2]Card

		for i := range hand {
			hand[i] = m.takeFromDeck()
		}
		s.Player.Cards = hand
	}
}

func (m *Match) PassTurn() *Seat {
	var nextSeat *Seat
	i := m.SeatTurn.Index + 1
	length := len(m.Seats)

	for range length {
		if int(i) >= length {
			i = ZERO
		}
		if next := m.Seats[i]; next.Player != nil {
			nextSeat = next
			break
		}
		i++
	}
	m.SeatTurn = nextSeat

	return nextSeat
}

func (m *Match) AllTableCardsAreRevealed() bool {
	for _, c := range m.TableCards {
		if c == BACK {
			return false
		}
	}
	return true
}

func (m *Match) RevealNextTableCard() error {
	if m.AllTableCardsAreRevealed() {
		return errors.New("all table cards are already revealed")
	}

	nextCard := m.takeFromDeck()

	(*m.Deck)[nextCard] = true
	for i, card := range m.TableCards {
		if card == BACK {
			m.TableCards[i] = nextCard
		}
	}

	return nil
}

func (m *Match) takeFromDeck() Card {
	remainingCards := []Card{}
	for card, isRevealed := range *m.Deck {
		if !isRevealed {
			remainingCards = append(remainingCards, card)
		}
	}
	nextCard := remainingCards[rand.Intn(len(remainingCards))]

	(*m.Deck)[nextCard] = true

	return nextCard
}

func (m *Match) DoPotTransaction(v int, p *Player) error {
	if v > 0 {
		if p.Score < v {
			return errors.New("player has insufficient score to pay for value")
		}
		m.Pot += v
		p.Score -= v
	} else {
		if m.Pot < v {
			return errors.New("pot has insufficient amount to pay player")
		}
		m.Pot -= v
		p.Score += v
	}

	return nil
}

func (m *Match) Showdown() []*Player {
	winners := []*Player{}

	// TODO: reveal hands, resolve pot and reset lastBet, deck and table cards

	return winners
}
