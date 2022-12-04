package main

import (
	"fmt"
    "strconv"
    "strings"
    "../utils"
)

type Interval struct {
    left int
    right int
}

func (i Interval) Contains(j Interval) bool {
    return i.left <= j.left && j.right <= i.right
}


func (i Interval) Overlap(j Interval) bool {
    return i.right >= j.left && j.right >= i.left
}

func parseInterval(interval string) Interval {
    ends := strings.Split(interval, "-")
    left, _ := strconv.Atoi(ends[0])
    right, _ := strconv.Atoi(ends[1])
    return Interval{ left, right }
}

func checkIntervals(assignments []string, check func (i, j Interval) bool) {
    total := 0

    for _, line := range assignments {
        halves := strings.Split(line, ",")
        i, j := parseInterval(halves[0]), parseInterval(halves[1])

        if check(i, j) {
            total++
        }
    }

    fmt.Println(total)
}

func partOne(assignments []string) {
    checkIntervals(assignments, func (i, j Interval) bool {
        return i.Contains(j) || j.Contains(i)
    })
}

func partTwo(assignments []string) {
    checkIntervals(assignments, func (i, j Interval) bool {
        return i.Overlap(j)
    })
}

func tests() {
    tests := utils.ReadLines("tests", "4.txt")
    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "4.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
