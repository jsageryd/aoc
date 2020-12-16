package main

import "testing"

func TestPart1(t *testing.T) {
	var input = []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}

	if got, want := part1(input), 165; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestParseMask(t *testing.T) {
	for n, tc := range []struct {
		mask                string
		wantZeros, wantOnes int
	}{
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", 0, 0},
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 2, 64},
	} {
		gotZeros, gotOnes := parseMask(tc.mask)

		if gotZeros != tc.wantZeros || gotOnes != tc.wantOnes {
			t.Errorf(
				"[%d] parseMask(%q) = %d, %d, want %d, %d",
				n, tc.mask, gotZeros, gotOnes, tc.wantZeros, tc.wantOnes,
			)
		}
	}
}
