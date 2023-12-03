package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var digitTexts = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func findTextDigit(slice string) (int, bool) {
	for d := 0; d < len(digitTexts); d++ {
		digit := digitTexts[d]
		if !strings.HasPrefix(slice, digit) {
			continue
		}
		return d, true
	}
	return -1, false
}

func findNumberDigit(char string) (int, bool) {
	digit, err := strconv.Atoi(char)
	if err != nil {
		return -1, false
	}
	return digit, true
}

func main() {

	file, err := os.Open("./input.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		firstDigitInLine := -1
		lastDigitInLine := -1

		for i := 0; i < len(line); i++ {
			var digit int
			var tFound, nFound bool

			digit, nFound = findNumberDigit(line[i : i+1])

			if !nFound {
				digit, tFound = findTextDigit(line[i:])
			}

			if tFound || nFound {
				if firstDigitInLine < 0 {
					firstDigitInLine = digit
				}
				lastDigitInLine = digit
			}
		}
		if lastDigitInLine == -1 {
			log.Fatal("last digit not found")
		}
		if firstDigitInLine == -1 {
			log.Fatal("first digit not found")
		}
		totalSum += firstDigitInLine*10 + lastDigitInLine
	}

	fmt.Println("Part 2, sum of all of the calibration values", totalSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
