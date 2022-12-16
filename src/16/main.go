package main

import (
    "fmt"
    "regexp"
    "strings"
    "strconv"
    "../utils"
)

type valve struct {
    name string
    rate int
    adj []string
}

func parse(lines []string) []valve {
    re := regexp.MustCompile(
        `Valve ([[:alpha:]]+) has flow rate=([[:digit:]]+); ` +
            `tunnel[s]? lead[s]? to valve[s]? ([[:alpha:] ,]+)`)

    ret := make([]valve, 0, len(lines))

    for _, line := range lines {
        match := re.FindStringSubmatch(line)

        name := match[1]
        rate, _ := strconv.Atoi(match[2])
        adj := strings.Split(match[3], ", ")

        ret = append(ret, valve{name, rate, adj})
    }

    return ret
}

func make3d(x, y, z int) [][][]int {
    ret := make([][][]int, x)
    for i := 0; i < x; i++ {
        ret[i] = make([][]int, y)
        for j := 0; j < y; j++ {
            ret[i][j] = make([]int, z)
            for k := 0; k < z; k++ {
                ret[i][j][k] = -1
            }
        }
    }

    return ret
}

func extractPos(code, cur, total int) int {
    for i := 0; i < cur; i++ {
        code /= total
    }

    return code % total
}

func deltaPos(delta, cur, total int) int {
    for i := 0; i < cur; i++ {
        delta *= total
    }

    return delta
}

func solveDp(valves []valve, totalTime int, start []string) int {
    // enumerate nonzero valves
    nzValveCnt, valveIdx := 0, make([]int, len(valves))
    for i, valve := range valves {
        valveIdx[i] = -1

        if valve.rate > 0 {
            valveIdx[i] = nzValveCnt
            nzValveCnt++
        }
    }

    totalMask := 1 << uint(nzValveCnt)

    actors := len(start)
    totalPos, startPos := 1, 0
    for _, curStart := range start {
        totalPos *= len(valves)

        for j, valve := range valves {
            if valve.name == curStart {
                startPos = startPos * len(valves) + j
                break
            }
        }
    }

    // dp[time][cur][mask] - (totalTime*actors + 1) * valveCnt**actors[=totalPos] * 2**nzValveCnt
    // storing only two adjacent layers
    dp := make3d(2, totalPos, totalMask)

    dp[0][startPos][0] = 0

    for i := 0; i < totalTime * actors; i++ {
        cActor, cTime := i % actors, i / actors

        layer := i & 1

        // open valve
        for j := 0; j < totalPos; j++ {
            pos := extractPos(j, cActor, len(valves))
            vidx := valveIdx[pos]

            if vidx == -1 {
                continue
            }

            mask := 1 << uint(vidx)
            for k := 0; k < totalMask; k++ {
                if dp[layer][j][k] != -1 && (mask & k) == 0 {
                    nk := mask | k
                    totalRate := (totalTime - cTime - 1) * valves[pos].rate
                    dp[layer^1][j][nk] = utils.Max(dp[layer^1][j][nk], dp[layer][j][k] + totalRate)
                }
            }
        }

        // just move somewhere
        for j := 0; j < totalPos; j++ {
            // redundant work for part 2, but don't really care
            pos := extractPos(j, cActor, len(valves))

            for _, adjName := range valves[pos].adj {
                idx := -1

                for ci, valve := range valves {
                    if valve.name == adjName {
                        idx = ci
                        break
                    }
                }

                npos := deltaPos(idx - pos, cActor, len(valves)) + j

                for k := 0; k < totalMask; k++ {
                    dp[layer^1][npos][k] = utils.Max(dp[layer^1][npos][k], dp[layer][j][k])
                }
           }
        }
    }

    ans := -1
    for j := 0; j < totalPos; j++ {
        for k := 0; k < totalMask; k++ {
            ans = utils.Max(ans, dp[(totalTime * actors) & 1][j][k])
        }
    }

    return ans
}

func partOne(lines []string) {
    valves := parse(lines)
    maxRate := solveDp(valves, 30, []string{"AA"})
    fmt.Println(maxRate)
}

func partTwo(lines []string) {
    valves := parse(lines)
    maxRate := solveDp(valves, 26, []string{"AA", "AA"})
    fmt.Println(maxRate)
}

func tests() {
    tests := utils.ReadLines("tests", "16.txt")
    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "16.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
