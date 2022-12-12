package main

import (
	"container/ring"
	"fmt"
	"strings"
)

func main() {
	var players, lastMarble int
	fmt.Scanf("%d players; last marble is worth %d points", &players, &lastMarble)

	highestScore := play(players, lastMarble)
	highestScore2 := play(players, lastMarble*100)

	fmt.Printf("Part 1: %d\n", highestScore)
	fmt.Printf("Part 2: %d\n", highestScore2)
}

type circle struct {
	data *ring.Ring
}

// place places the given marble into the circle between the marbles that are 1
// and 2 marbles clockwise of the current marble.
func (c *circle) place(marble int) {
	if c.data == nil {
		c.data = ring.New(1)
		c.data.Value = marble
		return
	}
	r := ring.New(1)
	r.Value = marble
	c.data = c.data.Next()
	c.data.Link(r)
	c.data = c.data.Next()
}

// delete deletes the marble at the given offset from the current marble and
// returns its value.
func (c *circle) delete(offset int) (marble int) {
	c.data = c.data.Move(offset - 1)
	defer func() { c.data = c.data.Next() }()
	return c.data.Unlink(1).Value.(int)
}

func (c *circle) String() string {
	if c.data == nil {
		return ""
	}
	cur := c.data.Value.(int)
	var data []int
	c.data.Do(func(v any) {
		data = append(data, v.(int))
	})
	for data[0] != 0 {
		zeroth := data[0]
		for n := 0; n < len(data)-1; n++ {
			data[n] = data[n+1]
		}
		data[len(data)-1] = zeroth
	}
	s := make([]string, 0, len(data))
	for _, d := range data {
		format := " %d "
		if d == cur {
			format = "[%d]"
		}
		m := fmt.Sprintf(format, d)
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
