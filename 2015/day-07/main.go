package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var instructions []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	c := newCircuit(instructions)

	a := c.measure("a")

	fmt.Printf("Part 1: %d\n", a)

	c = newCircuit(instructions)
	c.signals["b"] = &signal{v: a}

	fmt.Printf("Part 2: %d\n", c.measure("a"))
}

type circuit struct {
	signals map[string]*signal
}

func newCircuit(instructions []string) *circuit {
	c := &circuit{
		signals: make(map[string]*signal),
	}

	for _, inst := range instructions {
		s := strings.Fields(inst)
		target := s[len(s)-1]
		s = s[:len(s)-2]
		switch len(s) {
		case 1: // signal (int) or other wire (identifier)
			if i, err := strconv.Atoi(s[0]); err == nil {
				c.signals[target] = newSignal(func() uint16 { return uint16(i) })
			} else {
				c.signals[target] = newSignal(func() uint16 { return c.signals[s[0]].val() })
			}
		case 2: // NOT
			if i, err := strconv.Atoi(s[1]); err == nil {
				c.signals[target] = newSignal(func() uint16 { return ^uint16(i) })
			} else {
				c.signals[target] = newSignal(func() uint16 { return ^c.signals[s[1]].val() })
			}
		case 3: // other operators
			var left, right func() uint16

			if i, err := strconv.Atoi(s[0]); err == nil {
				left = func() uint16 { return uint16(i) }
			} else {
				left = func() uint16 { return c.signals[s[0]].val() }
			}

			if i, err := strconv.Atoi(s[2]); err == nil {
				right = func() uint16 { return uint16(i) }
			} else {
				right = func() uint16 { return c.signals[s[2]].val() }
			}

			switch s[1] {
			case "AND":
				c.signals[target] = newSignal(func() uint16 { return left() & right() })
			case "OR":
				c.signals[target] = newSignal(func() uint16 { return left() | right() })
			case "LSHIFT":
				c.signals[target] = newSignal(func() uint16 { return left() << right() })
			case "RSHIFT":
				c.signals[target] = newSignal(func() uint16 { return left() >> right() })
			}
		}
	}

	return c
}

func (c *circuit) measure(wire string) uint16 {
	if s, ok := c.signals[wire]; ok {
		return s.val()
	}
	return 0
}

func (c *circuit) measureAll() (signals map[string]uint16) {
	signals = make(map[string]uint16)

	for wire, s := range c.signals {
		signals[wire] = s.val()
	}

	return signals
}

type signal struct {
	f func() uint16
	v uint16
}

func newSignal(f func() uint16) *signal {
	return &signal{f: f}
}

func (s *signal) val() uint16 {
	if s.f != nil {
		s.v, s.f = s.f(), nil
	}

	return s.v
}
