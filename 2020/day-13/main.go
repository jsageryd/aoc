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
	offsets := make(map[int]int) // bus ID -> offset
	for offset, idStr := range schedule {
		if id, err := strconv.Atoi(idStr); err == nil {
			offsets[id] = offset
		}
	}

	freq := make(map[int]int) // timestamp -> number of times seen
	for multiplier := 1; ; multiplier++ {
		for id, offset := range offsets {
			t := id*multiplier - offset
			freq[t]++
			if freq[t] == len(offsets) {
				return t
			}
		}
	}
}
