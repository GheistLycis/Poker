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

func hasRoyalFlush(h [7]Card) bool {
	highestIsAce := slices.Contains(
		[]Card{CLUB_14, DIAMOND_14, HEART_14, SPADE_14},
		getHighest(h[:]),
	)

	return highestIsAce && hasStraightFlush(h)
}

func hasStraightFlush(h [7]Card) bool {
	suitMap := map[Suit][]Card{}
	for _, c := range h {
		if c != BACK {
			suit := getSuit(c)
			suitMap[suit] = append(suitMap[suit], c)
		}
	}

	for _, cards := range suitMap {
		if len(cards) >= 5 {
			hand := [7]Card{}
			copy(hand[:], cards)
			for i := range hand {
				if hand[i] == "" {
					hand[i] = BACK
				}
			}
			if hasStraight(hand) {
				return true
			}
		}
	}

	return false
}

func hasFourOfAKind(h [7]Card) bool {
	return hasNOfAKind(h[:], 4)
}

func hasFullHouse(h [7]Card) bool {
	hasTriple := false
	pairsOrBetter := 0
	for _, count := range powerCounts(h[:]) {
		if count >= 3 {
			hasTriple = true
		}
		if count >= 2 {
			pairsOrBetter++
		}
	}

	return hasTriple && pairsOrBetter >= 2
}

func hasFlush(h [7]Card) bool {
	suitMap := map[Suit]int{}
	for _, c := range h {
		suit := getSuit(c)
		suitMap[suit]++
		if suitMap[suit] == 5 {
			return true
		}
	}

	return false
}

func hasStraight(h [7]Card) bool {
	powerSet := map[int]bool{}
	for _, c := range h {
		if c != BACK {
			powerSet[getPower(c)] = true
		}
	}

	powers := make([]int, 0, len(powerSet))
	for p := range powerSet {
		powers = append(powers, p)
	}
	slices.Sort(powers)

	run := 1
	for i := 1; i < len(powers); i++ {
		if powers[i] != powers[i-1]+1 {
			run = 1
			continue
		}
		run++
		if run >= 5 {
			return true
		}
	}

	return false
}

func hasThreeOfAKind(h [7]Card) bool {
	return hasNOfAKind(h[:], 3)
}

func hasTwoPairs(h [7]Card) bool {
	pairs := 0
	for _, count := range powerCounts(h[:]) {
		if count >= 2 {
			pairs++
		}
	}

	return pairs >= 2
}

func hasOnePair(h [7]Card) bool {
	return hasNOfAKind(h[:], 2)
}
