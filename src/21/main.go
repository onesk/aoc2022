package main

import (
    "fmt"
	"math/big"
	"strings"
	"strconv"
	"../utils"
)

type monkey struct {
	isValue bool
	lhs, rhs, op string
	value int
}

func parse(lines []string) map[string]monkey {
	ret := make(map[string]monkey)

	for _, line := range lines {
		parts := strings.SplitN(line, ": ", 2)
		name, ops := parts[0], strings.SplitN(parts[1], " ", 3)
		value, err := strconv.Atoi(parts[1])

		var m monkey
		if err == nil {
			m = monkey{true, "", "", "", value}

		} else {
			lhs, op, rhs := ops[0], ops[1], ops[2]
			m = monkey{false, lhs, rhs, op, -1}

		}

		ret[name] = m
	}

	return ret
}

func eval(ms map[string]monkey, name string) *big.Rat {
	m := ms[name]

	if m.isValue {
		return big.NewRat(int64(m.value), 1)
	}

	lhs := eval(ms, m.lhs)
	rhs := eval(ms, m.rhs)

	switch m.op {
	case "+":
		lhs.Add(lhs, rhs)
	case "-":
		lhs.Sub(lhs, rhs)
	case "*":
		lhs.Mul(lhs, rhs)
	case "/":
		lhs.Mul(lhs, rhs.Inv(rhs))
	}

	return lhs
}

func find(ms map[string]monkey, name, needle string) []string {
	m := ms[name]

	if name == needle {
		return []string{name}
	}

	if m.isValue {
		return nil
	}


	if lhsv := find(ms, m.lhs, needle); lhsv != nil {
		return append([]string{name}, lhsv...)
	}

	if rhsv := find(ms, m.rhs, needle); rhsv != nil {
		return append([]string{name}, rhsv...)
	}

	return nil
}

func solveFor(ms map[string]monkey, tgt *big.Rat, name string, path []string) *big.Rat {
	if len(path) == 0 {
		return tgt
	}

	m := ms[name]
	isLhs := m.lhs == path[0]

	if m.op == "+" && isLhs {
		return solveFor(ms, tgt.Sub(tgt, eval(ms, m.rhs)), m.lhs, path[1:])
	}

	if m.op == "+" && !isLhs {
		return solveFor(ms, tgt.Sub(tgt, eval(ms, m.lhs)), m.rhs, path[1:])
	}

	if m.op == "-" && isLhs {
		return solveFor(ms, tgt.Add(tgt, eval(ms, m.rhs)), m.lhs, path[1:])
	}

	if m.op == "-" && !isLhs {
		lhs := eval(ms, m.lhs)
		return solveFor(ms, lhs.Sub(lhs, tgt), m.rhs, path[1:])
	}

	if m.op == "*" && isLhs {
		rhs := eval(ms, m.rhs)
		return solveFor(ms, tgt.Mul(tgt, rhs.Inv(rhs)), m.lhs, path[1:])
	}

	if m.op == "*" && !isLhs {
		lhs := eval(ms, m.lhs)
		return solveFor(ms, tgt.Mul(tgt, lhs.Inv(lhs)), m.rhs, path[1:])
	}

	if m.op == "/" && isLhs {
		rhs := eval(ms, m.rhs)
		return solveFor(ms, tgt.Mul(tgt, rhs), m.lhs, path[1:])
	}

	if m.op == "/" && !isLhs {
		lhs := eval(ms, m.lhs)
		return solveFor(ms, lhs.Mul(lhs, tgt.Inv(tgt)), m.rhs, path[1:])
	}

	return nil
}

func partOne(lines []string) {
	monkeys := parse(lines)
	ans := eval(monkeys, "root")
	fmt.Println(ans)
}

func partTwo(lines []string) {
	monkeys := parse(lines)
	root, path := monkeys["root"], find(monkeys, "root", "humn")

	var humnv *big.Rat
	if root.lhs == path[1] {
		humnv = solveFor(monkeys, eval(monkeys, root.rhs), root.lhs, path[2:])
	} else {
		humnv = solveFor(monkeys, eval(monkeys, root.lhs), root.rhs, path[2:])
	}

	fmt.Println(humnv)
}

func tests() {
    tests := utils.ReadLines("tests", "21.txt")
    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "21.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
