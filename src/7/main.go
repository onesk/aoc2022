package main

import (
    "fmt"
    "strings"
    "strconv"
    "../utils"
)

func basePath(path string) string {
    lastSlash := strings.LastIndex(path, "/")

    if lastSlash == -1 {
        lastSlash = 0
    }

    return path[:lastSlash]
}

func incrementAlongPath(sizes map[string]int, path string, size int) {
    for true {
        sizes[path] += size

        if path == "" {
            return
        }

        path = basePath(path)
    }
}

func simulate(cmds []string) map[string]int {
    sizes := make(map[string]int)

    cwd := ""

    for _, cmd := range cmds {
        parts := strings.SplitN(cmd, " ", 3)

        if len(parts) == 3 && parts[0] == "$" && parts[1] == "cd" {
            // cd command
            arg := parts[2]

            switch arg {
            case "/":
                cwd = ""
            case "..":
                cwd = basePath(cwd)
            default:
                cwd += "/" + arg
            }

        } else if len(parts) == 2 {
            // listing line
            dirOrSize := parts[0]

            if dirOrSize == "dir" {
                // do nothing for directories
                continue
            }

            size, _ := strconv.Atoi(dirOrSize)

            incrementAlongPath(sizes, cwd, size)
        }

        // "$ ls" is noneffectual, just skip
    }

    return sizes
}

func partOne(cmds []string) {
    dirTree := simulate(cmds)

    sum := 0
    for _, size := range dirTree {
        if size <= 100000 {
            sum += size
        }
    }

    fmt.Println(sum)
}

func partTwo(cmds []string) {
    dirTree := simulate(cmds)

    curTotal := dirTree[""]
    threshold := 30000000 - (70000000 - curTotal)

    min := 1000000000
    for _, size := range dirTree {
        if size > threshold {
            min = utils.Min(min, size)
        }
    }

    fmt.Println(min)
}

func tests() {
    tests := utils.ReadLines("tests", "7.txt")
    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "7.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
