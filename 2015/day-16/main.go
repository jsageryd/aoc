package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	detected := map[string]int{
		"cars":        2,
		"cats":        7,
		"children":    3,
		"goldfish":    5,
		"perfumes":    1,
		"pomeranians": 3,
		"samoyeds":    2,
		"trees":       3,
	}

	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	sues := parseInput(input)

	fmt.Printf("Part 1: %d\n", findSue(detected, sues))
	fmt.Printf("Part 2: %d\n", findRealSue(detected, sues))
}

type sue struct {
	id     int
	things map[string]int
}

func parseInput(input []string) []sue {
	var sues []sue

	for _, line := range input {
		s := sue{things: make(map[string]int)}
		head, tail, _ := strings.Cut(line, ": ")
		fmt.Sscanf(head, "Sue %d", &s.id)
		things := strings.Split(tail, ", ")
		for _, t := range things {
			name, kindsStr, _ := strings.Cut(t, ": ")
			kinds, _ := strconv.Atoi(kindsStr)
			s.things[name] = kinds
		}
		sues = append(sues, s)
	}

	return sues
}

func findSue(detected map[string]int, sues []sue) int {
	scores := make(map[int]int) // id -> score

	for _, s := range sues {
		for thing, kinds := range s.things {
			if kinds == detected[thing] {
				scores[s.id]++
			}
		}
	}

	var sueWithMaxScore int

	for id, score := range scores {
		if score > scores[sueWithMaxScore] {
			sueWithMaxScore = id
		}
	}

	return sueWithMaxScore
}

func findRealSue(detected map[string]int, sues []sue) int {
	scores := make(map[int]int) // id -> score

	for _, s := range sues {
		for thing, kinds := range s.things {
			var match bool

			switch thing {
			case "cats", "trees":
				match = kinds > detected[thing]
			case "pomeranians", "goldfish":
				match = kinds < detected[thing]
			default:
				match = kinds == detected[thing]
			}

			if match {
				scores[s.id]++
			}
		}
	}

	var sueWithMaxScore int

	for id, score := range scores {
		if score > scores[sueWithMaxScore] {
			sueWithMaxScore = id
		}
	}

	return sueWithMaxScore
}
