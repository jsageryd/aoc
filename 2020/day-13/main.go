package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	estimate, buses := parseInput(os.Stdin)

	busID, waitTime := earliestBus(estimate, buses)

	fmt.Printf("Part 1: %d\n", busID*waitTime)
}

func parseInput(r io.Reader) (estimate int, buses []int) {
	fmt.Fscanln(r, &estimate)
	var busesStr string
	fmt.Fscanln(r, &busesStr)
	for _, b := range strings.Split(busesStr, ",") {
		if n, err := strconv.Atoi(b); err == nil {
			buses = append(buses, n)
		}
	}
	return estimate, buses
}

func earliestBus(notBefore int, buses []int) (id int, waitTime int) {
	waitTime = notBefore
	for _, b := range buses {
		wt := b*(notBefore/b+1) - notBefore
		if wt < waitTime {
			id, waitTime = b, wt
		}
	}
	return id, waitTime
}
