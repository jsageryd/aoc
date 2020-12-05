package main

import "testing"

func TestHighestSeatID(t *testing.T) {
	specs := []string{
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	}

	if got, want := highestSeatID(specs), 820; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestParseSpec(t *testing.T) {
	for n, tc := range []struct {
		spec         string
		row, col, id int
	}{
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	} {
		row, col, id := parseSpec(tc.spec)

		if row != tc.row || col != tc.col || id != tc.id {
			t.Errorf(
				"[%d] parseSpec(%q) = %d, %d, %d, want %d, %d, %d",
				n, tc.spec, row, col, id, tc.row, tc.col, tc.id,
			)
		}
	}
}
