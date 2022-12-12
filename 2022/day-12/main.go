package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"runtime"
	"sync"
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
	coords := make(map[coord]byte)

	var start, goal coord

	for y := range input {
		for x := range input[y] {
			elevation := input[y][x]
			switch input[y][x] {
			case 'S':
				start = coord{x, y}
				elevation = 'a'
			case 'E':
				goal = coord{x, y}
				elevation = 'z'
			}
			coords[coord{x, y}] = elevation
		}
	}

	neighbours := func(c coord) []coord {
		var ns []coord

		for _, n := range []coord{
			{c.x - 1, c.y},
			{c.x + 1, c.y},
			{c.x, c.y - 1},
			{c.x, c.y + 1},
		} {
			if neighbour, ok := coords[n]; ok {
				if neighbour <= coords[c]+1 {
					ns = append(ns, n)
				}
			}
		}

		return ns
	}

	cost := func(a, b coord) int {
		return 1
	}

	heuristic := func(c coord) int {
		return manhattanDistance(c, goal)
	}

	path, found := aStar(start, goal, neighbours, cost, heuristic)
	if !found {
		return 0
	}

	return len(path) - 1
}

func part2(input []string) int {
	coords := make(map[coord]byte)

	var starts []coord
	var goal coord

	for y := range input {
		for x := range input[y] {
			elevation := input[y][x]
			switch input[y][x] {
			case 'a', 'S':
				starts = append(starts, coord{x, y})
				elevation = 'a'
			case 'E':
				goal = coord{x, y}
				elevation = 'z'
			}
			coords[coord{x, y}] = elevation
		}
	}

	neighbours := func(c coord) []coord {
		var ns []coord

		for _, n := range []coord{
			{c.x - 1, c.y},
			{c.x + 1, c.y},
			{c.x, c.y - 1},
			{c.x, c.y + 1},
		} {
			if neighbour, ok := coords[n]; ok {
				if neighbour <= coords[c]+1 {
					ns = append(ns, n)
				}
			}
		}

		return ns
	}

	cost := func(a, b coord) int {
		return 1
	}

	heuristic := func(c coord) int {
		return manhattanDistance(c, goal)
	}

	workers := runtime.NumCPU()

	var wg sync.WaitGroup
	wg.Add(workers)

	in := make(chan coord, len(starts))
	for _, start := range starts {
		in <- start
	}
	close(in)

	out := make(chan []coord)

	for n := 0; n < workers; n++ {
		go func() {
			for start := range in {
				path, found := aStar(start, goal, neighbours, cost, heuristic)
				if found {
					out <- path
				}
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	shortestPath := <-out

	for path := range out {
		if len(path) < len(shortestPath) {
			shortestPath = path
		}
	}

	return len(shortestPath) - 1
}

type coord struct {
	x, y int
}

// aStar returns the shortest path between start and goal.
//
// neighbours returns the neighbouring coordinates of c.
//
// cost returns the cost of moving from coordinate a to coordinate b, where a
// and b are neighbours.
//
// heuristic returns the cost of the estimated shortest possible path between
// the current coordinate c and the goal. The cost must not be overestimated.
//
// https://en.wikipedia.org/wiki/A*_search_algorithm
func aStar(
	start, goal coord,
	neighbours func(c coord) []coord,
	cost func(a, b coord) int,
	heuristic func(c coord) (cost int),
) (path []coord, found bool) {
	prev := make(map[coord]coord)
	accCost := map[coord]int{start: 0}

	nexts := &pqCoords{
		coords: []coord{start},
		less: func(a, b coord) bool {
			return accCost[a]+heuristic(a) < accCost[b]+heuristic(b)
		},
	}

	for nexts.Len() > 0 {
		cur := heap.Pop(nexts).(coord)

		if cur == goal {
			found = true
			break
		}

		for _, neighbour := range neighbours(cur) {
			neighbourCost := accCost[cur] + cost(cur, neighbour)
			_, ok := accCost[neighbour]
			if !ok || neighbourCost < accCost[neighbour] {
				accCost[neighbour] = neighbourCost
				prev[neighbour] = cur
				heap.Push(nexts, neighbour)
			}
		}
	}

	if !found {
		return nil, false
	}

	for cur, ok := goal, true; ok; cur, ok = prev[cur] {
		path = append(path, cur)
	}

	// reverse
	for n := 0; n < len(path)/2; n++ {
		path[n], path[len(path)-n-1] = path[len(path)-n-1], path[n]
	}

	return path, true
}

type pqCoords struct {
	coords []coord
	less   func(a, b coord) bool
}

func (q *pqCoords) Len() int {
	return len(q.coords)
}

func (q *pqCoords) Less(i, j int) bool {
	return q.less(q.coords[i], q.coords[j])
}

func (q *pqCoords) Swap(i, j int) {
	q.coords[i], q.coords[j] = q.coords[j], q.coords[i]
}

func (q *pqCoords) Push(x interface{}) {
	q.coords = append(q.coords, x.(coord))
}

func (q *pqCoords) Pop() interface{} {
	c := q.coords[len(q.coords)-1]
	q.coords = q.coords[:len(q.coords)-1]
	return c
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
