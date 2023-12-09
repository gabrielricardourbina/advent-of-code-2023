package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

type FuncPiece struct {
	from int
	to   int
	a    int
}

type SortBy []FuncPiece

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].from < a[j].from }

func mappingIntoPiecewise(sourceText string, label string) []FuncPiece {
	start := strings.Index(sourceText, label) + len(label) + len(" map:")
	end := strings.Index(sourceText[start:], "\n\n") + start
	lines := strings.Split(strings.TrimSpace(sourceText[start:end]), "\n")
	parsed := make([]FuncPiece, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		destination := stringToInt(fields[0])
		source := stringToInt(fields[1])
		rangeSize := stringToInt(fields[2])
		parsed[i].from = source
		parsed[i].to = source + rangeSize
		parsed[i].a = destination - source
	}
	sort.Sort(SortBy(parsed))
	filled := make([]FuncPiece, len(lines)*2+1)

	for i, piece := range parsed {
		if i+1 < len(parsed) {
			nextPiece := parsed[i+1]
			filled[(i+1)*2-1] = piece
			filled[(i+1)*2] = FuncPiece{piece.to, nextPiece.from, 0}
		} else {
			filled[(i+1)*2-1] = piece
		}
	}
	filled[0] = FuncPiece{0, filled[1].from, 0}
	filled[len(filled)-1] = FuncPiece{filled[len(filled)-2].to, math.MaxInt, 0}

	return filled
}

func MapPieceRangeToTargetDomain(
	sourcePiece FuncPiece,
	targetFunc []FuncPiece,
	output *[]FuncPiece,
) {
	for _, targetPiece := range targetFunc {
		if sourcePiece.from >= sourcePiece.to {
			break
		}

		FrMax, FrMin := sourcePiece.to+sourcePiece.a, sourcePiece.from+sourcePiece.a
		FdMax := sourcePiece.to
		GdMax, GdMin := targetPiece.to, targetPiece.from

		nextPiece := FuncPiece{sourcePiece.from, -1, sourcePiece.a + targetPiece.a}
		if FrMin >= GdMin && FrMin < GdMax {
			if FrMax <= GdMax {
				nextPiece.to = FdMax
			} else if FrMax > GdMax {
				nextPiece.to = GdMax - sourcePiece.a
			}
			sourcePiece.from = nextPiece.to
			*output = append(*output, nextPiece)
		}
	}

}

func ComposePieceWiseFuncs(targetFunc []FuncPiece, sources []FuncPiece) []FuncPiece {
	results := make([]FuncPiece, 0, len(sources))
	for _, sourcePiece := range sources {
		MapPieceRangeToTargetDomain(sourcePiece, targetFunc, &results)
	}
	return results
}

type MinInt int

func MakeMinInt() MinInt {
	return -1
}

func (self *MinInt) set(val int) {
	if *self == -1 || val < int(*self) {
		*self = MinInt(val)
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

	seedToSoil := mappingIntoPiecewise(text, "seed-to-soil")
	soilToFertilizer := mappingIntoPiecewise(text, "soil-to-fertilizer")
	fertilizerToWater := mappingIntoPiecewise(text, "fertilizer-to-water")
	waterToLight := mappingIntoPiecewise(text, "water-to-light")
	lightToTemperature := mappingIntoPiecewise(text, "light-to-temperature")
	temperatureToHumidity := mappingIntoPiecewise(text, "temperature-to-humidity")
	humidityToLocation := mappingIntoPiecewise(text, "humidity-to-location")

	C := ComposePieceWiseFuncs
	seedToLocationPieceWise := C(humidityToLocation,
		C(temperatureToHumidity,
			C(lightToTemperature,
				C(waterToLight,
					C(fertilizerToWater,
						C(soilToFertilizer,
							seedToSoil))))))

	minLocP1 := MakeMinInt()
	for _, input := range seeds {
		for _, piece := range seedToLocationPieceWise {
			if input >= piece.from && input < piece.to {
				minLocP1.set(input + piece.a)
			}
		}
	}
	fmt.Println("Part 1, lowest location number that corresponds to any of the initial seed numbers:", minLocP1)

	minLocP2 := MakeMinInt()

	for i := 0; i < len(seeds); i += 2 {
		rangeFrom := seeds[i]
		rangeTo := rangeFrom + seeds[i+1]
		for _, piece := range seedToLocationPieceWise {
			if rangeFrom >= rangeTo {
				break
			}
			if rangeFrom >= piece.from && rangeFrom < piece.to {
				minLocP2.set(rangeFrom + piece.a)
				if rangeTo <= piece.to {
					rangeFrom = rangeTo
					minLocP2.set(rangeTo - 1 + piece.a)
				} else {
					rangeFrom = piece.to
				}
			}
		}
	}
	fmt.Println("Part 2, lowest location number that corresponds to any of the initial seed ranges:", minLocP2)
}
