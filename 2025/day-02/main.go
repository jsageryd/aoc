package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	inputBytes, _ := io.ReadAll(os.Stdin)
	input := string(inputBytes)

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input string) int {
	var sum int

	for _, r := range parse(input) {
		for n := r[0]; n <= r[1]; n++ {
			m := 1

			for range digits(n) / 2 {
				m *= 10
			}

			if n/m == n%m {
				sum += n
			}
		}
	}

	return sum
}

func digits(n int) int {
	if n == 0 {
		return 1
	}

	var digits int

	for ; n > 0; n /= 10 {
		digits++
	}

	return digits
}

func parse(input string) [][2]int {
	var ranges [][2]int

	for _, rStr := range strings.Split(input, ",") {
		var r [2]int
		fmt.Sscanf(rStr, "%d-%d", &r[0], &r[1])
		ranges = append(ranges, r)
	}

	return ranges
}
