package main

import "testing"

func TestPart1(t *testing.T) {
	values := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	if got, want := part1(values), 37; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	values := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	if got, want := part2(values), 168; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestMedian(t *testing.T) {
	for n, tc := range []struct {
		values []int
		median int
	}{
		{values: []int{}, median: 0},
		{values: []int{1}, median: 1},
		{values: []int{1, 2}, median: 2},
		{values: []int{1, 2, 3}, median: 2},
		{values: []int{1, 2, 3, 4}, median: 3},
		{values: []int{1, 2, 3, 4, 5}, median: 3},
		{values: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, median: 2},
	} {
		if got, want := median(tc.values), tc.median; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestAlignmentCost(t *testing.T) {
	for n, tc := range []struct {
		values  []int
		alignAt int
		cost    int
	}{
		{values: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, alignAt: 1, cost: 41},
		{values: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, alignAt: 2, cost: 37},
		{values: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, alignAt: 3, cost: 39},
		{values: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, alignAt: 10, cost: 71},
	} {
		if got, want := alignmentCost(tc.values, tc.alignAt), tc.cost; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestAlignmentCost2(t *testing.T) {
	for n, tc := range []struct {
		values  []int
		alignAt int
		cost    int
	}{
		{values: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, alignAt: 2, cost: 206},
		{values: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, alignAt: 5, cost: 168},
	} {
		if got, want := alignmentCost2(tc.values, tc.alignAt), tc.cost; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
