package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Scanln(&input)

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
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

func part2(input string) int {
	for n := 13; n < len(input); n++ {
		m := map[byte]struct{}{}
		m[input[n]] = struct{}{}
		m[input[n-1]] = struct{}{}
		m[input[n-2]] = struct{}{}
		m[input[n-3]] = struct{}{}
		m[input[n-4]] = struct{}{}
		m[input[n-5]] = struct{}{}
		m[input[n-6]] = struct{}{}
		m[input[n-7]] = struct{}{}
		m[input[n-8]] = struct{}{}
		m[input[n-9]] = struct{}{}
		m[input[n-10]] = struct{}{}
		m[input[n-11]] = struct{}{}
		m[input[n-12]] = struct{}{}
		m[input[n-13]] = struct{}{}
		if len(m) == 14 {
			return n + 1
		}
	}
	return 0
}
