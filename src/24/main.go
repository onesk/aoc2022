package main

import (
    "fmt"
    "../utils"
)

type coord struct {
    x, y int
}

type bliz struct {
    c coord
    dx, dy, span int
}

type level struct {
    w, h int
    bs []bliz
}

func parse(lines []string) level {
    w, h := len(lines[0]), len(lines)
    bs := make([]bliz, 0)

    for y, row := range lines {
        for x, c := range row {
            if c == '.' || c == '#' {
                continue
            }

            var dx, dy, span int

            switch c {
            case '<':
                dx, dy, span = -1, 0, w
            case '>':
                dx, dy, span = +1, 0, w
            case '^':
                dx, dy, span = 0, -1, h
            case 'v':
                dx, dy, span = 0, +1, h
            }

            bs = append(bs, bliz{coord{x, y}, dx, dy, span-2})
        }
    }

    return level{w, h, bs}
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    }

    return gcd(b, a % b)
}

func lcm(a, b int) int {
    return a / gcd(a, b) * b
}

type qelem struct {
    c coord
    mod, len int
}

var dirs = [5]coord {
    {0, 0},
    {-1, 0},
    {+1, 0},
    {0, -1},
    {0, +1},
}

func (q qelem) key() string {
    return fmt.Sprintf("%v:%v", q.c, q.mod)
}

func hitBlizzard(l level, c coord, time int) bool {
    for _, b := range l.bs {
        nc := coord{b.c.x + b.dx * time, b.c.y + b.dy * time}

        for nc.x <= 0 || nc.y <= 0 || nc.x >= l.w-1 || nc.y >= l.h-1 {
            nc = coord{nc.x - b.dx * b.span, nc.y - b.dy * b.span}
        }

        if c == nc {
            return true
        }
    }

    return false
}

func bfs(l level, sx, sy, ex, ey, startTime int) int {
    mod := lcm(l.w-2, l.h-2)

    q := make([]qelem, l.w * l.h * mod)
    q[0] = qelem{coord{sx, sy}, startTime % mod, 0}

    qs, qe := 0, 1
    vis := make(map[string]bool)

    for qs < qe {
        qhead := q[qs]
        qs++

        if qhead.c.x == ex && qhead.c.y == ey {
            return qhead.len
        }

        for _, d := range dirs {
            np := coord{qhead.c.x + d.x, qhead.c.y + d.y}

            isStartEnd := (np.x == sx && np.y == sy) || (np.x == ex && np.y == ey)

            if (np.x <= 0 || np.x >= l.w-1 || np.y <= 0 || np.y >= l.h-1) && !isStartEnd {
                continue
            }

            nqh := qelem{np, (qhead.mod+1) % mod, qhead.len+1}

            if hitBlizzard(l, np, nqh.mod) || vis[nqh.key()] {
                continue
            }

            q[qe] = nqh
            qe++

            vis[nqh.key()] = true
        }
    }

    return -1
}

func forward(lines []string) (int, level) {
    l := parse(lines)
    ans := bfs(l, 1, 0, l.w-2, l.h-1, 0)

    return ans, l
}

func partOne(lines []string) {
    ans, _ := forward(lines)

    fmt.Println(ans)
}

func partTwo(lines []string) {
    forward, l := forward(lines)
    backward := bfs(l, l.w-2, l.h-1, 1, 0, forward)
    againForward := bfs(l, 1, 0, l.w-2, l.h-1, forward + backward)

    fmt.Println(forward + backward + againForward)
}

func tests() {
    test := utils.ReadLines("tests", "24.txt")

    partOne(test)
    partTwo(test)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "24.txt")

    partOne(puzzle)
    partTwo(puzzle)
}
