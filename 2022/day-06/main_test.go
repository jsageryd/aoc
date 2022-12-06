package main

import "testing"

func TestPart1(t *testing.T) {
	for n, tc := range []struct {
		input string
		start int
	}{
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	} {
		if got, want := part1(tc.input), tc.start; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestPart2(t *testing.T) {
	for n, tc := range []struct {
		input string
		start int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	} {
		if got, want := part2(tc.input), tc.start; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
