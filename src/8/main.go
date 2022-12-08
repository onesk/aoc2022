package main

import (
    "fmt"
    "../utils"
)

type trees [][]int

func parse(rows []string) trees {
    ret := make([][]int, 0)

    for _, row := range rows {
        trees := make([]int, 0, len(row))

        for _, ch := range row {
            trees = append(trees, int(ch) - int('0'))
        }

        ret = append(ret, trees)
    }

    return trees(ret)
}

type dir struct {
    di, dj int
}

var dirs [4]dir = [4]dir{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}

func inrange(trees trees, i, j int) bool {
    if i < 0 || i >= len(trees) {
        return false
    }

    if j < 0 || len(trees) == 0 || j >= len(trees[0]) {
        return false
    }

    return true
}

func cast(trees trees, i, j, h int, d dir) (visible int, anyHigher bool) {
    visible, anyHigher = 0, false

    for true {
        i, j = i + d.di, j + d.dj

        if !inrange(trees, i, j) {
            break
        }

        visible++

        if trees[i][j] >= h {
            anyHigher = true
            break
        }
    }

    return
}

func visibleFromOutside(trees trees, i, j, h int) bool {
    for _, d := range dirs {
        _, anyHigher := cast(trees, i, j, h, d)

        if !anyHigher {
            return true
        }
    }

    return false
}

func scenicScore(trees trees, i, j, h int) int {
    product := 1

    for _, d := range dirs {
        visible, _ := cast(trees, i, j, h, d)
        product *= visible
    }

    return product
}

func partOne(rows []string) {
    trees := parse(rows)

    sum := 0
    for i, hrow := range trees {
        for j, h := range hrow {
            if visibleFromOutside(trees, i, j, h) {
                sum += 1
            }
        }
    }

    fmt.Println(sum)
}

func partTwo(rows []string) {
    trees := parse(rows)

    score := 0
    for i, hrow := range trees {
        for j, h := range hrow {
            score = utils.Max(score, scenicScore(trees, i, j, h))
        }
    }

    fmt.Println(score)
}

func tests() {
    tests := utils.ReadLines("tests", "8.txt")
    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "8.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
