package main

import (
	"fmt"
	"strings"
)

func main() {
	var players, lastMarble int
	fmt.Scanf("%d players; last marble is worth %d points", &players, &lastMarble)

	highestScore := play(players, lastMarble)

	fmt.Printf("Part 1: %d\n", highestScore)
}

type circle struct {
	curIdx int
	data   []int
}

// place places the given marble into the circle between the marbles that are 1
// and 2 marbles clockwise of the current marble.
func (c *circle) place(marble int) {
	idx := 0
	if len(c.data) > 0 {
		idx = (c.curIdx + 2) % len(c.data)
	}
	c.data = append(c.data, 0)
	if idx == 0 {
		idx = len(c.data) - 1
	}
	copy(c.data[idx+1:], c.data[idx:])
	c.data[idx] = marble
	c.curIdx = idx
}

// delete deletes the marble at the given offset from the current marble and
// returns its value.
func (c *circle) delete(offset int) (marble int) {
	idx := (c.curIdx + offset)
	for idx < 0 {
		idx += len(c.data)
	}
	marble = c.data[idx]
	c.data = append(c.data[:idx], c.data[idx+1:]...)
	c.curIdx = idx % len(c.data)
	return marble
}

func (c *circle) String() string {
	s := make([]string, 0, len(c.data))
	for n := range c.data {
		format := " %d "
		if n == c.curIdx {
			format = "[%d]"
		}
		m := fmt.Sprintf(format, c.data[n])
		s = append(s, m)
	}
	return strings.Join(s, "")
}

func play(players, lastMarble int) int {
	remainingMarbles := make([]int, 0, lastMarble)
	for n := lastMarble; n >= 1; n-- {
		remainingMarbles = append(remainingMarbles, n)
	}

	currentPlayer := -1
	playerScores := map[int]int{}

	c := &circle{}
	c.place(0)

	for len(remainingMarbles) > 0 {
		currentPlayer = (currentPlayer + 1) % players

		nextMarble := remainingMarbles[len(remainingMarbles)-1]
		remainingMarbles = remainingMarbles[:len(remainingMarbles)-1]

		if nextMarble%23 == 0 {
			playerScores[currentPlayer] += nextMarble
			playerScores[currentPlayer] += c.delete(-7)
		} else {
			c.place(nextMarble)
		}
	}

	highestScore := 0
	for _, score := range playerScores {
		if score > highestScore {
			highestScore = score
		}
	}

	return highestScore
}
