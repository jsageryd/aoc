package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	var sum int

	for _, items := range input {
		sum += priority(duplicateItem(items))
	}

	return sum
}

func part2(input []string) int {
	var sum int

	for n := 0; n < len(input); n += 3 {
		sum += priority(commonItem(input[n], input[n+1], input[n+2]))
	}

	return sum
}

func duplicateItem(items string) string {
	a, b := split(items)
	for _, aa := range a {
		for _, bb := range b {
			if aa == bb {
				return string(aa)
			}
		}
	}
	return ""
}

func split(items string) (a, b string) {
	return items[:len(items)/2], items[len(items)/2:]
}

func commonItem(a, b, c string) string {
	for _, aa := range a {
		for _, bb := range b {
			for _, cc := range c {
				if aa == bb && bb == cc {
					return string(aa)
				}
			}
		}
	}
	return ""
}

func priority(item string) int {
	if item[0]&32 == 0 {
		return int(item[0] - 'A' + 27)
	}
	return int(item[0] - 'a' + 1)
}
