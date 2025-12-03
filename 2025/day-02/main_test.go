package main

import (
	"slices"
	"testing"
)

var input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

func TestPart1(t *testing.T) {
	if got, want := part1(input), 1227775554; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 4174379265; got != want {
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

func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		n     int
		parts int
		want  []int
	}{
		{1, 1, []int{1}},
		{10, 1, []int{10}},
		{10, 2, []int{1, 0}},
		{102, 1, []int{102}},
		{102, 3, []int{1, 0, 2}},
		{1020, 1, []int{1020}},
		{1020, 2, []int{10, 20}},
		{1020, 4, []int{1, 0, 2, 0}},
		{10203, 1, []int{10203}},
		{10203, 5, []int{1, 0, 2, 0, 3}},
		{12, 1, []int{12}},
		{12, 2, []int{1, 2}},
		{123, 1, []int{123}},
		{123, 2, nil},
		{123, 3, []int{1, 2, 3}},
		{1234, 1, []int{1234}},
		{1234, 2, []int{12, 34}},
		{1234, 3, nil},
		{1234, 4, []int{1, 2, 3, 4}},
		{12345, 1, []int{12345}},
		{12345, 2, nil},
		{12345, 3, nil},
		{12345, 4, nil},
		{12345, 5, []int{1, 2, 3, 4, 5}},
		{123456, 1, []int{123456}},
		{123456, 2, []int{123, 456}},
		{123456, 3, []int{12, 34, 56}},
		{123456, 4, nil},
		{123456, 5, nil},
		{123456, 6, []int{1, 2, 3, 4, 5, 6}},
	} {
		if got, want := split(tc.n, tc.parts), tc.want; !slices.Equal(got, want) {
			t.Errorf("split(%d, %d) = %d, want %d", tc.n, tc.parts, got, want)
		}
	}
}
