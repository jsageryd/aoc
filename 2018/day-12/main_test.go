package main

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func TestStep(t *testing.T) {
	rules := []string{
		"...## => #",
		"..#.. => #",
		".#... => #",
		".#.#. => #",
		".#.## => #",
		".##.. => #",
		".#### => #",
		"#.#.# => #",
		"#.### => #",
		"##.#. => #",
		"##.## => #",
		"###.. => #",
		"###.# => #",
		"####. => #",
	}

	t.Run("Simple", func(t *testing.T) {
		state := "#..#.#..##......###...###..........."

		state, offset := step(state, rules, 20)

		if got, want := sumOfIndicesWithPots(state, offset), 325; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Equilibrium", func(t *testing.T) {
		state := "#..#.#..##......###...###..........."

		state, offset := step(state, rules, 200)

		if got, want := sumOfIndicesWithPots(state, offset), 3374; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func BenchmarkStep(b *testing.B) {
	rules := []string{
		"...## => #",
		"..#.. => #",
		".#... => #",
		".#.#. => #",
		".#.## => #",
		".##.. => #",
		".#### => #",
		"#.#.# => #",
		"#.### => #",
		"##.#. => #",
		"##.## => #",
		"###.. => #",
		"###.# => #",
		"####. => #",
	}

	for n := 0; n < b.N; n++ {
		step("#..#.#..##......###...###...........", rules, 100000)
	}
}

func TestStateConversion(t *testing.T) {
	b := func(s string) byte {
		n, _ := strconv.ParseUint(s, 2, 8)
		return byte(n)
	}
	ss := func(b []byte) []string {
		ss := []string{}
		for _, v := range b {
			ss = append(ss, fmt.Sprintf("%05b", v))
		}
		return ss
	}

	for n, tc := range []struct {
		state    string
		internal []byte
	}{
		{"", []byte{}},
		{".", []byte{b("00000")}},
		{"..", []byte{b("00000"), b("00000")}},
		{"...", []byte{b("00000"), b("00000"), b("00000")}},
		{"#", []byte{b("00100")}},
		{"##", []byte{b("00110"), b("01100")}},
		{"###", []byte{b("00111"), b("01110"), b("11100")}},
		{"####", []byte{b("00111"), b("01111"), b("11110"), b("11100")}},
		{"#####", []byte{b("00111"), b("01111"), b("11111"), b("11110"), b("11100")}},
		{"######", []byte{b("00111"), b("01111"), b("11111"), b("11111"), b("11110"), b("11100")}},
		{".#.#.", []byte{b("00010"), b("00101"), b("01010"), b("10100"), b("01000")}},
		{"#.#.#", []byte{b("00101"), b("01010"), b("10101"), b("01010"), b("10100")}},
	} {
		if got, want := stateStringToInternal(tc.state), tc.internal; !bytes.Equal(got, want) {
			t.Errorf("[%d] stateStringToInternal(%q) -> %v, want %v", n, tc.state, ss(got), ss(want))
		}

		if got, want := internalToStateString(tc.internal), tc.state; got != want {
			t.Errorf("[%d] internalToStateString(%v) -> %q, want %q", n, ss(tc.internal), got, want)
		}
	}
}

func TestRuleToInternal(t *testing.T) {
	b := func(s string) byte {
		n, _ := strconv.ParseUint(s, 2, 8)
		return byte(n)
	}

	for n, tc := range []struct {
		rule     string
		internal byte
	}{
		{"####. => .", b("11110")},
		{"..#.. => .", b("00100")},
		{"#.#.. => .", b("10100")},
		{".##.. => .", b("01100")},
		{"##... => .", b("11000")},
		{"#.##. => .", b("10110")},
		{"####. => #", b("11110")},
		{"..#.. => #", b("00100")},
		{"#.#.. => #", b("10100")},
		{".##.. => #", b("01100")},
		{"##... => #", b("11000")},
		{"#.##. => #", b("10110")},
	} {
		if got, want := ruleToInternal(tc.rule), tc.internal; got != want {
			t.Errorf("[%d] ruleToInternal(%q) -> %05b, want %05b", n, tc.rule, got, want)
		}
	}
}

func TestResultToInternal(t *testing.T) {
	b := func(s string) byte {
		n, _ := strconv.ParseUint(s, 2, 8)
		return byte(n)
	}
	ss := func(b []byte) []string {
		ss := []string{}
		for _, v := range b {
			ss = append(ss, fmt.Sprintf("%05b", v))
		}
		return ss
	}

	for n, tc := range []struct {
		result   []byte
		internal []byte
	}{
		{[]byte{}, []byte{}},
		{[]byte{b("00000")}, []byte{b("00000")}},
		{[]byte{b("00100")}, []byte{b("00100")}},
		{[]byte{b("11111")}, []byte{b("00100")}},
		{[]byte{b("00100"), b("00100")}, []byte{b("00110"), b("01100")}},
		{[]byte{b("00100"), b("00100"), b("00100")}, []byte{b("00111"), b("01110"), b("11100")}},
	} {
		if got, want := resultToInternal(tc.result), tc.internal; !bytes.Equal(got, want) {
			t.Errorf("[%d] resultToInternal(%v) -> %v, want %v", n, ss(tc.result), ss(got), ss(want))
		}
	}
}

func TestRecenterAndAdjustLength(t *testing.T) {
	for n, tc := range []struct {
		a         []byte
		b         []byte
		offsetIn  int
		offsetOut int
	}{
		{[]byte{}, []byte{0, 0, 0, 0, 0}, 0, 2},
		{[]byte{1}, []byte{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, 0, 5},
		{[]byte{4}, []byte{0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0}, 0, 5},
		{[]byte{0, 0, 4, 0, 0, 0}, []byte{0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0}, 2, 5},
		{[]byte{0, 0, 0, 4, 0, 0}, []byte{0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0}, 3, 5},
		{[]byte{0, 0, 0, 0, 0, 0, 4, 0, 0}, []byte{0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0}, 6, 5},
		{[]byte{0, 0, 4, 0, 0, 0, 0, 0, 0}, []byte{0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0}, 2, 5},
		{[]byte{0, 0, 0, 0, 0, 4}, []byte{0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0}, 5, 5},
		{[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4}, []byte{0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0}, 10, 5},
		{[]byte{4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []byte{0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0}, 0, 5},
		{[]byte{0, 0, 4, 0, 0, 0, 0, 0, 4, 0, 0}, []byte{0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0}, 2, 5},
		{[]byte{0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0}, []byte{0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0}, 6, 5},
	} {
		out, offset := recenterAndAdjustLength(tc.a, tc.offsetIn)

		if got, want := out, tc.b; !bytes.Equal(got, want) {
			t.Errorf("[%d] slice = %v, want %v", n, got, want)
		}

		if got, want := offset, tc.offsetOut; got != want {
			t.Errorf("[%d] offset = %d, want %d", n, got, want)
		}
	}
}
