package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func stringToInt(text string) int {
	number, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal("should be a number", text)
	}
	return number
}

func nextInteger(num float64) float64 {
	ceilled := math.Ceil(num)
	if ceilled == num {
		return ceilled + 1
	}
	return ceilled
}

func prevInteger(num float64) float64 {
	floored := math.Floor(num)
	if floored == num {
		return floored - 1
	}
	return floored
}

func MinMaxRecord(duration int, record int) (int, int) {
	root := math.Sqrt(math.Pow(float64(duration), 2) - (4 * float64(record)))
	return int(nextInteger((float64(duration) - root) / 2)), int(prevInteger((float64(duration) + root) / 2))
}

func main() {
	file, err := os.ReadFile("./input.dat")
	if err != nil {
		log.Fatal(err)
	}
	text := string(append(file, 10))
	lines := strings.Split(text, "\n")
	raceDuration := strings.Fields(lines[0][strings.Index(lines[0], ":")+1:])
	travelRecord := strings.Fields(lines[1][strings.Index(lines[1], ":")+1:])

	nBeats := 1
	for i := 0; i < len(raceDuration); i++ {
		duration := stringToInt(raceDuration[i])
		record := stringToInt(travelRecord[i])
		min, max := MinMaxRecord(duration, record)
		nBeats *= 1 + max - min
	}

	fmt.Println("Part 1, number of ways the records can be beaten:", nBeats)

	duration := stringToInt(strings.Join(raceDuration, ""))
	record := stringToInt(strings.Join(travelRecord, ""))
	min, max := MinMaxRecord(duration, record)
	nBeatsP2 := 1 + max - min

	fmt.Println("Part 2, number of ways the long race record can be beaten:", nBeatsP2)
}
