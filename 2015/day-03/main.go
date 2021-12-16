package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	fmt.Printf("Part 1: %d\n", len(findHouses(input, 1)))
	fmt.Printf("Part 2: %d\n", len(findHouses(input, 2)))
}

type coord struct {
	x, y int
}

func findHouses(input []byte, santas int) map[coord]int {
	m := map[coord]int{{0, 0}: santas}

	var currentSanta int
	cs := make([]coord, santas)

	for _, b := range input {
		currentSanta = (currentSanta + 1) % santas
		switch b {
		case '>':
			cs[currentSanta].x++
		case '<':
			cs[currentSanta].x--
		case '^':
			cs[currentSanta].y--
		case 'v':
			cs[currentSanta].y++
		}
		m[cs[currentSanta]]++
	}

	return m
}
