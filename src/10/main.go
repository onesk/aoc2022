package main

import (
    "fmt"
    "strings"
    "strconv"
    "../utils"
)

type cpu struct {
    tick, x int
    cb func (tick, x int)
}

func newCpu(cb func (tick, x int)) cpu {
    return cpu{1, 1, cb}
}

func (c *cpu) tryCb() {
    c.cb(c.tick, c.x)
}

func (c *cpu) noop() {
    c.tryCb()
    c.tick++
}

func (c *cpu) addx(x int) {
    c.tryCb()
    c.tick++

    c.tryCb()
    c.tick++

    c.x += x
}

func simulate(cmds []string, cb func (tick, x int)) {
    cpu := newCpu(cb)

    for _, line := range cmds {
        parts := strings.SplitN(line, " ", 2)

        switch parts[0] {
        case "noop":
            cpu.noop()

        case "addx":
            x, _ := strconv.Atoi(parts[1])
            cpu.addx(x)
        }
    }
}

func partOne(cmds []string) {
    ans := 0

    simulate(cmds, func (tick, x int) {
        if tick % 40 == 20 {
            ans += tick * x
        }
    })

    fmt.Println(ans)
}

func partTwo(cmds []string) {
    var crt [6][40]bool

    simulate(cmds, func (tick, x int) {
        j := (tick - 1) % 40
        i := ((tick - 1) / 40) % 6

        if x - 1 <= j && j <= x + 1 {
            crt[i][j] = true
        }
    })

    for _, row := range crt {
        s := ""
        for _, f := range row {
            if f {
                s += "#"
            } else {
                s += "."
            }
        }

        fmt.Println(s)
    }
}

func tests() {
    tests1 := utils.ReadLines("tests", "10_1.txt")
    partOne(tests1)
    partTwo(tests1)


    tests2 := utils.ReadLines("tests", "10_2.txt")
    partOne(tests2)
    partTwo(tests2)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "10.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
