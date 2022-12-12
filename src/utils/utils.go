package utils

import (
    "os"
    "bufio"
    "path/filepath"
)

func ReadLines(directory string, filename string) ([]string) {
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

func Max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func Min(a, b int) int {
    if a < b {
        return a
    }

    return b
}

func Abs(a int) int {
    if a < 0 {
        return -a
    }

    return a
}
