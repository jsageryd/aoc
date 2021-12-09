package main

import (
	"fmt"
	"sort"
	"testing"
)

var input = []string{
	"2199943210",
	"3987894921",
	"9856789892",
	"8767896789",
	"9899965678",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 15; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 1134; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFindLowPoints(t *testing.T) {
	gotLowPoints := findLowPoints(parse(input))

	sort.Slice(gotLowPoints, func(i, j int) bool {
		ic, jc := gotLowPoints[i], gotLowPoints[j]
		return ic.y*10+ic.x < jc.y*10+jc.x
	})

	wantLowPoints := []coord{{1, 0}, {9, 0}, {2, 2}, {6, 4}}

	if fmt.Sprint(gotLowPoints) != fmt.Sprint(wantLowPoints) {
		t.Errorf("got %v, want %v", gotLowPoints, wantLowPoints)
	}
}
