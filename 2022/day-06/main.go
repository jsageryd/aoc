package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Scanln(&input)

	fmt.Printf("Part 1: %d\n", startOfMessage(input, 4))
	fmt.Printf("Part 2: %d\n", startOfMessage(input, 14))
}

func startOfMessage(input string, markerLength int) int {
next:
	for n := markerLength - 1; n < len(input); n++ {
		for i := 0; i < markerLength; i++ {
			for j := 0; j < markerLength; j++ {
				if i != j && input[n-i] == input[n-j] {
					continue next
				}
			}
		}
		return n + 1
	}
	return 0
}
