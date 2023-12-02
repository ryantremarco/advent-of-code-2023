package main

import (
	_ "embed"
	"testing"
)

//go:embed example
var example string

func TestPartOne(t *testing.T) {
	expected := 8
	got := partOne(parseInput(example))

	if got != expected {
		t.Fatalf("expected %d but got %d", expected, got)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 2286
	got := partTwo(parseInput(example))

	if got != expected {
		t.Fatalf("expected %d but got %d", expected, got)
	}
}
