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

func TestFindAnySizeSquare(t *testing.T) {
	for n, tc := range []struct {
		gridSerialNumber, x, y, side, totalPower int
	}{
		{18, 90, 269, 16, 113},
		{42, 232, 251, 12, 119},
	} {
		grid := makeGrid(tc.gridSerialNumber)
		x, y, side, totalPower := findAnySizeSquare(grid)

		if x != tc.x || y != tc.y || side != tc.side || totalPower != tc.totalPower {
			t.Errorf(
				"[%d] findAnySizeSquare(%d) = %d, %d, %d, %d; want %d, %d, %d, %d",
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
	var grid [301][301]int

	for y := 1; y < 300; y++ {
		for x := 1; x < 300; x++ {
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

	check(0, 0, 0)
	check(0, 1, 0)
	check(0, 2, 0)
	check(0, 3, 0)
	check(1, 1, 1)
	check(1, 2, 2)
	check(1, 3, 3)
	check(2, 1, 2)
	check(2, 2, 4)
	check(2, 3, 6)
	check(3, 1, 3)
	check(3, 2, 6)
	check(3, 3, 9)
}

func BenchmarkFind3x3Square(b *testing.B) {
	for n := 0; n < b.N; n++ {
		find3x3Square(makeGrid(1308))
	}
}

func BenchmarkFindAnySizeSquare(b *testing.B) {
	for n := 0; n < b.N; n++ {
		findAnySizeSquare(makeGrid(1308))
	}
}

func BenchmarkPowerLevel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for y := 1; y < 300; y++ {
			for x := 1; x < 300; x++ {
				powerLevel(1308, x, y)
			}
		}
	}
}
