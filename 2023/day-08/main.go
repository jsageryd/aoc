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
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	m := make(map[string]map[byte]string) // AAA -> {L: BBB, R: CCC}

	for _, line := range input[2:] {
		m[line[0:3]] = map[byte]string{
			'L': line[7:10],
			'R': line[12:15],
		}
	}

	var steps int

	cur := "AAA"
	inst := input[0]

	for cur != "ZZZ" {
		cur = m[cur][inst[steps%len(inst)]]
		steps++
	}

	return steps
}

func part2(input []string) int {
	m := make(map[string]map[byte]string) // AAA -> {L: BBB, R: CCC}

	for _, line := range input[2:] {
		m[line[0:3]] = map[byte]string{
			'L': line[7:10],
			'R': line[12:15],
		}
	}

	var curs []string

	for _, line := range input[2:] {
		if line[2] == 'A' {
			curs = append(curs, line[0:3])
		}
	}

	var steps int

	inst := input[0]
	cycles := make(map[int]int)

	for len(cycles) < len(curs) {
		for n := range curs {
			if curs[n][2] == 'Z' {
				cycles[n] = steps
			}
		}

		for n := range curs {
			curs[n] = m[curs[n]][inst[steps%len(inst)]]
		}

		steps++
	}

	minCycle := cycles[0]

	for _, cycle := range cycles {
		if cycle < minCycle {
			minCycle = cycle
		}
	}

	var arrived bool

	for {
		arrived = true
		for n := range cycles {
			if steps%cycles[n] != 0 {
				arrived = false
			}
		}
		if arrived {
			break
		}
		steps += minCycle
	}

	return steps
}
