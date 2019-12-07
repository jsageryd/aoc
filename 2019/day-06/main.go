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

	com := parse(input)

	fmt.Printf("Part 1: %d\n", totalOrbits(com))
}

func totalOrbits(o *Object) int {
	var total int

	var f func(o *Object, depth int)

	f = func(o *Object, depth int) {
		total += depth
		for _, orbit := range o.Orbits {
			f(orbit, depth+1)
		}
	}

	f(o, 0)

	return total
}

func parse(input []string) *Object {
	m := make(map[string]*Object)

	for n := range input {
		pair := strings.Split(input[n], ")")
		object, ok := m[pair[0]]
		if !ok {
			object = &Object{Name: pair[0]}
			m[object.Name] = object
		}
		orbit, ok := m[pair[1]]
		if !ok {
			orbit = &Object{Name: pair[1]}
			m[orbit.Name] = orbit
		}
		object.Orbits = append(object.Orbits, orbit)
	}

	for _, v := range m {
		sort.Slice(v.Orbits, func(i, j int) bool {
			return v.Orbits[i].Name < v.Orbits[j].Name
		})
	}

	return m["COM"]
}

type Object struct {
	Name   string
	Orbits []*Object
}

func (o *Object) String() string {
	orbitStrs := make([]string, 0, len(o.Orbits))
	for _, orbit := range o.Orbits {
		orbitStrs = append(orbitStrs, orbit.String())
	}
	if len(orbitStrs) > 0 {
		return fmt.Sprintf("%s(%s)", o.Name, strings.Join(orbitStrs, ","))
	}
	return o.Name
}
