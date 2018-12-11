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
