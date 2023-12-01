package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input []string

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		input = append(input, s.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	var sum int

	for _, line := range input {
		var digits []int

		for _, r := range line {
			if r >= '0' && r <= '9' {
				digits = append(digits, int(r-'0'))
			}
		}

		sum += digits[0]*10 + digits[len(digits)-1]
	}

	return sum
}
