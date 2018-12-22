package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var count int
	opsMap := map[int]map[string]struct{}{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		if strings.HasPrefix(str, "Before") {
			var before Reg
			fmt.Sscanf(str, "Before: [%d, %d, %d, %d]", &before[0], &before[1], &before[2], &before[3])

			scanner.Scan()
			str = scanner.Text()
			inst := parseInst(str)

			scanner.Scan()
			str = scanner.Text()
			var after Reg
			fmt.Sscanf(str, "After: [%d, %d, %d, %d]", &after[0], &after[1], &after[2], &after[3])

			mOps := matchingOps(inst, before, after)

			if len(mOps) >= 3 {
				count++
			}

			if _, ok := opsMap[inst.Op]; !ok {
				opsMap[inst.Op] = make(map[string]struct{})
			}

			for _, mOp := range mOps {
				opsMap[inst.Op][mOp] = struct{}{}
			}

			scanner.Scan()
		} else {
			break
		}
	}

	knownOps := map[int]string{}

	for len(knownOps) < len(ops) {
		for op, candidates := range opsMap {
			if len(candidates) == 1 {
				var opStr string
				for c := range candidates {
					opStr = c
				}
				knownOps[op] = opStr
				for op := range opsMap {
					delete(opsMap[op], opStr)
				}
				delete(opsMap, op)
				break
			}
		}
	}

	var reg Reg

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		inst := parseInst(scanner.Text())

		ops[knownOps[inst.Op]](&reg, inst.A, inst.B, inst.C)
	}

	fmt.Printf("Part 1: %d\n", count)
	fmt.Printf("Part 2: %d\n", reg[0])
}

func matchingOps(inst Inst, before, after Reg) []string {
	var mOps []string

	for op, f := range ops {
		reg := before
		f(&reg, inst.A, inst.B, inst.C)
		if reg == after {
			mOps = append(mOps, op)
		}
	}

	sort.Strings(mOps)

	return mOps
}

type Inst struct {
	Op int
	A  int
	B  int
	C  int
}

func parseInst(s string) Inst {
	var inst Inst
	fmt.Sscanf(s, "%d %d %d %d", &inst.Op, &inst.A, &inst.B, &inst.C)
	return inst
}

type Reg [4]int

var ops = map[string]func(reg *Reg, a, b, c int){
	"addr": func(reg *Reg, a, b, c int) { reg[c] = reg[a] + reg[b] },
	"addi": func(reg *Reg, a, b, c int) { reg[c] = reg[a] + b },
	"mulr": func(reg *Reg, a, b, c int) { reg[c] = reg[a] * reg[b] },
	"muli": func(reg *Reg, a, b, c int) { reg[c] = reg[a] * b },
	"banr": func(reg *Reg, a, b, c int) { reg[c] = reg[a] & reg[b] },
	"bani": func(reg *Reg, a, b, c int) { reg[c] = reg[a] & b },
	"borr": func(reg *Reg, a, b, c int) { reg[c] = reg[a] | reg[b] },
	"bori": func(reg *Reg, a, b, c int) { reg[c] = reg[a] | b },
	"setr": func(reg *Reg, a, b, c int) { reg[c] = reg[a] },
	"seti": func(reg *Reg, a, b, c int) { reg[c] = a },
	"gtir": func(reg *Reg, a, b, c int) {
		if a > reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	},
	"gtri": func(reg *Reg, a, b, c int) {
		if reg[a] > b {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	},
	"gtrr": func(reg *Reg, a, b, c int) {
		if reg[a] > reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	},
	"eqir": func(reg *Reg, a, b, c int) {
		if a == reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	},
	"eqri": func(reg *Reg, a, b, c int) {
		if reg[a] == b {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	},
	"eqrr": func(reg *Reg, a, b, c int) {
		if reg[a] == reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	},
}
