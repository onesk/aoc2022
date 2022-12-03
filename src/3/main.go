package main

import (
    "fmt"
    "strings"
    "../utils"
)

func priority(ch rune) int {
    if 'a' <= ch && ch <= 'z' {
        return int(ch) - int('a') + 1
    }

    return int(ch) - int('A') + 27
}

func duplicatePriority(rucksack string) int {
    half := len(rucksack) / 2
    left, right := rucksack[half:], rucksack[:half]

    for _, ch := range left {
        if strings.Contains(right, string(ch)) {
            return priority(ch)
        }
    }

    return 1e9
}

func commonPriority(rucksacks []string) int {
    first, rest := rucksacks[0], rucksacks[1:]

    for _, ch := range first {
        allContain := true

        for _, rucksack := range rest {
            allContain = allContain && strings.Contains(rucksack, string(ch))
        }

        if allContain {
            return priority(ch)
        }
    }

    return 1e9
}

func partOne(rucksacks []string) {
    sum := 0

    for _, rucksack := range rucksacks {
        sum += duplicatePriority(rucksack)
    }

    fmt.Println(sum)
}

func partTwo(rucksacks []string) {
    sum := 0

    for i := 0; i < len(rucksacks); i += 3 {
        sum += commonPriority(rucksacks[i:i+3])
    }

    fmt.Println(sum)
}

func tests() {
    tests := utils.ReadLines("tests", "3.txt")
    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "3.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
