package main

import (
	"fmt"
	"strings"
	"testing"
)

var input = []string{
	"Monkey 0:",
	"  Starting items: 79, 98",
	"  Operation: new = old * 19",
	"  Test: divisible by 23",
	"    If true: throw to monkey 2",
	"    If false: throw to monkey 3",
	"",
	"Monkey 1:",
	"  Starting items: 54, 65, 75, 74",
	"  Operation: new = old + 6",
	"  Test: divisible by 19",
	"    If true: throw to monkey 2",
	"    If false: throw to monkey 0",
	"",
	"Monkey 2:",
	"  Starting items: 79, 60, 97",
	"  Operation: new = old * old",
	"  Test: divisible by 13",
	"    If true: throw to monkey 1",
	"    If false: throw to monkey 3",
	"",
	"Monkey 3:",
	"  Starting items: 74",
	"  Operation: new = old + 3",
	"  Test: divisible by 17",
	"    If true: throw to monkey 0",
	"    If false: throw to monkey 1",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 10605; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 2713310158; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestParseMonkeys(t *testing.T) {
	gotMonkeys := parseMonkeys(input)

	wantMonkeys := []*monkey{
		{items: []int{79, 98}, op: "old * 19", testDiv: 23, trueMonkey: 2, falseMonkey: 3, friends: gotMonkeys},
		{items: []int{54, 65, 75, 74}, op: "old + 6", testDiv: 19, trueMonkey: 2, falseMonkey: 0, friends: gotMonkeys},
		{items: []int{79, 60, 97}, op: "old * old", testDiv: 13, trueMonkey: 1, falseMonkey: 3, friends: gotMonkeys},
		{items: []int{74}, op: "old + 3", testDiv: 17, trueMonkey: 0, falseMonkey: 1, friends: gotMonkeys},
	}

	var gotMonkeysStr, wantMonkeysStr []string
	for _, m := range gotMonkeys {
		gotMonkeysStr = append(gotMonkeysStr, fmt.Sprint(m))
	}
	for _, m := range wantMonkeys {
		wantMonkeysStr = append(wantMonkeysStr, fmt.Sprint(m))
	}

	if fmt.Sprint(gotMonkeysStr) != fmt.Sprint(wantMonkeysStr) {
		t.Errorf(
			"got:\n%s\n\nwant:\n%s",
			strings.Join(gotMonkeysStr, "\n"),
			strings.Join(wantMonkeysStr, "\n"),
		)
	}
}
