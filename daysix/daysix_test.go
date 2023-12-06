package main

import (
	_ "embed"
	"testing"
)

//go:embed example
var example string

func TestPartOne(t *testing.T) {
	expected := 288
	got := partOne(parseInput(example))

	if got != expected {
		t.Fatalf("expected %d but got %d", expected, got)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 71503
	got := partTwo(parseInput(example))

	if got != expected {
		t.Fatalf("expected %d but got %d", expected, got)
	}
}
