package internal

import (
	"testing"

	days "github.com/jdiamond95/aoc/2023/internal"
)

func TestDay5(t *testing.T) {
	f := "./inputs/day05.txt"
	targetA := 323142486
	targetB := 5
	resultA, resultB := days.Day5(f)

	if targetA != resultA {
		t.Fatalf("Day 5 Part A failed (%s): wanted %d, got %d", f, resultA, targetA)
	}

	if targetB != resultB {
		t.Fatalf("Day 5 Part B failed (%s): wanted %d, got %d", f, resultB, targetB)
	}
}
