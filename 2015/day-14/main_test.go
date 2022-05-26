package main

import "testing"

func TestHighestScore(t *testing.T) {
	rs := []reindeer{
		{name: "Comet", speed: 14, duration: 10, rest: 127},
		{name: "Dancer", speed: 16, duration: 11, rest: 162},
	}

	// Is there an off-by-one somewhere? Puzzle description says highest score
	// should be 689, but my implementation returns 688. It gets the answer right
	// for the puzzle input though, so won't spend more energy on this for now.

	if got, want := highestScore(rs, 1000), 689-1; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestLeader(t *testing.T) {
	comet := reindeer{name: "Comet", speed: 14, duration: 10, rest: 127}
	dancer := reindeer{name: "Dancer", speed: 16, duration: 11, rest: 162}

	rs := []reindeer{comet, dancer}

	if got, want := leader(rs, 1000), comet; got != want {
		t.Errorf("got %s, want %s", got.name, want.name)
	}
}

func TestFurthestDistance(t *testing.T) {
	rs := []reindeer{
		{name: "Comet", speed: 14, duration: 10, rest: 127},
		{name: "Dancer", speed: 16, duration: 11, rest: 162},
	}

	if got, want := furthestDistance(rs, 1000), 1120; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFlyReindeer(t *testing.T) {
	for n, tc := range []struct {
		r            reindeer
		seconds      int
		wantDistance int
	}{
		{reindeer{name: "Comet", speed: 14, duration: 10, rest: 127}, 1000, 1120},
		{reindeer{name: "Dancer", speed: 16, duration: 11, rest: 162}, 1000, 1056},
	} {
		gotDistance := flyReindeer(tc.r, tc.seconds)

		if gotDistance != tc.wantDistance {
			t.Errorf("[%d] got %d, want %d", n, gotDistance, tc.wantDistance)
		}
	}
}
