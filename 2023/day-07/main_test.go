package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}

	if got, want := part1(input), 6440; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestHandType(t *testing.T) {
	for n, tc := range []struct {
		hand string
		want int
	}{
		{hand: "AAAAA", want: 0}, // Five of a kind
		{hand: "AA8AA", want: 1}, // Four of a kind
		{hand: "23332", want: 2}, // Full house
		{hand: "TTT98", want: 3}, // Three of a kind
		{hand: "23432", want: 4}, // Two pair
		{hand: "A23A4", want: 5}, // One pair
		{hand: "23456", want: 6}, // High card
	} {
		if got, want := handType(tc.hand), tc.want; got != want {
			t.Errorf("[%d] handType(%q) = %d, want %d", n, tc.hand, got, want)
		}
	}
}
