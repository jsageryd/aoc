package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	list1, list2 := parse(input)

	slices.Sort(list1)
	slices.Sort(list2)

	var sum int

	for n := range list1 {
		sum += abs(list1[n] - list2[n])
	}

	return sum
}

func part2(input []string) int {
	list1, list2 := parse(input)

	freq := make(map[int]int)

	for n := range list2 {
		freq[list2[n]]++
	}

	var sum int

	for n := range list1 {
		sum += list1[n] * freq[list1[n]]
	}

	return sum
}

func parse(input []string) (list1, list2 []int) {
	for _, line := range input {
		var n1, n2 int

		fmt.Sscanf(line, "%d %d", &n1, &n2)

		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	return list1, list2
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
