package main

import (
    "fmt"
    "sort"
    "../utils"
)

type packet struct {
    value int

    // nil for integer leafs
    list *[]packet
}

type packetPair struct {
    left, right packet
}

func listify(p packet) []packet {
    if p.list == nil {
        inner := packet{p.value, nil}
        return []packet{inner}
    }

    return *p.list
}

// -1,0,1 return value
func cmp(left packet, right packet) int {
    // both integers
    if left.list == nil && right.list == nil {
        li, ri := left.value, right.value

        if li < ri {
            return -1
        }

        if li > ri {
            return 1
        }

        return 0
    }

    ll, rl := listify(left), listify(right)

    for i := 0; i < len(ll) || i < len(rl); i++ {
        if i == len(ll) {
            return -1
        }

        if i == len(rl) {
            return 1
        }

        if cmpv := cmp(ll[i], rl[i]); cmpv != 0 {
            return cmpv
        }
    }

    return 0
}

func firstDigit(s string) int {
    firstCharCode := int(s[0])

    if int('0') <= firstCharCode && firstCharCode <= int('9') {
        return firstCharCode - int('0')
    }

    return -1
}

func parseInt(s string) (value int, eol string) {
    value, eol = 0, s

    for firstDigit(eol) != -1 {
        value, eol = value * 10 + firstDigit(eol), eol[1:]
    }

    return
}

func parsePacket(line string) (p packet, eol string) {

    if firstDigit(line) != -1 {
        var value int
        value, eol = parseInt(line)
        p = packet{value, nil}
        return
    }

    list := make([]packet, 0)

    // skip '['
    eol = line[1:]

    for true {
        sep := eol[0]

        if sep == ']' {
            eol = eol[1:]
            break
        }

        if sep == ',' {
            eol = eol[1:]
            continue
        }

        var inner packet
        inner, eol = parsePacket(eol)

        list = append(list, inner)
    }

    p = packet{0, &list}
    return
}

func parsePairs(lines []string) []packetPair {
    pairs := make([]packetPair, 0, len(lines)/3)

    for i := 0; i < len(lines); i += 3 {
        left, _ := parsePacket(lines[i])
        right, _ := parsePacket(lines[i+1])

        pairs = append(pairs, packetPair{left, right})
    }

    return pairs
}

func partOne(lines []string) {
    pairs := parsePairs(lines)

    sum := 0
    for i, pair := range pairs {
        if cmp(pair.left, pair.right) == -1 {
            sum += i + 1
        }
    }

    fmt.Println(sum)
}

type allPackets []packet

func (a allPackets) Len() int {
    return len(a)
}

func (a allPackets) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a allPackets) Less(i, j int) bool {
    return cmp(a[i], a[j]) == -1
}

func partTwo(lines []string) {
    pairs := parsePairs(lines)
    all := allPackets(make([]packet, 0, (len(pairs)+1) * 2))

    for _, pair := range pairs {
        all = append(all, pair.left, pair.right)
    }

    a, _ := parsePacket("[[2]]")
    b, _ := parsePacket("[[6]]")
    all = append(all, a, b)

    sort.Sort(all)

    li := sort.Search(len(all), func (i int) bool { return cmp(all[i], a) >= 0 })
    lj := sort.Search(len(all), func (i int) bool { return cmp(all[i], b) >= 0 })

    fmt.Println((li + 1) * (lj + 1))
}

func tests() {
    tests := utils.ReadLines("tests", "13.txt")
    partOne(tests)
    partTwo(tests)
}

func main() {
    tests()

    puzzle := utils.ReadLines("inputs", "13.txt")
    partOne(puzzle)
    partTwo(puzzle)
}
