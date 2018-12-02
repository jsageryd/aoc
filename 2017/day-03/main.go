package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Scanln(&input)

	fmt.Printf("Part 1: %d\n", spiralDistance(input))
	fmt.Printf("Part 2: %d\n", spiralFirstLargerAdjacentSum(input))
}

func spiralDistance(target int) int {
	n := 0
	distX, distY := 0, 0

	spiral(
		func(x, y int) bool {
			n++
			distX = x
			distY = y
			return n < target
		},
	)

	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	return abs(distX) + abs(distY)
}

func spiralFirstLargerAdjacentSum(target int) int {
	spiralCells := map[int]map[int]int{
		0: map[int]int{0: 1},
	}

	adjSum := func(x, y int) int {
		sum := 0
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				sum += spiralCells[y+dy][x+dx]
			}
		}
		return sum
	}

	var sum int

	spiral(
		func(x, y int) bool {
			sum = adjSum(x, y)

			if _, ok := spiralCells[y]; !ok {
				spiralCells[y] = make(map[int]int)
			}
			spiralCells[y][x] = sum

			return sum <= target
		},
	)

	return sum
}

// spiral walks in a spiral, calling stepFn with the current coordinates for
// each step. The spiral continues as long as stepFn returns true.
func spiral(stepFn func(x, y int) bool) {
	x, y := 0, 0
	dx, dy := 1, 0
	radius := 0

	for {
		for quadrant := 0; quadrant < 4; quadrant++ {
			switch quadrant % 4 {
			case 0:
				radius++
				dx, dy = 1, 0
			case 1:
				dx, dy = 0, -1
			case 2:
				radius++
				dx, dy = -1, 0
			case 3:
				dx, dy = 0, 1
			}

			for step := 0; step < radius; step++ {
				if !stepFn(x, y) {
					return
				}
				x += dx
				y += dy
			}
		}
	}
}

/*
dx  dy   x   y
--------------
 1   0   1   0  |
                |
 0  -1   1  -1  |
                |
-1   0   0  -1  |  one revolution
-1   0  -1  -1  |
                |
 0   1  -1   0  |
 0   1  -1   1  |
--------------
+1   0   0   1
+1   0   1   1
+1   0   2   1

 0  -1   2   0
 0  -1   2  -1
 0  -1   2  -2

-1   0   1  -2
-1   0   0  -2
-1   0  -1  -2
-1   0  -2  -2

 0   1  -2  -1
 0   1  -2   0
 0   1  -2   1
 0   1  -2   2
--------------
 1   0  -1   2
 1   0   0   2
 1   0   1   2
 1   0   2   2
 1   0   3   2
*/
