package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", totalScore(input))
}

func totalScore(input []string) int {
	xyzToAbc := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	var totalScore int

	for _, line := range input {
		theirs, ours, _ := strings.Cut(line, " ")
		totalScore += play(xyzToAbc[ours], theirs)
	}

	return totalScore
}

func play(ours, theirs string) (score int) {
	return map[string]int{
		"AA": 4, "AB": 1, "AC": 7,
		"BA": 8, "BB": 5, "BC": 2,
		"CA": 3, "CB": 9, "CC": 6,
	}[ours+theirs]
}
