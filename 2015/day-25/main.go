package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	inputBytes, _ := io.ReadAll(os.Stdin)
	input := string(inputBytes)

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input string) int {
	code := 20151125

	row, column := parse(input)

	for range coordToIndex(column, row) - 1 {
		code *= 252533
		code %= 33554393
	}

	return code
}

func coordToIndex(x, y int) int {
	idx := 1

	dx := 1

	for range x - 1 {
		dx++
		idx += dx
	}

	dy := dx - 1

	for range y - 1 {
		dy++
		idx += dy
	}

	return idx
}

func parse(input string) (row, column int) {
	idx := strings.Index(input, "row")
	if idx == -1 {
		return -1, -1
	}

	fmt.Sscanf(input[idx:], "row %d, column %d", &row, &column)

	return row, column
}
