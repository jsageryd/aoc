package main

import (
	"bufio"
	"bytes"
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
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	g := parse(input)

	var totalFlashes int

	for i := 0; i < 100; i++ {
		step(g)
		totalFlashes += g.flashCount()
	}

	return totalFlashes
}

func part2(input []string) int {
	g := parse(input)

	var stepCount int

	for g.flashCount() != len(g) {
		stepCount++
		step(g)
	}

	return stepCount
}

func step(g grid) {
	flash := make(map[coord]struct{})

	for c, o := range g {
		o.flash = false
		o.energy++
		if o.energy > 9 {
			flash[c] = struct{}{}
		}
	}

	for len(flash) > 0 {
		var c coord
		for c = range flash {
			break
		}
		delete(flash, c)
		o := g[c]
		o.flash = true
		o.energy = 0
		for _, adj := range c.adj() {
			if o, ok := g[adj]; ok && !o.flash {
				o.energy++
				if o.energy > 9 {
					flash[adj] = struct{}{}
				}
			}
		}
	}
}

func parse(input []string) grid {
	g := make(grid)
	for y := range input {
		for x := range input[y] {
			n, err := strconv.Atoi(string(input[y][x]))
			if err != nil {
				log.Fatal(err)
			}
			g[coord{x, y}] = &octopus{energy: n}
		}
	}
	return g
}

type grid map[coord]*octopus

func (g grid) String() string {
	var maxX, maxY int
	for c := range g {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}

	rows := make([][]byte, maxY+1)
	for y := 0; y <= maxY; y++ {
		rows[y] = make([]byte, maxX+1)
		for x := 0; x <= maxX; x++ {
			v := strconv.Itoa(g[coord{x, y}].energy)
			rows[y][x] = v[len(v)-1]
		}
	}

	return string(bytes.Join(rows, []byte{'\n'}))
}

func (g grid) flashCount() int {
	var count int
	for _, o := range g {
		if o.flash {
			count++
		}
	}
	return count
}

type coord struct {
	x, y int
}

func (c coord) adj() []coord {
	return []coord{
		{c.x, c.y - 1},
		{c.x - 1, c.y - 1},
		{c.x - 1, c.y},
		{c.x - 1, c.y + 1},
		{c.x, c.y + 1},
		{c.x + 1, c.y + 1},
		{c.x + 1, c.y},
		{c.x + 1, c.y - 1},
	}
}

type octopus struct {
	energy int
	flash  bool
}
