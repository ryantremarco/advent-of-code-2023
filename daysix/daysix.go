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

type Race struct {
	MaxTime        int
	RecordDistance int
}

func parseInput(input string) []Race {
	var races []Race
	for strings.Contains(input, "  ") {
		input = strings.ReplaceAll(input, "  ", " ")
	}

	lines := strings.Split(input, newline)

	getValues := func(line string) []int {
		colonSplit := strings.Split(line, ": ")
		valuesSplit := strings.Split(colonSplit[1], " ")
		var values []int
		for _, value := range valuesSplit {
			values = append(values, toInt(value))
		}
		return values
	}

	times := getValues(lines[0])
	distances := getValues(lines[1])

	for i := range times {
		races = append(races, Race{
			MaxTime:        times[i],
			RecordDistance: distances[i],
		})
	}

	return races
}

func toInt[T ~string | ~rune | ~byte](t T) int {
	i, err := strconv.Atoi(string(t))
	if err != nil {
		panic(err)
	}
	return i
}

// distance = chargeTime * (maxTime - chargeTime)
// solving for chargeTime gives a quadratic
// chargeTime = (maxTime +- sqrt(maxTime^2 - 4distance))/2
func partOne(races []Race) int {
	marginMulti := 1

	for _, race := range races {
		maxTime := float64(race.MaxTime)
		distance := float64(race.RecordDistance)
		lowerChargeTime := (maxTime - math.Sqrt(maxTime*maxTime-4*distance)) / 2
		upperChargeTime := (maxTime + math.Sqrt(maxTime*maxTime-4*distance)) / 2

		roundedLower := int(lowerChargeTime)
		roundedUpper := int(upperChargeTime + .99)

		errorMargin := roundedUpper - roundedLower - 1

		marginMulti *= errorMargin
	}

	return marginMulti
}

func partTwo(races []Race) int {
	var timeStr, distanceStr string

	for _, race := range races {
		timeStr += strconv.Itoa(race.MaxTime)
		distanceStr += strconv.Itoa(race.RecordDistance)
	}

	realRace := Race{
		MaxTime:        toInt(timeStr),
		RecordDistance: toInt(distanceStr),
	}

	return partOne([]Race{realRace})
}
