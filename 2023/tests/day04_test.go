package internal

import (
	"testing"

	days "github.com/jdiamond95/aoc/2023/internal"
)

func TestDay4(t *testing.T) {
	f := "./inputs/day04.txt"
	targetA := 24175
	targetB := 18846301
	resultA, resultB := days.Day4(f)

	if targetA != resultA {
		t.Fatalf("Day 4 Part A failed (%s): wanted %d, got %d", f, resultA, targetA)
	}

	if targetB != resultB {
		t.Fatalf("Day 4 Part B failed (%s): wanted %d, got %d", f, resultB, targetB)
	}
}
