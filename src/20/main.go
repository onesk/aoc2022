package main

import (
    "fmt"
    "strconv"
    "../utils"
)

func parse(lines []string, key int) []int {
    ret := make([]int, 0, len(lines))

    for _, line := range lines {
        n, _ := strconv.Atoi(line)
        ret = append(ret, n * key)
    }

    return ret
}

func indexOf(slice []int, needle int) int {
    for i, e := range slice {
        if e == needle {
            return i
        }
    }

    return -1
}

func sgn(v int) int {
    if v > 0 {
        return 1
    }

    if v < 0 {
        return -1
    }

    return 0
}

func aswap(s, xref []int, pos, dir int) (npos int) {
    if dir > 0 {
        if pos == len(s)-1 {
            e, xe := s[pos], xref[pos]
            for j := pos; j >= 1; j-- {
                s[j], xref[j] = s[j-1], xref[j-1]
            }

            pos = 0
            s[pos], xref[pos] = e, xe
        }

        npos = pos + 1

    } else if dir < 0 {
        if pos == 0 {
            e, xe := s[0], xref[0]
            for j := 0; j+1 < len(s); j++ {
                s[j], xref[j] = s[j+1], xref[j+1]
            }

            pos = len(s)-1
            s[pos], xref[pos] = e, xe
        }

        npos = pos - 1

    } else {
        npos = pos

    }

    s[npos],    s[pos]    = s[pos],    s[npos]
    xref[npos], xref[pos] = xref[pos], xref[npos]

    return
}

func mix(nums []int, rounds int) []int {
    c := make([]int, len(nums))
    copy(c[:], nums[:])

    xref := make([]int, len(nums))
    for j := 0; j < len(xref); j++ {
        xref[j] = j
    }

    for r := 0; r < rounds; r++ {
        for i := 0; i < len(nums); i++ {
            pos  := indexOf(xref, i)
            cm   := c[pos]

            dir  := sgn(cm)
            reps := utils.Abs(cm) % (len(c) - 1)

            for j := 0; j < reps; j++ {
                pos = aswap(c, xref, pos, dir)
            }
        }
    }

    return c
}

func extractCoords(mixed []int) int {
    si := indexOf(mixed, 0)

    ans := 0
    for d := 1; d <= 3; d++ {
        ans += mixed[(si + 1000 * d) % len(mixed)]
    }

    return ans
}

func partOne(lines []string) {
    numbers := parse(lines, 1)
    mixed := mix(numbers, 1)
    ans := extractCoords(mixed)

    fmt.Println(ans)
}

func partTwo(lines []string) {
    numbers := parse(lines, 811589153)
    mixed := mix(numbers, 10)
    ans := extractCoords(mixed)

    fmt.Println(ans)
}

func tests() {
    test := utils.ReadLines("tests", "20.txt")

    partOne(test)
    partTwo(test)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "20.txt")

    partOne(puzzle)
    partTwo(puzzle)
}
