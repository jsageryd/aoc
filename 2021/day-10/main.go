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

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	scores := map[byte]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	var score int
	for _, line := range input {
		if c, ok := findIncorrectChar(line); ok {
			score += scores[c]
		}
	}
	return score
}

func findIncorrectChar(line string) (c byte, ok bool) {
	matching := map[byte]byte{')': '(', ']': '[', '}': '{', '>': '<'}
	var stack []byte
	for i := range line {
		switch c := line[i]; c {
		case '(', '[', '{', '<':
			stack = append(stack, c)
		case ')', ']', '}', '>':
			if stack[len(stack)-1] != matching[c] {
				return c, true
			}
			stack = stack[:len(stack)-1]
		}
	}
	return 0, false
}
