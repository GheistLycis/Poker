package app

import (
	"fmt"
	"testing"
)

type testCase struct {
	expected bool
	hand     [7]Card
}

func run(t *testing.T, tests []testCase) {
	for _, tt := range tests {
		tName := fmt.Sprint(tt.hand)

		t.Run(tName, func(t *testing.T) {
			if res := hasRoyalFlush(tt.hand); res != tt.expected {
				t.Errorf("%v: got %v; expected %v", tt.hand, res, tt.expected)
			}
		})
	}
}

func TestHasRoyalFlush(t *testing.T) {
	tests := []testCase{
		{
			expected: true,
			hand:     [7]Card{SPADE_14, CLUB_2, SPADE_11, DIAMOND_5, SPADE_13, SPADE_10, SPADE_12},
		},
		{
			expected: false,
			hand:     [7]Card{HEART_5, HEART_9, CLUB_2, HEART_7, HEART_6, HEART_8, DIAMOND_3},
		},
		{
			expected: false,
			hand:     [7]Card{DIAMOND_2, DIAMOND_4, DIAMOND_7, DIAMOND_9, DIAMOND_14, CLUB_3, HEART_5},
		},
		{
			expected: false,
			hand:     [7]Card{SPADE_10, HEART_11, CLUB_12, DIAMOND_13, SPADE_14, CLUB_2, HEART_3},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_2, DIAMOND_5, HEART_9, SPADE_11, CLUB_13, DIAMOND_3, HEART_7},
		},
	}

	run(t, tests)
}

func TestHasStraightFlush(t *testing.T) {
	tests := []testCase{
		{
			expected: true,
			hand:     [7]Card{HEART_9, CLUB_2, HEART_5, DIAMOND_3, HEART_7, HEART_6, HEART_8},
		},
		{
			expected: true,
			hand:     [7]Card{SPADE_14, CLUB_2, SPADE_11, DIAMOND_5, SPADE_13, SPADE_10, SPADE_12},
		},
		{
			expected: false,
			hand:     [7]Card{DIAMOND_2, DIAMOND_4, DIAMOND_7, DIAMOND_9, DIAMOND_14, CLUB_3, HEART_5},
		},
		{
			expected: false,
			hand:     [7]Card{SPADE_10, HEART_11, CLUB_12, DIAMOND_13, SPADE_14, CLUB_2, HEART_3},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_2, DIAMOND_5, HEART_9, SPADE_11, CLUB_13, DIAMOND_3, HEART_7},
		},
	}

	run(t, tests)
}

func TestHasFourOfAKind(t *testing.T) {
	tests := []testCase{
		{
			expected: true,
			hand:     [7]Card{CLUB_9, CLUB_2, DIAMOND_9, DIAMOND_5, HEART_9, HEART_7, SPADE_9},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_7, DIAMOND_5, DIAMOND_7, HEART_7, SPADE_2, CLUB_11, HEART_13},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_4, DIAMOND_4, HEART_9, SPADE_9, CLUB_2, DIAMOND_5, HEART_11},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_2, DIAMOND_5, HEART_9, SPADE_11, CLUB_13, DIAMOND_3, HEART_7},
		},
	}

	run(t, tests)
}

func TestHasFullHouse(t *testing.T) {
	tests := []testCase{
		{
			expected: true,
			hand:     [7]Card{CLUB_2, CLUB_5, DIAMOND_2, SPADE_9, DIAMOND_5, HEART_2, HEART_10},
		},
		{
			expected: true,
			hand:     [7]Card{CLUB_5, CLUB_9, DIAMOND_5, SPADE_2, DIAMOND_9, HEART_5, HEART_9},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_7, CLUB_2, DIAMOND_7, SPADE_5, DIAMOND_11, HEART_7, HEART_13},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_4, CLUB_2, DIAMOND_4, SPADE_9, DIAMOND_5, HEART_9, HEART_11},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_9, CLUB_2, DIAMOND_9, DIAMOND_5, HEART_9, HEART_7, SPADE_9},
		},
	}

	run(t, tests)
}

