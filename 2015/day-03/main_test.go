package main

import "testing"

func TestFindHouses(t *testing.T) {
	for n, tc := range []struct {
		input  []byte
		santas int
		houses map[coord]int
	}{
		{
			input:  []byte(">"),
			santas: 1,
			houses: map[coord]int{
				{0, 0}: 1,
				{1, 0}: 1,
			},
		},
		{
			input:  []byte("^>v<"),
			santas: 1,
			houses: map[coord]int{
				{0, 0}:  2,
				{0, -1}: 1,
				{1, -1}: 1,
				{1, 0}:  1,
			},
		},
		{
			input:  []byte("^v^v^v^v^v"),
			santas: 1,
			houses: map[coord]int{
				{0, 0}:  6,
				{0, -1}: 5,
			},
		},
		{
			input:  []byte("^v"),
			santas: 2,
			houses: map[coord]int{
				{0, 0}:  2,
				{0, -1}: 1,
				{0, 1}:  1,
			},
		},
		{
			input:  []byte("^>v<"),
			santas: 2,
			houses: map[coord]int{
				{0, 0}:  4,
				{0, -1}: 1,
				{1, 0}:  1,
			},
		},
		{
			input:  []byte("^v^v^v^v^v"),
			santas: 2,
			houses: map[coord]int{
				{0, 0}:  2,
				{0, -1}: 1,
				{0, -2}: 1,
				{0, -3}: 1,
				{0, -4}: 1,
				{0, -5}: 1,
				{0, 1}:  1,
				{0, 2}:  1,
				{0, 3}:  1,
				{0, 4}:  1,
				{0, 5}:  1,
			},
		},
	} {
		gotHouses := findHouses(tc.input, tc.santas)

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
