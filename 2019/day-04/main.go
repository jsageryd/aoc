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

	fmt.Printf("Part 1: %d\n", numberOfValidPasswords(min, max, false))
	fmt.Printf("Part 2: %d\n", numberOfValidPasswords(min, max, true))
}

func numberOfValidPasswords(min, max int, strictDouble bool) int {
	var count int

	for pass := min; pass <= max; pass++ {
		if validPassword(pass, strictDouble) {
			count++
		}
	}

	return count
}

func validPassword(pass int, strictDouble bool) bool {
	var digits int

	counts := make(map[int]int)

	for last := pass % 10; pass > 0; pass /= 10 {
		digits++
		cur := pass % 10
		counts[cur]++
		if cur > last {
			return false
		}
		last = cur
	}

	var double bool

	for _, v := range counts {
		if v >= 2 && (!strictDouble || v == 2) {
			double = true
			break
		}
	}

	return double && digits == 6
}
