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

	fmt.Printf("Part 1: %s\n", part1(input))
	fmt.Printf("Part 2: %s\n", part2(input))
}

func part1(input []string) string {
	stacks, steps := parseInput(input)

	moveCrates(stacks, steps)

	s := make([]rune, len(stacks))

	for idx, stack := range stacks {
		s[idx-1] = rune(stack.pop()[0])
	}

	return string(s)
}

func part2(input []string) string {
	stacks, steps := parseInput(input)

	moveCrates2(stacks, steps)

	s := make([]rune, len(stacks))

	for idx, stack := range stacks {
		s[idx-1] = rune(stack.pop()[0])
	}

	return string(s)
}

func parseInput(input []string) (stacks map[int]stack, steps []step) {
	for n := range input {
		if input[n] == "" {
			return parseStacks(input[:n]), parseSteps(input[n+1:])
		}
	}
	return nil, nil
}

func parseStacks(stacksStrs []string) map[int]stack {
	stacks := make(map[int]stack)

	for n := len(stacksStrs) - 1; n >= 0; n-- {
		for m := range stacksStrs[n] {
			if stacksStrs[n][m] == '[' {
				idx := m/4 + 1
				s := stacks[idx]
				s.push(string(stacksStrs[n][m+1]))
				stacks[idx] = s
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

func moveCrates(stacks map[int]stack, steps []step) {
	for _, step := range steps {
		from, to := stacks[step.from], stacks[step.to]

		for n := 0; n < step.count; n++ {
			to.push(from.pop())
		}

		stacks[step.from], stacks[step.to] = from, to
	}
}

func moveCrates2(stacks map[int]stack, steps []step) {
	for _, step := range steps {
		from, to := stacks[step.from], stacks[step.to]

		var buf stack

		for n := 0; n < step.count; n++ {
			buf.push(from.pop())
		}

		for n := 0; n < step.count; n++ {
			to.push(buf.pop())
		}

		stacks[step.from], stacks[step.to] = from, to
	}
}

type stack []string

func (s *stack) push(item string) {
	*s = append(*s, item)
}

func (s *stack) pop() string {
	*s = (*s)[:len(*s)-1]
	return (*s)[:len(*s)+1][len(*s)]
}

type step struct {
	count    int
	from, to int
}
