package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", calibrate(0, input))
	fmt.Printf("Part 2: %d\n", calibrateToFirstSeenTwice(0, input))
}

func calibrate(val int, calibrations []string) int {
	for _, c := range calibrations {
		chg, _ := strconv.Atoi(c)
		val += chg
	}

	return val
}

func calibrateToFirstSeenTwice(val int, calibrations []string) int {
	seen := map[int]bool{}

	for n := 0; !seen[val]; n = (n + 1) % len(calibrations) {
		seen[val] = true
		chg, _ := strconv.Atoi(calibrations[n])
		val += chg
	}

	return val
}
