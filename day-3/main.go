package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type P struct {
	i int
	j int
}

func isSymbol(char string) bool {
	if _, err := strconv.Atoi(char); err == nil {
		return false
	}
	return char != "."
}

func main() {
	file, err := os.ReadFile("./input.dat")
	if err != nil {
		log.Fatal(err)
	}
	schematics := string(file)
	tableView := strings.Split(schematics, "\n")
	capturedSum := 0
	gearMap := make(map[P][]int)
	for i, row := range tableView {
		for j := 0; j < len(row); j++ {
			capturedNum := -1

			hasAdjacentSymbol := false
			capturedGears := make(map[P]bool)

			setAdjacentSybol2 := func(i int, j int) {
				char := tableView[i][j : j+1]
				hasAdjacentSymbol = hasAdjacentSymbol || isSymbol(char)

				if char == "*" {
					capturedGears[P{i, j}] = true
				}
			}
			for k := j; k < len(row); k++ {
				digits := row[j : k+1]
				number, err := strconv.Atoi(digits)
				if err != nil {
					j = k
					break
				} else {
					capturedNum = number
					hasLeft := k > 0
					hasRight := k < len(row)-1
					if i > 0 {
						if hasLeft {
							setAdjacentSybol2(i-1, k-1)
						}
						setAdjacentSybol2(i-1, k)
						if hasRight {
							setAdjacentSybol2(i-1, k+1)
						}
					}
					if i < len(tableView)-1 {
						if hasLeft {
							setAdjacentSybol2(i+1, k-1)
						}
						setAdjacentSybol2(i+1, k)
						if hasRight {
							setAdjacentSybol2(i+1, k+1)
						}
					}
					if hasLeft {
						setAdjacentSybol2(i, k-1)
					}
					if hasRight {
						setAdjacentSybol2(i, k+1)
					}
				}
			}
			if capturedNum >= 0 {
				for point := range capturedGears {
					gearMap[point] = append(gearMap[point], capturedNum)
				}
				if hasAdjacentSymbol {
					capturedSum += capturedNum
				}
			}
		}
	}
	gearRatiosSum := 0
	for _, ratios := range gearMap {
		if len(ratios) == 2 {
			gearRatiosSum += ratios[0] * ratios[1]
		}
	}
	fmt.Println("Part 1,sum of all of the part numbers in the engine schematic", capturedSum)
	fmt.Println("Part 2, sum of all of the gear ratios in the engine schematic", gearRatiosSum)

}
