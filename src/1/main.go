package main

import (
    "fmt"
    "sort"
    "strconv"
	"../utils"
)

func groupCalories(lines []string) []int {
    var grouped []int

    total := 0
    lines = append(lines, "")

    for _, line := range lines {
        if line == "" {
            grouped = append(grouped, total)
            total = 0

        } else {
            calories, _ := strconv.Atoi(line)
            total += calories
        }
    }

    return grouped
}

func partOne(lines []string) {
    maxCalories := 0

    for _, totalCalories := range groupCalories(lines) {
        maxCalories = utils.Max(maxCalories, totalCalories)
    }

    fmt.Println(maxCalories)
}

func partTwo(lines []string) {
    grouped := groupCalories(lines)

    sort.Ints(grouped)

    sumTop := 0
    for _, elemTop := range grouped[len(grouped)-3:] {
        sumTop += elemTop
    }

    fmt.Println(sumTop)
}

func tests() {
    test := utils.ReadLines("tests", "1.txt")

    partOne(test)
    partTwo(test)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "1.txt")

    partOne(puzzle)
    partTwo(puzzle)
}
