package main

import "fmt"

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

	fmt.Printf("Part 1: %d\n", sumOfMetadata(root))
	fmt.Printf("Part 2: %d\n", valueOfRootNode(root))
}

type node struct {
	parent   *node
	children []*node
	metadata []int
}

func readTree(input []int) *node {
	in := make(chan int, len(input))
	for _, n := range input {
		in <- n
	}
	close(in)

	return readTreeRecursively(nil, in)
}

func readTreeRecursively(parent *node, input <-chan int) *node {
	childrenCount, metadataCount := <-input, <-input
	n := &node{parent: parent}
	for i := 0; i < childrenCount; i++ {
		n.children = append(n.children, readTreeRecursively(n, input))
	}
	for i := 0; i < metadataCount; i++ {
		n.metadata = append(n.metadata, <-input)
	}
	return n
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
