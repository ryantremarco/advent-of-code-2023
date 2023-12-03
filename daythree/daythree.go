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

var adjacents = []struct {
	x int
	y int
}{
	{1, -1},
	{1, 0},
	{1, 1},
	{0, -1},
	{0, 1},
	{-1, -1},
	{-1, 0},
	{-1, 1},
}

func main() {
	parsedInput := parseInput(input)
	fmt.Println("Part One:", partOne(parsedInput))
	fmt.Println("Part Two:", partTwo(parsedInput))
}

func parseInput(input string) [][]rune {
	var lines [][]rune
	for _, line := range strings.Split(input, newline) {
		if line == "" {
			continue
		}
		var runes []rune
		for _, char := range line {
			runes = append(runes, char)
		}
		lines = append(lines, runes)
	}

	return lines
}

func isGear[T ~string | ~rune | ~byte](t T) bool {
	return string(t) == "*"
}

func isSymbol[T ~string | ~rune | ~byte](t T) bool {
	if isInt(t) {
		return false
	}

	if string(t) == "." {
		return false
	}

	return true
}

func isInt[T ~string | ~rune | ~byte](t T) bool {
	_, err := strconv.Atoi(string(t))
	if err != nil {
		return false
	}
	return true
}

func toInt[T ~string | ~rune | ~byte](t T) int {
	i, err := strconv.Atoi(string(t))
	if err != nil {
		panic(err)
	}
	return i
}

func partOne(lines [][]rune) int {
	sum := 0

	for i, line := range lines {
		currNumStr := ""
		symbolAdjacent := false
		// first version of this was getting confused by numbers at the end of each line.
		// just added a . to each line to compensate rather than deal with changing the logic :)
		for j, char := range append([]rune(line), '.') {
			if !isInt(char) {
				if currNumStr != "" && symbolAdjacent {
					sum += toInt(currNumStr)
					symbolAdjacent = false
				}

				currNumStr = ""

				continue
			}

			currNumStr = currNumStr + string(char)

			if !symbolAdjacent {
				for _, adjacent := range adjacents {
					y := i + adjacent.y
					x := j + adjacent.x

					if x < 0 || x >= len(line) || y < 0 || y >= len(lines) {
						continue
					}

					if isSymbol(lines[y][x]) {
						symbolAdjacent = true
						break
					}
				}
			}
		}
	}

	return sum
}

func partTwo(lines [][]rune) int {
	sum := 0

	for i, line := range lines {
		for j, char := range line {
			gearRatio := 1
			adjacentPartCount := 0
			if !isGear(char) {
				continue
			}

			var linesCopy [][]rune
			for _, line := range lines {
				var lineCopy []rune
				for _, c := range line {
					lineCopy = append(lineCopy, c)
				}
				linesCopy = append(linesCopy, lineCopy)
			}

			for _, adjacent := range adjacents {
				y := i + adjacent.y
				x := j + adjacent.x

				if x < 0 || x >= len(line) || y < 0 || y >= len(linesCopy) {
					continue
				}

				if point := linesCopy[y][x]; isInt(point) {
					adjacentPartCount += 1

					currNumStr := string(point)

					left := 1
					right := 1

					for x-left >= 0 {
						point := linesCopy[y][x-left]
						if !isInt(point) {
							break
						}
						currNumStr = string(point) + currNumStr
						linesCopy[y][x-left] = '.'
						left += 1
					}

					for x+right < len(linesCopy[y]) {
						point := linesCopy[y][x+right]
						if !isInt(point) {
							break
						}
						currNumStr += string(point)
						linesCopy[y][x+right] = '.'
						right += 1
					}

					gearRatio *= toInt(currNumStr)
				}
			}

			if adjacentPartCount == 2 {
				sum += gearRatio
				continue
			}
		}
	}
	return sum
}
