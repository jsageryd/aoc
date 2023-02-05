package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var claims []claim

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		claims = append(claims, parseClaim(scanner.Text()))
	}

	fmt.Printf("Part 1: %d\n", overlappingArea(claims))
	fmt.Printf("Part 2: %d\n", nonOverlappingClaim(claims).ID)
}

func parseClaim(s string) claim {
	var id, x, y, w, h int
	fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
	return claim{ID: id, X: x, Y: y, W: w, H: h}
}

type claim struct {
	ID   int
	X, Y int
	W, H int
}

func overlappingArea(claims []claim) int {
	fabric := drawClaims(claims)

	overlap := 0

	for y := range fabric {
		for x := range fabric[y] {
			if fabric[y][x] == 'X' {
				overlap++
			}
		}
	}

	return overlap
}

func nonOverlappingClaim(claims []claim) claim {
	fabric := drawClaims(claims)

	var nonOverlappingClaim claim

next:
	for _, c := range claims {
		for y := c.Y; y < c.Y+c.H; y++ {
			for x := c.X; x < c.X+c.W; x++ {
				if fabric[y][x] == 'X' {
					continue next
				}
			}
		}
		nonOverlappingClaim = c
	}

	return nonOverlappingClaim
}

func drawClaims(claims []claim) [][]byte {
	var fabricW, fabricH int

	for _, c := range claims {
		if fabricW < c.X+c.W {
			fabricW = c.X + c.W
		}
		if fabricH < c.Y+c.H {
			fabricH = c.Y + c.H
		}
	}

	fabric := make([][]byte, fabricH)
	for y := range fabric {
		fabric[y] = make([]byte, fabricW)
	}

	for _, c := range claims {
		for y := c.Y; y < c.Y+c.H; y++ {
			for x := c.X; x < c.X+c.W; x++ {
				switch fabric[y][x] {
				case 0: // first to add fabric here, so add a #
					fabric[y][x] = '#'
				case '#', 'X': // # means there is fabric; X means there is overlapping fabric
					fabric[y][x] = 'X'
				}
			}
		}
	}

	return fabric
}
