package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func StringToInt(text string) int {
	number, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal("should be a number", text)
	}
	return number
}
func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

type Node struct {
	id uint16
	L  int
	R  int
}

func CompressNodeId(id string) uint16 {
	return (uint16(id[0]-65) << 0) | (uint16(id[1]-65) << 5) | (uint16(id[2]-65) << 10)
}

func DecompressNodeId(id uint16) string {
	fiveBitsMask := uint16(1<<5) - 1
	return string([]rune{
		rune((((id >> 0) & fiveBitsMask) + 65)),
		rune((((id >> 5) & fiveBitsMask) + 65)),
		rune((((id >> 10) & fiveBitsMask) + 65)),
	})
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a int, steps ...int) int {
	if len(steps) == 0 {
		return a
	}
	b := steps[0]
	tail := steps[1:]
	result := a * b / GCD(a, b)

	for i := 0; i < len(tail); i++ {
		result = LCM(result, tail[i])
	}

	return result
}

var A = []rune("A")[0]
var L = []rune("L")[0]
var Z = []rune("Z")[0]

func CountSteps(lines []string, isSource func(nodeId uint16) bool, isDestination func(nodeId uint16) bool) int {
	path := lines[0]
	textNodes := lines[2:]
	nodes := make([]Node, len(textNodes))
	sources := make([]int, 0, 26*26)
	blank := -1

	for i, line := range textNodes {
		nodes[i].id = CompressNodeId(line[0:3])
		nodes[i].L = blank
		nodes[i].R = blank
		if isSource(nodes[i].id) {
			sources = append(sources, i)
		}
	}
	for i, line := range textNodes {
		left := CompressNodeId(line[7:10])
		right := CompressNodeId(line[12:15])
		for j := range nodes {
			if nodes[i].R != blank && nodes[i].L != blank {
				break
			}
			if nodes[j].id == left {
				nodes[i].L = j
			}
			if nodes[j].id == right {
				nodes[i].R = j
			}
		}
	}

	stepsPerSource := make([]int, len(sources))
	for q, source := range sources {
	ready:
		for {
			for _, dir := range path {
				stepsPerSource[q]++

				if dir == L {
					source = nodes[source].L
				} else {
					source = nodes[source].R
				}

				if isDestination(nodes[source].id) {
					break ready
				}
			}
		}
	}
	return LCM(stepsPerSource[0], stepsPerSource[1:]...)

}
func main() {
	file, err := os.ReadFile("./input.dat")
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)
	lines := strings.Split(text, "\n")

	aaaEncoded := CompressNodeId("AAA")
	zzzEncoded := CompressNodeId("ZZZ")

	part1 := CountSteps(
		lines,
		func(id uint16) bool { return id == aaaEncoded },
		func(id uint16) bool { return id == zzzEncoded },
	)
	fmt.Println("Part 1, how many steps are required to reach ZZZ from AAA:", part1)

	fiveBitsMask := uint16(1<<5) - 1

	part2 := CountSteps(
		lines,
		func(id uint16) bool { return rune((((id >> 10) & fiveBitsMask) + 65)) == A },
		func(id uint16) bool { return rune((((id >> 10) & fiveBitsMask) + 65)) == Z },
	)
	fmt.Println("Part 2, how many steps are required to reach **Z from **A:", part2)
}
