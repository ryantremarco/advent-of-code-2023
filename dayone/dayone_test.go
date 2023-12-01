package main

import (
	_ "embed"
	"testing"
)

//go:embed partone_example
var partOneExample string

//go:embed parttwo_example
var partTwoExample string

func TestPartOne(t *testing.T) {
	expected := 142
	got := partOne(partOneExample)

	if got != expected {
		t.Fatalf("expected %d but got %d", expected, got)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 281
	got := partTwo(partTwoExample)

	if got != expected {
		t.Fatalf("expected %d but got %d", expected, got)
	}
}
