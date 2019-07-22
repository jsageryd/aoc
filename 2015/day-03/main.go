package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	fmt.Printf("Part 1: %d\n", len(findHouses(input)))
}

type coord struct {
	x, y int
}

func findHouses(input []byte) map[coord]int {
	var c coord

	m := map[coord]int{c: 1}

	for _, b := range input {
		switch b {
		case '>':
			c.x++
		case '<':
			c.x--
		case '^':
			c.y--
		case 'v':
			c.y++
		}
		m[c]++
	}

	return m
}
