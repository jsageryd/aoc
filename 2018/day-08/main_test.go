package main

import "testing"

func TestSumOfMetadata(t *testing.T) {
	input := []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}

	root := readTree(input)

	if got, want := sumOfMetadata(root), 138; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestValueOfRootNode(t *testing.T) {
	input := []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}

	root := readTree(input)

	if got, want := valueOfRootNode(root), 66; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
