package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	attendees, happiness := parse(input)

	totalHappiness := optimalTotalHappiness(attendees, happiness)

	fmt.Printf("Part 1: %d\n", totalHappiness)

	attendees = append(attendees, "self")
	sort.Strings(attendees)

	totalHappinessWithSelf := optimalTotalHappiness(attendees, happiness)

	fmt.Printf("Part 2: %d\n", totalHappinessWithSelf)
}

func optimalTotalHappiness(attendees []string, happiness map[string]map[string]int) int {
	var optimalTotal int

	// this tests more permutations than it needs to, but I already had the perm
	// function and it's fast enough.
	for ok := true; ok; ok = perm(attendees) {
		var total int

		for i := range attendees {
			attendee := attendees[i]
			left := attendees[(i+len(attendees)-1)%len(attendees)]
			right := attendees[(i+1)%len(attendees)]
			total += happiness[attendee][left]
			total += happiness[attendee][right]
		}

		if total > optimalTotal {
			optimalTotal = total
		}
	}

	return optimalTotal
}

// parse returns the list of (unique) attendees, and a happiness map showing the
// happiness of an attendee if places next to a neighbour, e.g. Alice would gain
// 54 happiness units by sitting next to Bob -> happiness["Alice"]["Bob"] == 54.
func parse(input []string) (attendees []string, happiness map[string]map[string]int) {
	attendeeMap := make(map[string]struct{})
	happiness = make(map[string]map[string]int)

	for _, line := range input {
		var (
			attendee   string
			neighbour  string
			gainOrLose string
			units      int
		)

		fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s", &attendee, &gainOrLose, &units, &neighbour)

		attendeeMap[attendee] = struct{}{}

		neighbour = strings.TrimRight(neighbour, ".")

		var change int

		if gainOrLose == "gain" {
			change = units
		} else {
			change = -units
		}

		if _, ok := happiness[attendee]; !ok {
			happiness[attendee] = make(map[string]int)
		}

		happiness[attendee][neighbour] = change
	}

	for a := range attendeeMap {
		attendees = append(attendees, a)
	}

	sort.Strings(attendees)

	return attendees, happiness
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
