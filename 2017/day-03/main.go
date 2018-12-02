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
	x, y := 0, 0
	dx, dy := 1, 0
	radius := 0
	n := 0

outer:
	for revolution := 0; ; revolution++ {
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
				n++
				if n == target {
					break outer
				}
				x += dx
				y += dy
			}
		}
	}

	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	return abs(x) + abs(y)
}

func spiralFirstLargerAdjacentSum(target int) int {
	x, y := 0, 0
	dx, dy := 1, 0
	radius := 0

	spiral := map[int]map[int]int{
		0: map[int]int{0: 1},
	}

	adjSum := func(x, y int) int {
		sum := 0
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				sum += spiral[y+dy][x+dx]
			}
		}
		return sum
	}

	for revolution := 0; ; revolution++ {
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
				sum := adjSum(x, y)

				if sum > target {
					return sum
				}

				if _, ok := spiral[y]; !ok {
					spiral[y] = make(map[int]int)
				}
				spiral[y][x] = sum

				x += dx
				y += dy
			}
		}
	}

	return 0
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
