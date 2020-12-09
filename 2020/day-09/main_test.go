package main

import "testing"

var input = []int{
	35,
	20,
	15,
	25,
	47,
	40,
	62,
	55,
	65,
	95,
	102,
	117,
	150,
	182,
	127,
	219,
	299,
	277,
	309,
	576,
}

func TestFindNumber(t *testing.T) {
	if got, want := findNumber(input, 5), 127; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFindWeakness(t *testing.T) {
	if got, want := findWeakness(input, 127), 62; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
