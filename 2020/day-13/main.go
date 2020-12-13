package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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
	var busIDs []int
	for id := range offsets {
		busIDs = append(busIDs, id)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(busIDs)))

	var count int

	var rec func(startT, idIdx int) int

	rec = func(startT, idIdx int) int {
		log.Printf("recursing to idx %d/%d, currently at timestamp %d", idIdx, len(busIDs), startT)
		id := busIDs[idIdx]
	next:
		for t := startT + offsets[id]; ; t += id {
			for i := 0; i < idIdx; i++ {
				count++
				if (t-offsets[id]+offsets[busIDs[i]])%busIDs[i] != 0 {
					continue next
				}
			}
			if idIdx == len(busIDs)-1 {
				return t - offsets[id]
			} else {
				return rec(t-offsets[id], idIdx+1)
			}
		}
	}

	return rec(0, 0)
}
