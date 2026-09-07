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

func hasRoyalFlush(h [7]Card) (bool, [5]Card) {
	matches, straightFlush := hasStraightFlush(h)
	if !matches {
		return false, [5]Card{}
	}

	highestIsAce := slices.Contains(
		[]Card{CLUB_14, DIAMOND_14, HEART_14, SPADE_14},
		getHighest(straightFlush[:]),
	)
	if !highestIsAce {
		return false, [5]Card{}
	}

	return true, straightFlush
}

func hasStraightFlush(h [7]Card) (bool, [5]Card) {
	for _, cards := range mapBySuit(h[:]) {
		if len(cards) >= 5 {
			fakeHand := [7]Card{}
			copy(fakeHand[:], cards)
			for i := len(cards); i < 7; i++ {
				fakeHand[i] = BACK
			}
			if matches, straightFlush := hasStraight(fakeHand); matches {
				return true, straightFlush
			}
		}
	}

	return false, [5]Card{}
}

func hasFourOfAKind(h [7]Card) (bool, [4]Card) {
	if matches, foaks := hasNOfAKind(h[:], 4); matches {
		highestFoak := getHighestOfAKind(foaks)

		return true, [4]Card(sortByPower(highestFoak, false))
	}

	return false, [4]Card{}
}

func hasFullHouse(h [7]Card) (bool, [5]Card) {
	matches, threeOfAKind := hasThreeOfAKind(h)
	if !matches {
		return false, [5]Card{}
	}
	matches, pairs := hasNOfAKind(h[:], 2)
	if !matches {
		return false, [5]Card{}
	}
	pairs = slices.DeleteFunc(pairs, func(pair []Card) bool {
		for _, c := range pair {
			if slices.Contains(threeOfAKind[:], c) {
				return true
			}
		}
		return false
	})
	if len(pairs) == 0 {
		return false, [5]Card{}
	}
	highestPair := getHighestOfAKind(pairs)
	cappedPair := [2]Card(sortByPower(highestPair, false))
	result := slices.Concat(threeOfAKind[:], cappedPair[:])

	return true, [5]Card(result)
}

func hasFlush(h [7]Card) (bool, [5]Card) {
	suitMap := mapBySuit(h[:])
	for _, cards := range suitMap {
		if len(cards) >= 5 {
			return true, [5]Card(sortByPower(cards, false))
		}
	}

	return false, [5]Card{}
}

func hasStraight(h [7]Card) (bool, [5]Card) {
	powerMap := mapByPower(h[:])
	if len(powerMap) < 5 {
		return false, [5]Card{}
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
			for j := i - 4; j <= i; j++ {
				power := powerSequence[j]
				card := powerMap[power][0]
				result = append(result, card)
			}

			return true, [5]Card(result)
		}
	}

	return false, [5]Card{}
}

func hasThreeOfAKind(h [7]Card) (bool, [3]Card) {
	if matches, toaks := hasNOfAKind(h[:], 3); matches {
		highestToak := getHighestOfAKind(toaks)

		return true, [3]Card(sortByPower(highestToak, false))
	}

	return false, [3]Card{}
}

func hasTwoPairs(h [7]Card) (bool, [4]Card) {
	powerMap := mapByPower(h[:])
	powers := []int{}
	for p, cards := range powerMap {
		if len(cards) >= 2 {
			powers = append(powers, p)
		}
	}
	if len(powers) < 2 {
		return false, [4]Card{}
	}
	slices.Sort(powers)
	slices.Reverse(powers)

	result := []Card{}
	result = append(result, powerMap[powers[0]][:2]...)
	result = append(result, powerMap[powers[1]][:2]...)

	return true, [4]Card(result)
}

func hasOnePair(h [7]Card) (bool, [2]Card) {
	if matches, pairs := hasNOfAKind(h[:], 2); matches {
		highestPair := getHighestOfAKind(pairs)

		return true, [2]Card(sortByPower(highestPair, false))
	}

	return false, [2]Card{}
}
