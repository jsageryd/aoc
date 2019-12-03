package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	wire1 := drawWire(input[0])
	wire2 := drawWire(input[1])

	closestByManhattan := closestIntersection(wire1, wire2)
	closestByWire := closestIntersectionByWire(wire1, wire2)

	fmt.Printf("Part 1: %d\n", distance(coord{}, closestByManhattan))
	fmt.Printf("Part 2: %d\n", idxOf(closestByWire, wire1)+idxOf(closestByWire, wire2))
}

type coord struct {
	x, y int
}

func drawWire(path string) []coord {
	dist := func(inst string) int {
		d, _ := strconv.Atoi(inst[1:])
		return d
	}

	var cur coord

	wire := []coord{cur}

	deltas := map[byte]coord{
		'U': {0, -1},
		'D': {0, 1},
		'L': {-1, 0},
		'R': {1, 0},
	}

	for _, inst := range strings.Split(path, ",") {
		d := deltas[inst[0]]
		for n := 0; n < dist(inst); n++ {
			cur.x += d.x
			cur.y += d.y
			wire = append(wire, cur)
		}
	}

	return wire
}

func intersections(w1, w2 []coord) []coord {
	var cs []coord

	m1 := make(map[coord]struct{})
	m2 := make(map[coord]struct{})

	for _, c := range w1 {
		m1[c] = struct{}{}
	}

	for _, c := range w2 {
		m2[c] = struct{}{}
	}

	for c1 := range m1 {
		if c1 == (coord{}) {
			continue
		}
		if _, ok := m2[c1]; ok {
			cs = append(cs, c1)
		}
	}

	sort.Slice(cs, func(i, j int) bool {
		return fmt.Sprint(cs[i]) < fmt.Sprint(cs[j])
	})

	return cs
}

func distance(c1, c2 coord) int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	return abs(c2.x-c1.x) + abs(c2.y-c1.y)
}

func closestIntersection(w1, w2 []coord) coord {
	var closest coord

	for _, c := range intersections(w1, w2) {
		if closest == (coord{}) || distance(coord{}, c) < distance(coord{}, closest) {
			closest = c
		}
	}

	return closest
}

func closestIntersectionByWire(w1, w2 []coord) coord {
	var closest coord
	shortestWires := -1

	for _, c := range intersections(w1, w2) {
		d1 := idxOf(c, w1)
		d2 := idxOf(c, w2)

		if shortestWires == -1 || d1+d2 < shortestWires {
			closest = c
			shortestWires = d1 + d2
		}
	}

	return closest
}

func idxOf(c coord, wire []coord) int {
	for idx := range wire {
		if wire[idx] == c {
			return idx
		}
	}
	return -1
}
