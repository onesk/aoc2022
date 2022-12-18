package main

import (
    "fmt"
    "strings"
    "strconv"
    "../utils"
)

type coord struct {
    x, y, z int
}

func (s coord) key() string {
    return fmt.Sprintf("%v;%v;%v", s.x, s.y, s.z)
}

var adjs = [6]coord{
    {-1, 0, 0},
    {+1, 0, 0},
    {0, -1, 0},
    {0, +1, 0},
    {0, 0, -1},
    {0, 0, +1},
}

func parse(lines []string) []coord {
    ret := make([]coord, 0, len(lines))

    for _, line := range lines {
        parts := strings.Split(line, ",")

        x, _ := strconv.Atoi(parts[0])
        y, _ := strconv.Atoi(parts[1])
        z, _ := strconv.Atoi(parts[2])

        ret = append(ret, coord{x, y, z})
    }

    return ret
}

const inf = 1000000000

func dfs(space map[string]bool, min, max, c coord) {
    if c.x < min.x || c.x > max.x || c.y < min.y || c.y > max.y || c.z < min.z || c.z > max.z {
        return
    }

    _, touched := space[c.key()]
    if touched {
        return
    }

    space[c.key()] = false

    for _, adj := range adjs {
        dfs(space, min, max, coord{c.x + adj.x, c.y + adj.y, c.z + adj.z})
    }
}

func surface(coords []coord, insides bool) int {
    space := make(map[string]bool)
    for _, c := range coords {
        space[c.key()] = true
    }

    if !insides {
        min, max := coord{inf, inf, inf}, coord{-inf, -inf, -inf}

        for _, c := range coords {
            min.x, min.y, min.z = utils.Min(min.x, c.x-1), utils.Min(min.y, c.y-1), utils.Min(min.z, c.z-1)
            max.x, max.y, max.z = utils.Max(max.x, c.x+1), utils.Max(max.y, c.y+1), utils.Max(max.z, c.z+1)
        }

        dfs(space, min, max, min)
    }

    surface := 0
    for _, c := range coords {
        for _, adj := range adjs {
            nc := coord{c.x + adj.x, c.y + adj.y, c.z + adj.z}
            occ, touched := space[nc.key()]
            if !occ && (insides || touched) {
                surface++
            }
        }
    }

    return surface
}

func partOne(lines []string) {
    coords := parse(lines)
    surface := surface(coords, true)
    fmt.Println(surface)
}

func partTwo(lines []string) {
    coords := parse(lines)
    surface := surface(coords, false)
    fmt.Println(surface)
}

func tests() {
    tests := utils.ReadLines("tests", "18.txt")
    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "18.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
