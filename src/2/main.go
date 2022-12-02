package main

import (
    "fmt"
    "../utils"
)

// RPS
const xyz = "XYZ"
const abc = "ABC"

func outcome(your int, other int) int {
    return (4 + your - other) % 3
}

func index(label byte, labels string) int {
    for i, char := range labels {
        if label == byte(char) {
            return i
        }
    }

    return 1e9 // let it be an error value
}

func roundScore(your int, other int) int {
    return (your + 1) + 3*outcome(your, other)
}

func partOne(strats []string) {
    totalScore := 0

    for _, strat := range strats {
        your := index(strat[2], xyz)
        other := index(strat[0], abc)

        totalScore += roundScore(your, other)
    }

    fmt.Println(totalScore)
}

func partTwo(strats []string) {
    totalScore := 0

    for _, strat := range strats {
        neededOutcome := index(strat[2], xyz)
        other := index(strat[0], abc)

        for your := 0; your <= 2; your++ {
            if outcome(your, other) == neededOutcome {
                totalScore += roundScore(your, other)
            }
        }
    }

    fmt.Println(totalScore)
}

func tests() {
    tests := utils.ReadLines("tests", "2.txt")

    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "2.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
