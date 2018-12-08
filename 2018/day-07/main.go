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

	fmt.Printf("Part 1: %s\n", orderedSteps(parseInput(input)))
	fmt.Printf("Part 2: %d\n", timeToCompleteSteps(parseInput(input), 5, 60))
}

func parseInput(input []string) *graph {
	g := &graph{}
	for _, s := range input {
		var from, to string
		fmt.Sscanf(s, "Step %s must be finished before step %s can begin.", &from, &to)
		if _, ok := g.findNode(from); !ok {
			g.addNode(from)
		}
		if _, ok := g.findNode(to); !ok {
			g.addNode(to)
		}
		g.addEdge(from, to)
	}
	return g
}

func orderedSteps(g *graph) string {
	res := []string{}
	for {
		next := g.nextNode()
		if next == "" {
			break
		}
		g.removeNode(next)
		res = append(res, next)
	}
	return strings.Join(res, "")
}

func timeToCompleteSteps(g *graph, workers, costbase int) int {
	seconds := 0

	slots := make([]string, workers)
	progress := map[string]int{}

	for {
		for n := range slots {
			if slots[n] == "" {
				node := g.nextNode()
				slots[n] = node
			}
		}

		for n := range slots {
			if slots[n] == "" {
				continue
			}
			progress[slots[n]]++
			if progress[slots[n]] >= cost(slots[n], costbase) {
				g.removeNode(slots[n])
				slots[n] = ""
			}
		}

		seconds++

		allWorkersIdle := true
		for _, node := range slots {
			if node != "" {
				allWorkersIdle = false
			}
		}

		if allWorkersIdle && g.nodeCount() == 0 {
			break
		}
	}

	return seconds
}

func cost(step string, costbase int) (seconds int) {
	return int(step[0]) - 'A' + costbase + 1
}

type graph struct {
	servedNodes []*node
	nodes       []*node
	edges       []*edge
}

type node struct {
	name string
}

func (n *node) String() string {
	return n.name
}

type edge struct {
	from *node
	to   *node
}

func (e *edge) String() string {
	return fmt.Sprintf("%s -> %s", e.from, e.to)
}

func (g *graph) addNode(name string) {
	g.nodes = append(g.nodes, &node{name: name})
}

func (g *graph) removeNode(name string) {
	idx := -1
	for i := range g.nodes {
		if g.nodes[i].name == name {
			idx = i
		}
	}
	if idx > -1 {
		var edgesToRemove []*edge
		for i := range g.edges {
			if g.edges[i].from == g.nodes[idx] || g.edges[i].to == g.nodes[idx] {
				edgesToRemove = append(edgesToRemove, g.edges[i])
			}
		}
		for _, e := range edgesToRemove {
			g.removeEdge(e)
		}
		g.nodes = append(g.nodes[:idx], g.nodes[idx+1:]...)
	}
}

func (g *graph) addEdge(from, to string) {
	fromNode, _ := g.findNode(from)
	toNode, _ := g.findNode(to)
	g.edges = append(g.edges, &edge{from: fromNode, to: toNode})
}

func (g *graph) findNode(name string) (*node, bool) {
	for n := range g.nodes {
		if g.nodes[n].name == name {
			return g.nodes[n], true
		}
	}
	return nil, false
}

func (g *graph) removeEdge(e *edge) {
	idx := -1
	for i := range g.edges {
		if g.edges[i] == e {
			idx = i
		}
	}
	if idx > -1 {
		g.edges = append(g.edges[:idx], g.edges[idx+1:]...)
	}
}

func (g *graph) startingNodes() []*node {
	var nodes []*node
next:
	for _, n := range g.nodes {
		for _, e := range g.edges {
			if e.to == n {
				continue next
			}
		}
		nodes = append(nodes, n)
	}
	return nodes
}

func (g *graph) nextNode() string {
	candidates := g.startingNodes()
	if len(candidates) > 0 {
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].name < candidates[j].name
		})
		for n := range candidates {
			if !g.everServed(candidates[n].name) {
				node := candidates[n]
				g.servedNodes = append(g.servedNodes, node)
				return node.name
			}
		}
	}
	return ""
}

func (g *graph) nodeCount() int {
	return len(g.nodes)
}

func (g *graph) everServed(name string) bool {
	for n := range g.servedNodes {
		if g.servedNodes[n].name == name {
			return true
		}
	}
	return false
}

func (g *graph) String() string {
	lines := []string{"digraph {"}
	sort.Slice(g.nodes, func(i, j int) bool {
		return g.nodes[i].name < g.nodes[j].name
	})
	for _, n := range g.nodes {
		lines = append(lines, fmt.Sprintf("  %s;", n.name))
	}
	sort.Slice(g.edges, func(i, j int) bool {
		return g.edges[i].String() < g.edges[j].String()
	})
	for _, e := range g.edges {
		lines = append(lines, fmt.Sprintf("  %s -> %s;", e.from.name, e.to.name))
	}
	lines = append(lines, "}")
	return strings.Join(lines, "\n")
}
