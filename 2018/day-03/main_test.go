package main

import (
	"testing"
)

func TestParseClaim(t *testing.T) {
	for n, tc := range []struct {
		in  string
		out claim
	}{
		{"#123 @ 3,2: 5x4", claim{ID: 123, X: 3, Y: 2, W: 5, H: 4}},
	} {
		if got, want := parseClaim(tc.in), tc.out; got != want {
			t.Errorf("[%d] parseClaim(%q) = %+v, want %+v", n, tc.in, got, want)
		}
	}
}

func TestOverlappigArea(t *testing.T) {
	for n, tc := range []struct {
		in  []claim
		out int
	}{
		{
			in: []claim{
				{ID: 1, X: 1, Y: 3, W: 4, H: 4},
				{ID: 2, X: 3, Y: 1, W: 4, H: 4},
				{ID: 3, X: 5, Y: 5, W: 2, H: 2},
			},
			out: 4,
		},
	} {
		if got, want := overlappingArea(tc.in), tc.out; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestNonOverlappigClaim(t *testing.T) {
	c1 := claim{ID: 1, X: 1, Y: 3, W: 4, H: 4}
	c2 := claim{ID: 2, X: 3, Y: 1, W: 4, H: 4}
	c3 := claim{ID: 3, X: 5, Y: 5, W: 2, H: 2}

	for n, tc := range []struct {
		in  []claim
		out claim
	}{
		{
			in:  []claim{c1, c2, c3},
			out: c3,
		},
	} {
		if got, want := nonOverlappingClaim(tc.in), tc.out; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
