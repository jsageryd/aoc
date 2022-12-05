package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %s\n", part1(input))
	fmt.Printf("Part 2: %s\n", part2(input))
}

func part1(input []string) string {
	stacks, steps := parseInput(input)

	moveCrates(stacks, steps)

	var s string

	for _, stack := range stacks {
		s += stack.pop()
	}

	return s
}

func part2(input []string) string {
	stacks, steps := parseInput(input)

	moveCrates2(stacks, steps)

	var s string

	for _, stack := range stacks {
		s += stack.pop()
	}

	return s
}

func parseInput(input []string) (stacks []stack, steps []step) {
	for n := range input {
		if input[n] == "" {
			return parseStacks(input[:n]), parseSteps(input[n+1:])
		}
	}
	return nil, nil
}

func parseStacks(stacksStrs []string) []stack {
	stackCount := len(strings.Fields(stacksStrs[len(stacksStrs)-1]))
	stacks := make([]stack, stackCount)

	for n := len(stacksStrs) - 1; n >= 0; n-- {
		for m := range stacksStrs[n] {
			if stacksStrs[n][m] == '[' {
				stacks[m/4].push(string(stacksStrs[n][m+1]))
			}
		}
	}

	return stacks
}

func parseSteps(stepStrs []string) []step {
	var steps []step

	for _, stepStr := range stepStrs {
		var s step
		fmt.Sscanf(stepStr, "move %d from %d to %d", &s.count, &s.from, &s.to)
		steps = append(steps, s)
	}

	return steps
}

func moveCrates(stacks []stack, steps []step) {
	for _, step := range steps {
		for n := 0; n < step.count; n++ {
			stacks[step.to-1].push(stacks[step.from-1].pop())
		}
	}
}

func moveCrates2(stacks []stack, steps []step) {
	for _, step := range steps {
		var buf stack

		for n := 0; n < step.count; n++ {
			buf.push(stacks[step.from-1].pop())
		}

		for n := 0; n < step.count; n++ {
			stacks[step.to-1].push(buf.pop())
		}
	}
}

type stack []string

func (s *stack) push(item string) {
	*s = append(*s, item)
}

func (s *stack) pop() (item string) {
	item, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return item
}

type step struct {
	count    int
	from, to int
}
