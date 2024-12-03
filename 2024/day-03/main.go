package main

import (
	"bufio"
	"cmp"
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
	fmt.Printf("Part 2: %d\n", part2(input))
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

func part2(input []string) int {
	var sum int

	allRe := regexp.MustCompile(`(do)\(\)|(don't)\(\)|(mul)\((\d{1,3}),(\d{1,3})\)`)

	mulEnabled := true

	for n := range input {
		for _, match := range allRe.FindAllStringSubmatch(input[n], -1) {
			switch cmp.Or(match[1], match[2], match[3]) {
			case "do":
				mulEnabled = true
			case "don't":
				mulEnabled = false
			case "mul":
				if mulEnabled {
					a, _ := strconv.Atoi(match[4])
					b, _ := strconv.Atoi(match[5])

					sum += a * b
				}
			}
		}
	}

	return sum
}
