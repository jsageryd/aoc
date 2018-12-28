package main

import (
	"strings"
	"testing"
)

func TestStep(t *testing.T) {
	inputArea := stringToArea(`
    .#.#...|#.
    .....#|##|
    .|..|...#.
    ..|#.....#
    #.#|||#|#|
    ...#.||...
    .|....|...
    ||...#|.#|
    |.||||..|.
    ...#.|..|.
  `)

	for n, tc := range []struct {
		steps    int
		wantArea Area
	}{
		{
			steps:    0,
			wantArea: inputArea,
		},
		{
			steps: 1,
			wantArea: stringToArea(`
        .......##.
        ......|###
        .|..|...#.
        ..|#||...#
        ..##||.|#|
        ...#||||..
        ||...|||..
        |||||.||.|
        ||||||||||
        ....||..|.
      `),
		},
		{
			steps: 2,
			wantArea: stringToArea(`
        .......#..
        ......|#..
        .|.|||....
        ..##|||..#
        ..###|||#|
        ...#|||||.
        |||||||||.
        ||||||||||
        ||||||||||
        .|||||||||
      `),
		},
		{
			steps: 3,
			wantArea: stringToArea(`
        .......#..
        ....|||#..
        .|.||||...
        ..###|||.#
        ...##|||#|
        .||##|||||
        ||||||||||
        ||||||||||
        ||||||||||
        ||||||||||
      `),
		},
		{
			steps: 4,
			wantArea: stringToArea(`
        .....|.#..
        ...||||#..
        .|.#||||..
        ..###||||#
        ...###||#|
        |||##|||||
        ||||||||||
        ||||||||||
        ||||||||||
        ||||||||||
      `),
		},
		{
			steps: 5,
			wantArea: stringToArea(`
        ....|||#..
        ...||||#..
        .|.##||||.
        ..####|||#
        .|.###||#|
        |||###||||
        ||||||||||
        ||||||||||
        ||||||||||
        ||||||||||
      `),
		},
		{
			steps: 6,
			wantArea: stringToArea(`
        ...||||#..
        ...||||#..
        .|.###|||.
        ..#.##|||#
        |||#.##|#|
        |||###||||
        ||||#|||||
        ||||||||||
        ||||||||||
        ||||||||||
      `),
		},
		{
			steps: 7,
			wantArea: stringToArea(`
        ...||||#..
        ..||#|##..
        .|.####||.
        ||#..##||#
        ||##.##|#|
        |||####|||
        |||###||||
        ||||||||||
        ||||||||||
        ||||||||||
      `),
		},
		{
			steps: 8,
			wantArea: stringToArea(`
        ..||||##..
        ..|#####..
        |||#####|.
        ||#...##|#
        ||##..###|
        ||##.###||
        |||####|||
        ||||#|||||
        ||||||||||
        ||||||||||
      `),
		},
		{
			steps: 9,
			wantArea: stringToArea(`
        ..||###...
        .||#####..
        ||##...##.
        ||#....###
        |##....##|
        ||##..###|
        ||######||
        |||###||||
        ||||||||||
        ||||||||||
      `),
		},
		{
			steps: 10,
			wantArea: stringToArea(`
        .||##.....
        ||###.....
        ||##......
        |##.....##
        |##.....##
        |##....##|
        ||##.####|
        ||#####|||
        ||||#|||||
        ||||||||||
      `),
		},
	} {
		gotArea := step(inputArea, tc.steps)

		if len(gotArea) != len(tc.wantArea) {
			t.Fatalf("[%d] (%d steps)\ngot:\n%v\n\nwant:\n%v", n, tc.steps, gotArea, tc.wantArea)
		}

		for y := range tc.wantArea {
			if len(gotArea[y]) != len(tc.wantArea[y]) {
				t.Fatalf("[%d] (%d steps)\ngot:\n%v\n\nwant:\n%v", n, tc.steps, gotArea, tc.wantArea)
			}

			for x := range tc.wantArea[y] {
				if gotArea[y][x] != tc.wantArea[y][x] {
					t.Fatalf("[%d] (%d steps)\ngot:\n%v\n\nwant:\n%v", n, tc.steps, gotArea, tc.wantArea)
				}
			}
		}
	}
}

func BenchmarkStep(b *testing.B) {
	inputArea := stringToArea(`
    .#.#...|#.
    .....#|##|
    .|..|...#.
    ..|#.....#
    #.#|||#|#|
    ...#.||...
    .|....|...
    ||...#|.#|
    |.||||..|.
    ...#.|..|.
  `)

	for n := 0; n < b.N; n++ {
		step(inputArea, 100000)
	}
}

func stringToArea(s string) Area {
	var area Area
	for _, row := range strings.Split(strings.TrimSpace(s), "\n") {
		area = append(area, []byte(strings.TrimSpace(row)))
	}
	return area
}
