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

	fish := spawn(parse(input), 80)
	fmt.Printf("Part 1: %d\n", len(fish))
}

func spawn(fish []int, days int) []int {
	for day := 0; day < days; day++ {
		var newFishCount int
		for i := range fish {
			if fish[i] == 0 {
				newFishCount++
				fish[i] = 6
			} else {
				fish[i]--
			}
		}
		for i := 0; i < newFishCount; i++ {
			fish = append(fish, 8)
		}
	}
	return fish
}

func parse(input string) []int {
	strs := strings.Split(input, ",")
	ints := make([]int, 0, len(strs))
	for _, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, n)
	}
	return ints
}
