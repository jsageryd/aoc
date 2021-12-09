package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	if len(input) == 0 {
		return 0
	}

	heightMap := parse(input)

	var riskLevelSum int

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			riskLevelSum += riskLevel(x, y, heightMap)
		}
	}

	return riskLevelSum
}

func parse(input []string) map[coord]int {
	heightMap := make(map[coord]int)
	for y := range input {
		for x := range input[y] {
			n, err := strconv.Atoi(string(input[y][x]))
			if err != nil {
				log.Fatal(err)
			}
			heightMap[coord{x, y}] = n
		}
	}
	return heightMap
}

func riskLevel(x, y int, heightMap map[coord]int) int {
	height := heightMap[coord{x, y}]

	for _, adjCoord := range []coord{
		{x, y - 1},
		{x - 1, y},
		{x, y + 1},
		{x + 1, y},
	} {
		if adjHeight, ok := heightMap[adjCoord]; ok {
			if adjHeight <= height {
				return 0
			}
		}
	}

	return height + 1
}

type coord struct {
	x, y int
}
