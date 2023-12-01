package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input []string

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		input = append(input, s.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
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

func part2(input []string) int {
	var sum int

	words := []string{
		"one", "two", "three",
		"four", "five", "six",
		"seven", "eight", "nine",
	}

	for _, line := range input {
		var digits []int

		for n := range line {
			if line[n] >= '0' && line[n] <= '9' {
				digits = append(digits, int(line[n]-'0'))
				continue
			}

			for m := range words {
				if strings.HasPrefix(line[n:], words[m]) {
					digits = append(digits, m+1)
					break
				}
			}
		}

		sum += digits[0]*10 + digits[len(digits)-1]
	}

	return sum
}
