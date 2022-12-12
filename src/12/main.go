package main

import (
    "fmt"
    "../utils"
)

type hmap struct {
    elev [][]int
    w, h int
}

type coord struct {
    i, j int
}

var dirs = [4]coord{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}

func parse(lines []string) (heights hmap, s, e coord) {
    w, h := len(lines), len(lines[0])
    elev := make([][]int, w)

    for i, line := range lines {

        elev[i] = make([]int, h)

        for j, ch := range line {
            if ch == 'S' {
                s, elev[i][j] = coord{i, j}, 0

            } else if ch == 'E' {
                e, elev[i][j] = coord{i, j}, 25

            } else {
                // a..z
                elev[i][j] = int(ch) - 'a'
            }
        }
    }

    return hmap{elev, w, h}, s, e
}

func bfs(heights hmap, s []coord, e coord) int {
    queue, qs, qe := make([]coord, heights.w * heights.h), 0, len(s)
    copy(queue[:len(s)], s)

    dist := make([][]int, heights.w)
    for i := 0; i < heights.w; i++ {
        dist[i] = make([]int, heights.h)
    }

    for _, sc := range s {
        dist[sc.i][sc.j] = 1
    }

    for qs < qe {
        c := queue[qs]
        qs++

        if c.i == e.i && c.j == e.j {
            return dist[c.i][c.j] - 1
        }

        for _, dc := range dirs {
            ni, nj := c.i + dc.i, c.j + dc.j

            if ni < 0 || ni >= heights.w || nj < 0 || nj >= heights.h {
                continue
            }

            delev := heights.elev[ni][nj] - heights.elev[c.i][c.j]

            if delev > 1 || dist[ni][nj] != 0 {
                continue
            }

            dist[ni][nj] = dist[c.i][c.j] + 1
            queue[qe] = coord{ni, nj}
            qe++
        }
    }

    return -1
}

func partOne(lines []string) {
    heights, s, e := parse(lines)
    dist := bfs(heights, []coord{s}, e)

    fmt.Println(dist)
}

func partTwo(lines []string) {
    heights, _, e := parse(lines)

    s := make([]coord, 0)
    for i, row := range heights.elev {
        for j, h := range row {
            if h == 0 {
                s = append(s, coord{i, j})
            }
        }
    }

    dist := bfs(heights, s, e)

    fmt.Println(dist)
}

func tests() {
    tests := utils.ReadLines("tests", "12.txt")
    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "12.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
