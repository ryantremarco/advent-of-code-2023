package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input
var input string

const newline = "\n"

func main() {
	parsedInput := parseInput(input)
	fmt.Println("Part One:", partOne(parsedInput))
	fmt.Println("Part Two:", partTwo(parsedInput))
}

type Card struct {
	ID             int
	Count          int
	WinningNumbers []int
	ScratchNumbers []int
}

func parseInput(input string) []Card {
	var cards []Card

	for strings.Contains(input, "  ") {
		input = strings.ReplaceAll(input, "  ", " ")
	}

	for _, line := range strings.Split(input, newline) {
		if line == "" {
			continue
		}

		card := Card{Count: 1}

		idSplit := strings.Split(line, ": ")
		card.ID = toInt(strings.Split(idSplit[0], " ")[1])

		numbersSplit := strings.Split(idSplit[1], " | ")
		winningStrs := strings.Split(numbersSplit[0], " ")
		scratchStrs := strings.Split(numbersSplit[1], " ")

		for _, winning := range winningStrs {
			card.WinningNumbers = append(card.WinningNumbers, toInt(winning))
		}

		for _, scratch := range scratchStrs {
			card.ScratchNumbers = append(card.ScratchNumbers, toInt(scratch))
		}

		cards = append(cards, card)
	}

	return cards
}

func toInt[T ~string | ~rune | ~byte](t T) int {
	i, err := strconv.Atoi(string(t))
	if err != nil {
		panic(err)
	}
	return i
}

func contains[T comparable](haystack []T, needle T) bool {
	for _, hay := range haystack {
		if hay == needle {
			return true
		}
	}

	return false
}

func partOne(cards []Card) int {
	winnings := 0

	for _, card := range cards {
		winners := 0
		for _, scratch := range card.ScratchNumbers {
			if contains(card.WinningNumbers, scratch) {
				winners++
			}
		}

		if winners > 0 {
			winnings += int(math.Pow(2, float64(winners-1)))
		}
	}

	return winnings
}

func partTwo(cards []Card) int {
	totalCards := 0

	for i, card := range cards {
		totalCards += card.Count

		winners := 0
		for _, scratch := range card.ScratchNumbers {
			if contains(card.WinningNumbers, scratch) {
				winners++
			}
		}

		for j := 1; i+j < len(cards) && j <= winners; j++ {
			cards[i+j].Count += card.Count
		}
	}

	return totalCards
}
