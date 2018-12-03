package main

import "testing"

func TestMazeSteps(t *testing.T) {
	for n, tc := range []struct {
		in  []int
		out int
	}{
		{[]int{0, 3, 0, 1, -3}, 5},
	} {
		if got, want := mazeSteps(tc.in), tc.out; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestMazeStepsAlternate(t *testing.T) {
	for n, tc := range []struct {
		in  []int
		out int
	}{
		{[]int{0, 3, 0, 1, -3}, 10},
	} {
		if got, want := mazeStepsAlternate(tc.in), tc.out; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
