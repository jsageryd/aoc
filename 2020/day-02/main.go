package main

import (
	"fmt"
)

func main() {
	var (
		min, max int
		letter   byte
		password string
	)

	format := "%d-%d %c: %s"

	var validCount int

	for {
		if _, err := fmt.Scanf(format, &min, &max, &letter, &password); err != nil {
			break
		}

		if valid(min, max, letter, password) {
			validCount++
		}
	}

	fmt.Printf("Part 1: %d\n", validCount)
}

func valid(min, max int, letter byte, password string) bool {
	var count int
	for n := range password {
		if password[n] == letter {
			count++
		}
	}
	return count >= min && count <= max
}
