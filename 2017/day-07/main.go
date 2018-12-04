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

	root := rootProgram(input)

	unbalanced, diff := findUnbalancedChild(root)

	fmt.Printf("Part 1: %s\n", root.Name)
	fmt.Printf("Part 2: %d\n", unbalanced.Weight-diff)
}

type Program struct {
	Name        string     `json:"name"`
	Weight      int        `json:"weight"`
	TotalWeight int        `json:"total_weight"`
	Parent      *Program   `json:"-"`
	Children    []*Program `json:"children"`
}

func rootProgram(input []string) *Program {
	for _, p := range parseProgramTree(input) {
		if p.Parent == nil {
			return p
		}
	}
	return nil
}

func setTotalWeights(root *Program) {
	root.TotalWeight = root.Weight
	for _, c := range root.Children {
		setTotalWeights(c)
		root.TotalWeight += c.TotalWeight
	}
}

func findUnbalancedChild(root *Program) (child *Program, diff int) {
	if len(root.Children) < 3 {
		return nil, 0
	}

	sort.Slice(root.Children, func(i, j int) bool {
		return root.Children[i].TotalWeight >= root.Children[j].TotalWeight
	})

	if c1, c2 := root.Children[0], root.Children[1]; c1.TotalWeight != c2.TotalWeight {
		if c, d := findUnbalancedChild(c1); d != 0 {
			return c, d
		}
		return c1, c1.TotalWeight - c2.TotalWeight
	}

	return nil, 0
}

func parseProgramTree(input []string) []*Program {
	programMap := map[string]*Program{}

	for _, entry := range input {
		var name string
		var weight int
		fmt.Sscanf(entry, "%s (%d)", &name, &weight)
		parent, ok := programMap[name]
		if !ok {
			parent = &Program{Name: name}
			programMap[name] = parent
		}
		parent.Weight = weight
		if idx := strings.Index(entry, "->"); idx > -1 {
			childNames := strings.Split(entry[idx+3:], ", ")
			for _, name := range childNames {
				child, ok := programMap[name]
				if !ok {
					child = &Program{Name: name}
					programMap[name] = child
				}
				child.Parent = parent
			}
		}
	}

	programs := make([]*Program, 0, len(programMap))

	var root *Program

	for _, p := range programMap {
		if p.Parent != nil {
			p.Parent.Children = append(p.Parent.Children, p)
		} else {
			root = p
		}
		programs = append(programs, p)
	}

	setTotalWeights(root)

	return programs
}