func TestHasFlush(t *testing.T) {
	tests := []testCase{
		{
			expected: true,
			hand:     [7]Card{HEART_2, CLUB_9, HEART_5, DIAMOND_3, HEART_9, HEART_11, HEART_14},
		},
		{
			expected: false,
			hand:     [7]Card{HEART_2, CLUB_9, HEART_5, DIAMOND_3, HEART_9, HEART_11, SPADE_14},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_2, DIAMOND_5, HEART_9, SPADE_11, CLUB_13, DIAMOND_3, HEART_7},
		},
	}

	run(t, tests)
}

func TestHasStraight(t *testing.T) {
	tests := []testCase{
		{
			expected: true,
			hand:     [7]Card{CLUB_5, DIAMOND_2, DIAMOND_6, HEART_3, HEART_7, SPADE_8, CLUB_9},
		},
		{
			expected: true,
			hand:     [7]Card{CLUB_5, DIAMOND_6, HEART_7, SPADE_8, CLUB_9, DIAMOND_9, HEART_2},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_2, DIAMOND_4, HEART_7, SPADE_9, CLUB_11, DIAMOND_13, HEART_3},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_2, DIAMOND_3, HEART_4, SPADE_5, CLUB_9, DIAMOND_10, HEART_13},
		},
		{
			expected: false,
			hand:     [7]Card{SPADE_14, CLUB_2, DIAMOND_3, HEART_4, SPADE_5, CLUB_9, DIAMOND_11},
		},
	}

	run(t, tests)
}

func TestHasThreeOfAKind(t *testing.T) {
	tests := []testCase{
		{
			expected: true,
			hand:     [7]Card{CLUB_7, CLUB_2, DIAMOND_7, SPADE_5, DIAMOND_11, HEART_7, HEART_13},
		},
		{
			expected: true,
			hand:     [7]Card{CLUB_9, CLUB_2, DIAMOND_9, DIAMOND_5, HEART_9, HEART_7, SPADE_9},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_4, CLUB_2, DIAMOND_4, SPADE_9, DIAMOND_5, HEART_9, HEART_11},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_2, DIAMOND_5, HEART_9, SPADE_11, CLUB_13, DIAMOND_3, HEART_7},
		},
	}

	run(t, tests)
}

func TestHasTwoPairs(t *testing.T) {
	tests := []testCase{
		{
			expected: true,
			hand:     [7]Card{CLUB_4, CLUB_2, DIAMOND_4, SPADE_9, DIAMOND_5, HEART_9, HEART_11},
		},
		{
			expected: true,
			hand:     [7]Card{CLUB_3, DIAMOND_3, HEART_7, SPADE_9, CLUB_11, DIAMOND_13, HEART_13},
		},
		{
			expected: true,
			hand:     [7]Card{CLUB_2, CLUB_5, DIAMOND_2, SPADE_9, DIAMOND_5, HEART_2, HEART_10},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_6, DIAMOND_6, HEART_2, SPADE_5, CLUB_9, DIAMOND_11, HEART_13},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_7, CLUB_2, DIAMOND_7, SPADE_5, DIAMOND_11, HEART_7, HEART_13},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_2, DIAMOND_5, HEART_9, SPADE_11, CLUB_13, DIAMOND_3, HEART_7},
		},
	}

	run(t, tests)
}

func TestHasOnePair(t *testing.T) {
	tests := []testCase{
		{
			expected: true,
			hand:     [7]Card{CLUB_6, CLUB_9, DIAMOND_6, SPADE_5, HEART_11, HEART_13, DIAMOND_2},
		},
		{
			expected: true,
			hand:     [7]Card{CLUB_9, CLUB_2, DIAMOND_9, DIAMOND_5, HEART_9, HEART_7, SPADE_9},
		},
		{
			expected: false,
			hand:     [7]Card{CLUB_2, DIAMOND_5, HEART_9, SPADE_11, CLUB_13, DIAMOND_3, HEART_7},
		},
	}

	run(t, tests)
}
