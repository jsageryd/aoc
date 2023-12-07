package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input []string

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		input = append(input, s.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	timesStr := strings.Fields(input[0])[1:]
	distancesStr := strings.Fields(input[1])[1:]

	multipliedWaysToWin := 1

	for n := range timesStr {
		raceTime, _ := strconv.Atoi(timesStr[n])
		recordDistance, _ := strconv.Atoi(distancesStr[n])

		var waysToWin int

		for speed := 1; speed < raceTime; speed++ {
			remainingTime := raceTime - speed
			distance := speed * remainingTime

			if distance > recordDistance {
				waysToWin++
			}
		}

		multipliedWaysToWin *= waysToWin
	}

	return multipliedWaysToWin
}
