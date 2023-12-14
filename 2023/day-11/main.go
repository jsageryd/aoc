package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	var input []string

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		input = append(input, s.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	galaxies := parse(input)

	expand(galaxies)

	var indices []int
	for n := range galaxies {
		indices = append(indices, n)
	}

	var sum int

	combinations(indices, 2, func(comb []int) bool {
		sum += manhattanDistance(galaxies[comb[0]], galaxies[comb[1]])
		return true
	})

	return sum
}

type coord struct {
	x, y int
}

func parse(input []string) (galaxies []coord) {
	for y := range input {
		for x := range input[y] {
			if input[y][x] == '#' {
				galaxies = append(galaxies, coord{x, y})
			}
		}
	}

	return galaxies
}

func bounds(coords []coord) (topLeft, bottomRight coord) {
	topLeft = coords[0]
	bottomRight = coords[0]

	for _, g := range coords {
		topLeft.x = min(g.x, topLeft.x)
		topLeft.y = min(g.y, topLeft.y)
		bottomRight.x = max(g.x, bottomRight.x)
		bottomRight.y = max(g.y, bottomRight.y)
	}

	return topLeft, bottomRight
}

func expand(galaxies []coord) {
	topLeft, bottomRight := bounds(galaxies)

	var emptyRows, emptyCols []int

	for y := topLeft.y; y <= bottomRight.y; y++ {
		rowEmpty := true

		for x := topLeft.x; x <= bottomRight.x; x++ {
			if slices.Contains(galaxies, coord{x, y}) {
				rowEmpty = false
				break
			}
		}

		if rowEmpty {
			emptyRows = append(emptyRows, y)
		}
	}

	for x := topLeft.x; x <= bottomRight.x; x++ {
		colEmpty := true

		for y := topLeft.y; y <= bottomRight.y; y++ {
			if slices.Contains(galaxies, coord{x, y}) {
				colEmpty = false
				break
			}
		}

		if colEmpty {
			emptyCols = append(emptyCols, x)
		}
	}

	slices.Reverse(emptyRows)
	slices.Reverse(emptyCols)

	for n := range galaxies {
		for _, y := range emptyRows {
			if galaxies[n].y > y {
				galaxies[n].y++
			}
		}
		for _, x := range emptyCols {
			if galaxies[n].x > x {
				galaxies[n].x++
			}
		}
	}
}

// combinations picks k elements from s and calls f for each combination, until
// there are no more combinations or the given function returns false.
func combinations(s []int, k int, f func(comb []int) bool) {
	cont := true

	comb := make([]int, k)

	var rec func(ss []int, cc []int)

	rec = func(ss []int, cc []int) {
		for n := 0; n <= len(ss)-len(cc) && cont; n++ {
			cc[0] = ss[n]
			if len(cc) > 1 {
				rec(ss[n+1:], cc[1:])
			} else {
				cont = cont && f(comb)
			}
		}
	}

	rec(s, comb)

	/*
		Algorithm example

		abcdef pick 3
		^^^
		^^ ^
		^^  ^
		^^   ^
		^ ^^
		^ ^ ^
		^ ^  ^
		^  ^^
		^  ^ ^
		^   ^^
		 ^^^
		 ^^ ^
		 ^^  ^
		 ^ ^^
		 ^ ^ ^
		 ^  ^^
		  ^^^
		  ^^ ^
		  ^ ^^
		   ^^^
	*/
}

// manhattanDistance returns the Manhattan distance between the given
// coordinates.
//
// Manhattan distances from center point:
//
//	43234
//	32123
//	21012
//	32123
//	43234
//
// https://en.wikipedia.org/wiki/Manhattan_distance
func manhattanDistance(a, b coord) int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	return abs(a.x-b.x) + abs(a.y-b.y)
}

func printUniverse(galaxies []coord) {
	topLeft, bottomRight := bounds(galaxies)

	for y := topLeft.y; y <= bottomRight.y; y++ {
		for x := topLeft.x; x <= bottomRight.x; x++ {
			if slices.Contains(galaxies, coord{x, y}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
