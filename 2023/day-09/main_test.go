package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}

	if got, want := part1(input), 114; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}

	if got, want := part2(input), 2; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestNextValue(t *testing.T) {
	for n, tc := range []struct {
		line string
		want int
	}{
		{"0 3 6 9 12 15", 18},
		{"1 3 6 10 15 21", 28},
		{"10 13 16 21 30 45", 68},
	} {
		if got, want := nextValue(tc.line), tc.want; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestPrevValue(t *testing.T) {
	for n, tc := range []struct {
		line string
		want int
	}{
		{"0 3 6 9 12 15", -3},
		{"1 3 6 10 15 21", 0},
		{"10 13 16 21 30 45", 5},
	} {
		if got, want := prevValue(tc.line), tc.want; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
