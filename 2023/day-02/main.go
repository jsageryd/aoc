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
	fmt.Printf("Part 2: %d\n", part2(input))
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

func part2(input []string) int {
	var sum int

	for _, line := range input {
		_, sets := parse(line)

		highest := make(map[string]int)

		for _, set := range sets {
			for colour, count := range set {
				if highest[colour] < count {
					highest[colour] = count
				}
			}
		}

		power := 1
		for _, count := range highest {
			power *= count
		}

		sum += power
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
