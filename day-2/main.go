package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isImposible(color string, count int) bool {
	return (color == "red" && count > 12) || (color == "blue" && count > 14) || (color == "green" && count > 13)
}

func main() {
	file, err := os.Open("./input.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sumGames := 0
	sumPower := 0
	gameId := 0

	for scanner.Scan() {
		gameId += 1
		line := scanner.Text()
		i0 := strings.Index(line, ":") + 1
		posibleGame := true
		maxGreen, maxRed, maxBlue := 0, 0, 0
		for _, game := range strings.Split(line[i0:], ";") {

			for _, cubeSet := range strings.Split(game, ",") {
				trimmed := strings.Trim(cubeSet, " ")
				splitted := strings.Split(trimmed, " ")
				color := splitted[1]
				count, countErr := strconv.Atoi(splitted[0])

				if countErr != nil {
					log.Fatal(err)
				}

				if color == "green" && count > maxGreen {
					maxGreen = count
				}

				if color == "blue" && count > maxBlue {
					maxBlue = count
				}

				if color == "red" && count > maxRed {
					maxRed = count
				}

				if isImposible(color, count) {
					posibleGame = false
				}
			}
		}
		sumPower += (maxGreen * maxRed * maxBlue)
		if posibleGame {
			sumGames += gameId
		}
	}
	fmt.Println("Part 1, sum of posible game Ids:", sumGames)
	fmt.Println("Part 2, sum of min set posible powers:", sumPower)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
