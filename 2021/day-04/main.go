package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numbers := stringsToInts(strings.Split(scanner.Text(), ","))

	var boards []*board

	var rows [][]int
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}
		rows = append(rows, stringsToInts(strings.Fields(scanner.Text())))
		if len(rows) == 5 {
			boards = append(boards, newBoard(rows))
			rows = [][]int{}
		}
	}

	fmt.Printf("Part 1: %d\n", firstWinningBoard(numbers, boards).score())
	fmt.Printf("Part 2: %d\n", lastWinningBoard(numbers, boards).score())
}

func firstWinningBoard(numbers []int, boards []*board) *board {
	for _, n := range numbers {
		for _, b := range boards {
			b.mark(n)
			if b.bingo() {
				return b
			}
		}
	}
	return nil
}

func lastWinningBoard(numbers []int, boards []*board) *board {
	var lastBoard *board
	done := make(map[*board]bool)

	for _, n := range numbers {
		for _, b := range boards {
			if !done[b] {
				b.mark(n)
				if b.bingo() {
					lastBoard = b
					done[b] = true
				}
			}
		}
	}

	return lastBoard
}

type board struct {
	grid       [][]int
	marked     map[int]bool
	lastMarked int
}

func newBoard(grid [][]int) *board {
	return &board{
		grid:   grid,
		marked: make(map[int]bool),
	}
}

func (b *board) mark(n int) {
	for y := range b.grid {
		for x := range b.grid[y] {
			if b.grid[y][x] == n {
				b.marked[n] = true
				b.lastMarked = n
			}
		}
	}
}

func (b *board) bingo() bool {
	if len(b.grid) == 0 || len(b.grid[0]) == 0 {
		return false
	}

	// horizontal
	for y := 0; y < len(b.grid); y++ {
		var xCount int
		for x := 0; x < len(b.grid[0]); x++ {
			if b.marked[b.grid[y][x]] {
				xCount++
			}
		}
		if xCount == 5 {
			return true
		}
	}

	// vertical
	for x := 0; x < len(b.grid[0]); x++ {
		var yCount int
		for y := 0; y < len(b.grid); y++ {
			if b.marked[b.grid[y][x]] {
				yCount++
			}
		}
		if yCount == 5 {
			return true
		}
	}

	return false
}

func (b *board) score() int {
	var sumUnmarked int
	for y := range b.grid {
		for x := range b.grid[y] {
			if !b.marked[b.grid[y][x]] {
				sumUnmarked += b.grid[y][x]
			}
		}
	}
	return sumUnmarked * b.lastMarked
}

func (b *board) String() string {
	if b == nil {
		return ""
	}

	var buf bytes.Buffer

	tw := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.AlignRight)
	for y := range b.grid {
		for x := range b.grid[y] {
			if b.marked[b.grid[y][x]] {
				fmt.Fprintf(tw, "[%d]\t", b.grid[y][x])
			} else {
				fmt.Fprintf(tw, "%d\t", b.grid[y][x])
			}
		}
		fmt.Fprint(tw, "\n")
	}
	tw.Flush()

	var rows []string

	for _, row := range strings.Split(buf.String(), "\n") {
		if len(row) > 0 {
			rows = append(rows, row[1:])
		}
	}

	return strings.Join(rows, "\n")
}

func stringsToInts(strings []string) []int {
	ints := make([]int, 0, len(strings))
	for _, s := range strings {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, n)
	}
	return ints
}
