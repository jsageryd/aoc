package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
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

func part2(input []string) int {
	if len(input) == 0 {
		return 0
	}

	heightMap := parse(input)

	basinSizes := findBasinSizes(heightMap)

	if len(basinSizes) < 3 {
		log.Fatalf("found only %d basins", len(basinSizes))
	}

	sort.Ints(basinSizes)

	prod := 1

	for _, size := range basinSizes[len(basinSizes)-3:] {
		prod *= size
	}

	return prod
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

func findLowPoints(heightMap map[coord]int) []coord {
	var lowPoints []coord

next:
	for c := range heightMap {
		height := heightMap[c]

		for _, adjCoord := range []coord{
			{c.x, c.y - 1},
			{c.x - 1, c.y},
			{c.x, c.y + 1},
			{c.x + 1, c.y},
		} {
			if adjHeight, ok := heightMap[adjCoord]; ok {
				if adjHeight <= height {
					continue next
				}
			}
		}

		lowPoints = append(lowPoints, c)
	}

	return lowPoints
}

func findBasinSizes(heightMap map[coord]int) []int {
	var basinSizes []int

	var maxX, maxY int
	for c := range heightMap {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}

	for _, c := range findLowPoints(heightMap) {
		seen := make(map[coord]bool)
		var floodFind func(c coord)
		floodFind = func(c coord) {
			if c.x < 0 || c.x > maxX || c.y < 0 || c.y > maxY ||
				seen[c] || heightMap[c] == 9 {
				return
			}
			seen[c] = true
			for _, adjC := range []coord{
				{c.x, c.y - 1},
				{c.x - 1, c.y},
				{c.x, c.y + 1},
				{c.x + 1, c.y},
			} {
				floodFind(adjC)
			}
		}
		floodFind(c)
		basinSizes = append(basinSizes, len(seen))
	}

	return basinSizes
}

type coord struct {
	x, y int
}
