package main

import (
	"slices"
	"testing"
)

var input = []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 99; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 44; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCombTargetSum(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	target := 7

	var gotCombs [][]int

	combTargetSum(s, target, func(comb []int) bool {
		gotCombs = append(gotCombs, slices.Clone(comb))
		return true
	})

	wantCombs := [][]int{
		{1, 2, 4},
		{1, 6},
		{2, 5},
		{3, 4},
		{7},
	}

	if !slices.EqualFunc(gotCombs, wantCombs, slices.Equal) {
		t.Errorf("got %d, want %d", gotCombs, wantCombs)
	}
}
