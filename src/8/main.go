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

func sameShapeInts(trees trees, num int) [][]int {
    ret := make([][]int, 0, len(trees))

    for _, row := range(trees) {
        newRow := make([]int, len(row))

        for j, _ := range newRow {
            newRow[j] = num
        }

        ret = append(ret, newRow)
    }

    return ret
}

const horStackCap = 10
const maxHeight = 10000

// positions are 1-indexed
// h decrease, pos increase
type stackItem struct {
    h, pos int
}

type horstack []stackItem

func newHorstack() horstack {
    return make([]stackItem, 0, horStackCap)
}

func (s *horstack) rightmost() stackItem {
    l := len(*s)

    if l == 0 {
        return stackItem{maxHeight, 0}
    }

    return (*s)[l-1]
}

func (s *horstack) pop() {
    *s = (*s)[:len(*s)-1]
}

func (s *horstack) update(h, pos int) (vis int, out bool) {
    for s.rightmost().h < h {
        s.pop()
    }

    vis, out = pos - s.rightmost().pos, len(*s) == 0

    *s = append(*s, stackItem{h, pos})
    return
}

func mapAccum(trees trees, each func (i, j, vis int, out bool)) {
    rows, cols := len(trees), 0

    // map each row, both directions
    for i, row := range trees {
        cols = len(row)

        forward, backward := newHorstack(), newHorstack()

        for j, h := range row {
            vis, out := forward.update(h, j)
            each(i, j, vis, out)

            backJ := cols-j-1
            vis, out = backward.update(row[backJ], j)
            each(i, backJ, vis, out)
        }
    }

    // map each column, both directions
    for j := 0; j < cols; j++ {
        forward, backward := newHorstack(), newHorstack()

        for i := 0; i < rows; i++ {
            vis, out := forward.update(trees[i][j], i)
            each(i, j, vis, out)

            backI := rows-1-i
            vis, out = backward.update(trees[backI][j], i)
            each(backI, j, vis, out)
        }
    }
}

func partOne(rows []string) {
    trees := parse(rows)
    outs := sameShapeInts(trees, 0)

    mapAccum(trees, func (i, j, vis int, out bool) {
        if out {
            outs[i][j] = 1
        }
    })

    sum := 0
    for _, row := range outs {
        for _, out := range row {
            sum += out
        }
    }

    fmt.Println(sum)
}

func partTwo(rows []string) {
    trees := parse(rows)
    outs := sameShapeInts(trees, 1)

    mapAccum(trees, func (i, j, vis int, out bool) {
        outs[i][j] *= vis
    })

    max := 0
    for _, row := range outs {
        for _, out := range row {
            max = utils.Max(max, out)
        }
    }

    fmt.Println(max)
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
