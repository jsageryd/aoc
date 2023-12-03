package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input [][]byte

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		input = append(input, []byte(s.Text()))
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input [][]byte) int {
	var sum int

	symbols, digits := parse(input)

	for sc := range symbols {
		for _, n := range neighbouringNumbers(sc, digits) {
			sum += n
		}
	}

	return sum
}

func part2(input [][]byte) int {
	var sum int

	symbols, digits := parse(input)

	for sc, sb := range symbols {
		if sb != '*' {
			continue
		}

		nn := neighbouringNumbers(sc, digits)
		if len(nn) < 2 {
			continue
		}

		ratio := 1
		for _, n := range nn {
			ratio *= n
		}

		sum += ratio
	}

	return sum
}

func parse(input [][]byte) (symbols, digits map[coord]byte) {
	symbols = make(map[coord]byte)
	digits = make(map[coord]byte)

	for y := range input {
		for x := range input[y] {
			c := coord{x, y}
			b := input[y][x]

			switch {
			case isSymbol(b):
				symbols[c] = b
			case isDigit(b):
				digits[c] = b
			}
		}
	}

	return symbols, digits
}

func isSymbol(b byte) bool {
	return !isDigit(b) && b != '.'
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func neighbouringNumbers(c coord, digits map[coord]byte) []int {
	var nums []int

	seen := make(map[coord]bool)
	for _, nc := range neighbours(c) {
		if _, ok := digits[nc]; ok {
			numC, n := findNumber(nc, digits)
			if !seen[numC] {
				nums = append(nums, n)
				seen[numC] = true
			}
		}
	}

	return nums
}

func findNumber(c coord, digits map[coord]byte) (firstDigit coord, n int) {
	for {
		if _, ok := digits[coord{c.x + 1, c.y}]; !ok {
			break
		}
		c.x++
	}

	var num int
	m := 1

	for d, ok := digits[c]; ok; d, ok = digits[c] {
		c.x--
		num += m * (int(d) - '0')
		m *= 10
	}

	return c, num
}

type coord struct {
	x, y int
}

func neighbours(c coord) []coord {
	return []coord{
		{c.x + 1, c.y}, {c.x + 1, c.y - 1},
		{c.x, c.y - 1}, {c.x - 1, c.y - 1},
		{c.x - 1, c.y}, {c.x - 1, c.y + 1},
		{c.x, c.y + 1}, {c.x + 1, c.y + 1},
	}
}
