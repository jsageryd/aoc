package main

import (
	"fmt"
	"io"
	"log"
	"sort"
)

type reindeer struct {
	name     string
	speed    int
	duration int
	rest     int
}

func main() {
	var rs []reindeer

	for {
		var r reindeer

		if _, err := fmt.Scanf(
			"%s can fly %d km/s for %d seconds, but then must rest for %d seconds.\n",
			&r.name, &r.speed, &r.duration, &r.rest,
		); err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		rs = append(rs, r)
	}

	fmt.Printf("Part 1: %d\n", furthestDistance(rs, 2503))
	fmt.Printf("Part 2: %d\n", highestScore(rs, 2503))
}

func highestScore(rs []reindeer, seconds int) (score int) {
	scores := make(map[reindeer]int)

	for s := 1; s <= seconds; s++ {
		l := leader(rs, s)
		scores[l]++
	}

	var maxScore int

	for _, s := range scores {
		if s > maxScore {
			maxScore = s
		}
	}

	return maxScore
}

func leader(rs []reindeer, seconds int) reindeer {
	var lead reindeer
	var leadDistance int

	for _, r := range rs {
		distance := flyReindeer(r, seconds)
		if distance > leadDistance {
			lead = r
			leadDistance = distance
		}
	}

	return lead
}

func furthestDistance(rs []reindeer, seconds int) (distance int) {
	var distances []int

	for _, r := range rs {
		distances = append(distances, flyReindeer(r, seconds))
	}

	sort.Ints(distances)

	return distances[len(distances)-1]
}

func flyReindeer(r reindeer, seconds int) (distance int) {
	iterations := seconds / (r.duration + r.rest)
	remainingDuration := seconds % (r.duration + r.rest)

	if remainingDuration > r.duration {
		remainingDuration = r.duration
	}

	return iterations*r.speed*r.duration + r.speed*remainingDuration
}
