package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", len(fullyOverlappingPairs(input)))
	fmt.Printf("Part 1: %d\n", len(partiallyOverlappingPairs(input)))
}

func fullyOverlappingPairs(input []string) []string {
	var pairs []string

	for _, line := range input {
		rangeA, rangeB := parsePair(line)
		if rangeA.contains(rangeB) || rangeB.contains(rangeA) {
			pairs = append(pairs, line)
		}
	}

	return pairs
}

func partiallyOverlappingPairs(input []string) []string {
	var pairs []string

	for _, line := range input {
		rangeA, rangeB := parsePair(line)
		if rangeA.overlaps(rangeB) {
			pairs = append(pairs, line)
		}
	}

	return pairs
}

func parsePair(pair string) (rangeA, rangeB rang) {
	fmt.Sscanf(
		pair, "%d-%d,%d-%d",
		&rangeA.from, &rangeA.to,
		&rangeB.from, &rangeB.to,
	)
	return rangeA, rangeB
}

type rang struct {
	from, to int
}

func (r rang) contains(other rang) bool {
	return r.from <= other.from && r.to >= other.to
}

func (r rang) overlaps(other rang) bool {
	return !(r.from < other.from && r.to < other.from) &&
		!(r.to > other.to && r.from > other.to)
}
