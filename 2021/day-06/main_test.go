package main

import (
	"fmt"
	"testing"
)

func TestSpawn(t *testing.T) {
	freq := func(ints []int) []int {
		f := make([]int, 9)
		for _, i := range ints {
			f[i]++
		}
		return f
	}

	for n, tc := range []struct {
		input []int
		days  int
		want  []int
	}{
		{input: freq([]int{3, 4, 3, 1, 2}), days: 1, want: freq([]int{2, 3, 2, 0, 1})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 2, want: freq([]int{1, 2, 1, 6, 0, 8})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 3, want: freq([]int{0, 1, 0, 5, 6, 7, 8})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 4, want: freq([]int{6, 0, 6, 4, 5, 6, 7, 8, 8})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 5, want: freq([]int{5, 6, 5, 3, 4, 5, 6, 7, 7, 8})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 6, want: freq([]int{4, 5, 4, 2, 3, 4, 5, 6, 6, 7})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 7, want: freq([]int{3, 4, 3, 1, 2, 3, 4, 5, 5, 6})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 8, want: freq([]int{2, 3, 2, 0, 1, 2, 3, 4, 4, 5})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 9, want: freq([]int{1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 8})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 10, want: freq([]int{0, 1, 0, 5, 6, 0, 1, 2, 2, 3, 7, 8})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 11, want: freq([]int{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 7, 8, 8, 8})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 12, want: freq([]int{5, 6, 5, 3, 4, 5, 6, 0, 0, 1, 5, 6, 7, 7, 7, 8, 8})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 13, want: freq([]int{4, 5, 4, 2, 3, 4, 5, 6, 6, 0, 4, 5, 6, 6, 6, 7, 7, 8, 8})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 14, want: freq([]int{3, 4, 3, 1, 2, 3, 4, 5, 5, 6, 3, 4, 5, 5, 5, 6, 6, 7, 7, 8})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 15, want: freq([]int{2, 3, 2, 0, 1, 2, 3, 4, 4, 5, 2, 3, 4, 4, 4, 5, 5, 6, 6, 7})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 16, want: freq([]int{1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 1, 2, 3, 3, 3, 4, 4, 5, 5, 6, 8})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 17, want: freq([]int{0, 1, 0, 5, 6, 0, 1, 2, 2, 3, 0, 1, 2, 2, 2, 3, 3, 4, 4, 5, 7, 8})},
		{input: freq([]int{3, 4, 3, 1, 2}), days: 18, want: freq([]int{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8})},
	} {
		spawn(tc.input, tc.days)

		if got, want := tc.input, tc.want; fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("[%d] got %d fish %v, want %d fish %v", n, countFish(got), got, countFish(want), want)
		}
	}
}
