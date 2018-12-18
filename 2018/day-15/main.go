package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	var input [][]byte

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Bytes())
	}

	c := newCave(input)

	outcome, _ := fight(c)
	outcome2 := replayFightUntilElvesStrongEnough(input)

	fmt.Printf("Part 1: %d\n", outcome)
	fmt.Printf("Part 2: %d\n", outcome2)
}

func replayFightUntilElvesStrongEnough(input [][]byte) (outcome int) {
	for elfAttackPower := 0; ; elfAttackPower++ {
		c := newCave(input)

		for _, u := range c.Units {
			if u.Kind == 'E' {
				u.AP = elfAttackPower
			}
		}

		outcome, elvesDead := fight(c)

		if elvesDead == 0 {
			return outcome
		}
	}
}

func (c *Cave) FirstEnemyNearLocation(me *Unit, loc Coord) *Unit {
	adjC := loc.Adjacent()
	sort.Slice(adjC, func(i, j int) bool {
		if adjC[i].Y == adjC[j].Y {
			return adjC[i].X < adjC[j].X
		}
		return adjC[i].Y < adjC[j].Y
	})

	var candidates []*Unit
	for n := range adjC {
		if adjC[n] == me.Location {
			continue
		}
		for _, adjU := range c.Units {
			if adjU.HP <= 0 {
				continue
			}
			if adjU.Location == adjC[n] {
				if adjU.Kind != me.Kind {
					candidates = append(candidates, adjU)
					break
				}
			}
		}
	}
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].HP == candidates[j].HP {
			if candidates[i].Location.Y == candidates[j].Location.Y {
				return candidates[i].Location.X < candidates[j].Location.X
			}
			return candidates[i].Location.Y < candidates[j].Location.Y
		}
		return candidates[i].HP < candidates[j].HP
	})
	if len(candidates) > 0 {
		return candidates[0]
	}

	return nil
}

func fight(c *Cave) (outcome, elvesDead int) {
	var round int

	for {
		c.SortUnits()

		unitsLooped := 0

		for _, unit := range c.Units {
			unitsLooped++

			if unit.HP <= 0 {
				continue
			}

			attackableEnemy := c.FirstEnemyNearLocation(unit, unit.Location)

			if attackableEnemy != nil {
				attackableEnemy.HP -= unit.AP
				if attackableEnemy.HP <= 0 {
					c.Open[attackableEnemy.Location] = struct{}{}
				}
			} else {
				lnes := c.LocationsNearEnemies(unit)
				dm := c.DistanceMap(unit.Location)
				reachableLnes := make([]Coord, 0, len(lnes))
				for _, lne := range lnes {
					if _, ok := dm[lne]; ok {
						reachableLnes = append(reachableLnes, lne)
					}
				}
				sort.Slice(reachableLnes, func(i, j int) bool {
					if dm[reachableLnes[i]] == dm[reachableLnes[j]] {
						if reachableLnes[i].Y == reachableLnes[j].Y {
							return reachableLnes[i].X < reachableLnes[j].X
						}
						return reachableLnes[i].Y < reachableLnes[j].Y
					}
					return dm[reachableLnes[i]] < dm[reachableLnes[j]]
				})

				if len(reachableLnes) == 0 {
					continue
				}

				path := findPath(reachableLnes[0], dm)
				c.Open[unit.Location] = struct{}{}
				unit.Location = path[0]
				delete(c.Open, unit.Location)

				if attackableEnemy := c.FirstEnemyNearLocation(unit, unit.Location); attackableEnemy != nil {
					attackableEnemy.HP -= unit.AP
					if attackableEnemy.HP <= 0 {
						c.Open[attackableEnemy.Location] = struct{}{}
					}
				}
			}

			var elvesExist, goblinsExist bool
			for _, u := range c.Units {
				if u.HP > 0 {
					switch u.Kind {
					case 'E':
						elvesExist = true
					case 'G':
						goblinsExist = true
					}
				}
			}

			if !elvesExist || !goblinsExist {
				break
			}
		}

		if unitsLooped == len(c.Units) {
			round++
		}

		var elvesDead int
		var elfHP, goblinHP int
		for _, u := range c.Units {
			if u.HP > 0 {
				switch u.Kind {
				case 'E':
					elfHP += u.HP
				case 'G':
					goblinHP += u.HP
				}
			} else {
				if u.Kind == 'E' {
					elvesDead++
				}
			}
		}

		if elfHP == 0 || goblinHP == 0 {
			var totalHP int
			if elfHP > 0 {
				totalHP = elfHP
			} else if goblinHP > 0 {
				totalHP = goblinHP
			}

			return round * totalHP, elvesDead
		}
	}

	return 0, 0
}

