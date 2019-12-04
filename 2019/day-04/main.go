package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string

	fmt.Scan(&input)

	minmax := strings.Split(input, "-")
	min, _ := strconv.Atoi(minmax[0])
	max, _ := strconv.Atoi(minmax[1])

	fmt.Printf("Part 1: %d\n", numberOfValidPasswords(min, max))
}

func numberOfValidPasswords(min, max int) int {
	var count int

	for pass := min; pass <= max; pass++ {
		if validPassword(pass) {
			count++
		}
	}

	return count
}

func validPassword(pass int) bool {
	var digits int
	var double bool

	for last := pass%10 + 1; pass > 0; pass /= 10 {
		digits++
		cur := pass % 10
		if cur > last {
			return false
		}
		if cur == last {
			double = true
		}
		last = cur
	}

	return double && digits == 6
}
