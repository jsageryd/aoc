package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var tracks [][]byte

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tracks = append(tracks, []byte(scanner.Text()))
	}

	carts := extractCarts(tracks)

	firstCollider, lastSurvivor := race(tracks, carts)

	fmt.Printf("Part 1: %d,%d\n", firstCollider.x, firstCollider.y)
	fmt.Printf("Part 2: %d,%d\n", lastSurvivor.x, lastSurvivor.y)
}

func race(tracks [][]byte, carts []*cart) (firstCollider *cart, lastSurvivor *cart) {
	for {
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y == carts[j].y {
				return carts[i].x < carts[j].x
			}
			return carts[i].y < carts[j].y
		})

		var colliders []*cart

		for _, cart := range carts {
			switch cart.heading {
			case north:
				cart.y--
			case west:
				cart.x--
			case south:
				cart.y++
			case east:
				cart.x++
			}

			for _, otherCart := range carts {
				if cart != otherCart && cart.x == otherCart.x && cart.y == otherCart.y {
					if firstCollider == nil {
						firstCollider = cart
					}
					colliders = append(colliders, cart)
					colliders = append(colliders, otherCart)
				}
			}

			switch tracks[cart.y][cart.x] {
			case '\\':
				switch cart.heading {
				case north:
					cart.heading = west
				case west:
					cart.heading = north
				case south:
					cart.heading = east
				case east:
					cart.heading = south
				}
			case '/':
				switch cart.heading {
				case north:
					cart.heading = east
				case west:
					cart.heading = south
				case south:
					cart.heading = west
				case east:
					cart.heading = north
				}
			case '+':
				cart.heading = applyTurn(cart.heading, cart.nextTurn)
				switch cart.nextTurn {
				case left:
					cart.nextTurn = straight
				case straight:
					cart.nextTurn = right
				case right:
					cart.nextTurn = left
				}
			}
		}

		var remainingCarts []*cart
	next:
		for _, cart := range carts {
			for _, collider := range colliders {
				if cart == collider {
					continue next
				}
			}
			remainingCarts = append(remainingCarts, cart)
		}
		carts = remainingCarts

		switch len(carts) {
		case 0:
			return firstCollider, nil
		case 1:
			return firstCollider, carts[0]
		}
	}
}

const (
	north = iota
	west
	south
	east
)

const (
	left = iota
	straight
	right
)

func applyTurn(heading int, turn int) (newHeading int) {
	switch turn {
	case left:
		switch heading {
		case north:
			return west
		case west:
			return south
		case south:
			return east
		case east:
			return north
		}
	case right:
		switch heading {
		case north:
			return east
		case west:
			return north
		case south:
			return west
		case east:
			return south
		}
	}
	return heading
}

type cart struct {
	id       int
	heading  int
	x, y     int
	nextTurn int
}

func (c *cart) String() string {
	return fmt.Sprintf("[%d] at %d,%d heading %s, turning %s next", c.id, c.x, c.y, headingToStr(c.heading), turnToStr(c.nextTurn))
}

func extractCarts(tracks [][]byte) []*cart {
	var carts []*cart
	for y := 0; y < len(tracks); y++ {
		for x := 0; x < len(tracks[y]); x++ {
			switch tracks[y][x] {
			case '^':
				tracks[y][x] = '|'
				carts = append(carts, &cart{heading: north, x: x, y: y})
			case '<':
				tracks[y][x] = '-'
				carts = append(carts, &cart{heading: west, x: x, y: y})
			case 'v':
				tracks[y][x] = '|'
				carts = append(carts, &cart{heading: south, x: x, y: y})
			case '>':
				tracks[y][x] = '-'
				carts = append(carts, &cart{heading: east, x: x, y: y})
			}
		}
	}
	for n := range carts {
		carts[n].id = n
	}
	return carts
}

func headingToStr(heading int) string {
	switch heading {
	case north:
		return "north"
	case west:
		return "west"
	case south:
		return "south"
	case east:
		return "east"
	default:
		return "unknown"
	}
}

func turnToStr(turn int) string {
	switch turn {
	case left:
		return "left"
	case straight:
		return "straight"
	case right:
		return "right"
	default:
		return "unknown"
	}
}
