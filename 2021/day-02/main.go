package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	var s submarine
	s.pilot(input)

	fmt.Printf("Part 1: %d\n", s.horizontal*s.depth)
}

type submarine struct {
	horizontal int
	depth      int
}

func (s *submarine) pilot(commands []string) {
	for _, c := range commands {
		var dir string
		var dist int

		if _, err := fmt.Sscanf(c, "%s %d", &dir, &dist); err != nil {
			log.Fatal(err)
		}

		switch dir {
		case "forward":
			s.horizontal += dist
		case "up":
			s.depth -= dist
		case "down":
			s.depth += dist
		}
	}
}
