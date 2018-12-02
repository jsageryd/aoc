package main

import (
	"testing"
)

func TestSpiralDistance(t *testing.T) {
	for n, tc := range []struct {
		in  int
		out int
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	} {
		if got, want := spiralDistance(tc.in), tc.out; got != want {
			t.Errorf("[%d] spiralDistance(%d) = %d, want %d", n, tc.in, got, want)
		}
	}
}

func TestSpiralFirstLargerAdjacentSum(t *testing.T) {
	for n, tc := range []struct {
		in  int
		out int
	}{
		{1, 2},
		{2, 4},
		{3, 4},
		{4, 5},
		{5, 10},
		{6, 10},
		{7, 10},
		{8, 10},
		{9, 10},
		{10, 11},
		{11, 23},
		{23, 25},
		{25, 26},
		{26, 54},
		{54, 57},
		{57, 59},
		{59, 122},
		{122, 133},
		{133, 142},
		{142, 147},
		{147, 304},
		{304, 330},
		{330, 351},
		{351, 362},
		{362, 747},
		{747, 806},
	} {
		if got, want := spiralFirstLargerAdjacentSum(tc.in), tc.out; got != want {
			t.Errorf("[%d] spiralFirstLargerAdjacentSum(%d) = %d, want %d", n, tc.in, got, want)
		}
	}
}
