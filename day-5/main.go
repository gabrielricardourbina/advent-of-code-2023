package main

import (
	"fmt"
	"log"
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

type mapping struct {
	destination int
	source      int
	rangeSize   int
}

func parseMapping(sourceText string, label string) func(source int) int {
	start := strings.Index(sourceText, label) + len(label) + len(" map:")
	end := strings.Index(sourceText[start:], "\n\n") + start
	lines := strings.Split(strings.TrimSpace(sourceText[start:end]), "\n")
	parsed := make([]mapping, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		parsed[i].destination = stringToInt(fields[0])
		parsed[i].source = stringToInt(fields[1])
		parsed[i].rangeSize = stringToInt(fields[2])
	}
	return func(source int) int {
		for _, line := range parsed {
			if source >= line.source && source < line.source+line.rangeSize {
				return line.destination + (source - line.source)
			}
		}

		return source
	}
}

func main() {
	file, err := os.ReadFile("./input.dat")
	if err != nil {
		log.Fatal(err)
	}
	text := string(append(file, 10, 10))
	eol0 := strings.Index(text, "\n")
	seedsAsString := strings.Fields(text[len("seeds:"):eol0])
	seeds := make([]int, len(seedsAsString))

	for i, seed := range seedsAsString {
		seeds[i] = stringToInt(seed)
	}

	seedToSoil := parseMapping(text, "seed-to-soil")
	soilToFertilizer := parseMapping(text, "soil-to-fertilizer")
	fertilizerToWater := parseMapping(text, "fertilizer-to-water")
	waterToLight := parseMapping(text, "water-to-light")
	lightToTemperature := parseMapping(text, "light-to-temperature")
	temperatureToHumidity := parseMapping(text, "temperature-to-humidity")
	humidityToLocation := parseMapping(text, "humidity-to-location")

	seedToLocation := func(seed int) int {
		return humidityToLocation(temperatureToHumidity(lightToTemperature(waterToLight(fertilizerToWater(soilToFertilizer(seedToSoil(seed)))))))
	}

	minLocP1 := -1
	for _, seed := range seeds {
		location := seedToLocation(seed)
		if minLocP1 == -1 || location < minLocP1 {
			minLocP1 = location
		}
	}

	minLocP2 := -1
	for i := 0; i < len(seeds); i += 2 {
		start, rangeSize := seeds[i], seeds[i+1]
		for j := 0; j < rangeSize; j++ {
			location := seedToLocation(start + j)
			if minLocP2 == -1 || location < minLocP2 {
				minLocP2 = location
			}
		}
	}
	fmt.Println("Part 1, lowest location number that corresponds to any of the initial seed numbers:", minLocP1)
	fmt.Println("Part 2, lowest location number that corresponds to any of the initial seed ranges:", minLocP2)
}
