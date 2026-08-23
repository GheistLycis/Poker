package app

import "slices"

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

func getHandStats(h [5]Card) ([5]int, [5]Suit) {
	powers := [5]int{}
	suits := [5]Suit{}
	for i, c := range h {
		powers[i] = getPower(c)
		suits[i] = getSuit(c)
	}

	return powers, suits
}

func hasRoyalFlush(h [5]Card) bool {
	powers, suits := getHandStats(h)
	slices.Sort(powers[:])

	if isSequence(powers[:]) && powers[4] == 14 && isAllTheSameSuit(suits[:]) {
		return true
	}

	return false
}

func hasStraightFlush(h [5]Card) bool {
	powers, suits := getHandStats(h)

	if isSequence(powers[:]) && isAllTheSameSuit(suits[:]) {
		return true
	}

	return false
}

func hasFourOfAKind(h [5]Card) bool {
	powers, _ := getHandStats(h)
	slices.Sort(powers[:])

	if isAllTheSamePower(powers[:4]) || isAllTheSamePower(powers[1:]) {
		return true
	}

	return false
}

func hasFullHouse(h [5]Card) bool {
	powers, _ := getHandStats(h)
	slices.Sort(powers[:])

	if (isAllTheSamePower(powers[:2]) && isAllTheSamePower(powers[2:])) ||
		(isAllTheSamePower(powers[:3]) && isAllTheSamePower(powers[3:])) {
		return true
	}

	return false
}

func hasFlush(h [5]Card) bool {
	_, suits := getHandStats(h)

	if isAllTheSameSuit(suits[:]) {
		return true
	}

	return false
}

func hasStraight(h [5]Card) bool {
	powers, _ := getHandStats(h)

	if isSequence(powers[:]) {
		return true
	}

	return false
}

func hasThreeOfAKind(h [5]Card) bool {
	powers, _ := getHandStats(h)
	slices.Sort(powers[:])

	if isAllTheSamePower(powers[:3]) ||
		isAllTheSamePower(powers[1:4]) ||
		isAllTheSamePower(powers[2:]) {
		return true
	}

	return false
}

func hasTwoPairs(h [5]Card) bool {
	powers, _ := getHandStats(h)
	slices.Sort(powers[:])
	pairValues := map[int]bool{}

	for i := range len(powers) - 1 {
		if powers[i] == powers[i+1] {
			pairValues[powers[i]] = true
		}
	}

	return len(pairValues) >= 2
}

func hasOnePair(h [5]Card) bool {
	powers, _ := getHandStats(h)
	slices.Sort(powers[:])

	if isAllTheSamePower(powers[:2]) {
		return true
	}
	if isAllTheSamePower(powers[1:3]) {
		return true
	}
	if isAllTheSamePower(powers[2:4]) {
		return true
	}
	if isAllTheSamePower(powers[3:]) {
		return true
	}

	return false
}
