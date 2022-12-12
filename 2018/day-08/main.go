package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"
)

func main() {
	var input []int

	for {
		var n int
		if _, err := fmt.Scan(&n); err != nil {
			break
		}
		input = append(input, n)
	}

	root := readTree(input)

	outputDot := flag.Bool("dot", false, "Output GraphViz dot")
	flag.Parse()

	if *outputDot {
		fmt.Println(dot(root))
		return
	}

	fmt.Printf("Part 1: %d\n", sumOfMetadata(root))
	fmt.Printf("Part 2: %d\n", valueOfRootNode(root))
}

type node struct {
	children []*node
	metadata []int
}

func readTree(input []int) *node {
	in := make(chan int, len(input))
	for _, n := range input {
		in <- n
	}
	close(in)

	var recurse func(*node) *node

	recurse = func(root *node) *node {
		childrenCount, metadataCount := <-in, <-in
		n := &node{}
		for i := 0; i < childrenCount; i++ {
			n.children = append(n.children, recurse(n))
		}
		for i := 0; i < metadataCount; i++ {
			n.metadata = append(n.metadata, <-in)
		}
		return n
	}

	return recurse(nil)
}

func sumOfMetadata(root *node) int {
	sum := 0
	for _, n := range root.metadata {
		sum += n
	}
	for _, c := range root.children {
		sum += sumOfMetadata(c)
	}
	return sum
}

func valueOfRootNode(root *node) int {
	sum := 0

	if len(root.children) == 0 {
		for _, n := range root.metadata {
			sum += n
		}
		return sum
	}

	for _, n := range root.metadata {
		if n-1 < len(root.children) {
			sum += valueOfRootNode(root.children[n-1])
		}
	}

	return sum
}

func dot(root *node) string {
	nodeIDs := map[*node]int{}

	nodeID := 0

	var recurseChildren func(root *node) [][2]int

	recurseChildren = func(root *node) [][2]int {
		var res [][2]int
		if _, ok := nodeIDs[root]; !ok {
			nodeID++
			nodeIDs[root] = nodeID
		}
		for _, n := range root.children {
			if _, ok := nodeIDs[n]; !ok {
				nodeID++
				nodeIDs[n] = nodeID
			}
			res = append(res, [2]int{nodeIDs[root], nodeIDs[n]})
			res = append(res, recurseChildren(n)...)
		}
		return res
	}

	var recurseMetadata func(root *node) [][2]int

	recurseMetadata = func(root *node) [][2]int {
		var res [][2]int
		for _, n := range root.metadata {
			if n-1 < len(root.children) {
				res = append(res, [2]int{nodeIDs[root], nodeIDs[root.children[n-1]]})
			}
		}
		for _, n := range root.children {
			res = append(res, recurseMetadata(n)...)
		}
		return res
	}

	childEdges := recurseChildren(root)
	sort.Slice(childEdges, func(i, j int) bool {
		if childEdges[i][0] == childEdges[j][0] {
			return childEdges[i][1] < childEdges[j][1]
		}
		return childEdges[i][0] < childEdges[j][0]
	})

	metadataEdges := recurseMetadata(root)
	sort.Slice(metadataEdges, func(i, j int) bool {
		if metadataEdges[i][0] == metadataEdges[j][0] {
			return metadataEdges[i][1] < metadataEdges[j][1]
		}
		return metadataEdges[i][0] < metadataEdges[j][0]
	})

	lines := []string{"digraph {"}

	nodeIDsSorted := make([]int, 0, len(nodeIDs))
	for _, id := range nodeIDs {
		nodeIDsSorted = append(nodeIDsSorted, id)
	}
	sort.Ints(nodeIDsSorted)
	for _, id := range nodeIDsSorted {
		lines = append(lines, fmt.Sprintf("  %d;", id))
	}
	for _, e := range childEdges {
		lines = append(lines, fmt.Sprintf("  %d -> %d;", e[0], e[1]))
	}
	for _, e := range metadataEdges {
		lines = append(lines, fmt.Sprintf("  %d -> %d [style = dotted];", e[0], e[1]))
	}
	lines = append(lines, "}")

	return strings.Join(lines, "\n")
}
