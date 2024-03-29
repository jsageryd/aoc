package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"H => HO",
		"H => OH",
		"O => HH",
		"",
		"HOH",
	}

	if got, want := part1(input), 4; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	for n, tc := range []struct {
		input []string
		want  int
	}{
		{[]string{"e => H", "e => O", "H => HO", "H => OH", "O => HH", "", "HOH"}, 3},
		{[]string{"e => H", "e => O", "H => HO", "H => OH", "O => HH", "", "HOHOHO"}, 6},
	} {
		if got, want := part2(tc.input), tc.want; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestIndices(t *testing.T) {
	for n, tc := range []struct {
		s, sep string
		want   []int
	}{
		{s: "", sep: "", want: []int{0}},
		{s: "", sep: "H", want: nil},
		{s: "HOH", sep: "", want: []int{0, 1, 2, 3}},
		{s: "HOH", sep: "H", want: []int{0, 2}},
		{s: "HOH", sep: "O", want: []int{1}},
		{s: "HOHOHO", sep: "HOHO", want: []int{0, 2}},
	} {
		if got, want := indices(tc.s, tc.sep), tc.want; fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("[%d] got %v, want %v", n, got, want)
		}
	}
}
