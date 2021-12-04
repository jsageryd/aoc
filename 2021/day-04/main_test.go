package main

import (
	"testing"
)

func TestFirstWinningBoard(t *testing.T) {
	numbers := []int{
		7, 4, 9, 5, 11, 17, 23, 2, 0,
		14, 21, 24, 10, 16, 13, 6, 15, 25,
		12, 22, 18, 20, 8, 19, 3, 26, 1,
	}

	boards := []*board{
		newBoard([][]int{
			{22, 13, 17, 11, 0},
			{8, 2, 23, 4, 24},
			{21, 9, 14, 16, 7},
			{6, 10, 3, 18, 5},
			{1, 12, 20, 15, 19},
		}),
		newBoard([][]int{
			{3, 15, 0, 2, 22},
			{9, 18, 13, 17, 5},
			{19, 8, 7, 25, 23},
			{20, 11, 10, 24, 4},
			{14, 21, 16, 12, 6},
		}),
		newBoard([][]int{
			{14, 21, 17, 24, 4},
			{10, 16, 15, 9, 19},
			{18, 8, 23, 26, 20},
			{22, 11, 13, 6, 5},
			{2, 0, 12, 3, 7},
		}),
	}

	board := firstWinningBoard(numbers, boards)

	if got, want := board, boards[2]; got.String() != want.String() {
		t.Errorf("got board:\n%s\n\nwant:\n%s", got, want)
	}

	if got, want := board.lastMarked, 24; got != want {
		t.Errorf("got last marked %d, want %d", got, want)
	}

	if got, want := board.score(), 4512; got != want {
		t.Errorf("got score %d, want %d", got, want)
	}
}

func TestLastWinningBoard(t *testing.T) {
	numbers := []int{
		7, 4, 9, 5, 11, 17, 23, 2, 0,
		14, 21, 24, 10, 16, 13, 6, 15, 25,
		12, 22, 18, 20, 8, 19, 3, 26, 1,
	}

	boards := []*board{
		newBoard([][]int{
			{22, 13, 17, 11, 0},
			{8, 2, 23, 4, 24},
			{21, 9, 14, 16, 7},
			{6, 10, 3, 18, 5},
			{1, 12, 20, 15, 19},
		}),
		newBoard([][]int{
			{3, 15, 0, 2, 22},
			{9, 18, 13, 17, 5},
			{19, 8, 7, 25, 23},
			{20, 11, 10, 24, 4},
			{14, 21, 16, 12, 6},
		}),
		newBoard([][]int{
			{14, 21, 17, 24, 4},
			{10, 16, 15, 9, 19},
			{18, 8, 23, 26, 20},
			{22, 11, 13, 6, 5},
			{2, 0, 12, 3, 7},
		}),
	}

	board := lastWinningBoard(numbers, boards)

	if got, want := board, boards[1]; got.String() != want.String() {
		t.Errorf("got board:\n%s\n\nwant:\n%s", got, want)
	}

	if got, want := board.lastMarked, 13; got != want {
		t.Errorf("got last marked %d, want %d", got, want)
	}

	if got, want := board.score(), 1924; got != want {
		t.Errorf("got score %d, want %d", got, want)
	}
}

func TestBoard_Bingo(t *testing.T) {
	for _, tc := range []struct {
		desc    string
		numbers []int
		board   *board
		bingo   bool
	}{
		{
			desc:    "Not bingo",
			numbers: []int{1, 2, 3, 4, 5},
			board: newBoard([][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			}),
			bingo: false,
		},
		{
			desc:    "Horizontal",
			numbers: []int{1, 2, 3, 4, 5},
			board: newBoard([][]int{
				{0, 0, 0, 0, 0},
				{1, 2, 3, 4, 5},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			}),
			bingo: true,
		},
		{
			desc:    "Vertical",
			numbers: []int{1, 2, 3, 4, 5},
			board: newBoard([][]int{
				{0, 1, 0, 0, 0},
				{0, 2, 0, 0, 0},
				{0, 3, 0, 0, 0},
				{0, 4, 0, 0, 0},
				{0, 5, 0, 0, 0},
			}),
			bingo: true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			for _, n := range tc.numbers {
				tc.board.mark(n)
			}

			if got, want := tc.board.bingo(), tc.bingo; got != want {
				t.Errorf("got %t, want %t", got, want)
			}
		})
	}
}
