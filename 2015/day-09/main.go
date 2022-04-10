package main

import (
	"fmt"
	"sort"
)

func main() {
	distances := map[string]map[string]int{} // from -> to -> distance

	for {
		var (
			from, to string
			distance int
		)
		if _, err := fmt.Scanf("%s to %s = %d", &from, &to, &distance); err != nil {
			break
		}
		if _, ok := distances[from]; !ok {
			distances[from] = map[string]int{}
		}
		if _, ok := distances[to]; !ok {
			distances[to] = map[string]int{}
		}
		distances[from][to] = distance
		distances[to][from] = distance
	}

	_, shortestDistance := findShortestRoute(distances)
	_, longestDistance := findLongestRoute(distances)

	fmt.Printf("Part 1: %d\n", shortestDistance)
	fmt.Printf("Part 2: %d\n", longestDistance)
}

func findShortestRoute(distances map[string]map[string]int) (route []string, distance int) {
	locations := make([]string, 0, len(distances))
	for loc := range distances {
		locations = append(locations, loc)
	}

	route = make([]string, len(locations))
	sort.Strings(locations)

	for ok := true; ok; ok = perm(locations) {
		distCandidate := 0
		for i := 0; i < len(locations)-1; i++ {
			from, to := locations[i], locations[i+1]
			distCandidate += distances[from][to]
		}
		if distance == 0 || distCandidate < distance {
			copy(route, locations)
			distance = distCandidate
		}
	}

	return route, distance
}

func findLongestRoute(distances map[string]map[string]int) (route []string, distance int) {
	locations := make([]string, 0, len(distances))
	for loc := range distances {
		locations = append(locations, loc)
	}

	route = make([]string, len(locations))
	sort.Strings(locations)

	for ok := true; ok; ok = perm(locations) {
		distCandidate := 0
		for i := 0; i < len(locations)-1; i++ {
			from, to := locations[i], locations[i+1]
			distCandidate += distances[from][to]
		}
		if distCandidate > distance {
			copy(route, locations)
			distance = distCandidate
		}
	}

	return route, distance
}

// perm permutates s in-place into its next lexicographical permutation and
// returns true if successful, or false if there are no more permutations.
//
// https://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func perm(s []string) bool {
	var k int
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] < s[i+1] {
			k = i
			break
		}
	}
	for l := len(s) - 1; l > k; l-- {
		if s[k] < s[l] {
			s[k], s[l] = s[l], s[k]
			rev(s[k+1:])
			return true
		}
	}
	return false
}

// rev reverses the elements of s.
func rev(s []string) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
