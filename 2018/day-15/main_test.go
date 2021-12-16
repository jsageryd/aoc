package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestFight(t *testing.T) {
	for n, tc := range []struct {
		plan    string
		outcome int
	}{
		{`
#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######
`, 27730,
		},
		{`
#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######
`, 36334,
		},

		{`
#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######
`, 39514,
		},

		{`
#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######
`, 27755,
		},

		{`
#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######
`, 28944,
		},

		{`
#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########
`, 18740,
		},
	} {
		var plan [][]byte

		for _, line := range strings.Split(strings.TrimSpace(tc.plan), "\n") {
			plan = append(plan, []byte(line))
		}

		c := newCave(plan)

		gotOutcome, _ := fight(c)

		if got, want := gotOutcome, tc.outcome; got != want {
			t.Errorf("[%d] outcome = %d, want %d", n, got, want)
		}
	}
}

func TestCoord_Adjacent(t *testing.T) {
	c := Coord{0, 0}

	adjCoords := c.Adjacent()

	wantAdjCoords := []Coord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	if len(adjCoords) != len(wantAdjCoords) {
		t.Fatalf("got %v, want %v", adjCoords, wantAdjCoords)
	}

	for n := range wantAdjCoords {
		if adjCoords[n] != wantAdjCoords[n] {
			t.Errorf("got %v, want %v", adjCoords, wantAdjCoords)
		}
	}
}

func TestCave_LocationsNearTarget(t *testing.T) {
	c := newCave(bytes.Split(bytes.TrimSpace([]byte(`
########
#.G....#
#..E.G.#
#......#
########
	`)), []byte{'\n'}))

	//   01234567
	// 0 ########
	// 1 #.G?.?.#
	// 2 #.?E?G?#
	// 3 #..?.?.#
	// 4 ########
	wantLocations := []Coord{
		{3, 1},
		{5, 1},
		{2, 2},
		{4, 2},
		{6, 2},
		{3, 3},
		{5, 3},
	}

	me := &Unit{Location: Coord{2, 1}}

	gotLocations := c.LocationsNearEnemies(me)

	if len(gotLocations) != len(wantLocations) {
		t.Fatalf("got %v, want %v", gotLocations, wantLocations)
	}

	for n := range wantLocations {
		if gotLocations[n] != wantLocations[n] {
			t.Fatalf("got %v, want %v", gotLocations, wantLocations)
		}
	}
}

func TestCave_DistanceMap(t *testing.T) {
	c := newCave(bytes.Split(bytes.TrimSpace([]byte(`
########
#.G....#
#..E.G.#
#......#
########
	`)), []byte{'\n'}))

	gotDistanceMap := c.DistanceMap(Coord{3, 2})

	//   01234567
	// 0 ########
	// 1 #.G....#
	// 2 #..E.G.#
	// 3 #......#
	// 4 ########
	wantDistanceMap := map[Coord]int{
		{3, 2}: 0,
		{1, 1}: 3, {3, 1}: 1, {4, 1}: 2, {5, 1}: 3, {6, 1}: 4,
		{1, 2}: 2, {2, 2}: 1, {4, 2}: 1, {6, 2}: 5,
		{1, 3}: 3, {2, 3}: 2, {3, 3}: 1, {4, 3}: 2, {5, 3}: 3, {6, 3}: 4,
	}

	if got, want := len(gotDistanceMap), len(wantDistanceMap); got != want {
		t.Fatalf("got %d distances, want %d", got, want)
	}

	for c := range wantDistanceMap {
		if got, want := gotDistanceMap[c], wantDistanceMap[c]; got != want {
			t.Errorf("%v -> %d, want %d", c, got, want)
		}
	}
}

func TestFindPath(t *testing.T) {
	c := newCave(bytes.Split(bytes.TrimSpace([]byte(`
########
#.E....#
####...#
#.G....#
########
	`)), []byte{'\n'}))

	gotPath := findPath(Coord{3, 3}, c.DistanceMap(Coord{2, 1}))

	//   01234567
	// 0 ########
	// 1 #.E....#
	// 2 ####...#
	// 3 #.G....#
	// 4 ########
	wantPath := []Coord{{3, 1}, {4, 1}, {4, 2}, {4, 3}, {3, 3}}

	if len(gotPath) != len(wantPath) {
		t.Fatalf("got %v, want %v", gotPath, wantPath)
	}

	for n := range wantPath {
		if gotPath[n] != wantPath[n] {
			t.Fatalf("got %v, want %v", gotPath, wantPath)
		}
	}
}
