package main

import (
	"testing"
)

func TestCalibrate(t *testing.T) {
	for n, tc := range []struct {
		start int
		in    []int
		out   int
	}{
		{0, []int{+1, +1, +1}, 3},
		{0, []int{+1, +1, -2}, 0},
		{0, []int{-1, -2, -3}, -6},
	} {
		if got, want := calibrate(tc.start, tc.in), tc.out; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestCalibrateToFirstSeenTwice(t *testing.T) {
	for n, tc := range []struct {
		start int
		in    []int
		out   int
	}{
		{0, []int{+1, -1}, 0},
		{0, []int{+3, +3, +4, -2, -4}, 10},
		{0, []int{-6, +3, +8, +5, -6}, 5},
		{0, []int{+7, +7, -2, -7, -4}, 14},
	} {
		if got, want := calibrateToFirstSeenTwice(tc.start, tc.in), tc.out; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
