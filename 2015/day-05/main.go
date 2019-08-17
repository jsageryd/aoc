package main

import (
	"bufio"
	"fmt"
	"os"
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
	const vowels = "aeiou"

	var (
		hasThreeVowels     bool
		hasDoubleLetter    bool
		hasForbiddenString bool
	)

	var vowelCount int
	var lastR rune

	for _, r := range s {
		for _, v := range vowels {
			if r == v {
				vowelCount++
				break
			}
		}

		if r == lastR {
			hasDoubleLetter = true
		}

		if lastR == 'a' && r == 'b' ||
			lastR == 'c' && r == 'd' ||
			lastR == 'p' && r == 'q' ||
			lastR == 'x' && r == 'y' {
			hasForbiddenString = true
		}

		lastR = r
	}

	hasThreeVowels = vowelCount >= 3

	return hasThreeVowels && hasDoubleLetter && !hasForbiddenString
}

func nice2(s string) bool {
	return hasTwoNonOverlappingPairs(s) && hasRepeatedLetterWithOneInbetween(s)
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
