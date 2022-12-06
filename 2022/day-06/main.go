package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Scanln(&input)

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input string) int {
	for n := 3; n < len(input); n++ {
		m := map[byte]struct{}{}
		m[input[n]] = struct{}{}
		m[input[n-1]] = struct{}{}
		m[input[n-2]] = struct{}{}
		m[input[n-3]] = struct{}{}
		if len(m) == 4 {
			return n + 1
		}
	}
	return 0
}
