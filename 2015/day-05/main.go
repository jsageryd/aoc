package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var niceStrings int
	var niceStrings2 int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if nice(scanner.Text()) {
			niceStrings++
		}

		if nice2(scanner.Text()) {
			niceStrings2++
		}
	}

	fmt.Printf("Part 1: %d\n", niceStrings)
	fmt.Printf("Part 2: %d\n", niceStrings2)
}

func nice(s string) bool {
	return hasThreeVowels(s) &&
		hasDoubleLetter(s) &&
		!hasForbiddenString(s, []string{"ab", "cd", "pq", "xy"})
}

func nice2(s string) bool {
	return hasTwoNonOverlappingPairs(s) && hasRepeatedLetterWithOneInbetween(s)
}

func hasThreeVowels(s string) bool {
	const vowels = "aiueo"
	var count int

	for _, r := range s {
		for _, v := range vowels {
			if r == v {
				count++
			}
			if count >= 3 {
				return true
			}
		}
	}

	return false
}

func hasDoubleLetter(s string) bool {
	var lastR rune

	for _, r := range s {
		if lastR == r {
			return true
		}

		lastR = r
	}

	return false
}

func hasForbiddenString(s string, forbidden []string) bool {
	for _, f := range forbidden {
		if strings.Contains(s, f) {
			return true
		}
	}

	return false
}

func hasTwoNonOverlappingPairs(s string) bool {
	pairs := map[string]struct{}{}

	var (
		lastR    rune
		lastPair string
	)

	for n, r := range s {
		if n != 0 {
			pair := string([]rune{lastR, r})

			if _, ok := pairs[pair]; ok {
				return true
			}

			pairs[lastPair] = struct{}{}
			lastPair = pair
		}

		lastR = r
	}

	return false
}

func hasRepeatedLetterWithOneInbetween(s string) bool {
	var lastR, lastLastR rune

	for _, r := range s {
		if r == lastLastR {
			return true
		}
		lastR, lastLastR = r, lastR
	}

	return false
}
