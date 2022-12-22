package main

import (
    "fmt"
    "../utils"
)

type move struct {
    head int // -1 L 0 C +1 R
    steps int
}

const kVoid int = 0
const kEmpty int = 1
const kWall int = 2

type level [][]int

func parseLevel(lines []string) level {
    width := 0
    for _, line := range lines {
        width = utils.Max(width, len(line))
    }

    ret := make([][]int, 0, len(lines))
    for _, line := range lines {
        row := make([]int, width)
        for j, c := range line {
            if c == '.' {
                row[j] = kEmpty

            } else if c == '#' {
                row[j] = kWall

            }
        }

        ret = append(ret, row)
    }

    return level(ret)
}

func parseInt(line string) (int, string) {
    value := 0

    for len(line) > 0 && line[0] >= '0' && line[0] <= '9' {
        value = value * 10 + int(line[0]) - int('0')
        line = line[1:]
    }

    return value, line
}

func parseMoves(line string) []move {
    ret := make([]move, 0)

    head := 0
    for {
        var steps int
        steps, line = parseInt(line)
        ret = append(ret, move{head, steps})

        if line == "" {
            break
        }

        if line[0] == 'L' {
            head = -1
        }

        if line[0] == 'R' {
            head = +1
        }

        line = line[1:]
    }

    return ret
}

func parse(lines []string) ([][]int, []move) {
    level := parseLevel(lines[:len(lines)-2])
    moves := parseMoves(lines[len(lines)-1])
    return level, moves
}

type coord struct {
    x, y int
}

func (c coord) answer(head int) int {
    return 1000 * c.y + 4 * c.x + head
}

func start(l level) coord {
    for y, row := range l {
        for x, c := range row {
            if c == kEmpty {
                return coord{x + 1, y + 1}
            }
        }
    }

    return coord{}
}

// RDLU
var dirs = [4]coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func tryMoveFlat(l level, s coord, dx, dy int) (coord, bool) {
    c := s

    for {
        ny := (c.y - 1 + dy + len(l)) % len(l) + 1
        row := l[ny-1]
        nx := (c.x - 1 + dx + len(row)) % len(row) + 1

        c = coord{nx, ny}

        if l[ny - 1][nx - 1] == kEmpty {
            break
        }

        if l[ny - 1][nx - 1] == kWall {
            return s, true
        }
    }

    return c, false
}

func applyMoves(l level, s coord, head int, ms []move) (coord, int) {
    for _, m := range ms {
        head = (head + 4 + m.head) % 4

        for i := 0; i < m.steps; i++ {
            var hit bool
            s, hit = tryMoveFlat(l, s, dirs[head].x, dirs[head].y)

            if hit {
                break
            }
        }
    }

    return s, head
}

func partOne(lines []string) {
    level, moves := parse(lines)

    s := start(level)
    e, head := applyMoves(level, s, 0, moves)

    fmt.Println(e.answer(head))
}

type vec3 struct {
    x, y, z int
}

func (a vec3) cross(b vec3) vec3 {
    return vec3{a.y * b.z - a.z * b.y, a.z * b.x - a.x * b.z, a.x * b.y - a.y * b.x}
}

func (a vec3) add(b vec3) vec3 {
    return vec3{a.x + b.x, a.y + b.y, a.z + b.z}
}

func (a vec3) scale(m int) vec3 {
    return vec3{a.x * m, a.y * m, a.z * m}
}

func (a vec3) neg() vec3{
    return vec3{-a.x, -a.y, -a.z}
}

func (a vec3) incube(n int) bool {
    return a.x >= 0 && a.x < n && a.y >= 0 && a.y < n && a.z >= 0 && a.z < n
}

type tile struct {
    p, n vec3
}

func (t tile) key() string {
    return fmt.Sprintf("%v", t)
}

type basis struct {
    dx, dy vec3
}

type cubeIter struct {
    t tile
    b basis
}

type fold struct {
    c coord
    l int
}

// cube of side n - [0, n-1]^3

func newCubeIter() cubeIter {
    return cubeIter{tile{vec3{0, 0, 0}, vec3{0, 0, 1}}, basis{vec3{1, 0, 0}, vec3{0, 1, 0}}}
}

// only one of x/y can be +-1
func (c cubeIter) move(n, x, y int) cubeIter {
    if np := c.t.p.add(c.b.dx.scale(x)).add(c.b.dy.scale(y)); np.incube(n) {
        return cubeIter{tile{np, c.t.n}, c.b}
    }

    // flip basis, normal should point inward

    ndx, ndy := c.b.dx, c.b.dy

    if x != 0 {
        ndx = c.t.n

        if x < 0 {
            ndx = ndx.neg()
        }
    }

    if y != 0 {
        ndy = c.t.n

        if y < 0 {
            ndy = ndy.neg()
        }
    }

    nn := ndx.cross(ndy)

    if tp := c.t.p.add(nn); !tp.incube(n) {
        nn = nn.neg()
    }

    return cubeIter{tile{c.t.p, nn}, basis{ndx, ndy}}
}

func dfs(n int, l level, vis [][]bool, lookup map[string]fold, c cubeIter, s coord) {
    vis[s.y-1][s.x-1] = true

    lookup[c.t.key()] = fold{s, l[s.y-1][s.x-1]}

    for _, d := range dirs {
        ns := coord{s.x + d.x, s.y + d.y}

        if ns.y < 1 || ns.y > len(l) {
            continue
        }

        row := l[ns.y-1]

        if ns.x < 1 || ns.x > len(row) {
            continue
        }

        if !vis[ns.y - 1][ns.x - 1] && l[ns.y - 1][ns.x - 1] != kVoid {
            dfs(n, l, vis, lookup, c.move(n, d.x, d.y), ns)
        }
    }
}

func cubeFold(n int, l level, s coord) map[string]fold {
    vis := make([][]bool, len(l))
    lookup := make(map[string]fold)

    for i, row := range l {
        vis[i] = make([]bool, len(row))
    }

    dfs(n, l, vis, lookup, newCubeIter(), s)

    return lookup
}

func tryMoveFlatCubeIter(n int, l level, lookup map[string]fold, c cubeIter, dx, dy int) (cubeIter, bool) {
    nc := c.move(n, dx, dy)

    if lookup[nc.t.key()].l == kWall {
        return c, true
    }

    return nc, false
}

func applyMovesCubeIter(n int, l level, lookup map[string]fold, c cubeIter, head int, ms []move) (coord, int) {
    var pc coord

    for _, m := range ms {
        head = (head + 4 + m.head) % 4

        for i := 0; i < m.steps; i++ {
            var hit bool

            cc := c
            c, hit = tryMoveFlatCubeIter(n, l, lookup, c, dirs[head].x, dirs[head].y)
            // fmt.Println(c, hit, lookup[c.t.key()])

            if hit {
                break

            } else {
                pc = lookup[cc.t.key()].c
            }
        }
    }

    lc := lookup[c.t.key()].c
    ldir := coord{lc.x - pc.x, lc.y - pc.y}

    for lhead, d := range dirs {
        if d == ldir {
            return lc, lhead
        }
    }

    return lc, 1000000000
}

func partTwo(lines []string, side int) {
    level, moves := parse(lines)

    s := start(level)
    lookup := cubeFold(side, level, s)

    e, head := applyMovesCubeIter(side, level, lookup, newCubeIter(), 0, moves)

    fmt.Println(e.answer(head))
}

func tests() {
    tests := utils.ReadLines("tests", "22.txt")
    partOne(tests)
    partTwo(tests, 4)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "22.txt")
    partOne(puzzle)
    partTwo(puzzle, 50)
}
