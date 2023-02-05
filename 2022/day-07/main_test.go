package main

import "testing"

var input = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 95437; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 24933642; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestEntry_TotalSize(t *testing.T) {
	entry := &Entry{
		children: []*Entry{
			{size: 1},
			{size: 2},
			{
				children: []*Entry{
					{size: 3},
				},
			},
			{size: 4},
			{
				children: []*Entry{
					{
						children: []*Entry{
							{size: 5},
							{size: 6},
						},
					},
					{size: 7},
					{size: 8},
				},
			},
		},
	}

	if got, want := entry.totalSize(), 36; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestBuildTree(t *testing.T) {
	tree := buildTree(input)

	wantTreeStr := `- / (dir)
  - a (dir)
    - e (dir)
      - i (file, size=584)
    - f (file, size=29116)
    - g (file, size=2557)
    - h.lst (file, size=62596)
  - b.txt (file, size=14848514)
  - c.dat (file, size=8504156)
  - d (dir)
    - j (file, size=4060174)
    - d.log (file, size=8033020)
    - d.ext (file, size=5626152)
    - k (file, size=7214296)`

	if got, want := tree.String(), wantTreeStr; got != want {
		t.Errorf("got:\n%s\n\nwant:\n%s", got, want)
	}
}
