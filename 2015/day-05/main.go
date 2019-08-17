package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var niceStrings int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if nice(scanner.Text()) {
			niceStrings++
		}
	}

	fmt.Printf("Part 1: %d\n", niceStrings)
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
