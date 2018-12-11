package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var input int
	fmt.Scanf("%d", &input)

	grid := makeGrid(input)

	var x, y, side int

	x, y, _ = find3x3Square(grid)
	fmt.Printf("Part 1: %d,%d\n", x, y)

	x, y, side, _ = findNxNSquare(grid)
	fmt.Printf("Part 2: %d,%d,%d\n", x, y, side)
}

func find3x3Square(grid [300][300]int) (x, y, totalPower int) {
	for yy := 1; yy <= 300-2; yy++ {
		for xx := 1; xx <= 300-2; xx++ {
			sum := 0
			for oy := 0; oy < 3; oy++ {
				for ox := 0; ox < 3; ox++ {
					sum += grid[yy+oy-1][xx+ox-1]
				}
			}
			if sum > totalPower {
				x, y, totalPower = xx, yy, sum
			}
		}
	}
	return x, y, totalPower
}

func findNxNSquare(grid [300][300]int) (x, y, side, totalPower int) {
	workers := runtime.NumCPU()

	in := make(chan int, 300)
	for squareSide := 1; squareSide <= 300; squareSide++ {
		in <- squareSide
	}
	close(in)

	out := make(chan []int, workers)

	var wg sync.WaitGroup
	wg.Add(workers)

	for n := 0; n < workers; n++ {
		go func() {
			var x, y, side, totalPower int
			for squareSide := range in {
				for yy := 1; yy <= 300-squareSide+1; yy++ {
					for xx := 1; xx <= 300-squareSide+1; xx++ {
						sum := 0
						for oy := 0; oy < squareSide; oy++ {
							for ox := 0; ox < squareSide; ox++ {
								sum += grid[yy+oy-1][xx+ox-1]
							}
						}
						if sum > totalPower {
							x, y, side, totalPower = xx, yy, squareSide, sum
						}
					}
				}
			}
			out <- []int{x, y, side, totalPower}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for res := range out {
		if res[3] > totalPower {
			x, y, side, totalPower = res[0], res[1], res[2], res[3]
		}
	}

	return x, y, side, totalPower
}

func makeGrid(serialNumber int) [300][300]int {
	var grid [300][300]int
	for y := 0; y < 300; y++ {
		for x := 0; x < 300; x++ {
			grid[y][x] = powerLevel(serialNumber, x+1, y+1)
		}
	}
	return grid
}

func powerLevel(gridSerialNumber, x, y int) int {
	rackID := x + 10
	powerLevel := rackID * y
	powerLevel += gridSerialNumber
	powerLevel *= rackID
	powerLevel = powerLevel / 100 % 10
	powerLevel -= 5
	return powerLevel
}
