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
	for h := 1; ; h++ {
		var p int

		for e := 1; e <= h; e++ {
			if h%e == 0 {
				p += 10 * e
			}
		}

		if p >= input {
			return h
		}
	}

	return 0
}
