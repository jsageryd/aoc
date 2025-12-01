package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	var count int

	dial := 50

	for _, v := range parse(input) {
		for v < 0 {
			v += 100
		}

		dial = (dial + v) % 100

		if dial == 0 {
			count++
		}
	}

	return count
}

func part2(input []string) int {
	var count int

	dial := 50

	for _, v := range parse(input) {
		left := v < 0

		if v < 0 {
			v = -v
		}

		for range v {
			c := 1

			if left {
				c = -1
			}

			for c < 0 {
				c += 100
			}

			dial = (dial + c) % 100

			if dial == 0 {
				count++
			}
		}
	}

	return count
}

func parse(input []string) []int {
	var vs []int

	for _, line := range input {
		v, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		if line[0] == 'L' {
			v = -v
		}

		vs = append(vs, v)
	}

	return vs
}
