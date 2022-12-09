package main

import (
    "fmt"
    "strings"
    "strconv"
    "../utils"
)

type coord struct {
    x, y int
}

func (c *coord) moveDir(dir coord) {
    c.x += dir.x
    c.y += dir.y
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}

func sgn(x int) int {
    if x < 0 {
        return -1
    }

    if x > 0 {
        return 1
    }

    return 0
}

func (tail *coord) followHead(head coord) {
    dx, dy := head.x - tail.x, head.y - tail.y

    if abs(dx) <= 1 && abs(dy) <= 1 {
        return
    }

    tail.x += sgn(dx)
    tail.y += sgn(dy)
}

func (c coord) asString() string {
    return fmt.Sprintf("%v", c)
}

type move struct {
    dir coord
    reps int
}

var dirs =  map[string]coord {
    "U": coord{0, +1},
    "D": coord{0, -1},
    "L": coord{-1, 0},
    "R": coord{+1, 0},
}

func parse(cmds []string) []move {
    ret := make([]move, 0, len(cmds))

    for _, line := range cmds {
        parts := strings.SplitN(line, " ", 2)

        dir, _ := dirs[parts[0]]
        reps, _ := strconv.Atoi(parts[1])

        ret = append(ret, move{dir, reps})
    }

    return ret
}

func simulate(moves []move, tails int) int {
    visited := make(map[string]bool)

    rope := make([]coord, tails+1)
    for i := 0; i < len(rope); i++ {
        rope[i] = coord{0, 0}
    }

    visited[coord{0, 0}.asString()] = true

    for _, move := range moves {
        for i := 0; i < move.reps; i++ {
            rope[0].moveDir(move.dir)

            for j := 1; j < len(rope); j++ {
                rope[j].followHead(rope[j-1])
            }

            visited[rope[len(rope)-1].asString()] = true
        }
    }

    return len(visited)
}

func partOne(cmds []string) {
    moves := parse(cmds)
    visited := simulate(moves, 1)
    fmt.Println(visited)
}

func partTwo(cmds []string) {
    moves := parse(cmds)
    visited := simulate(moves, 9)
    fmt.Println(visited)
}

func tests() {
    tests := utils.ReadLines("tests", "9.txt")
    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "9.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
