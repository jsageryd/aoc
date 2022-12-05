package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"    [D]",
		"[N] [C]",
		"[Z] [M] [P]",
		" 1   2   3",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	if got, want := part1(input), "CMZ"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"    [D]",
		"[N] [C]",
		"[Z] [M] [P]",
		" 1   2   3",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	if got, want := part2(input), "MCD"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestParseInput(t *testing.T) {
	input := []string{
		"    [D]",
		"[N] [C]",
		"[Z] [M] [P]",
		" 1   2   3",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	gotStacks, gotSteps := parseInput(input)

	wantStacks := []stack{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}

	wantSteps := []step{
		{count: 1, from: 2, to: 1},
		{count: 3, from: 1, to: 3},
		{count: 2, from: 2, to: 1},
		{count: 1, from: 1, to: 2},
	}

	if fmt.Sprint(gotStacks) != fmt.Sprint(wantStacks) {
		t.Errorf("got %v, want %v", gotStacks, wantStacks)
	}

	if fmt.Sprint(gotSteps) != fmt.Sprint(wantSteps) {
		t.Errorf("got %v, want %v", gotSteps, wantSteps)
	}
}

func TestParseStacks(t *testing.T) {
	stacksStrs := []string{
		"    [D]",
		"[N] [C]",
		"[Z] [M] [P]",
		" 1   2   3",
	}

	gotStacks := parseStacks(stacksStrs)

	wantStacks := []stack{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}

	if fmt.Sprint(gotStacks) != fmt.Sprint(wantStacks) {
		t.Errorf("got %v, want %v", gotStacks, wantStacks)
	}
}

func TestParseSteps(t *testing.T) {
	stepStrs := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	gotSteps := parseSteps(stepStrs)

	wantSteps := []step{
		{count: 1, from: 2, to: 1},
		{count: 3, from: 1, to: 3},
		{count: 2, from: 2, to: 1},
		{count: 1, from: 1, to: 2},
	}

	if fmt.Sprint(gotSteps) != fmt.Sprint(wantSteps) {
		t.Errorf("got %v, want %v", gotSteps, wantSteps)
	}
}

func TestMoveCrates(t *testing.T) {
	crates := []stack{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}

	steps := []step{
		{count: 1, from: 2, to: 1},
		{count: 3, from: 1, to: 3},
		{count: 2, from: 2, to: 1},
		{count: 1, from: 1, to: 2},
	}

	moveCrates(crates, steps)

	wantCrates := []stack{
		{"C"},
		{"M"},
		{"P", "D", "N", "Z"},
	}

	if fmt.Sprint(crates) != fmt.Sprint(wantCrates) {
		t.Errorf("got %v, want %v", crates, wantCrates)
	}
}

func TestMoveCrates2(t *testing.T) {
	crates := []stack{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}

	steps := []step{
		{count: 1, from: 2, to: 1},
		{count: 3, from: 1, to: 3},
		{count: 2, from: 2, to: 1},
		{count: 1, from: 1, to: 2},
	}

	moveCrates2(crates, steps)

	wantCrates := []stack{
		{"M"},
		{"C"},
		{"P", "Z", "N", "D"},
	}

	if fmt.Sprint(crates) != fmt.Sprint(wantCrates) {
		t.Errorf("got %v, want %v", crates, wantCrates)
	}
}

func TestStackPush(t *testing.T) {
	var s stack

	s.push("a")
	s.push("b")
	s.push("c")

	if got, want := s, []string{"a", "b", "c"}; fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestStackPop(t *testing.T) {
	var s stack

	s.push("a")
	s.push("b")
	s.push("c")

	item := s.pop()

	if got, want := s, []string{"a", "b"}; fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("stack is %v, want %v", got, want)
	}

	if got, want := item, "c"; got != want {
		t.Errorf("item is %q, want %q", got, want)
	}
}
