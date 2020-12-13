package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	estimate, schedule := parseInput(os.Stdin)

	busID, waitTime := earliestBus(estimate, schedule)

	fmt.Printf("Part 1: %d\n", busID*waitTime)
	fmt.Printf("Part 2: %d\n", earliestTimestampForSubsequentDepartures(schedule))
}

func parseInput(r io.Reader) (estimate int, schedule []string) {
	var scheduleStr string
	fmt.Fscanln(r, &estimate)
	fmt.Fscanln(r, &scheduleStr)
	return estimate, strings.Split(scheduleStr, ",")
}

func earliestBus(notBefore int, schedule []string) (id int, waitTime int) {
	waitTime = notBefore
	for _, bStr := range schedule {
		if b, err := strconv.Atoi(bStr); err == nil {
			wt := b*(notBefore/b+1) - notBefore
			if wt < waitTime {
				id, waitTime = b, wt
			}
		}
	}
	return id, waitTime
}

func earliestTimestampForSubsequentDepartures(schedule []string) int {
	var scheduleInt []int
	for _, bStr := range schedule {
		b, err := strconv.Atoi(bStr)
		if err != nil {
			b = -1
		}
		scheduleInt = append(scheduleInt, b)
	}

next:
	for t := 0; ; t++ {
		for n, b := range scheduleInt {
			if b != -1 {
				if dep := t + n; dep%b != 0 {
					continue next
				}
			}
		}
		return t
	}
}
