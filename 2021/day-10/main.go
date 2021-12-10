package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
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

func part2(input []string) int {
	scoreMap := map[byte]int{')': 1, ']': 2, '}': 3, '>': 4}
	var scores []int
	for _, line := range input {
		if _, ok := findIncorrectChar(line); !ok {
			tail := findTail(line)
			var score int
			for i := range tail {
				score *= 5
				score += scoreMap[tail[i]]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
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

func findTail(line string) string {
	matching := map[byte]byte{'(': ')', '[': ']', '{': '}', '<': '>'}
	var stack []byte
	for i := range line {
		switch c := line[i]; c {
		case '(', '[', '{', '<':
			stack = append(stack, c)
		case ')', ']', '}', '>':
			stack = stack[:len(stack)-1]
		}
	}
	var tail []byte
	for i := len(stack) - 1; i >= 0; i-- {
		tail = append(tail, matching[stack[i]])
	}
	return string(tail)
}
