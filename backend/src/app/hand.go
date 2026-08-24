package app

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

func hasRoyalFlush(h [5]Card) bool {
	hand := h[:]
	highest := getHighest(hand)
	acePower := getPower(SPADE_14) // * could be any suit

	if isSequence(hand) && getPower(highest) == acePower && isAllTheSameSuit(hand) {
		return true
	}

	return false
}

func hasStraightFlush(h [5]Card) bool {
	hand := h[:]

	if isSequence(hand) && isAllTheSameSuit(hand) {
		return true
	}

	return false
}

func hasFourOfAKind(h [5]Card) bool {
	hand := h[:]
	sortByPower(hand)

	if isAllTheSamePower(hand[:4]) || isAllTheSamePower(hand[1:]) {
		return true
	}

	return false
}

func hasFullHouse(h [5]Card) bool {
	hand := h[:]
	sortByPower(hand)

	if (isAllTheSamePower(hand[:2]) && isAllTheSamePower(hand[2:])) ||
		(isAllTheSamePower(hand[:3]) && isAllTheSamePower(hand[3:])) {
		return true
	}

	return false
}

func hasFlush(h [5]Card) bool {
	hand := h[:]

	if isAllTheSameSuit(hand) {
		return true
	}

	return false
}

func hasStraight(h [5]Card) bool {
	hand := h[:]

	if isSequence(hand) {
		return true
	}

	return false
}

func hasThreeOfAKind(h [5]Card) bool {
	hand := h[:]
	sortByPower(hand)

	if isAllTheSamePower(hand[:3]) ||
		isAllTheSamePower(hand[1:4]) ||
		isAllTheSamePower(hand[2:]) {
		return true
	}

	return false
}

func hasTwoPairs(h [5]Card) bool {
	hand := h[:]
	sortByPower(hand)
	pairValues := map[int]bool{}

	for i := 0; i < len(hand)-1; i++ {
		power := getPower(hand[i])
		if power == getPower(hand[i+1]) {
			pairValues[power] = true
		}
	}

	return len(pairValues) >= 2
}

func hasOnePair(h [5]Card) bool {
	hand := h[:]
	sortByPower(hand)

	for i := range 3 {
		if isAllTheSamePower(hand[i : i+2]) {
			return true
		}
	}

	return false
}
