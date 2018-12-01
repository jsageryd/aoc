package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Scanln(&input)

	fmt.Printf("Part 1: %d\n", inverseCaptcha(input, 1))
	fmt.Printf("Part 2: %d\n", inverseCaptcha(input, len(input)/2))
}

func inverseCaptcha(seq string, lookahead int) int {
	sum := 0

	for n := range seq {
		if seq[n] == seq[(n+lookahead)%len(seq)] {
			d, _ := strconv.Atoi(string(seq[n]))
			sum += d
		}
	}

	return sum
}
