package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func doubleEachTime(n int) int {
	return ((1 << (n)) >> 1)
}
func main() {
	file, err := os.Open("./input.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	parsedCards := make([]int, 0, 1000)
	for scanner.Scan() {
		line := scanner.Text()
		i0 := strings.Index(line, ":") + 1
		cardContent := strings.Split(line[i0:], "|")
		winningNumbers := strings.Fields(cardContent[0])
		cardNumbers := strings.Fields(cardContent[1])
		winningCount := 0

		for _, number := range cardNumbers {
			if slices.Contains(winningNumbers, number) {
				winningCount++
			}
		}

		parsedCards = append(parsedCards, winningCount)
	}

	cardPoints := 0
	cardCopies := make([]int, len(parsedCards))

	for i, _ := range cardCopies {
		cardCopies[i] = 1
		cardPoints += doubleEachTime(parsedCards[i])

	}
	totalCards := 0
	for i, wins := range parsedCards {
		multiplier := cardCopies[i]
		totalCards += multiplier
		for j := 1; j <= wins; j++ {
			cardCopies[i+j] += multiplier
		}
	}

	fmt.Println("Part 1, cards are worth a total of:", cardPoints)
	fmt.Println("Part 2, total scratchcards:", totalCards)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
