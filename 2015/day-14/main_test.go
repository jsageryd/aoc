package main

import "testing"

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
