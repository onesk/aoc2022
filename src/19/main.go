package main

import (
    "fmt"
    "../utils"
)

const kinds = 4

const ore = 0
const clay = 1
const obsidian = 2
const geode = 3

type counts [kinds]int
type blueprint [kinds]counts

func (c counts) cloneCounts() (ret counts) {
    copy(ret[:], c[:])
    return
}

func (c *counts) addRate(rs counts) {
    for i, r := range rs {
        (*c)[i] += r
    }
}

func (c counts) increment(dest int) counts {
    ret := c.cloneCounts()
    ret[dest]++
    return ret
}

func (b blueprint) applyCounts(dest int, c counts) (counts, bool) {
    for i, prereq := range b[dest] {
        if prereq > c[i] {
            return c, false
        }
    }

    ret := c.cloneCounts()

    for i, prereq := range b[dest] {
        ret[i] = c[i] - prereq
    }

    return ret, true
}

// Parser writing laziness kicks in, again.

var testsInput = []blueprint {
    {{4, 0, 0, 0}, {2, 0, 0, 0}, {3, 14, 0, 0}, {2, 0, 7, 0}},
    {{2, 0, 0, 0}, {3, 0, 0, 0}, {3, 8, 0, 0}, {3, 0, 12, 0}},
}

var puzzleInput = []blueprint {
    {{4, 0, 0, 0}, {4, 0, 0, 0}, {2, 11, 0, 0}, {2, 0, 7,  0}},
    {{4, 0, 0, 0}, {4, 0, 0, 0}, {4, 20, 0, 0}, {2, 0, 8,  0}},
    {{3, 0, 0, 0}, {4, 0, 0, 0}, {3, 15, 0, 0}, {4, 0, 16, 0}},
    {{2, 0, 0, 0}, {4, 0, 0, 0}, {2, 15, 0, 0}, {3, 0, 16, 0}},
    {{4, 0, 0, 0}, {3, 0, 0, 0}, {4, 8,  0, 0}, {3, 0, 7,  0}},
    {{3, 0, 0, 0}, {4, 0, 0, 0}, {4, 18, 0, 0}, {3, 0, 8,  0}},
    {{2, 0, 0, 0}, {4, 0, 0, 0}, {2, 20, 0, 0}, {3, 0, 15, 0}},
    {{3, 0, 0, 0}, {4, 0, 0, 0}, {3, 15, 0, 0}, {3, 0, 20, 0}},
    {{2, 0, 0, 0}, {4, 0, 0, 0}, {3, 14, 0, 0}, {4, 0, 9,  0}},
    {{4, 0, 0, 0}, {4, 0, 0, 0}, {2, 9,  0, 0}, {3, 0, 15, 0}},
    {{2, 0, 0, 0}, {3, 0, 0, 0}, {3, 9,  0, 0}, {3, 0, 9,  0}},
    {{3, 0, 0, 0}, {4, 0, 0, 0}, {3, 16, 0, 0}, {3, 0, 14, 0}},
    {{3, 0, 0, 0}, {4, 0, 0, 0}, {3, 6,  0, 0}, {2, 0, 10, 0}},
    {{4, 0, 0, 0}, {3, 0, 0, 0}, {2, 7,  0, 0}, {3, 0, 8,  0}},
    {{4, 0, 0, 0}, {3, 0, 0, 0}, {2, 17, 0, 0}, {3, 0, 16, 0}},
    {{4, 0, 0, 0}, {4, 0, 0, 0}, {4, 15, 0, 0}, {4, 0, 17, 0}},
    {{3, 0, 0, 0}, {3, 0, 0, 0}, {3, 9,  0, 0}, {3, 0, 7,  0}},
    {{4, 0, 0, 0}, {4, 0, 0, 0}, {4, 8,  0, 0}, {3, 0, 19, 0}},
    {{4, 0, 0, 0}, {4, 0, 0, 0}, {3, 19, 0, 0}, {4, 0, 15, 0}},
    {{2, 0, 0, 0}, {3, 0, 0, 0}, {3, 14, 0, 0}, {3, 0, 19, 0}},
    {{3, 0, 0, 0}, {4, 0, 0, 0}, {2, 15, 0, 0}, {3, 0, 7,  0}},
    {{4, 0, 0, 0}, {4, 0, 0, 0}, {4, 7,  0, 0}, {2, 0, 16, 0}},
    {{2, 0, 0, 0}, {2, 0, 0, 0}, {2, 7,  0, 0}, {2, 0, 14, 0}},
    {{3, 0, 0, 0}, {3, 0, 0, 0}, {3, 19, 0, 0}, {3, 0, 19, 0}},
    {{2, 0, 0, 0}, {3, 0, 0, 0}, {3, 11, 0, 0}, {3, 0, 14, 0}},
    {{4, 0, 0, 0}, {3, 0, 0, 0}, {3, 7,  0, 0}, {3, 0, 9,  0}},
    {{2, 0, 0, 0}, {4, 0, 0, 0}, {2, 16, 0, 0}, {2, 0, 9,  0}},
    {{3, 0, 0, 0}, {4, 0, 0, 0}, {4, 19, 0, 0}, {4, 0, 11, 0}},
    {{4, 0, 0, 0}, {3, 0, 0, 0}, {2, 5,  0, 0}, {2, 0, 10, 0}},
    {{2, 0, 0, 0}, {4, 0, 0, 0}, {3, 20, 0, 0}, {2, 0, 17, 0}},
}

func recMaxGeodes(b blueprint, totalTime int, oldPosMap, cTime int, cnts, rates counts) int {
    if totalTime == cTime {
        return cnts[geode]
    }

    posMap := 0
    for dest := 0; dest < kinds; dest++ {
        _, canAfford := b.applyCounts(dest, cnts)

        if canAfford {
            posMap |= (1 << uint(dest))
        }
    }

    high, low, ignorePosMap := geode, ore, false

    if (posMap & 8) != 0 {
        high, low, ignorePosMap = geode, geode, true
    }

    best := 0
    for dest := low; dest <= high; dest++ {
        newCounts, applied := b.applyCounts(dest, cnts)

        if !applied || ((oldPosMap & (1 << uint(dest))) != 0 && !ignorePosMap) {
            continue
        }

        newCounts.addRate(rates)

        recBest := recMaxGeodes(b, totalTime, 0, cTime+1, newCounts, rates.increment(dest))
        best = utils.Max(best, recBest)
    }

    skipCounts := cnts.cloneCounts()
    skipCounts.addRate(rates)

    skipBest := recMaxGeodes(b, totalTime, posMap, cTime+1, skipCounts, rates)
    best = utils.Max(best, skipBest)

    return best
}

func maxGeodes(b blueprint, totalTime int) int {
    cnts := counts{0, 0, 0, 0}
    rates := counts{1, 0, 0, 0}

    return recMaxGeodes(b, totalTime, 0, 0, cnts, rates)
}

func partOne(input []blueprint) {
    ans := 0
    for i, b := range input {
        bBest := maxGeodes(b, 24)

        fmt.Println(i + 1, bBest)
        ans += (i + 1) * bBest
    }

    fmt.Println(ans)
}

func partTwo(input []blueprint) {
    ans := 1
    for i, b := range input[:utils.Min(3, len(input))] {
        bBest := maxGeodes(b, 32)

        fmt.Println(i + 1, bBest)
        ans *= bBest
    }

    fmt.Println(ans)
}

func tests() {
    partOne(testsInput)
    partTwo(testsInput)
}

func main() {
    tests()

    partOne(puzzleInput)
    partTwo(puzzleInput)
}
