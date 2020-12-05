package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", highestSeatID(input))
}

func highestSeatID(specs []string) int {
	var highestID int

	for _, spec := range specs {
		_, _, id := parseSpec(spec)
		if id > highestID {
			highestID = id
		}
	}

	return highestID
}

func parseSpec(spec string) (row, col, id int) {
	rp, cp := 64, 4

	for _, r := range spec {
		switch r {
		case 'F':
			rp /= 2
		case 'B':
			row |= rp
			rp /= 2
		case 'L':
			cp /= 2
		case 'R':
			col |= cp
			cp /= 2
		}
	}

	return row, col, row*8 + col
}
