package main

import "testing"

func TestFind3x3Square(t *testing.T) {
	for n, tc := range []struct {
		gridSerialNumber, x, y, totalPower int
	}{
		{18, 33, 45, 29},
		{42, 21, 61, 30},
	} {
		grid := makeGrid(tc.gridSerialNumber)
		x, y, totalPower := find3x3Square(grid)

		if x != tc.x || y != tc.y || totalPower != tc.totalPower {
			t.Errorf(
				"[%d] find3x3Square(%d) = %d, %d, %d; want %d, %d, %d",
				n, tc.gridSerialNumber, x, y, totalPower, tc.x, tc.y, tc.totalPower,
			)
		}
	}
}

func TestFindNxNSquare(t *testing.T) {
	for n, tc := range []struct {
		gridSerialNumber, x, y, side, totalPower int
	}{
		{18, 90, 269, 16, 113},
		{42, 232, 251, 12, 119},
	} {
		grid := makeGrid(tc.gridSerialNumber)
		x, y, side, totalPower := findNxNSquare(grid)

		if x != tc.x || y != tc.y || side != tc.side || totalPower != tc.totalPower {
			t.Errorf(
				"[%d] findNxNSquare(%d) = %d, %d, %d, %d; want %d, %d, %d, %d",
				n, tc.gridSerialNumber, x, y, side, totalPower, tc.x, tc.y, tc.side, tc.totalPower,
			)
		}
	}
}

func TestPowerLevel(t *testing.T) {
	for n, tc := range []struct {
		gridSerialNumber, x, y, powerLevel int
	}{
		{8, 3, 5, 4},
		{57, 122, 79, -5},
		{39, 217, 196, 0},
		{71, 101, 153, 4},
	} {
		powerLevel := powerLevel(tc.gridSerialNumber, tc.x, tc.y)

		if powerLevel != tc.powerLevel {
			t.Errorf(
				"[%d] powerLevel(%d, %d, %d) = %d; want %d",
				n, tc.gridSerialNumber, tc.x, tc.y, powerLevel, tc.powerLevel,
			)
		}
	}
}

func TestSummedGrid(t *testing.T) {
	var grid [300][300]int

	for y := 0; y < 300; y++ {
		for x := 0; x < 300; x++ {
			grid[y][x] = 1
		}
	}

	sGrid := summedGrid(grid)

	check := func(x, y, sum int) {
		if sGrid[y][x] != sum {
			t.Errorf("%d,%d = %d, want %d", y, x, sGrid[y][x], sum)
		}
	}

	/*
	 1  1  1      1  2  3
	 1  1  1  ->  2  4  6
	 1  1  1      3  6  9
	*/

	check(0, 0, 1)
	check(0, 1, 2)
	check(0, 2, 3)
	check(1, 0, 2)
	check(1, 1, 4)
	check(1, 2, 6)
	check(2, 0, 3)
	check(2, 1, 6)
	check(2, 2, 9)
}
