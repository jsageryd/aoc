package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func main() {
	inputBytes, _ := io.ReadAll(os.Stdin)
	input := string(inputBytes)

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
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

func part2(input string) int {
	var sum int

	for _, r := range parse(input) {
		for n := r[0]; n <= r[1]; n++ {
			for i := 2; i <= digits(n); i++ {
				if len(slices.Compact(split(n, i))) == 1 {
					sum += n
					break
				}
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

func split(n, parts int) []int {
	d := digits(n)

	if d%parts != 0 {
		return nil
	}

	partSize := d / parts

	var s []int

	for exp := d - partSize; exp >= 0; exp -= partSize {
		m := 1

		for range exp {
			m *= 10
		}

		s = append(s, n/m)

		n %= m
	}

	return s
}
