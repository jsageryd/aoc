package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		`5483143223`,
		`2745854711`,
		`5264556173`,
		`6141336146`,
		`6357385478`,
		`4167524645`,
		`2176841721`,
		`6882881134`,
		`4846848554`,
		`5283751526`,
	}

	if got, want := part1(input), 1656; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		`5483143223`,
		`2745854711`,
		`5264556173`,
		`6141336146`,
		`6357385478`,
		`4167524645`,
		`2176841721`,
		`6882881134`,
		`4846848554`,
		`5283751526`,
	}

	if got, want := part2(input), 195; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestStep(t *testing.T) {
	t.Run("Grid 1", func(t *testing.T) {
		input := []string{
			`11111`,
			`19991`,
			`19191`,
			`19991`,
			`11111`,
		}

		for _, tc := range []struct {
			steps int
			state []string
		}{
			{
				steps: 1,
				state: []string{
					`34543`,
					`40004`,
					`50005`,
					`40004`,
					`34543`,
				},
			},
			{
				steps: 2,
				state: []string{
					`45654`,
					`51115`,
					`61116`,
					`51115`,
					`45654`,
				},
			},
		} {
			gotState := parse(input)

			for i := 0; i < tc.steps; i++ {
				step(gotState)
			}

			wantState := parse(tc.state)

			if gotState.String() != wantState.String() {
				t.Errorf("[step %d] got:\n%s\n\nwant:\n%s", tc.steps, gotState, wantState)
			}
		}
	})

	t.Run("Grid 2", func(t *testing.T) {
		input := []string{
			`5483143223`,
			`2745854711`,
			`5264556173`,
			`6141336146`,
			`6357385478`,
			`4167524645`,
			`2176841721`,
			`6882881134`,
			`4846848554`,
			`5283751526`,
		}

		for _, tc := range []struct {
			steps int
			state []string
		}{
			{
				steps: 1,
				state: []string{
					`6594254334`,
					`3856965822`,
					`6375667284`,
					`7252447257`,
					`7468496589`,
					`5278635756`,
					`3287952832`,
					`7993992245`,
					`5957959665`,
					`6394862637`,
				},
			},
			{
				steps: 2,
				state: []string{
					`8807476555`,
					`5089087054`,
					`8597889608`,
					`8485769600`,
					`8700908800`,
					`6600088989`,
					`6800005943`,
					`0000007456`,
					`9000000876`,
					`8700006848`,
				},
			},
			{
				steps: 100,
				state: []string{
					`0397666866`,
					`0749766918`,
					`0053976933`,
					`0004297822`,
					`0004229892`,
					`0053222877`,
					`0532222966`,
					`9322228966`,
					`7922286866`,
					`6789998766`,
				},
			},
		} {
			gotState := parse(input)

			for i := 0; i < tc.steps; i++ {
				step(gotState)
			}

			wantState := parse(tc.state)

			if gotState.String() != wantState.String() {
				t.Errorf("[step %d] got:\n%s\n\nwant:\n%s", tc.steps, gotState, wantState)
			}
		}
	})
}
