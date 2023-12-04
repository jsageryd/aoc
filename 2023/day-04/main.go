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

	for _, line := range input {
		c := parse(line)
		sum += (1 << matches(c)) >> 1
	}

	return sum
}

func part2(input []string) int {
	cards := make(map[int]card) // card number -> card

	var cardIDs []int
	var scratchedCardIDs []int

	for n, line := range input {
		c := parse(line)
		cards[n+1] = c
		cardIDs = append(cardIDs, n+1)
	}

	for len(cardIDs) > 0 {
		c := cards[cardIDs[0]]

		for n := 1; n <= matches(c); n++ {
			cardIDs = append(cardIDs, cardIDs[0]+n)
		}

		scratchedCardIDs = append(scratchedCardIDs, cardIDs[0])
		cardIDs = cardIDs[1:]
	}

	return len(scratchedCardIDs)
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

func matches(c card) int {
	var matches int

	for _, winningN := range c.winning {
		if slices.Contains(c.have, winningN) {
			matches++
		}
	}

	return matches
}
