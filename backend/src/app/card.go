package app

import (
	"slices"
	"strings"
)

/*
const (

	BACK       Card = "BACK"
	CLUB_2     Card = "CLUB_2"
	CLUB_3     Card = "CLUB_3"
	CLUB_4     Card = "CLUB_4"
	CLUB_5     Card = "CLUB_5"
	CLUB_6     Card = "CLUB_6"
	CLUB_7     Card = "CLUB_7"
	CLUB_8     Card = "CLUB_8"
	CLUB_9     Card = "CLUB_9"
	CLUB_10    Card = "CLUB_10"
	CLUB_11    Card = "CLUB_JACK"
	CLUB_12    Card = "CLUB_QUEEN"
	CLUB_13    Card = "CLUB_KING"
	CLUB_14    Card = "CLUB_A"
	DIAMOND_2  Card = "DIAMOND_2"
	DIAMOND_3  Card = "DIAMOND_3"
	DIAMOND_4  Card = "DIAMOND_4"
	DIAMOND_5  Card = "DIAMOND_5"
	DIAMOND_6  Card = "DIAMOND_6"
	DIAMOND_7  Card = "DIAMOND_7"
	DIAMOND_8  Card = "DIAMOND_8"
	DIAMOND_9  Card = "DIAMOND_9"
	DIAMOND_10 Card = "DIAMOND_10"
	DIAMOND_11 Card = "DIAMOND_JACK"
	DIAMOND_12 Card = "DIAMOND_QUEEN"
	DIAMOND_13 Card = "DIAMOND_KING"
	DIAMOND_14 Card = "DIAMOND_A"
	HEART_2    Card = "HEART_2"
	HEART_3    Card = "HEART_3"
	HEART_4    Card = "HEART_4"
	HEART_5    Card = "HEART_5"
	HEART_6    Card = "HEART_6"
	HEART_7    Card = "HEART_7"
	HEART_8    Card = "HEART_8"
	HEART_9    Card = "HEART_9"
	HEART_10   Card = "HEART_10"
	HEART_11   Card = "HEART_JACK"
	HEART_12   Card = "HEART_QUEEN"
	HEART_13   Card = "HEART_KING"
	HEART_14   Card = "HEART_A"
	SPADE_2    Card = "SPADE_2"
	SPADE_3    Card = "SPADE_3"
	SPADE_4    Card = "SPADE_4"
	SPADE_5    Card = "SPADE_5"
	SPADE_6    Card = "SPADE_6"
	SPADE_7    Card = "SPADE_7"
	SPADE_8    Card = "SPADE_8"
	SPADE_9    Card = "SPADE_9"
	SPADE_10   Card = "SPADE_10"
	SPADE_11   Card = "SPADE_JACK"
	SPADE_12   Card = "SPADE_QUEEN"
	SPADE_13   Card = "SPADE_KING"
	SPADE_14   Card = "SPADE_A"

)
*/
type Card string

const (
	BACK       Card = "BACK"
	CLUB_2     Card = "CLUB_2"
	CLUB_3     Card = "CLUB_3"
	CLUB_4     Card = "CLUB_4"
	CLUB_5     Card = "CLUB_5"
	CLUB_6     Card = "CLUB_6"
	CLUB_7     Card = "CLUB_7"
	CLUB_8     Card = "CLUB_8"
	CLUB_9     Card = "CLUB_9"
	CLUB_10    Card = "CLUB_10"
	CLUB_11    Card = "CLUB_JACK"
	CLUB_12    Card = "CLUB_QUEEN"
	CLUB_13    Card = "CLUB_KING"
	CLUB_14    Card = "CLUB_A"
	DIAMOND_2  Card = "DIAMOND_2"
	DIAMOND_3  Card = "DIAMOND_3"
	DIAMOND_4  Card = "DIAMOND_4"
	DIAMOND_5  Card = "DIAMOND_5"
	DIAMOND_6  Card = "DIAMOND_6"
	DIAMOND_7  Card = "DIAMOND_7"
	DIAMOND_8  Card = "DIAMOND_8"
	DIAMOND_9  Card = "DIAMOND_9"
	DIAMOND_10 Card = "DIAMOND_10"
	DIAMOND_11 Card = "DIAMOND_JACK"
	DIAMOND_12 Card = "DIAMOND_QUEEN"
	DIAMOND_13 Card = "DIAMOND_KING"
	DIAMOND_14 Card = "DIAMOND_A"
	HEART_2    Card = "HEART_2"
	HEART_3    Card = "HEART_3"
	HEART_4    Card = "HEART_4"
	HEART_5    Card = "HEART_5"
	HEART_6    Card = "HEART_6"
	HEART_7    Card = "HEART_7"
	HEART_8    Card = "HEART_8"
	HEART_9    Card = "HEART_9"
	HEART_10   Card = "HEART_10"
	HEART_11   Card = "HEART_JACK"
	HEART_12   Card = "HEART_QUEEN"
	HEART_13   Card = "HEART_KING"
	HEART_14   Card = "HEART_A"
	SPADE_2    Card = "SPADE_2"
	SPADE_3    Card = "SPADE_3"
	SPADE_4    Card = "SPADE_4"
	SPADE_5    Card = "SPADE_5"
	SPADE_6    Card = "SPADE_6"
	SPADE_7    Card = "SPADE_7"
	SPADE_8    Card = "SPADE_8"
	SPADE_9    Card = "SPADE_9"
	SPADE_10   Card = "SPADE_10"
	SPADE_11   Card = "SPADE_JACK"
	SPADE_12   Card = "SPADE_QUEEN"
	SPADE_13   Card = "SPADE_KING"
	SPADE_14   Card = "SPADE_A"
)

