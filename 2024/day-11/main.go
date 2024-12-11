package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		input = s.Text()
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input string) int {
	ints := parse(input)

	for range 25 {
		var newInts []int

		for n := range ints {
			switch {
			case ints[n] == 0:
				newInts = append(newInts, 1)
			case digits(ints[n])%2 == 0:
				left, right := split(ints[n])
				newInts = append(newInts, left, right)
			default:
				newInts = append(newInts, ints[n]*2024)
			}
		}

		ints = newInts
	}

	return len(ints)
}

func parse(input string) []int {
	var ints []int

	for _, s := range strings.Fields(input) {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}

	return ints
}

func digits(n int) int {
	var digits int

	for ; n > 0; n /= 10 {
		digits++
	}

	return digits
}

func split(n int) (int, int) {
	m := 1

	for range digits(n) / 2 {
		m *= 10
	}

	return n / m, n % m
}
