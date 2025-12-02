package main

import (
	"slices"
	"testing"
)

func TestPart1(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	if got, want := part1(input), 1227775554; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestDigits(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{0, 1},
		{1, 1},
		{9, 1},
		{10, 2},
		{99, 2},
		{100, 3},
		{999, 3},
		{1000, 4},
		{9999, 4},
		{10000, 5},
	} {
		if got, want := digits(tc.n), tc.want; got != want {
			t.Errorf("digits(%d) = %d, want %d", tc.n, got, want)
		}
	}
}

func TestParse(t *testing.T) {
	gotRanges := parse("5-10,15-20,25-30")

	wantRanges := [][2]int{
		{5, 10},
		{15, 20},
		{25, 30},
	}

	if !slices.Equal(gotRanges, wantRanges) {
		t.Errorf("got %v, want %v", gotRanges, wantRanges)
	}
}
