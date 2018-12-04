package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	sort.Strings(input)

	guard, minute := guardMostAsleepWithMinute(input)
	guard2, minute2 := guardMostAsleepWithMinuteAlternate(input)

	fmt.Printf("Part 1: %d\n", guard*minute)
	fmt.Printf("Part 2: %d\n", guard2*minute2)
}

func guardMostAsleepWithMinute(input []string) (guard, minute int) {
	guardTotalTime := map[int]time.Duration{}
	guardMinuteFreq := map[int]map[int]int{}

	countMinutesFreq := func(guard int, t1, t2 time.Time) {
		if _, ok := guardMinuteFreq[guard]; !ok {
			guardMinuteFreq[guard] = make(map[int]int)
		}
		diffMinutes := int(t2.Sub(t1).Minutes())
		startMinute := t1.Minute()
		for n := startMinute; n < startMinute+diffMinutes; n++ {
			guardMinuteFreq[guard][n]++
		}
	}

	var start time.Time
	for n := range input {
		t, _ := time.Parse("2006-01-02 15:04", input[n][1:17])

		msgWords := strings.Split(input[n][19:], " ")
		switch msgWords[0] {
		case "Guard":
			if !start.IsZero() {
				guardTotalTime[guard] += t.Sub(start)
				countMinutesFreq(guard, start, t)
			}
			guard, _ = strconv.Atoi(msgWords[1][1:])
			start = time.Time{}
		case "falls":
			start = t
		case "wakes":
			guardTotalTime[guard] += t.Sub(start)
			countMinutesFreq(guard, start, t)
			start = time.Time{}
		default:
			log.Fatalf("unknown input: %v", msgWords)
		}
	}

	for g := range guardTotalTime {
		if guardTotalTime[g] > guardTotalTime[guard] {
			guard = g
		}
	}

	guardMostFreqMinute := -1
	for min := range guardMinuteFreq[guard] {
		if guardMinuteFreq[guard][min] > guardMinuteFreq[guard][guardMostFreqMinute] {
			guardMostFreqMinute = min
		}
	}

	return guard, guardMostFreqMinute
}

func guardMostAsleepWithMinuteAlternate(input []string) (guard, minute int) {
	guardMinuteFreq := map[int]map[int]int{}

	countMinutesFreq := func(guard int, t1, t2 time.Time) {
		if _, ok := guardMinuteFreq[guard]; !ok {
			guardMinuteFreq[guard] = make(map[int]int)
		}
		diffMinutes := int(t2.Sub(t1).Minutes())
		startMinute := t1.Minute()
		for n := startMinute; n < startMinute+diffMinutes; n++ {
			guardMinuteFreq[guard][n]++
		}
	}

	var start time.Time
	for n := range input {
		t, _ := time.Parse("2006-01-02 15:04", input[n][1:17])

		msgWords := strings.Split(input[n][19:], " ")
		switch msgWords[0] {
		case "Guard":
			if !start.IsZero() {
				countMinutesFreq(guard, start, t)
			}
			guard, _ = strconv.Atoi(msgWords[1][1:])
			start = time.Time{}
		case "falls":
			start = t
		case "wakes":
			countMinutesFreq(guard, start, t)
			start = time.Time{}
		default:
			log.Fatalf("unknown input: %v", msgWords)
		}
	}

	guardMostFreqMinuteAlt := -1
	for g := range guardMinuteFreq {
		for min := range guardMinuteFreq[g] {
			if guardMinuteFreq[g][min] > guardMinuteFreq[guard][guardMostFreqMinuteAlt] {
				guard = g
				guardMostFreqMinuteAlt = min
			}
		}
	}

	return guard, guardMostFreqMinuteAlt
}
