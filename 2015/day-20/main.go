package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Scanln(&input)

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input int) int {
	p := make([]int, input) // house -> presents

	for e := 1; e < len(p); e++ {
		for h := 0; h < len(p); h += e {
			p[h] += 10 * e
		}

		// p is house -> presents so p[e] looks like an error, but e is already
		// looping from 1 to len(p), so we can re-use that.
		if p[e] >= input {
			return e
		}
	}

	return 0
}
