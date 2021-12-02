package main

import "testing"

func TestSubmarine_Pilot(t *testing.T) {
	var s submarine

	s.pilot([]string{
		`forward 5`,
		`down 5`,
		`forward 8`,
		`up 3`,
		`down 8`,
		`forward 2`,
	})

	wantHorizontal := 15
	wantDepth := 10

	if s.horizontal != wantHorizontal || s.depth != wantDepth {
		t.Errorf(
			"got (%d, %d), want (%d, %d)",
			s.horizontal, s.depth, wantHorizontal, wantDepth,
		)
	}
}
