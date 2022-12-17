package main

import (
    "fmt"
    "../utils"
)

var shapes = [][][]int {
    {
        {1, 1, 1, 1},
    },

    {
        {0, 1, 0},
        {1, 1, 1},
        {0, 1, 0},
    },

    {
        {1, 1, 1},
        {0, 0, 1},
        {0, 0, 1},
    },

    {
        {1},
        {1},
        {1},
        {1},
    },

    {
        {1, 1},
        {1, 1},
    },
}

type state [][]bool

func (s state) printOut() {
    for y := len(s)-1; y >= 0; y-- {
        for _, f := range s[y] {
            if f {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }

        fmt.Printf("\n")
    }
}

func (s state) filledRows() int {
    for y := len(s)-1; y >= 0; y-- {
        for _, f := range s[y] {
            if f {
                return y+1
            }
        }
    }

    return 0
}

func (s *state) assureHeight(width, height int) {
    for len(*s) < height {
        *s = append(*s, make([]bool, width))
    }
}

func (s *state) hitsShape(shape [][]int, width, lx, by int) bool {
    for y, row := range shape {
        for x, f := range row {
            if f == 1 && (by + y < 0 || lx + x < 0 || lx + x >= width || (*s)[by + y][lx + x]) {
                return true;
            }
        }
    }

    return false
}

func (s *state) applyShape(shape [][]int, lx, by int) {
    for y, row := range shape {
        for x, f := range row {
            if f == 1 {
                (*s)[by + y][lx + x] = true;
            }
        }

    }
}

func (s *state) tryMoveShape(shape [][]int, width, lx, ty, dx, dy int) (nx, ny int, moved bool) {
    if !s.hitsShape(shape, width, lx+dx, ty+dy) {
        nx, ny, moved = lx+dx, ty+dy, true

    } else {
        nx, ny, moved = lx, ty, false

    }

    return
}

func simulate(width, dx, dy int, jets string, cb func (i, j int, state state) bool) {
    state := state(make([][]bool, 0))

    i, j := 0, 0

    for cb(i, j, state) {
        topY := state.filledRows()
        shape := shapes[i % len(shapes)]

        state.assureHeight(width, topY + dy + len(shape))

        cx, cy := dx, topY + dy

        for {
            jet := jets[j % len(jets)]
            j++

            if jet == '<' {
                cx, cy, _ = state.tryMoveShape(shape, width, cx, cy, -1, 0)
            } else {
                cx, cy, _ = state.tryMoveShape(shape, width, cx, cy, +1, 0)
            }

            var moved bool
            cx, cy, moved = state.tryMoveShape(shape, width, cx, cy, 0, -1)

            if !moved {
                state.applyShape(shape, cx, cy)
                break
            }
        }

        i++
    }
}

func partOne(lines []string) {
    insn := lines[0]

    var ans int
    simulate(7, 2, 3, insn, func (i, j int, state state) bool {
        if i == 2022 {
            ans = state.filledRows()
            return false
        }

        return true
    })

    fmt.Println(ans)
}

const keyRows = 50

type memoRow struct {
    i, height int
}

func partTwo(lines []string) {
    insn := lines[0]

    heights := make([]int, 0, 200)
    var prefix, cycle memoRow

    memo := make(map[string]memoRow)
    simulate(7, 2, 3, insn, func (i, j int, state state) bool {
        topFilled := state.filledRows()
        heights = append(heights, topFilled)

        key := ""
        for y := topFilled - 1; y >= topFilled - keyRows; y-- {
            if y >= 0 {
                key += fmt.Sprintf("%v;", state[y])
            } else {
                key += fmt.Sprintf("bedrock;")
            }
        }

        key += fmt.Sprintf("###;%d;%d", i % len(shapes), j % len(insn))

        mr, isCycle := memo[key]

        if isCycle {
            prefix = mr
            cycle = memoRow{ i - mr.i, topFilled - mr.height }
            return false
        }

        memo[key] = memoRow{ i, topFilled }
        return true
    })

    cnt := 1000000000000
    rep := cnt - prefix.i
    div, rem := rep / cycle.i, rep % cycle.i
    ans := prefix.height + div * cycle.height
    ans += heights[len(heights) - cycle.i + rem - 1] - heights[len(heights) - cycle.i - 1]

    fmt.Println(ans)
}

func tests() {
    tests := utils.ReadLines("tests", "17.txt")
    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "17.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
