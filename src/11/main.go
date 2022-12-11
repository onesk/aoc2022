package main

import (
    "fmt"
    "sort"
)

type monkey struct {
    items []int
    effect func (wl int) int
    test func (wl int) bool
    ifTrue int
    ifFalse int
}

// Too lazy to write a parser for that, Emacs macro works just fine.
// Yeah that's a shortcut, f*** if I care ;)

const primorial = 2 * 3 * 5 * 7 * 11 * 13 * 17 * 19 * 23

var test = []monkey{
    {[]int{79, 98},
        func (wl int) int { return wl * 19; },
        func (wl int) bool { return wl % 23 == 0 },
        2, 3},
    {[]int{54, 65, 75, 74},
        func (wl int) int { return wl + 6; },
        func (wl int) bool { return wl % 19 == 0 },
        2, 0},
    {[]int{79, 60, 97},
        func (wl int) int { return wl * wl; },
        func (wl int) bool { return wl % 13 == 0 },
        1, 3},
    {[]int{74},
        func (wl int) int { return wl + 3; },
        func (wl int) bool { return wl % 17 == 0 },
        0, 1},
}

var puzzle = []monkey{
    {[]int{89, 73, 66, 57, 64, 80},
        func (wl int) int { return wl * 3; },
        func (wl int) bool { return wl % 13 == 0 },
        6, 2},
    {[]int{83, 78, 81, 55, 81, 59, 69},
        func (wl int) int { return wl + 1; },
        func (wl int) bool { return wl % 3 == 0 },
        7, 4},
    {[]int{76, 91, 58, 85},
        func (wl int) int { return wl * 13; },
        func (wl int) bool { return wl % 7 == 0 },
        1, 4},
    {[]int{71, 72, 74, 76, 68},
        func (wl int) int { return wl * wl; },
        func (wl int) bool { return wl % 2 == 0 },
        6, 0},
    {[]int{98, 85, 84},
        func (wl int) int { return wl + 7; },
        func (wl int) bool { return wl % 19 == 0 },
        5, 7},
    {[]int{78},
        func (wl int) int { return wl + 8; },
        func (wl int) bool { return wl % 5 == 0 },
        3, 0},
    {[]int{86, 70, 60, 88, 88, 78, 74, 83},
        func (wl int) int { return wl + 4; },
        func (wl int) bool { return wl % 11 == 0 },
        1, 2},
    {[]int{81, 58},
        func (wl int) int { return wl + 5; },
        func (wl int) bool { return wl % 17 == 0 },
        3, 5},
}

func deepCopy(ms []monkey) []monkey {
    ret := make([]monkey, len(ms))
    copy(ret, ms)

    for i, _ := range ret {
        items := make([]int, len(ret[i].items))
        copy(items, ret[i].items)
        ret[i].items = items
    }

    return ret
}

func round(ms []monkey, modulo int, inspectCb func (monkey int)) {
    for i, _ := range ms {
        // to account for possible appends to the current stack
        for len(ms[i].items) > 0 {
            var item int
            item, ms[i].items = ms[i].items[0], ms[i].items[1:]

            inspectCb(i)
            item = ms[i].effect(item)

            if modulo == 0 {
                item /= 3
            } else {
                item %= modulo
            }

            dest := ms[i].ifFalse
            if ms[i].test(item) {
                dest = ms[i].ifTrue
            }

            ms[dest].items = append(ms[dest].items, item)
        }
    }
}

func solve(ms []monkey, modulo int, rounds int) {
    counts := make([]int, len(ms))

    for i := 0; i < rounds; i++ {
        round(ms, modulo, func (i int) {
            counts[i]++
        })
    }

    sort.Ints(counts)
    ans := counts[len(counts)-1] * counts[len(counts)-2]

    fmt.Println(ans)
}

func partOne(ms []monkey) {
    solve(ms, 0, 20)
}

func partTwo(ms []monkey) {
    solve(ms, primorial, 10000)
}

func tests() {
    partOne(deepCopy(test))
    partTwo(deepCopy(test))
}

func main() {
    tests()

    partOne(deepCopy(puzzle))
    partTwo(deepCopy(puzzle))
}
