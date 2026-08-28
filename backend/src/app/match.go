package app

import (
	"errors"
	"math/rand"
	"slices"
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
		CLUB_14:    false,
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
		DIAMOND_14: false,
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
		HEART_14:   false,
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
		SPADE_14:   false,
	}
	seats := [8]*Seat{}
	for i := range seats {
		seats[i] = &Seat{
			Index: SeatIndex(i),
		}
	}

	return &Match{
		Seats: seats,
		Deck:  deck,
	}
}

func (m *Match) InitRound() {
	m.LastBet = 0
	for card := range *m.Deck {
		(*m.Deck)[card] = false
	}
	m.TableCards = [...]Card{BACK, BACK, BACK, BACK, BACK}
	for range 3 {
		m.RevealNextTableCard()
	}
	for i := range m.RoundSeats {
		m.RoundSeats[i] = nil
	}
	for i, s := range m.Seats {
		if s.Player != nil {
			m.RoundSeats[i] = s
		}
	}
	for _, s := range m.RoundSeats {
		if s == nil || s.Player == nil {
			continue
		}

		var hand [2]Card

		for i := range hand {
			hand[i] = m.takeFromDeck()
		}
		s.Player.Cards = hand
	}
	m.SeatTurn = m.RoundSeats[0]
}

func (m *Match) PassTurn() *Seat {
	var nextSeat *Seat
	var currentRoundSeatIdx int
	n := len(m.RoundSeats)

	for i := 1; i <= n; i++ {
		next := m.RoundSeats[(currentRoundSeatIdx+i)%n]
		if next != nil && next.Player != nil {
			nextSeat = next
			break
		}
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

	for i, card := range m.TableCards {
		if card == BACK {
			m.TableCards[i] = nextCard
			break
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

/*
v < 0 transfers from pot to player; v > 0 transfers from player to pot

returns error if payer has unsufficient balance to pay
*/
func (m *Match) DoPotTransaction(v int, p *Player) error {
	if v > 0 {
		if p.Score < v {
			return errors.New("player has insufficient score to pay for value")
		}
		m.Pot += v
		p.Score -= v
	} else {
		value := -v
		if m.Pot < value {
			return errors.New("pot has insufficient amount to pay player")
		}
		m.Pot -= value
		p.Score += value
	}

	return nil
}

type playerHand struct {
	hand        Hand
	highestCard Card
}

func (m *Match) Showdown() []*Player {
	winners := []*Player{}

	handMap := map[*Player]playerHand{}
	for _, s := range m.RoundSeats {
		player := s.Player
		hand, highestCard := m.calculateHand(player.Cards)
		handMap[player] = playerHand{hand: hand, highestCard: highestCard}
	}

	highestHand := HIGH_CARD
	for _, h := range handMap {
		if slices.Index(HandRank[:], h.hand) > slices.Index(HandRank[:], highestHand) {
			highestHand = h.hand
		}
	}

	for p, h := range handMap {
		if h.hand == highestHand {
			winners = append(winners, p)
		}
	}

	if len(winners) > 1 {
		highestCard := BACK
		for _, p := range winners {
			playerHighestCard := handMap[p].highestCard
			if getPower(playerHighestCard) > getPower(highestCard) {
				highestCard = playerHighestCard
			}
		}

		playersWithHighestCard := []*Player{}
		for _, p := range winners {
			if handMap[p].highestCard == highestCard {
				playersWithHighestCard = append(playersWithHighestCard, p)
			}
		}

		winners = playersWithHighestCard
	}

	paymentAmount := m.Pot / len(winners)
	for _, w := range winners {
		m.DoPotTransaction(-paymentAmount, w)
	}

	return winners
}

func (m *Match) calculateHand(h [2]Card) (Hand, Card) {
	hand := [7]Card(slices.Concat(h[:], m.TableCards[:]))
	highestCard := getHighest(hand[:]) // TODO: should come from the matching hand, not the hand as a whole

	if hasRoyalFlush(hand) {
		return ROYAL_FLUSH, highestCard
	}
	if hasStraightFlush(hand) {
		return STRAIGHT_FLUSH, highestCard
	}
	if hasFourOfAKind(hand) {
		return FOUR_OF_A_KIND, highestCard
	}
	if hasFullHouse(hand) {
		return FULL_HOUSE, highestCard
	}
	if hasFlush(hand) {
		return FLUSH, highestCard
	}
	if hasStraight(hand) {
		return STRAIGHT, highestCard
	}
	if hasThreeOfAKind(hand) {
		return THREE_OF_A_KIND, highestCard
	}
	if hasTwoPairs(hand) {
		return TWO_PAIRS, highestCard
	}
	if hasOnePair(hand) {
		return ONE_PAIR, highestCard
	}
	return HIGH_CARD, highestCard
}
