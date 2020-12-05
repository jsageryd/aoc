package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanln(&input)

	part1Str := input
	for i := 0; i < 40; i++ {
		part1Str = lookAndSay(part1Str)
	}

	fmt.Printf("Part 1: %d\n", len(part1Str))
}

func lookAndSay(input string) string {
	var b strings.Builder

	var count int
	var lastR rune
	for _, r := range input {
		if count != 0 && r != lastR {
			fmt.Fprintf(&b, "%d%c", count, lastR)
			count = 0
		}
		count++
		lastR = r
	}
	fmt.Fprintf(&b, "%d%c", count, lastR)

	return b.String()
}
