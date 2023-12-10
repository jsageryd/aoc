package main

import "testing"

func TestPart1(t *testing.T) {
	for n, tc := range []struct {
		input []string
		want  int
	}{
		{
			input: []string{
				".....",
				".S-7.",
				".|.|.",
				".L-J.",
				".....",
			},
			want: 4,
		},
		{
			input: []string{
				"-L|F7",
				"7S-7|",
				"L|7||",
				"-L-J|",
				"L|-JF",
			},
			want: 4,
		},
		{
			input: []string{
				"..F7.",
				".FJ|.",
				"SJ.L7",
				"|F--J",
				"LJ...",
			},
			want: 8,
		},
		{
			input: []string{
				"7-F7-",
				".FJ|7",
				"SJLL7",
				"|F--J",
				"LJ.LJ",
			},
			want: 8,
		},
	} {
		if got, want := part1(tc.input), tc.want; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
