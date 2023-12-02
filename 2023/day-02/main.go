package main

import (
	"bufio"
	"fmt"
	"os"
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

	limit := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

nextGame:
	for _, line := range input {
		gameID, sets := parse(line)

		for _, set := range sets {
			for colour, count := range set {
				if count > limit[colour] {
					continue nextGame
				}
			}
		}

		sum += gameID
	}

	return sum
}

func parse(line string) (gameID int, sets []map[string]int) {
	head, tail, _ := strings.Cut(line, ": ")

	fmt.Sscanf(head, "Game %d", &gameID)

	for _, setStr := range strings.Split(tail, "; ") {
		set := make(map[string]int)
		for _, countAndColour := range strings.Split(setStr, ", ") {
			var count int
			var colour string
			fmt.Sscanf(countAndColour, "%d %s", &count, &colour)
			set[colour] = count
		}
		sets = append(sets, set)
	}

	return gameID, sets
}
