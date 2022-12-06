package main

import (
    "fmt"
    "sort"
    "../utils"
)

func firstDistinct(line string, size int) int {
    window := make([]byte, size)

    for i := 0; i + size < len(line); i++ {
        copy(window, line[i:i + size])

        sort.Slice(window, func (i, j int) bool { return window[i] < window[j] })

        distinct := true
        for j := 0; j + 1 < size; j++ {
            if window[j] == window[j+1] {
                distinct = false
                break
            }
        }

        if distinct {
            return i + size
        }
    }

    return -1
}

func partOne(lines []string) {
    line := lines[0]

    fmt.Println(firstDistinct(line, 4))
}


func partTwo(lines []string) {
    line := lines[0]

    fmt.Println(firstDistinct(line, 14))
}

func tests() {
    tests := utils.ReadLines("tests", "6.txt")
    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "6.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
