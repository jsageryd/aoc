package main

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	input := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}

	t.Run("Part 1", func(t *testing.T) {
		if got, want := run(input, step), 37; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		if got, want := run(input, stepV2), 26; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestStep(t *testing.T) {
	steps := [][]string{
		{
			"L.LL.LL.LL",
			"LLLLLLL.LL",
			"L.L.L..L..",
			"LLLL.LL.LL",
			"L.LL.LL.LL",
			"L.LLLLL.LL",
			"..L.L.....",
			"LLLLLLLLLL",
			"L.LLLLLL.L",
			"L.LLLLL.LL",
		},
		{
			"#.##.##.##",
			"#######.##",
			"#.#.#..#..",
			"####.##.##",
			"#.##.##.##",
			"#.#####.##",
			"..#.#.....",
			"##########",
			"#.######.#",
			"#.#####.##",
		},
		{
			"#.LL.L#.##",
			"#LLLLLL.L#",
			"L.L.L..L..",
			"#LLL.LL.L#",
			"#.LL.LL.LL",
			"#.LLLL#.##",
			"..L.L.....",
			"#LLLLLLLL#",
			"#.LLLLLL.L",
			"#.#LLLL.##",
		},
		{
			"#.##.L#.##",
			"#L###LL.L#",
			"L.#.#..#..",
			"#L##.##.L#",
			"#.##.LL.LL",
			"#.###L#.##",
			"..#.#.....",
			"#L######L#",
			"#.LL###L.L",
			"#.#L###.##",
		},
		{
			"#.#L.L#.##",
			"#LLL#LL.L#",
			"L.L.L..#..",
			"#LLL.##.L#",
			"#.LL.LL.LL",
			"#.LL#L#.##",
			"..L.L.....",
			"#L#LLLL#L#",
			"#.LLLLLL.L",
			"#.#L#L#.##",
		},
		{
			"#.#L.L#.##",
			"#LLL#LL.L#",
			"L.#.L..#..",
			"#L##.##.L#",
			"#.#L.LL.LL",
			"#.#L#L#.##",
			"..L.L.....",
			"#L#L##L#L#",
			"#.LLLLLL.L",
			"#.#L#L#.##",
		},
		{
			"#.#L.L#.##",
			"#LLL#LL.L#",
			"L.#.L..#..",
			"#L##.##.L#",
			"#.#L.LL.LL",
			"#.#L#L#.##",
			"..L.L.....",
			"#L#L##L#L#",
			"#.LLLLLL.L",
			"#.#L#L#.##",
		},
	}

	for i := 0; i < len(steps)-1; i++ {
		input := steps[i]
		next := steps[i+1]

		layout := makeLayoutMap(input)
		wantLayout := makeLayoutMap(next)

		step(layout)

		if fmt.Sprint(layout) != fmt.Sprint(wantLayout) {
			t.Errorf("[%d] got %v,\nwant %v", i, layout, wantLayout)
		}
	}

	if got, want := step(makeLayoutMap(steps[0])), true; got != want {
		t.Errorf("[modified layout] got %t, want %t", got, want)
	}

	if got, want := step(makeLayoutMap(steps[len(steps)-1])), false; got != want {
		t.Errorf("[unmodified layout] got %t, want %t", got, want)
	}
}

func TestStepV2(t *testing.T) {
	steps := [][]string{
		{
			"L.LL.LL.LL",
			"LLLLLLL.LL",
			"L.L.L..L..",
			"LLLL.LL.LL",
			"L.LL.LL.LL",
			"L.LLLLL.LL",
			"..L.L.....",
			"LLLLLLLLLL",
			"L.LLLLLL.L",
			"L.LLLLL.LL",
		},
		{
			"#.##.##.##",
			"#######.##",
			"#.#.#..#..",
			"####.##.##",
			"#.##.##.##",
			"#.#####.##",
			"..#.#.....",
			"##########",
			"#.######.#",
			"#.#####.##",
		},
		{
			"#.LL.LL.L#",
			"#LLLLLL.LL",
			"L.L.L..L..",
			"LLLL.LL.LL",
			"L.LL.LL.LL",
			"L.LLLLL.LL",
			"..L.L.....",
			"LLLLLLLLL#",
			"#.LLLLLL.L",
			"#.LLLLL.L#",
		},
		{
			"#.L#.##.L#",
			"#L#####.LL",
			"L.#.#..#..",
			"##L#.##.##",
			"#.##.#L.##",
			"#.#####.#L",
			"..#.#.....",
			"LLL####LL#",
			"#.L#####.L",
			"#.L####.L#",
		},
		{
			"#.L#.L#.L#",
			"#LLLLLL.LL",
			"L.L.L..#..",
			"##LL.LL.L#",
			"L.LL.LL.L#",
			"#.LLLLL.LL",
			"..L.L.....",
			"LLLLLLLLL#",
			"#.LLLLL#.L",
			"#.L#LL#.L#",
		},
		{
			"#.L#.L#.L#",
			"#LLLLLL.LL",
			"L.L.L..#..",
			"##L#.#L.L#",
			"L.L#.#L.L#",
			"#.L####.LL",
			"..#.#.....",
			"LLL###LLL#",
			"#.LLLLL#.L",
			"#.L#LL#.L#",
		},
		{
			"#.L#.L#.L#",
			"#LLLLLL.LL",
			"L.L.L..#..",
			"##L#.#L.L#",
			"L.L#.LL.L#",
			"#.LLLL#.LL",
			"..#.L.....",
			"LLL###LLL#",
			"#.LLLLL#.L",
			"#.L#LL#.L#",
		},
		{
			"#.L#.L#.L#",
			"#LLLLLL.LL",
			"L.L.L..#..",
			"##L#.#L.L#",
			"L.L#.LL.L#",
			"#.LLLL#.LL",
			"..#.L.....",
			"LLL###LLL#",
			"#.LLLLL#.L",
			"#.L#LL#.L#",
		},
	}

	for i := 0; i < len(steps)-1; i++ {
		input := steps[i]
		next := steps[i+1]

		layout := makeLayoutMap(input)
		wantLayout := makeLayoutMap(next)

		stepV2(layout)

		if fmt.Sprint(layout) != fmt.Sprint(wantLayout) {
			t.Errorf("[%d] got %v,\nwant %v", i, layout, wantLayout)
		}
	}

	if got, want := step(makeLayoutMap(steps[0])), true; got != want {
		t.Errorf("[modified layout] got %t, want %t", got, want)
	}

	if got, want := step(makeLayoutMap(steps[len(steps)-1])), false; got != want {
		t.Errorf("[unmodified layout] got %t, want %t", got, want)
	}
}
