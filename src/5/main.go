package main

import (
    "fmt"
    "regexp"
    "strconv"
    "../utils"
)

type Crates [][]string

type Move struct {
    cnt, from, to int
}

func (m Move) apply(crates Crates, reversed bool) {
    lenFrom := len(crates[m.from - 1])
    leftFrom, moved := crates[m.from - 1][:lenFrom - m.cnt], crates[m.from - 1][lenFrom - m.cnt:]

    if reversed {
        for i, j := 0, len(moved) - 1; i < j; i, j = i + 1, j - 1 {
            moved[i], moved[j] = moved[j], moved[i]
        }
    }

    crates[m.from - 1] = leftFrom
    crates[m.to - 1] = append(crates[m.to - 1], moved...)
}

func parseCrates(desc []string) Crates {
    lastLineLen := len(desc[len(desc)-1])
    crates := make([][]string, (lastLineLen+3)/4)

    for k := range desc {
        line := desc[len(desc)-k-1]

        for i, j := 0, 2; j < len(line); j += 4 {

            lbra, label, rbra := line[j-2], line[j-1], line[j]

            if lbra == '[' && rbra == ']' {
                crates[i] = append(crates[i], string(label))
            }

            i++
        }
    }

    return Crates(crates)
}

func parseMoves(desc []string) []Move {
    re := regexp.MustCompile(`move ([[:digit:]]+) from ([[:digit:]]+) to ([[:digit:]]+)`)
    moves := make([]Move, len(desc))

    for i, line := range desc {
        match := re.FindStringSubmatch(line)

        cnt,_   := strconv.Atoi(match[1])
        from, _ := strconv.Atoi(match[2])
        to, _   :=  strconv.Atoi(match[3])

        moves[i] = Move{cnt, from, to}
    }

    return moves
}

func parseInput(desc []string) (crates Crates, moves []Move){
    sepLine := 0

    for i, line := range desc {
        if line == "" {
            sepLine = i
            break
        }
    }

    crates = parseCrates(desc[:sepLine])
    moves = parseMoves(desc[sepLine+1:])
    return
}

func solve(desc []string, reversed bool) {
    crates, moves := parseInput(desc)

    for _, move := range moves {
        move.apply(crates, reversed)
    }

    tops := ""

    for _, stack := range crates {
        tops += stack[len(stack) - 1]
    }

    fmt.Println(tops)
}

func partOne(desc []string) {
    solve(desc, true)
}

func partTwo(desc []string) {
    solve(desc, false)
}

func tests() {
    tests := utils.ReadLines("tests", "5.txt")
    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "5.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
