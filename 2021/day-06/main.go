package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var input string

	fmt.Scanln(&input)

	fish := parse(input)
	spawn(fish, 80)
	fmt.Printf("Part 1: %d\n", countFish(fish))

	fish = parse(input)
	spawn(fish, 256)
	fmt.Printf("Part 2: %d\n", countFish(fish))
}

func spawn(fishFreq []int, days int) {
	for day := 0; day < days; day++ {
		zeros := fishFreq[0]
		copy(fishFreq, fishFreq[1:])
		fishFreq[6] += zeros
		fishFreq[8] = zeros
	}
}

func countFish(fishFreq []int) int {
	var sum int
	for _, f := range fishFreq {
		sum += f
	}
	return sum
}

func parse(input string) (fishFreq []int) {
	strs := strings.Split(input, ",")
	f := make([]int, 9)
	for _, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		f[n]++
	}
	return f
}
