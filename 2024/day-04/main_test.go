package main

import (
	"slices"
	"testing"
)

var input = []string{
	"MMMSXXMASM",
	"MSAMXMSMSA",
	"AMXSXMAAMM",
	"MSAMASMSMX",
	"XMASAMXAMM",
	"XXAMMXXAMA",
	"SMSMSASXSS",
	"SAXAMASAAA",
	"MAMMMXMMMM",
	"MXMXAXMASX",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 18; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 9; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestRotate(t *testing.T) {
	input := []string{
		"ABCDEF",
		"GHIJKL",
		"MNOPQR",
	}

	for _, tc := range []struct {
		direction int
		want      []string
	}{
		{direction: 0, want: []string{"ABCDEF", "GHIJKL", "MNOPQR"}},
		{direction: 1, want: []string{"F", "EL", "DKR", "CJQ", "BIP", "AHO", "GN", "M"}},
		{direction: 2, want: []string{"FLR", "EKQ", "DJP", "CIO", "BHN", "AGM"}},
		{direction: 3, want: []string{"R", "LQ", "FKP", "EJO", "DIN", "CHM", "BG", "A"}},
		{direction: 4, want: []string{"RQPONM", "LKJIHG", "FEDCBA"}},
		{direction: 5, want: []string{"M", "NG", "OHA", "PIB", "QJC", "RKD", "LE", "F"}},
		{direction: 6, want: []string{"MGA", "NHB", "OIC", "PJD", "QKE", "RLF"}},
		{direction: 7, want: []string{"A", "GB", "MHC", "NID", "OJE", "PKF", "QL", "R"}},
	} {
		if got, want := rotate(input, tc.direction), tc.want; !slices.Equal(got, want) {
			t.Errorf("[direction %d] got %v, want %v", tc.direction, got, want)
		}
	}
}
