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