/*
const (

	CLUB       Suit = "CLUB"
	DIAMOND       Suit = "DIAMOND"
	HEART       Suit = "HEART"
	SPADE       Suit = "SPADE"

)
*/
type Suit string

const (
	CLUB    Suit = "CLUB"
	DIAMOND Suit = "DIAMOND"
	HEART   Suit = "HEART"
	SPADE   Suit = "SPADE"
)

var CardRank = map[Suit][13]Card{
	CLUB: [...]Card{
		CLUB_2,
		CLUB_3,
		CLUB_4,
		CLUB_5,
		CLUB_6,
		CLUB_7,
		CLUB_8,
		CLUB_9,
		CLUB_10,
		CLUB_11,
		CLUB_12,
		CLUB_13,
		CLUB_14,
	},
	DIAMOND: [...]Card{
		DIAMOND_2,
		DIAMOND_3,
		DIAMOND_4,
		DIAMOND_5,
		DIAMOND_6,
		DIAMOND_7,
		DIAMOND_8,
		DIAMOND_9,
		DIAMOND_10,
		DIAMOND_11,
		DIAMOND_12,
		DIAMOND_13,
		DIAMOND_14,
	},
	HEART: [...]Card{
		HEART_2,
		HEART_3,
		HEART_4,
		HEART_5,
		HEART_6,
		HEART_7,
		HEART_8,
		HEART_9,
		HEART_10,
		HEART_11,
		HEART_12,
		HEART_13,
		HEART_14,
	},
	SPADE: [...]Card{
		SPADE_2,
		SPADE_3,
		SPADE_4,
		SPADE_5,
		SPADE_6,
		SPADE_7,
		SPADE_8,
		SPADE_9,
		SPADE_10,
		SPADE_11,
		SPADE_12,
		SPADE_13,
		SPADE_14,
	},
}

func getSuit(c Card) Suit {
	if c == BACK {
		return ""
	}

	return Suit(strings.Split(string(c), "_")[0])
}

func getPower(c Card) int {
	suit := getSuit(c)
	if suit == "" {
		return 0
	}
	suitRank := CardRank[suit]

	return slices.Index(suitRank[:], c) + 2
}

func getHighest(cards []Card) Card {
	highest := cards[0]
	for _, c := range cards {
		if getPower(c) > getPower(highest) {
			highest = c
		}
	}

	return highest
}

func getHighestWithSamePower(list [][]Card) []Card {
	highest := list[0]
	if len(list) > 1 {
		for _, cardsWithSamePower := range list[1:] {
			if getPower(cardsWithSamePower[0]) > getPower(highest[0]) {
				highest = cardsWithSamePower
			}
		}
	}

	return highest
}

func sortByPower(cards []Card) []Card {
	sorted := make([]Card, len(cards))
	copy(sorted, cards)

	slices.SortFunc(sorted, func(a, b Card) int {
		return getPower(a) - getPower(b)
	})

	return sorted
}

func mapByPower(h []Card) map[int][]Card {
	result := map[int][]Card{}
	for _, c := range h {
		power := getPower(c)
		result[power] = append(result[power], c)
	}

	return result
}

func mapBySuit(h []Card) map[Suit][]Card {
	result := map[Suit][]Card{}
	for _, c := range h {
		suit := getSuit(c)
		result[suit] = append(result[suit], c)
	}

	return result
}

func hasNOfAKind(h []Card, n int) (bool, [][]Card) {
	result := [][]Card{}
	for _, c := range mapByPower(h) {
		if len(c) >= n {
			result = append(result, c)
		}
	}

	return len(result) > 0, result
}
