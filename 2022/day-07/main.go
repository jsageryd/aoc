package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	fmt.Printf("Part 2: %d\n", part2(input))
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

func part2(input []string) int {
	const totalSpace = 70000000
	const spaceNeeded = 30000000

	root := buildTree(input)

	freeSpace := totalSpace - root.totalSize()
	deleteAtLeast := spaceNeeded - freeSpace

	var allDirs []*Entry

	var rec func(cur *Entry)

	rec = func(cur *Entry) {
		if cur.typ == "dir" {
			allDirs = append(allDirs, cur)
			for _, c := range cur.children {
				rec(c)
			}
		}
	}

	rec(buildTree(input))

	sort.Slice(allDirs, func(i, j int) bool {
		return allDirs[i].totalSize() < allDirs[j].totalSize()
	})

	for _, entry := range allDirs {
		if s := entry.totalSize(); s >= deleteAtLeast {
			return s
		}
	}

	return 0
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
