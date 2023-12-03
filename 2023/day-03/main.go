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

	s := parseSchematic(input)

	for y := range input {
		for x := range input[y] {
			if n, isPartNum := s.num(x, y); n != 0 {
				if isPartNum {
					sum += n
				}
			}
		}
	}

	return sum
}

func part2(input [][]byte) int {
	var sum int

	s := parseSchematic(input)
	heads, tails := s.numsMap()

	for y := range input {
		for x := range input[y] {
			if s[coord{x, y}] == '*' {
				nums := make(map[coord]int)

				for _, nc := range (coord{x, y}).neighbours() {
					if n, ok := heads[nc]; ok {
						nums[nc] = n
					}
					if tc, ok := tails[nc]; ok {
						nums[tails[tc]] = heads[tc]
					}
				}

				if len(nums) < 2 {
					continue
				}

				ratio := 1
				for _, n := range nums {
					ratio *= n
				}

				sum += ratio
			}
		}
	}

	return sum
}

func parseSchematic(input [][]byte) schematic {
	s := make(schematic)

	for y := range input {
		for x := range input[y] {
			s[coord{x, y}] = input[y][x]
		}
	}

	return s
}

type coord struct {
	x, y int
}

func (c coord) neighbours() []coord {
	return []coord{
		{c.x + 1, c.y},
		{c.x + 1, c.y - 1},
		{c.x, c.y - 1},
		{c.x - 1, c.y - 1},
		{c.x - 1, c.y},
		{c.x - 1, c.y + 1},
		{c.x, c.y + 1},
		{c.x + 1, c.y + 1},
	}
}

type schematic map[coord]byte

func (s schematic) num(x, y int) (n int, isPartNum bool) {
	isDigit := func(b byte) bool {
		return b >= '0' && b <= '9'
	}

	isSymbol := func(b byte) bool {
		return b != 0 && b != '.' && !isDigit(b)
	}

	// return if the coordinate is not a digit
	if !isDigit(s[coord{x, y}]) {
		return 0, false
	}

	// return if this is not the first digit in the number
	if isDigit(s[coord{x - 1, y}]) {
		return 0, false
	}

	var digits []int
	for xx := x; isDigit(s[coord{xx, y}]); xx++ {
		digits = append(digits, int(s[coord{xx, y}])-'0')

		for _, nc := range (coord{xx, y}).neighbours() {
			if isSymbol(s[nc]) {
				isPartNum = true
			}
		}
	}

	var num int

	m := 1
	for n := len(digits) - 1; n >= 0; n-- {
		num += m * digits[n]
		m *= 10
	}

	return num, isPartNum
}

// numsMap returns a map of all coordinates that represent the first digit
// (head) of a number, as well as a mapping of each of the following digits
// (tail) of each number to its head.
func (s schematic) numsMap() (heads map[coord]int, tails map[coord]coord) {
	isDigit := func(b byte) bool {
		return b >= '0' && b <= '9'
	}

	heads = make(map[coord]int)
	tails = make(map[coord]coord)

	for c := range s {
		if n, _ := s.num(c.x, c.y); n != 0 {
			heads[c] = n
			for x := c.x; isDigit(s[coord{x, c.y}]); x++ {
				tails[coord{x, c.y}] = c
			}
		}
	}

	return heads, tails
}
