package main

import "testing"

var input = []string{
	"pbga (66)",
	"xhth (57)",
	"ebii (61)",
	"havc (66)",
	"ktlj (57)",
	"fwft (72) -> ktlj, cntj, xhth",
	"qoyq (66)",
	"padx (45) -> pbga, havc, qoyq",
	"tknk (41) -> ugml, padx, fwft",
	"jptl (61)",
	"ugml (68) -> gyxo, ebii, jptl",
	"gyxo (61)",
	"cntj (57)",
}

func TestParseProgramTree(t *testing.T) {
	root := parseProgramTree(input)

	if got, want := root.Name, "tknk"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFindUnbalancedChild(t *testing.T) {
	root := parseProgramTree(input)

	p, diff := findUnbalancedChild(root)

	if p == nil {
		t.Fatal("no child found")
	}

	if got, want := p.Name, "ugml"; got != want {
		t.Errorf("program = %q, want %q", got, want)
	}

	if got, want := diff, 8; got != want {
		t.Errorf("weight diff = %d, want %d", got, want)
	}
}

func TestTotalWeight(t *testing.T) {
	/*
	   p0 (weight 1)
	   ├── p1 (weight 2)
	   │   ├── p3 (weight 4)
	   │   └── p4 (weight 5)
	   └── p2 (weight 3)
	*/

	p := make([]*Program, 5)

	p[4] = &Program{Weight: 5, Parent: p[1]}
	p[3] = &Program{Weight: 4, Parent: p[1]}
	p[2] = &Program{Weight: 3, Parent: p[0]}
	p[1] = &Program{Weight: 2, Parent: p[0], Children: []*Program{p[3], p[4]}}
	p[0] = &Program{Weight: 1, Children: []*Program{p[1], p[2]}}

	totalWeights := []int{15, 11, 3, 4, 5}

	setTotalWeights(p[0])

	for n := 0; n < 5; n++ {
		if got, want := p[n].TotalWeight, totalWeights[n]; got != want {
			t.Errorf("p[%d].TotalWeight = %d, want %d", n, got, want)
		}
	}
}