func findPath(to Coord, distanceMap map[Coord]int) []Coord {
	revPath := []Coord{to}

	coord := to

	for {
		dist, ok := distanceMap[coord]
		if !ok {
			break
		}

		adjC := coord.Adjacent()

		reachableAdjC := make([]Coord, 0, len(adjC))
		for n := range adjC {
			if _, ok := distanceMap[adjC[n]]; ok {
				reachableAdjC = append(reachableAdjC, adjC[n])
			}
		}

		sort.Slice(reachableAdjC, func(i, j int) bool {
			if distanceMap[reachableAdjC[i]] == distanceMap[reachableAdjC[j]] {
				if reachableAdjC[i].Y == reachableAdjC[j].Y {
					return reachableAdjC[i].X < reachableAdjC[j].X
				}
				return reachableAdjC[i].Y < reachableAdjC[j].Y
			}
			return distanceMap[reachableAdjC[i]] < distanceMap[reachableAdjC[j]]
		})

		if dist == 1 {
			break
		}

		coord = reachableAdjC[0]
		revPath = append(revPath, coord)
	}

	path := make([]Coord, 0, len(revPath))
	for n := len(revPath) - 1; n >= 0; n-- {
		path = append(path, revPath[n])
	}

	return path
}

type Coord struct {
	X, Y int
}

type Unit struct {
	AP       int
	HP       int
	Kind     byte
	Location Coord
}

func (c Coord) Adjacent() []Coord {
	cc := make([]Coord, 0, 4)
	cc = append(cc, Coord{c.X - 1, c.Y})
	cc = append(cc, Coord{c.X + 1, c.Y})
	cc = append(cc, Coord{c.X, c.Y - 1})
	cc = append(cc, Coord{c.X, c.Y + 1})
	return cc
}

type Cave struct {
	Open  map[Coord]struct{}
	Walls map[Coord]struct{}
	Units []*Unit
}

func newCave(plan [][]byte) *Cave {
	c := &Cave{
		Open:  make(map[Coord]struct{}),
		Walls: make(map[Coord]struct{}),
		Units: make([]*Unit, 0),
	}

	for y := range plan {
		for x := range plan[y] {
			switch plan[y][x] {
			case '.':
				c.Open[Coord{x, y}] = struct{}{}
			case '#':
				c.Walls[Coord{x, y}] = struct{}{}
			case 'E', 'G':
				c.Units = append(c.Units,
					&Unit{AP: 3, HP: 200, Kind: plan[y][x], Location: Coord{x, y}},
				)
			}
		}
	}

	return c
}

// LocationsNearEnemies returns a list of coordinates adjacent to enemies.
func (c *Cave) LocationsNearEnemies(me *Unit) []Coord {
	coordMap := map[Coord]struct{}{}

	for _, u := range c.Units {
		if u.HP <= 0 {
			continue
		}
		if u.Kind == me.Kind {
			continue
		}
		if u.Location == me.Location {
			continue
		}
		for _, coord := range u.Location.Adjacent() {
			if _, ok := c.Open[coord]; ok {
				coordMap[coord] = struct{}{}
			}
		}
	}

	cc := make([]Coord, 0, len(coordMap))

	for c, _ := range coordMap {
		cc = append(cc, c)
	}

	sort.Slice(cc, func(i, j int) bool {
		if cc[i].Y == cc[j].Y {
			return cc[i].X < cc[j].X
		}
		return cc[i].Y < cc[j].Y
	})

	return cc
}

func (c *Cave) DistanceMap(from Coord) map[Coord]int {
	distM := map[Coord]int{from: 0}

	var recurse func(Coord, int)

	recurse = func(rFrom Coord, curDist int) {
		for _, adjCoord := range rFrom.Adjacent() {
			if adjCoord == from {
				continue
			}
			if _, ok := c.Open[adjCoord]; !ok {
				continue
			}
			if d, ok := distM[adjCoord]; ok {
				if curDist < d {
					distM[adjCoord] = curDist
					recurse(adjCoord, curDist+1)
				}
			} else {
				distM[adjCoord] = curDist
				recurse(adjCoord, curDist+1)
			}
		}
	}

	recurse(from, 1)

	return distM
}

func (c *Cave) SortUnits() {
	sort.Slice(c.Units, func(i, j int) bool {
		if c.Units[i].Location.Y == c.Units[j].Location.Y {
			return c.Units[i].Location.X < c.Units[j].Location.X
		}
		return c.Units[i].Location.Y < c.Units[j].Location.Y
	})
}

func (c *Cave) String() string {
	var maxX, maxY int

	for wall := range c.Walls {
		if wall.X > maxX {
			maxX = wall.X
		}
		if wall.Y > maxY {
			maxY = wall.Y
		}
	}

	units := map[Coord]*Unit{}
	for _, u := range c.Units {
		if u.HP > 0 {
			units[u.Location] = u
		}
	}
	buf := new(bytes.Buffer)

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if u, ok := units[Coord{x, y}]; ok {
				buf.WriteByte(u.Kind)
				continue
			}
			if _, ok := c.Walls[Coord{x, y}]; ok {
				buf.WriteByte('#')
				continue
			}

			if _, ok := c.Open[Coord{x, y}]; ok {
				buf.WriteByte('.')
				continue
			}

			buf.WriteByte(' ')
		}
		fmt.Fprintln(buf)
	}

	return buf.String()
}
