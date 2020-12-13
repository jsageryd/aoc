package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
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

	workers := runtime.NumCPU()

	var wg sync.WaitGroup
	wg.Add(workers)

	in := make(chan int, 64)
	out := make(chan int, workers)
	done := make(chan struct{}, workers)

	sliceSize := 1_000_000

	go func() {
		for startAt := 0; ; startAt += sliceSize {
			select {
			case <-done:
				close(in)
				return
			case in <- startAt:
			}
		}
	}()

	for n := 0; n < workers; n++ {
		go func() {
			defer wg.Done()
			for startAt := range in {
			next:
				for t := startAt; t < startAt+sliceSize; t++ {
					for n, b := range scheduleInt {
						if b != -1 {
							if dep := t + n; dep%b != 0 {
								continue next
							}
						}
					}
					done <- struct{}{}
					out <- t
					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var ts []int
	for t := range out {
		ts = append(ts, t)
	}
	sort.Ints(ts)

	return ts[0]
}
