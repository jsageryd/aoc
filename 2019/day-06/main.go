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
	fmt.Printf("Part 2: %d\n", orbitalTransfers(find(com, "YOU").Parent, find(com, "SAN").Parent))
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

func orbitalTransfers(from, to *Object) int {
	var base *Object

	seen := make(map[*Object]struct{})

	for obj := from; obj != nil; obj = obj.Parent {
		seen[obj] = struct{}{}
	}

	for obj := to; obj != nil; obj = obj.Parent {
		if _, ok := seen[obj]; ok {
			base = obj
			break
		}
	}

	if base == nil {
		return -1
	}

	var transfers int

	for obj := from; obj != base; obj = obj.Parent {
		transfers++
	}

	for obj := to; obj != base; obj = obj.Parent {
		transfers++
	}

	return transfers
}

func find(root *Object, objName string) *Object {
	var f func(o *Object) *Object

	f = func(o *Object) *Object {
		if o.Name == objName {
			return o
		}
		for _, orbit := range o.Orbits {
			if o := f(orbit); o != nil {
				return o
			}
		}
		return nil
	}

	return f(root)
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
		orbit.Parent = object
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
	Parent *Object
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
