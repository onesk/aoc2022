package main

import (
    "fmt"
	"../utils"
)

func fromSnafu(s string) int {
    cur := 0

    for _, c := range s {
        cur *= 5

        switch c {
        case '=':
            cur += -2
        case '-':
            cur += -1
        case '0':
            cur += 0
        case '1':
            cur += 1
        case '2':
            cur += 2
        }
    }

    return cur
}

func toSnafu(v int) string {
    digits := make([]int, 0)

    for {
        digits = append(digits, v%5)
        v /= 5

        if v == 0 {
            break
        }
    }

    digits = append(digits, 0)

    for i := 0; i < len(digits) - 1; i++ {
        for digits[i] > 2 {
            digits[i] -= 5
            digits[i+1] += 1
        }
    }

    s, j := "", len(digits)

    for j > 0 && digits[j-1] == 0 {
        j--
    }

    for j > 0 {
        s += string("=-012"[digits[j-1]+2])
        j--
    }

    return s
}

func partOne(lines []string) {
    sum := 0

    for _, line := range lines {
        sum += fromSnafu(line)
    }

    fmt.Println(toSnafu(sum))
}

func tests() {
    test := utils.ReadLines("tests", "25.txt")

    partOne(test)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "25.txt")

    partOne(puzzle)
}
