package main

import (
	"fmt"
	"os"
	"testing"
)

func TestReduce(t *testing.T) {
	for n, tc := range []struct {
		in  string
		out string
	}{
		{"aA", ""},
		{"abBA", ""},
		{"abAB", "abAB"},
		{"aabAAB", "aabAAB"},
		{"dabAcCaCBAcCcaDA", "dabCBAcaDA"},
	} {
		if got, want := reduce(tc.in), tc.out; got != want {
			t.Errorf("[%d] reduce(%q) = %q, want %q", n, tc.in, got, want)
		}
	}
}

func TestReduceAlternate(t *testing.T) {
	for n, tc := range []struct {
		in  string
		out string
	}{
		{"dabAcCaCBAcCcaDA", "daDA"},
	} {
		if got, want := reduceAlternate(tc.in), tc.out; got != want {
			t.Errorf("[%d] reduceAlternate(%q) = %q, want %q", n, tc.in, got, want)
		}
	}
}

func TestAsciiToLower(t *testing.T) {
	input := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	want := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"

	for n := range input {
		if got, want := asciiToLower(input[n]), want[n]; got != want {
			t.Errorf("asciiToLower(%q) = %q, want %q", input[n], got, want)
		}
	}
}

func BenchmarkReduce(b *testing.B) {
	f, _ := os.Open("input")
	defer f.Close()

	var input string
	fmt.Fscanln(f, &input)

	for n := 0; n < b.N; n++ {
		reduce(input)
	}
}

func BenchmarkReduceAlternate(b *testing.B) {
	f, _ := os.Open("input")
	defer f.Close()

	var input string
	fmt.Fscanln(f, &input)

	for n := 0; n < b.N; n++ {
		reduceAlternate(input)
	}
}
