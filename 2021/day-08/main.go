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

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	var count int
	for _, entry := range input {
		sep := strings.IndexByte(entry, '|')
		values := strings.Fields(entry[sep:])
		for _, v := range values {
			switch len(v) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	return count
}
