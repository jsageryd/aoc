package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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

	for n := range input {
		if safe(parseReport(input[n])) {
			sum++
		}
	}

	return sum
}

func part2(input []string) int {
	var sum int

	for n := range input {
		report := parseReport(input[n])

		if safe(report) {
			sum++
			continue
		}

		for n := range report {
			r := slices.Clone(report)
			r = append(r[:n], r[n+1:]...)

			if safe(r) {
				sum++
				break
			}
		}
	}

	return sum
}

func safe(report []int) bool {
	if len(report) < 2 {
		return false
	}

	increasing := report[0] < report[1]

	for n := 1; n < len(report); n++ {
		switch {
		case increasing && report[n-1] > report[n]:
			return false
		case !increasing && report[n-1] < report[n]:
			return false
		case abs(report[n-1]-report[n]) < 1:
			return false
		case abs(report[n-1]-report[n]) > 3:
			return false
		}
	}

	return true
}

func parseReport(line string) []int {
	var report []int

	levelStrs := strings.Fields(line)

	for n := range levelStrs {
		level, _ := strconv.Atoi(levelStrs[n])
		report = append(report, level)
	}

	return report
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
