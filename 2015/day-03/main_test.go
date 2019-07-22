package main

import "testing"

func TestFindHouses(t *testing.T) {
	for n, tc := range []struct {
		input  []byte
		houses map[coord]int
	}{
		{
			input: []byte(">"),
			houses: map[coord]int{
				coord{0, 0}: 1,
				coord{1, 0}: 1,
			},
		},
		{
			input: []byte("^>v<"),
			houses: map[coord]int{
				coord{0, 0}:  2,
				coord{0, -1}: 1,
				coord{1, -1}: 1,
				coord{1, 0}:  1,
			},
		},
		{
			input: []byte("^v^v^v^v^v"),
			houses: map[coord]int{
				coord{0, 0}:  6,
				coord{0, -1}: 5,
			},
		},
	} {
		gotHouses := findHouses(tc.input)

		if len(gotHouses) != len(tc.houses) {
			t.Fatalf("[%d] got %v, want %v", n, gotHouses, tc.houses)
		}

		for k := range tc.houses {
			if gotHouses[k] != tc.houses[k] {
				t.Fatalf("[%d] got %v, want %v", n, gotHouses, tc.houses)
			}
		}
	}
}
