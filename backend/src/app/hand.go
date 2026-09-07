package app

import (
	"slices"
)

/*
const (

	ROYAL_FLUSH Hand = "ROYAL_FLUSH"
	STRAIGHT_FLUSH Hand = "STRAIGHT_FLUSH"
	FOUR_OF_A_KIND Hand = "FOUR_OF_A_KIND"
	FULL_HOUSE Hand = "FULL_HOUSE"
	FLUSH Hand = "FLUSH"
	STRAIGHT Hand = "STRAIGHT"
	THREE_OF_A_KIND Hand = "THREE_OF_A_KIND"
	TWO_PAIRS Hand = "TWO_PAIRS"
	ONE_PAIR Hand = "ONE_PAIR"
	HIGH_CARD Hand = "HIGH_CARD"

)
*/
type Hand string

const (
	ROYAL_FLUSH     Hand = "ROYAL_FLUSH"
	STRAIGHT_FLUSH  Hand = "STRAIGHT_FLUSH"
	FOUR_OF_A_KIND  Hand = "FOUR_OF_A_KIND"
	FULL_HOUSE      Hand = "FULL_HOUSE"
	FLUSH           Hand = "FLUSH"
	STRAIGHT        Hand = "STRAIGHT"
	THREE_OF_A_KIND Hand = "THREE_OF_A_KIND"
	TWO_PAIRS       Hand = "TWO_PAIRS"
	ONE_PAIR        Hand = "ONE_PAIR"
	HIGH_CARD       Hand = "HIGH_CARD"
)

var HandRank = [...]Hand{
	HIGH_CARD,
	ONE_PAIR,
	TWO_PAIRS,
	THREE_OF_A_KIND,
	STRAIGHT,
	FLUSH,
	FULL_HOUSE,
	FOUR_OF_A_KIND,
	STRAIGHT_FLUSH,
	ROYAL_FLUSH,
}

func hasRoyalFlush(h [7]Card) (bool, []Card) {
	highestIsAce := slices.Contains(
		[]Card{CLUB_14, DIAMOND_14, HEART_14, SPADE_14},
		getHighest(h[:]),
	)
	if !highestIsAce {
		return false, nil
	}

	return hasStraightFlush(h)
}

func hasStraightFlush(h [7]Card) (bool, []Card) {
	matches, flush := hasFlush(h)
	if !matches {
		return false, nil
	}
	fakeHand := [7]Card{}
	for i, c := range flush {
		fakeHand[i] = c
	}
	fakeHand[5] = BACK
	fakeHand[6] = BACK
	if matches, straightFlush := hasStraight(fakeHand); matches {
		return true, straightFlush
	}

	return false, nil
}

func hasFourOfAKind(h [7]Card) (bool, []Card) {
	if matches, foaks := hasNOfAKind(h[:], 4); matches {
		return true, getHighestWithSamePower(foaks)
	}

	return false, nil
}

func hasFullHouse(h [7]Card) (bool, []Card) {
	matches, threeOfAKind := hasThreeOfAKind(h)
	if !matches {
		return false, nil
	}
	matches, pairs := hasNOfAKind(h[:], 2)
	if !matches {
		return false, nil
	}
	pairs = slices.DeleteFunc(pairs, func(pair []Card) bool {
		for _, c := range pair {
			if slices.Contains(threeOfAKind, c) {
				return true
			}
		}
		return false
	})
	if len(pairs) == 0 {
		return false, nil
	}

	return true, slices.Concat(threeOfAKind, getHighestWithSamePower(pairs))
}

func hasFlush(h [7]Card) (bool, []Card) {
	suitMap := mapBySuit(h[:])
	for _, cards := range suitMap {
		if len(cards) >= 5 {
			return true, cards
		}
	}

	return false, nil
}

func hasStraight(h [7]Card) (bool, []Card) {
	powerMap := mapByPower(h[:])
	if len(powerMap) < 5 {
		return false, nil
	}

	powerSequence := []int{}
	for p := range powerMap {
		powerSequence = append(powerSequence, p)
	}
	slices.Sort(powerSequence)
	slices.Reverse(powerSequence)

	count := 1
	for i := 1; i < len(powerSequence); i++ {
		if powerSequence[i] != powerSequence[i-1]-1 {
			count = 1
			continue
		}
		count++
		if count >= 5 {
			result := []Card{}
			for j := i - 5; j < i; j++ {
				power := powerSequence[j]
				card := powerMap[power][0]
				result = append(result, card)
			}

			return true, result
		}
	}

	return false, nil
}

func hasThreeOfAKind(h [7]Card) (bool, []Card) {
	if matches, toaks := hasNOfAKind(h[:], 3); matches {
		return true, getHighestWithSamePower(toaks)
	}

	return false, nil
}

func hasTwoPairs(h [7]Card) (bool, []Card) {
	pairs := [][]Card{}
	for _, cards := range mapByPower(h[:]) {
		if len(cards) >= 2 {
			pairs = append(pairs, cards)
		}
	}

	if len(pairs) >= 2 {
		result := []Card{}
		for len(result) < 2 {
			highest := getHighestWithSamePower(pairs)
			for _, c := range highest {
				result = append(result, c)
			}
			pairs = slices.DeleteFunc(pairs, func(p []Card) bool {
				for i, c := range p {
					if highest[i] != c {
						return false
					}
				}

				return true
			})
		}

		return true, result
	}

	return false, nil
}

func hasOnePair(h [7]Card) (bool, []Card) {
	if matches, pairs := hasNOfAKind(h[:], 2); matches {
		return true, getHighestWithSamePower(pairs)
	}

	return false, nil
}
