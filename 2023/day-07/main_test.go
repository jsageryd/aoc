package main

import (
	"testing"
)

var input = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 6440; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 5905; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestHandType(t *testing.T) {
	t.Run("Regular", func(t *testing.T) {
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
			if got, want := handType(tc.hand, false), tc.want; got != want {
				t.Errorf("[%d] handType(%q, false) = %d, want %d", n, tc.hand, got, want)
			}
		}
	})

	t.Run("Joker rule", func(t *testing.T) {
		for n, tc := range []struct {
			hand string
			want int
		}{
			{hand: "32T3K", want: 5}, // One pair
			{hand: "T55J5", want: 1}, // Four of a kind
			{hand: "KK677", want: 4}, // Two pair
			{hand: "KTJJT", want: 1}, // Four of a kind
			{hand: "QQQJA", want: 1}, // Four of a kind
			{hand: "QJJQ2", want: 1}, // Four of a kind
			{hand: "JJJJJ", want: 0}, // Five of a kind
		} {
			if got, want := handType(tc.hand, true), tc.want; got != want {
				t.Errorf("[%d] handType(%q, true) = %d, want %d", n, tc.hand, got, want)
			}
		}
	})
}
