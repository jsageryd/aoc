package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var input []string

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		input = append(input, s.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	var sum int

	tests := parse(input)

	for _, t := range tests {
		results := calculate(t.numbers, []op{add, mul})

		if slices.Contains(results, t.value) {
			sum += t.value
		}
	}

	return sum
}

func part2(input []string) int {
	var sum int

	tests := parse(input)

	for _, t := range tests {
		results := calculate(t.numbers, []op{add, mul, cat})

		if slices.Contains(results, t.value) {
			sum += t.value
		}
	}

	return sum
}

func calculate(numbers []int, ops []op) []int {
	var res []int

	var rec func(acc int, numbers []int)

	rec = func(acc int, numbers []int) {
		if len(numbers) == 0 {
			res = append(res, acc)
			return
		}

		for _, op := range ops {
			rec(op(acc, numbers[0]), numbers[1:])
		}
	}

	rec(numbers[0], numbers[1:])

	return res
}

type op func(n1, n2 int) int

func add(n1, n2 int) int {
	return n1 + n2
}

func mul(n1, n2 int) int {
	return n1 * n2
}

func cat(n1, n2 int) int {
	n, _ := strconv.Atoi(strconv.Itoa(n1) + strconv.Itoa(n2))
	return n
}

type Test struct {
	value   int
	numbers []int
}

func parse(input []string) []Test {
	var tests []Test

	for n := range input {
		valueStr, numbersStr, _ := strings.Cut(input[n], ": ")
		value, _ := strconv.Atoi(valueStr)
		var numbers []int
		for _, numStr := range strings.Fields(numbersStr) {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}
		tests = append(tests, Test{value, numbers})
	}

	return tests
}
