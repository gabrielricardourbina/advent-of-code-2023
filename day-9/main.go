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

func ComputeMultipliers() [][]int {
	multipliers := make([][]int, 30)
	multipliers[0] = []int{1}
	multipliers[1] = []int{1, 1}
	for i := 2; i < len(multipliers); i++ {
		multipliers[i] = make([]int, i+1)
		multipliers[i][0] = 1
		for j := 1; j < i; j++ {
			multipliers[i][j] = (multipliers[i-1][j] + multipliers[i-1][j-1])
		}
		multipliers[i][i] = 1
	}

	for i := range multipliers {
		sign := 1
		for j := range multipliers[i] {
			multipliers[i][j] *= sign
			sign *= -1
		}
	}
	return multipliers
}

func ExtrapolatePrev(sequence []int, multipliers [][]int) int {
	prevValue := 0
	sign := 1

	for i := 0; i < len(sequence); i++ {
		rowSum := 0
		q := len(multipliers[i]) - 1
		for j, m := range multipliers[i] {
			rowSum += m * sequence[q-j]
		}
		prevValue += sign * rowSum
		sign *= -1
	}

	return prevValue
}

func ExtrapolateNext(sequence []int, multipliers [][]int) int {
	nextValue := 0
	n := len(sequence) - 1
	for i := 0; i < len(sequence); i++ {
		rowSum := 0
		for j, m := range multipliers[i] {
			rowSum += m * sequence[n-j]
		}
		nextValue += rowSum
	}

	return nextValue
}

func main() {
	file, err := os.ReadFile("./input.dat")
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)
	lines := strings.Split(text, "\n")
	histories := make([][]int, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		histories[i] = make([]int, len(fields))
		for j, value := range fields {
			histories[i][j] = StringToInt(value)
		}
	}
	multipliers := ComputeMultipliers()
	nextSum := 0
	prevSum := 0
	for i := range histories {
		nextValue := ExtrapolateNext(histories[i], multipliers)
		prevValue := ExtrapolatePrev(histories[i], multipliers)
		nextSum += nextValue
		prevSum += prevValue
	}

	fmt.Println("Part 1, sum of next values:", nextSum)
	fmt.Println("Part 2, sum of prev values:", prevSum)
}
