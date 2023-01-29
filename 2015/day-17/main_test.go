package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	containers := []int{20, 15, 10, 5, 5}

	if got, want := part1(containers, 25), 4; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCombinations(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	k := 3

	wantCombs := [][]int{
		{1, 2, 3},
		{1, 2, 4},
		{1, 2, 5},
		{1, 2, 6},
		{1, 3, 4},
		{1, 3, 5},
		{1, 3, 6},
		{1, 4, 5},
		{1, 4, 6},
		{1, 5, 6},
		{2, 3, 4},
		{2, 3, 5},
		{2, 3, 6},
		{2, 4, 5},
		{2, 4, 6},
		{2, 5, 6},
		{3, 4, 5},
		{3, 4, 6},
		{3, 5, 6},
		{4, 5, 6},
	}

	var gotCombs [][]int

	combinations(s, k, func(comb []int) bool {
		c := make([]int, k)
		copy(c, comb)
		gotCombs = append(gotCombs, c)
		return true
	})

	if fmt.Sprint(gotCombs) != fmt.Sprint(wantCombs) {
		t.Errorf("got:\n%v\nwant:\n%v", gotCombs, wantCombs)
	}
}
