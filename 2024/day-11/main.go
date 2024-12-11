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
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input string) int {
	return blink(parse(input), 25)
}

func part2(input string) int {
	return blink(parse(input), 75)
}

func blink(ints []int, n int) int {
	memory := make(map[[2]int]int)

	var rec func(n, iterations int) int

	rec = func(n, iterations int) int {
		if iterations == 0 {
			return 1
		}

		if sum, ok := memory[[2]int{n, iterations}]; ok {
			return sum
		}

		var sum int

		switch {
		case n == 0:
			sum += rec(1, iterations-1)
		case digits(n)%2 == 0:
			left, right := split(n)
			sum += rec(left, iterations-1)
			sum += rec(right, iterations-1)
		default:
			sum += rec(n*2024, iterations-1)
		}

		memory[[2]int{n, iterations}] = sum

		return sum
	}

	var sum int

	for _, v := range ints {
		sum += rec(v, n)
	}

	return sum
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
