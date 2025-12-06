package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
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

	for _, problem := range slice(input) {
		var op func(a, b int) int

		switch opStr := strings.TrimSpace(problem[len(problem)-1]); opStr {
		case "+":
			op = func(a, b int) int { return a + b }
		case "*":
			op = func(a, b int) int { return a * b }
		default:
			log.Fatalf("unknown op string %q", opStr)
		}

		var values []int

		for _, v := range problem[:len(problem)-1] {
			i, err := strconv.Atoi(strings.TrimSpace(v))
			if err != nil {
				log.Fatal(err)
			}

			values = append(values, i)
		}

		res := values[0]

		for _, v := range values[1:] {
			res = op(res, v)
		}

		sum += res
	}

	return sum
}

func slice(input []string) [][]string {
	var idxs []int

	// Find offsets based on last line
	for idx, r := range input[len(input)-1] {
		if r != ' ' {
			idxs = append(idxs, idx)
		}
	}

	ss := make([][]string, len(idxs))

	// Slice following the offsets
	for y := range input {
		from := idxs[0]

		for x, to := range idxs[1:] {
			ss[x] = append(ss[x], input[y][from:to-1])
			from = to
		}

		ss[len(ss)-1] = append(ss[len(ss)-1], input[y][from:])
	}

	// Add spacing to the last slice to make it look like the others
	lastSlice := ss[len(ss)-1]

	maxLen := len(slices.MaxFunc(lastSlice, func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}))

	for n := range lastSlice {
		lastSlice[n] += strings.Repeat(" ", maxLen-len(lastSlice[n]))
	}

	return ss
}
