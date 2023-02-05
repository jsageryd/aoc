package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	/*
	   House 1 got 10 presents.
	   House 2 got 30 presents.
	   House 3 got 40 presents.
	   House 4 got 70 presents.
	   House 5 got 60 presents.
	   House 6 got 120 presents.
	   House 7 got 80 presents.
	   House 8 got 150 presents.
	   House 9 got 130 presents.
	*/

	for n, tc := range []struct {
		input int
		want  int
	}{
		{10, 1},
		{20, 2},
		{30, 2},
		{60, 4},
		{80, 6},
		{130, 8},
		{150, 8},
	} {
		if got, want := part1(tc.input), tc.want; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
