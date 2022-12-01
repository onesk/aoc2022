package main

import (
    "os"
    "fmt"
    "sort"
    "bufio"
    "strconv"
    "path/filepath"
)

func readLines(directory string, filename string) ([]string) {
    path := filepath.Join(directory, filename)

    file, _ := os.Open(path)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    return lines
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

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
        maxCalories = max(maxCalories, totalCalories)
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
    partOne(readLines("tests", "1.txt"))
    partTwo(readLines("tests", "1.txt"))
}

func main() {
    tests()

    partOne(readLines("inputs", "1.txt"))
    partTwo(readLines("inputs", "1.txt"))
}
