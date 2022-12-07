package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	var sum int

	var rec func(cur *Entry)

	rec = func(cur *Entry) {
		if cur.typ == "dir" {
			if s := cur.totalSize(); s <= 100000 {
				sum += s
			}
			for _, c := range cur.children {
				rec(c)
			}
		}
	}

	rec(buildTree(input))

	return sum
}

type Entry struct {
	parent   *Entry
	children []*Entry
	typ      string
	name     string
	size     int
}

func (e *Entry) totalSize() int {
	var totalSize int

	var rec func(cur *Entry)

	rec = func(cur *Entry) {
		totalSize += cur.size
		for _, c := range cur.children {
			rec(c)
		}
	}

	rec(e)

	return totalSize
}

func (e *Entry) String() string {
	var rec func(cur *Entry, level int)

	var list []string

	rec = func(cur *Entry, level int) {
		var info string
		switch cur.typ {
		case "file":
			info = fmt.Sprintf("file, size=%d", cur.size)
		case "dir":
			info = "dir"
		}
		list = append(list, fmt.Sprintf(
			"%s- %s (%s)",
			strings.Repeat("  ", level),
			cur.name,
			info,
		))
		if cur.typ == "dir" {
			for _, c := range cur.children {
				rec(c, level+1)
			}
		}
	}

	rec(e, 0)

	return strings.Join(list, "\n")
}

func buildTree(input []string) *Entry {
	root := &Entry{
		typ:  "dir",
		name: "/",
	}

	pwd := root

	var ls bool

	for _, line := range input {
		fields := strings.Fields(line)

		if fields[0] == "$" {
			ls = false

			switch fields[1] {
			case "cd":
				switch fields[2] {
				case "/":
					pwd = root
				case "..":
					pwd = pwd.parent
				default:
					for _, c := range pwd.children {
						if c.name == fields[2] {
							pwd = c
							break
						}
					}
				}
			case "ls":
				ls = true
			}
		} else {
			if ls {
				entry := &Entry{parent: pwd}

				if fields[0] == "dir" {
					entry.typ = "dir"
					entry.name = fields[1]
				} else {
					entry.typ = "file"
					entry.size, _ = strconv.Atoi(fields[0])
					entry.name = fields[1]
				}

				pwd.children = append(pwd.children, entry)
			}
		}
	}

	return root
}
