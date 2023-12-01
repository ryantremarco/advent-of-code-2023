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

func main() {
	fmt.Println("Part One", partOne(input))
	fmt.Println("Part Two", partTwo(input))
}

func partTwo(input string) int {
	digitMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	total := 0
	for _, line := range strings.Split(input, newline) {
		var first, firstPos, last, lastPos = 0, len(line), 0, 0

		for k, v := range digitMap {
			vStr := strconv.Itoa(v)

			if pos := strings.Index(line, k); pos != -1 && pos <= firstPos {
				firstPos = pos
				first = v
			}
			if pos := strings.Index(line, vStr); pos != -1 && pos <= firstPos {
				firstPos = pos
				first = v
			}

			if pos := strings.LastIndex(line, k); pos != -1 && pos >= lastPos {
				lastPos = pos
				last = v
			}
			if pos := strings.LastIndex(line, vStr); pos != -1 && pos >= lastPos {
				lastPos = pos
				last = v
			}
		}
		value := first*10 + last
		total += value
	}

	return total
}

func partOne(input string) int {
	digits := []int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
	}

	total := 0
	for _, line := range strings.Split(input, newline) {
		var first, firstPos, last, lastPos = 0, len(line), 0, 0

		for _, v := range digits {
			vStr := strconv.Itoa(v)

			if pos := strings.Index(line, vStr); pos != -1 && pos <= firstPos {
				firstPos = pos
				first = v
			}

			if pos := strings.LastIndex(line, vStr); pos != -1 && pos >= lastPos {
				lastPos = pos
				last = v
			}
		}
		value := first*10 + last
		total += value
	}

	return total
}
