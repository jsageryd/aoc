package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	enemy := parse(input)

	minCost := 10000

	for _, c := range configurations() {
		player := contender{hitPoints: 100, damage: c.damage, armor: c.armor}
		boss := enemy

		if playerWins(&player, &boss) && c.cost < minCost {
			minCost = c.cost
		}
	}

	return minCost
}

func parse(input []string) contender {
	var c contender

	fmt.Sscanf(input[0], "Hit Points: %d", &c.hitPoints)
	fmt.Sscanf(input[1], "Damage: %d", &c.damage)
	fmt.Sscanf(input[2], "Armor: %d", &c.armor)

	return c
}

type contender struct {
	hitPoints int
	damage    int
	armor     int
}

type conf struct {
	cost   int
	damage int
	armor  int
}

func (c conf) add(other conf) conf {
	return conf{
		cost:   c.cost + other.cost,
		damage: c.damage + other.damage,
		armor:  c.armor + other.armor,
	}
}

func configurations() []conf {
	weapons := []conf{
		{cost: 8, damage: 4, armor: 0},  // dagger
		{cost: 10, damage: 5, armor: 0}, // shortsword
		{cost: 25, damage: 6, armor: 0}, // warhammer
		{cost: 40, damage: 7, armor: 0}, // longsword
		{cost: 74, damage: 8, armor: 0}, // greataxe
	}

	armor := []conf{
		{cost: 13, damage: 0, armor: 1},  // leather
		{cost: 31, damage: 0, armor: 2},  // chainmail
		{cost: 53, damage: 0, armor: 3},  // splintmail
		{cost: 75, damage: 0, armor: 4},  // bandedmail
		{cost: 102, damage: 0, armor: 5}, // platemail
	}

	rings := []conf{
		{cost: 25, damage: 1, armor: 0},  // damage +1
		{cost: 50, damage: 2, armor: 0},  // damage +2
		{cost: 100, damage: 3, armor: 0}, // damage +3
		{cost: 20, damage: 0, armor: 1},  // defense +1
		{cost: 40, damage: 0, armor: 2},  // defense +2
		{cost: 80, damage: 0, armor: 3},  // defense +3
	}

	var cs []conf

	for _, w := range weapons {
		// no armor, no rings
		cs = append(cs, w)

		// just armor
		for _, a := range armor {
			cs = append(cs, w.add(a))
		}

		// just one ring
		for _, r := range rings {
			cs = append(cs, w.add(r))
		}

		// armor and one ring
		for _, a := range armor {
			for _, r := range rings {
				cs = append(cs, w.add(a).add(r))
			}
		}

		// armor and two rings
		for _, a := range armor {
			for _, r1 := range rings {
				for _, r2 := range rings {
					if r1 != r2 {
						cs = append(cs, w.add(a).add(r1).add(r2))
					}
				}
			}
		}

		// two rings
		for _, r1 := range rings {
			for _, r2 := range rings {
				if r1 != r2 {
					cs = append(cs, w.add(r1).add(r2))
				}
			}
		}
	}

	return cs
}

// playerWins runs the fight and returns true if player wins.
func playerWins(player, enemy *contender) bool {
	cs := []*contender{player, enemy}

	for n := 0; ; n++ {
		a, b := cs[n%2], cs[(n+1)%2]
		b.hitPoints -= a.damage - b.armor
		if b.hitPoints <= 0 {
			break
		}
	}

	return enemy.hitPoints <= 0
}
