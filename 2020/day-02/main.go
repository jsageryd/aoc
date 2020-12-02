package main

import (
	"fmt"
)

func main() {
	var (
		n1, n2   int
		letter   byte
		password string
	)

	format := "%d-%d %c: %s"

	var validCount int
	var validPart2Count int

	for {
		if _, err := fmt.Scanf(format, &n1, &n2, &letter, &password); err != nil {
			break
		}

		if valid(n1, n2, letter, password) {
			validCount++
		}

		if validPart2(n1, n2, letter, password) {
			validPart2Count++
		}
	}

	fmt.Printf("Part 1: %d\n", validCount)
	fmt.Printf("Part 2: %d\n", validPart2Count)
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

func validPart2(pos1, pos2 int, letter byte, password string) bool {
	p1 := password[pos1-1] == letter
	p2 := password[pos2-1] == letter
	return (p1 || p2) && !(p1 && p2)
}
