package main

import "testing"

func TestMap_Step(t *testing.T) {
	input := []string{
		"x=495, y=2..7",
		"y=7, x=495..501",
		"x=501, y=3..7",
		"x=498, y=2..4",
		"x=506, y=1..2",
		"x=498, y=10..13",
		"x=504, y=10..13",
		"y=13, x=498..504",
	}

	m := newMap(input)

	for m.Step() {
	}

	var movingWater, stillWater int
	for c, v := range m.m {
		if c.y >= m.minY && c.y <= m.maxY {
			switch v {
			case '|':
				movingWater++
			case '~':
				stillWater++
			}
		}
	}

	if got, want := movingWater+stillWater, 57; got != want {
		t.Errorf("movingWater + stillWater = %d, want %d", got, want)
	}

	if got, want := stillWater, 29; got != want {
		t.Errorf("stillWater = %d, want %d", got, want)
	}
}
