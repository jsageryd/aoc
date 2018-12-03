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
	fmt.Printf("Part 1: %d\n", nonOverlappingClaim(claims).ID)
}

func parseClaim(s string) claim {
	var id, x, y, w, h int
	fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
	return claim{ID: id, Rect: rect{X: x, Y: y, W: w, H: h}}
}

type claim struct {
	ID   int
	Rect rect
}

type rect struct {
	X, Y int
	W, H int
}

func overlappingArea(claims []claim) int {
	var fabricW, fabricH int

	for _, c := range claims {
		if fabricW < c.Rect.X+c.Rect.W {
			fabricW = c.Rect.X + c.Rect.W
		}
		if fabricH < c.Rect.Y+c.Rect.H {
			fabricH = c.Rect.Y + c.Rect.H
		}
	}

	fabric := make([][]byte, fabricH)

	for _, c := range claims {
		for y := c.Rect.Y; y < c.Rect.Y+c.Rect.H; y++ {
			if fabric[y] == nil {
				fabric[y] = make([]byte, fabricW)
			}

			for x := c.Rect.X; x < c.Rect.X+c.Rect.W; x++ {
				switch fabric[y][x] {
				case 0: // first to add fabric here, so add a #
					fabric[y][x] = '#'
				case '#', 'X': // # means there is fabric; X means there is overlapping fabric
					fabric[y][x] = 'X'
				}
			}
		}
	}

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
	var fabricW, fabricH int

	for _, c := range claims {
		if fabricW < c.Rect.X+c.Rect.W {
			fabricW = c.Rect.X + c.Rect.W
		}
		if fabricH < c.Rect.Y+c.Rect.H {
			fabricH = c.Rect.Y + c.Rect.H
		}
	}

	fabric := make([][]byte, fabricH)

	for _, c := range claims {
		for y := c.Rect.Y; y < c.Rect.Y+c.Rect.H; y++ {
			if fabric[y] == nil {
				fabric[y] = make([]byte, fabricW)
			}

			for x := c.Rect.X; x < c.Rect.X+c.Rect.W; x++ {
				switch fabric[y][x] {
				case 0: // first to add fabric here, so add a #
					fabric[y][x] = '#'
				case '#', 'X': // # means there is fabric; X means there is overlapping fabric
					fabric[y][x] = 'X'
				}
			}
		}
	}

	var nonOverlappingClaim claim

next:
	for _, c := range claims {
		for y := c.Rect.Y; y < c.Rect.Y+c.Rect.H; y++ {
			for x := c.Rect.X; x < c.Rect.X+c.Rect.W; x++ {
				if fabric[y][x] == 'X' {
					continue next
				}
			}
		}
		nonOverlappingClaim = c
	}

	return nonOverlappingClaim
}
