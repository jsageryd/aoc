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

	for _, line := range input {
		c := parse(line)
		var matches int
		for _, winningN := range c.winning {
			if slices.Contains(c.have, winningN) {
				matches++
			}
		}
		sum += (1 << matches) >> 1
	}

	return sum
}

func parse(line string) card {
	var c card

	_, tail, _ := strings.Cut(line, ":")
	winningStr, haveStr, _ := strings.Cut(tail, "|")

	for _, s := range strings.Fields(winningStr) {
		n, _ := strconv.Atoi(s)
		c.winning = append(c.winning, n)
	}

	for _, s := range strings.Fields(haveStr) {
		n, _ := strconv.Atoi(s)
		c.have = append(c.have, n)
	}

	return c
}

type card struct {
	winning []int
	have    []int
}
