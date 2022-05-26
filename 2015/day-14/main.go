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
