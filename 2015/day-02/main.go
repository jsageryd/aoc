package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var paper, ribbon int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		paper += paperNeeded(scanner.Text())
		ribbon += ribbonNeeded(scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", paper)
	fmt.Printf("Part 2: %d\n", ribbon)
}

func paperNeeded(box string) int {
	s := make([]int, 3)
	fmt.Sscanf(box, "%dx%dx%d", &s[0], &s[1], &s[2])
	sort.Ints(s)
	return (2*s[0]*s[1] + 2*s[1]*s[2] + 2*s[0]*s[2]) + (s[0] * s[1])
}

func ribbonNeeded(box string) int {
	s := make([]int, 3)
	fmt.Sscanf(box, "%dx%dx%d", &s[0], &s[1], &s[2])
	sort.Ints(s)
	return (2*s[0] + 2*s[1]) + (s[0] * s[1] * s[2])
}
