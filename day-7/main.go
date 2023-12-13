package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

var CardValuesWithJoker = map[rune]int{
	[]rune("J")[0]: 0,
	[]rune("2")[0]: 1,
	[]rune("3")[0]: 2,
	[]rune("4")[0]: 3,
	[]rune("5")[0]: 4,
	[]rune("6")[0]: 5,
	[]rune("7")[0]: 6,
	[]rune("8")[0]: 7,
	[]rune("9")[0]: 8,
	[]rune("T")[0]: 9,
	[]rune("Q")[0]: 10,
	[]rune("K")[0]: 11,
	[]rune("A")[0]: 12,
}

var CardValues = map[rune]int{
	[]rune("2")[0]: 0,
	[]rune("3")[0]: 1,
	[]rune("4")[0]: 2,
	[]rune("5")[0]: 3,
	[]rune("6")[0]: 4,
	[]rune("7")[0]: 5,
	[]rune("8")[0]: 6,
	[]rune("9")[0]: 7,
	[]rune("T")[0]: 8,
	[]rune("J")[0]: 9,
	[]rune("Q")[0]: 10,
	[]rune("K")[0]: 11,
	[]rune("A")[0]: 12,
}
var Base13Scale = []int{
	IntPow(13, 0),
	IntPow(13, 1),
	IntPow(13, 2),
	IntPow(13, 3),
	IntPow(13, 4),
	IntPow(13, 5),
}

type Hand struct {
	cards string
	bid   int
}
type HandWithJoker struct {
	Hand
}

type CardSubStack struct {
	card  rune
	count int
}

func SortStack(a, b CardSubStack) int {
	if a.count > b.count {
		return -1
	} else if a.count < b.count {
		return 1
	}
	return 0
}
func GetCardStackWithJokers(hand Hand) []CardSubStack {
	countMap := make(map[rune]int, 5)

	for _, char := range hand.cards {
		countMap[char] += 1
	}

	jokerRune := []rune("J")[0]
	cardStack := make([]CardSubStack, 0, 5)
	jokerCount := 0

	for card, count := range countMap {
		if card == jokerRune {
			jokerCount += count
		} else {
			cardStack = append(cardStack, CardSubStack{card, count})
		}
	}

	slices.SortFunc(cardStack, SortStack)

	if len(cardStack) == 0 && jokerCount == 5 {
		cardStack = append(cardStack, CardSubStack{jokerRune, 5})
	} else {
		cardStack[0].count += jokerCount
	}
	return cardStack
}

func GetCardStack(hand Hand) []CardSubStack {
	countMap := make(map[rune]int, 5)

	for _, char := range hand.cards {
		countMap[char] += 1
	}

	cardStack := make([]CardSubStack, 0, 5)
	for card, count := range countMap {
		cardStack = append(cardStack, CardSubStack{card, count})
	}
	slices.SortFunc(cardStack, SortStack)

	return cardStack
}
func EvalCardStackStrength(cardStack []CardSubStack) int {
	if cardStack[0].count == 5 {
		return 6 // Five of a kind
	}
	if cardStack[0].count == 4 && cardStack[1].count == 1 {
		return 5 // Four of a kind
	}
	if cardStack[0].count == 3 && cardStack[1].count == 2 {
		return 4 // Full house
	}
	if cardStack[0].count == 3 && cardStack[1].count == 1 && cardStack[2].count == 1 {
		return 3 // Three of a kind
	}
	if cardStack[0].count == 2 && cardStack[1].count == 2 && cardStack[2].count == 1 {
		return 2 // Two pair
	}
	if cardStack[0].count == 2 {
		return 1 // One pair
	}
	return 0 // High card
}

func HandStrengthWithJoker(hand Hand) int {
	return EvalCardStackStrength(GetCardStackWithJokers(hand))
}

func HandStrength(hand Hand) int {
	return EvalCardStackStrength(GetCardStack(hand))

}

func Base13Value(tokens string, cardValues map[rune]int) int {
	base13Value := 0

	for k, char := range tokens {
		n := len(tokens) - 1 - k
		base13Value += Base13Scale[n] * cardValues[char]
	}

	return base13Value
}

func CalculateWininnings(lines []string, cardValues map[rune]int, handStrength func(hand Hand) int) int {
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		hands[i].cards = fields[0]
		hands[i].bid = StringToInt(fields[1])
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		aBase13Value := Base13Value(a.cards, cardValues) + Base13Scale[5]*handStrength(a)
		bBase13Value := Base13Value(b.cards, cardValues) + Base13Scale[5]*handStrength(b)

		if aBase13Value < bBase13Value {
			return -1
		} else if aBase13Value > bBase13Value {
			return 1
		}
		return 0
	})

	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.bid
	}
	return totalWinnings
}

func main() {
	file, err := os.ReadFile("./input.dat")
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)
	lines := strings.Split(text, "\n")

	fmt.Println("Part 1, total winnings:", CalculateWininnings(lines, CardValues, HandStrength), 251927063)

	fmt.Println("Part 2, total winnings with Jokers:", CalculateWininnings(lines, CardValuesWithJoker, HandStrengthWithJoker), 255632664)

}
