package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	g := newGrid(1000, 1000)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		g = applyInstruction(scanner.Text(), g)
	}

	fmt.Printf("Part 1: %d\n", lightsLit(g))
}

func lightsLit(g grid) int {
	var count int

	for y := range g {
		for x := range g[y] {
			if g[y][x] {
				count++
			}
		}
	}

	return count
}

func newGrid(sizeX, sizeY int) grid {
	g := make(grid, sizeY)
	for n := range g {
		g[n] = make([]bool, sizeX)
	}

	return g
}

func applyInstruction(instStr string, g grid) grid {
	inst := parseInstruction(instStr)

	for y := inst.from.y; y <= inst.to.y; y++ {
		for x := inst.from.x; x <= inst.to.x; x++ {
			switch inst.action {
			case actionOn:
				g[y][x] = true
			case actionOff:
				g[y][x] = false
			case actionToggle:
				g[y][x] = !g[y][x]
			}
		}
	}

	return g
}

func parseInstruction(s string) instruction {
	fields := strings.Fields(s)

	var action int

	switch fields[len(fields)-4] {
	case "on":
		action = actionOn
	case "off":
		action = actionOff
	case "toggle":
		action = actionToggle
	}

	return instruction{
		action: action,
		to:     parseCoord(fields[len(fields)-1]),
		from:   parseCoord(fields[len(fields)-3]),
	}
}

func parseCoord(s string) coord {
	var c coord
	fmt.Sscanf(s, "%d,%d", &c.x, &c.y)
	return c
}

const (
	actionUnknown = iota
	actionOn
	actionOff
	actionToggle
)

type grid [][]bool

func (g grid) String() string {
	rows := make([]string, len(g))
	for y := range g {
		row := make([]rune, len(g[y]))
		for x := range g[y] {
			if g[y][x] {
				row[x] = 'o'
			} else {
				row[x] = '-'
			}
		}
		rows[y] = string(row)
	}
	return strings.Join(rows, "\n")
}

type coord struct {
	x, y int
}

type instruction struct {
	action int
	from   coord
	to     coord
}
