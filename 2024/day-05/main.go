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
}

func part1(input []string) int {
	var sum int

	rules, updates := parse(input)

	for n := range updates {
		if slices.IsSortedFunc(updates[n], func(a, b int) int {
			for _, rule := range rules {
				switch rule {
				case [2]int{a, b}:
					return -1
				case [2]int{b, a}:
					return 1
				}
			}
			return 0
		}) {
			sum += updates[n][len(updates[n])/2]
		}
	}

	return sum
}

func parse(input []string) (rules [][2]int, updates [][]int) {
	sepIdx := slices.Index(input, "")
	if sepIdx == -1 {
		return nil, nil
	}

	for _, line := range input[:sepIdx] {
		n1Str, n2Str, _ := strings.Cut(line, "|")
		n1, _ := strconv.Atoi(n1Str)
		n2, _ := strconv.Atoi(n2Str)
		rules = append(rules, [2]int{n1, n2})
	}

	for _, line := range input[sepIdx+1:] {
		var update []int
		updateStrs := strings.Split(line, ",")
		for _, s := range updateStrs {
			i, _ := strconv.Atoi(s)
			update = append(update, i)
		}
		updates = append(updates, update)
	}

	return rules, updates
}
