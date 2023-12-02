package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

const newline = "\n"

type Round struct {
	Red   int
	Blue  int
	Green int
}

type Game struct {
	ID     int
	Rounds []Round
}

func main() {
	games := parseInput(input)
	fmt.Println("Part One:", partOne(games))
	fmt.Println("Part Two:", partTwo(games))
}

func parseInput(input string) []Game {
	var games []Game
	for _, line := range strings.Split(input, newline) {
		var game Game
		if line == "" {
			continue
		}

		idSplit := strings.Split(line, ": ")
		idStr := strings.ReplaceAll(idSplit[0], "Game ", "")
		game.ID = strToInt(idStr)

		for _, roundStr := range strings.Split(idSplit[1], "; ") {
			var round Round
			for _, colour := range strings.Split(roundStr, ", ") {
				split := strings.Split(colour, " ")

				switch split[1] {
				case "red":
					round.Red = strToInt(split[0])
				case "blue":
					round.Blue = strToInt(split[0])
				case "green":
					round.Green = strToInt(split[0])
				}
			}
			game.Rounds = append(game.Rounds, round)
		}

		games = append(games, game)
	}

	return games
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func partOne(games []Game) int {
	idsSum := 0

	redMax := 12
	blueMax := 14
	greenMax := 13

	for _, game := range games {
		possible := true

		for _, round := range game.Rounds {
			if round.Red > redMax || round.Blue > blueMax || round.Green > greenMax {
				possible = false
				break
			}
		}

		if possible {
			idsSum += game.ID
		}
	}

	return idsSum
}

func partTwo(games []Game) int {
	var powerSums int

	for _, game := range games {

		var redMax, greenMax, blueMax int

		for _, round := range game.Rounds {
			if round.Red > redMax {
				redMax = round.Red
			}

			if round.Blue > blueMax {
				blueMax = round.Blue
			}

			if round.Green > greenMax {
				greenMax = round.Green
			}
		}

		powerSums += redMax * greenMax * blueMax
	}

	return powerSums
}
