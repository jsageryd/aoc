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
	for n := markerLength - 1; n < len(input); n++ {
		m := make(map[byte]struct{})
		for i := 0; i < markerLength; i++ {
			m[input[n-i]] = struct{}{}
		}
		if len(m) == markerLength {
			return n + 1
		}
	}
	return 0
}
