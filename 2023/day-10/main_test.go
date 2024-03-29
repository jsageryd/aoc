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

func TestPart2(t *testing.T) {
	for n, tc := range []struct {
		input []string
		want  int
	}{
		{
			input: []string{
				"...........",
				".S-------7.",
				".|F-----7|.",
				".||.....||.",
				".||.....||.",
				".|L-7.F-J|.",
				".|..|.|..|.",
				".L--J.L--J.",
				"...........",
			},
			want: 4,
		},
		{
			input: []string{
				"..........",
				".S------7.",
				".|F----7|.",
				".||....||.",
				".||....||.",
				".|L-7F-J|.",
				".|..||..|.",
				".L--JL--J.",
				"..........",
			},
			want: 4,
		},
		{
			input: []string{
				".F----7F7F7F7F-7....",
				".|F--7||||||||FJ....",
				".||.FJ||||||||L7....",
				"FJL7L7LJLJ||LJ.L-7..",
				"L--J.L7...LJS7F-7L7.",
				"....F-J..F7FJ|L7L7L7",
				"....L7.F7||L7|.L7L7|",
				".....|FJLJ|FJ|F7|.LJ",
				"....FJL-7.||.||||...",
				"....L---J.LJ.LJLJ...",
			},
			want: 8,
		},
		{
			input: []string{
				"FF7FSF7F7F7F7F7F---7",
				"L|LJ||||||||||||F--J",
				"FL-7LJLJ||||||LJL-77",
				"F--JF--7||LJLJ7F7FJ-",
				"L---JF-JLJ.||-FJLJJ7",
				"|F|F-JF---7F7-L7L|7|",
				"|FFJF7L7F-JF7|JL---7",
				"7-L-JL7||F7|L7F-7F7|",
				"L.L7LFJ|||||FJL7||LJ",
				"L7JLJL-JLJLJL--JLJ.L",
			},
			want: 10,
		},
		{
			input: []string{
				"........",
				"|S----7.",
				"||..--|.",
				"|L----J.",
				"........",
			},
			want: 4,
		},
		{
			input: []string{
				"F------7",
				"|------|",
				"||..--||",
				"||-----|",
				"L------S",
			},
			want: 18,
		},
	} {
		if got, want := part2(tc.input), tc.want; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
