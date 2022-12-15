package main

import (
    "fmt"
    "regexp"
    "strconv"
    "../utils"
)

type coord struct {
    x, y int
}

func (a coord) mdist(b coord) int {
    return utils.Abs(a.x - b.x) + utils.Abs(a.y - b.y)
}

type reading struct {
    sensor, beacon coord
}

func (r reading) span() int {
    return r.beacon.mdist(r.sensor)
}

func (r reading) impossibru(c coord) bool {
    return c.mdist(r.sensor) <= r.span()
}

func parseReadings(lines []string) []reading {
    re := regexp.MustCompile(
        `Sensor at x=([-[:digit:]]+), y=([-[:digit:]]+): ` +
            `closest beacon is at x=([-[:digit:]]+), y=([-[:digit:]]+)`)

    ret := make([]reading, 0, len(lines))

    for _, line := range lines {
        var r reading
        match := re.FindStringSubmatch(line)

        r.sensor.x, _ = strconv.Atoi(match[1])
        r.sensor.y, _ = strconv.Atoi(match[2])
        r.beacon.x, _ = strconv.Atoi(match[3])
        r.beacon.y, _ = strconv.Atoi(match[4])

        ret = append(ret, r)
    }

    return ret
}

const inf = 1000000000

type span struct {
    min, max int
}

func (s *span) expand(value int) {
    s.min, s.max = utils.Min(s.min, value), utils.Max(s.max, value)
}

func partOne(lines []string, cy int) {
    readings := parseReadings(lines)

    s, maxd := span{inf, -inf}, 0
    for _, r := range readings {
        s.expand(r.sensor.x)
        s.expand(r.beacon.x)

        maxd = utils.Max(maxd, r.sensor.mdist(r.beacon))
    }

    s = span{s.min - maxd, s.max + maxd}

    ans := 0
    for cx := s.min; cx <= s.max; cx++ {
        c := coord{cx, cy}

        noBeacon, isBeacon := false, false
        for _, r := range readings {
            if r.impossibru(c) {
                noBeacon = true
            }

            if c == r.beacon {
                isBeacon = true
            }
        }

        if noBeacon && !isBeacon {
            ans++
        }
    }

    fmt.Println(ans)
}

func boundaryIter(r reading, cb func (c coord)) {
    dist := r.span() + 1

    for dx := 0; dx <= dist; dx++ {
        dy := dist - dx

        cb(coord{r.sensor.x + dx, r.sensor.y + dy})
        cb(coord{r.sensor.x + dx, r.sensor.y - dy})
        cb(coord{r.sensor.x - dx, r.sensor.y + dy})
        cb(coord{r.sensor.x - dx, r.sensor.y - dy})
    }
}

func partTwo(lines []string, maxc int) {
    readings := parseReadings(lines)

    var found coord

    for _, r := range readings {
        boundaryIter(r, func (c coord) {
            if c.x < 0 || c.y < 0 || c.x > maxc || c.y > maxc {
                return
            }

            for _, rs := range readings {
                if rs.impossibru(c) {
                    return
                }
            }

            found = c
        })
    }

    fmt.Println(found.x * 4000000 + found.y)
}

func tests() {
    tests := utils.ReadLines("tests", "15.txt")
    partOne(tests, 10)
    partTwo(tests, 20)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "15.txt")
    partOne(puzzle, 2000000)
    partTwo(puzzle, 4000000)
}
