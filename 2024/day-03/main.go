package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var input []string

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		input = append(input, s.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	var sum int

	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	for n := range input {
		for _, match := range mulRe.FindAllStringSubmatch(input[n], -1) {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])

			sum += a * b
		}
	}

	return sum
}
