package main

import (
    "fmt"
    "strings"
    "strconv"
    "../utils"
)

const inf = 1000000000
const extent = 500

type coord struct {
    x, y int
}

func (c coord) key() string {
    return fmt.Sprintf("%v:%v", c.x, c.y)
}

type aabb struct {
    tl, br coord
}

func emptyAABB() aabb {
    return aabb{coord{inf, inf}, coord{-inf, -inf}}
}

func (c aabb) addCoord(n coord) aabb {
    return aabb{
        coord{utils.Min(c.tl.x, n.x), utils.Min(c.tl.y, n.y)},
        coord{utils.Max(c.br.x, n.x), utils.Max(c.br.y, n.y)},
    }
}

func (c aabb) dist(n coord) int {
    return utils.Max(
        utils.Max(c.tl.x - n.x, c.tl.y - n.y),
        utils.Max(n.x - c.br.x, n.y - c.br.y),
    )
}

type lineStrip []coord

func parse(lines []string) []lineStrip {
    ret := make([]lineStrip, 0, len(lines))

    for _, line := range lines {
        parts := strings.Split(line, " -> ")
        strip := make(lineStrip, 0, len(parts))

        for _, part := range parts {
            coords := strings.Split(part, ",")

            x, _ := strconv.Atoi(coords[0])
            y, _ := strconv.Atoi(coords[1])

            strip = append(strip, coord{x, y})
        }

        ret = append(ret, strip)
    }

    return ret
}

// true if sand, false if wall
type cave map[string]bool

func sgn(x int) int {
    if x > 0 {
        return 1
    }

    if x < 0 {
        return -1
    }

    return 0
}

func buildWalls(strips []lineStrip) (nc cave, bounds aabb) {
    nc, bounds = make(cave), emptyAABB()

    for _, strip := range strips {
        for i := 0; i + 1 < len(strip); i++ {
            c, n := strip[i], strip[i+1]
            dx, dy := sgn(n.x - c.x), sgn(n.y - c.y)

            for {
                nc[c.key()] = false
                bounds = bounds.addCoord(c)

                if c == n {
                    break
                }

                c.x, c.y = c.x + dx, c.y + dy
            }
        }
    }

    return
}

func dropSand(cave cave, aabb aabb, maxDist int, s coord) bool {
    dxs := [3]int{0, -1, 1}

Outer:
    for aabb.dist(s) <= maxDist {
        for _, dx := range dxs {
            n := coord{s.x + dx, s.y + 1}
            _, occluder := cave[n.key()]

            if !occluder {
                s = n
                continue Outer
            }
        }

        cave[s.key()] = true
        return true
    }

    return false
}

func partOne(lines []string) {
    strips := parse(lines)
    cave, aabb := buildWalls(strips)

    startPoint := coord{500, 0}
    aabb = aabb.addCoord(startPoint)

    ans := 0
    for dropSand(cave, aabb, 1, startPoint) {
        ans++
    }

    fmt.Println(ans)
}

func partTwo(lines []string) {
    strips := parse(lines)
    _, aabb := buildWalls(strips)

    lineBelow := []coord{{aabb.tl.x - extent, aabb.br.y + 2}, {aabb.br.x + extent, aabb.br.y + 2}}
    strips = append(strips, lineBelow)

    var cave cave
    cave, aabb = buildWalls(strips)

    startPoint := coord{500, 0}
    aabb = aabb.addCoord(startPoint)

    ans := 0
    for {
        _, occluder := cave[startPoint.key()]
        if occluder {
            break
        }

        dropSand(cave, aabb, 1, startPoint)
        ans++
    }

    fmt.Println(ans)
}

func tests() {
    tests := utils.ReadLines("tests", "14.txt")
    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "14.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
