package main

import (
    "fmt"
    "../utils"
)

type coord struct {
    x, y int
}

func (c coord) key() string {
    return fmt.Sprintf("%v", c)
}

func parse(lines []string) []coord {
    ret := make([]coord, 0)

    for y, row := range lines {
        for x, c := range row {
            if c == '#' {
                ret = append(ret, coord{x, y})
            }
        }
    }

    return ret
}

type rule struct {
    dir coord
    probes []coord
}

var rules = [4]rule {
    {coord{0, -1}, []coord{{-1, -1}, {0, -1}, {+1, -1}}}, // north - N, NE, NW
    {coord{0, +1}, []coord{{-1, +1}, {0, +1}, {+1, +1}}}, // south - S, SE, SW
    {coord{-1, 0}, []coord{{-1, -1}, {-1, 0}, {-1, +1}}}, // west  - W, NW, SW
    {coord{+1, 0}, []coord{{+1, -1}, {+1, 0}, {+1, +1}}}, // east  - E, NE, SE
}

func simulate(es []coord, delta int) int {
    occ := make(map[string]bool)

    for _, e := range es {
        occ[e.key()] = true
    }

    props := make([]coord, len(es))
    copy(props, es)

    for i, e := range es {
        other := false

        for dx := -1; dx <= 1; dx++ {
            for dy := -1; dy <= 1; dy++ {
                if dx == 0 && dy == 0 {
                    continue
                }

                if occ[coord{e.x + dx, e.y + dy}.key()] {
                    other = true
                }
            }
        }

        if !other {
            continue
        }

        for j := 0; j < 4; j++ {
            r := rules[(j + delta) % 4]

            allFree := true
            for _, prb := range r.probes {
                if occ[coord{e.x + prb.x, e.y + prb.y}.key()] {
                    allFree = false
                }
            }

            if allFree {
                props[i] = coord{e.x + r.dir.x, e.y + r.dir.y}
                break
            }
        }
    }

    dedup := make(map[string]int)

    for _, p := range props {
        dedup[p.key()]++
    }

    moved := 0

    for i, p := range props {
        if dedup[p.key()] == 1 {
            if p != es[i] {
                moved++
            }

            es[i] = p
        }
    }

    return moved
}

type aabb struct {
    min, max coord
}

func (c aabb) area() int {
    return (c.max.x - c.min.x + 1) * (c.max.y - c.min.y + 1)
}

const inf = 1000000000

func getAabb(es []coord) aabb {
    min, max := coord{inf, inf}, coord{-inf, -inf}

    for _, e := range es {
        min.x, min.y = utils.Min(min.x, e.x), utils.Min(min.y, e.y)
        max.x, max.y = utils.Max(max.x, e.x), utils.Max(max.y, e.y)
    }

    return aabb{min, max}
}

func partOne(lines []string) {
    elves := parse(lines)

    for i := 0; i < 10; i++ {
        simulate(elves, i)
    }

    aabb := getAabb(elves)

    ans := aabb.area() - len(elves)
    fmt.Println(ans)
}

func partTwo(lines []string) {
    elves := parse(lines)

    ans := 0
    for simulate(elves, ans) > 0 {
        ans++
    }

    fmt.Println(ans + 1)
}

func tests() {
    tests := utils.ReadLines("tests", "23.txt")
    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "23.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
