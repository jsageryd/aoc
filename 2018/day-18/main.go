package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strings"
)

func main() {
	var area Area

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		area = append(area, scanner.Bytes())
	}

	areaAfter10Steps := step(area, 10)
	_, trees, lumberyards := areaAfter10Steps.Count()
	fmt.Printf("Part 1: %d\n", trees*lumberyards)

	areaAfter1000000000Steps := step(area, 1000000000)
	_, trees, lumberyards = areaAfter1000000000Steps.Count()
	fmt.Printf("Part 2: %d\n", trees*lumberyards)
}

type Area [][]byte

func (a Area) String() string {
	var ss []string
	for y := range a {
		ss = append(ss, string(a[y]))
	}
	return strings.Join(ss, "\n")
}

func (a Area) Count() (open, trees, lumberyards int) {
	for y := range a {
		for x := range a[y] {
			switch a[y][x] {
			case '.':
				open++
			case '|':
				trees++
			case '#':
				lumberyards++
			}
		}
	}
	return open, trees, lumberyards
}

func (a Area) Hash() []byte {
	h := md5.New()
	for y := range a {
		h.Write(a[y])
	}
	return h.Sum(nil)
}

func step(a Area, steps int) Area {
	cur := make(Area, len(a))
	next := make(Area, len(a))
	for y := range a {
		cur[y] = make([]byte, len(a[y]))
		copy(cur[y], a[y])
		next[y] = make([]byte, len(a[y]))
		copy(next[y], a[y])
	}

	buf := []Area{cur, next}

	seen := map[string]int{string(cur.Hash()): 0}
	skippedAhead := false

	for n := 0; n < steps; n++ {
		cur = buf[n%2]
		next = buf[(n+1)%2]

		for y := range cur {
			for x := range cur[y] {
				_, trees, lumberyards := adjacent(cur, x, y)
				switch cur[y][x] {
				case '.':
					if trees >= 3 {
						next[y][x] = '|'
					} else {
						next[y][x] = '.'
					}
				case '|':
					if lumberyards >= 3 {
						next[y][x] = '#'
					} else {
						next[y][x] = '|'
					}
				case '#':
					if lumberyards >= 1 && trees >= 1 {
						next[y][x] = '#'
					} else {
						next[y][x] = '.'
					}
				}
			}
		}

		if !skippedAhead {
			nh := string(next.Hash())
			if nn, ok := seen[nh]; ok {
				odd := false
				if n%2 != 0 {
					odd = true
				}
				cycle := n - nn
				n += cycle * ((steps - n) / cycle)
				if odd && n%2 == 0 {
					n--
				}
				skippedAhead = true
			} else {
				seen[nh] = n
			}
		}
	}

	return next
}

func adjacent(a Area, x, y int) (open, trees, lumberyards int) {
	open = 8
	for yy := y - 1; yy <= y+1; yy++ {
		if yy < 0 || yy > len(a)-1 {
			continue
		}
		for xx := x - 1; xx <= x+1; xx++ {
			if xx < 0 || xx > len(a[yy])-1 {
				continue
			}
			if xx == x && yy == y {
				continue
			}
			switch a[yy][xx] {
			case '|':
				open--
				trees++
			case '#':
				open--
				lumberyards++
			}
		}
	}
	return open, trees, lumberyards
}
