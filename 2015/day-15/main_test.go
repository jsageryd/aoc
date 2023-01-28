package main

import (
	"fmt"
	"testing"
)

var input = []string{
	"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
	"Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 62842880; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestParseInput(t *testing.T) {
	wantOutput := []ingredient{
		{name: "Butterscotch", capacity: -1, durability: -2, flavor: 6, texture: 3},
		{name: "Cinnamon", capacity: 2, durability: 3, flavor: -2, texture: -1},
	}

	if got, want := parseInput(input), wantOutput; fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDistribute(t *testing.T) {
	for n, tc := range []struct {
		n         int
		max       int
		wantDists [][]int
	}{
		{
			n:   1,
			max: 0,
			wantDists: [][]int{
				{0},
			},
		},
		{
			n:   1,
			max: 3,
			wantDists: [][]int{
				{3},
			},
		},
		{
			n:   2,
			max: 3,
			wantDists: [][]int{
				{3, 0},
				{2, 1},
				{1, 2},
				{0, 3},
			},
		},
		{
			n:   3,
			max: 2,
			wantDists: [][]int{
				{2, 0, 0},
				{1, 1, 0},
				{1, 0, 1},
				{0, 2, 0},
				{0, 1, 1},
				{0, 0, 2},
			},
		},
	} {
		var gotDists [][]int

		distribute(tc.n, tc.max, func(dist []int) {
			gotDist := make([]int, len(dist))
			copy(gotDist, dist)
			gotDists = append(gotDists, gotDist)
		})

		if fmt.Sprint(gotDists) != fmt.Sprint(tc.wantDists) {
			t.Errorf("[%d] got %v, want %v", n, gotDists, tc.wantDists)
		}
	}
}
