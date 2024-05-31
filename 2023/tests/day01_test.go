package internal

import (
	"testing"

	days "github.com/jdiamond95/aoc/2023/internal"
)

func TestDay1Sample(t *testing.T) {
	f := "./inputs/day01.txt"
	targetA := 56042
	targetB := 55358
	resultA, resultB := days.Day1(f)

	if targetA != resultA {
		t.Fatalf("Day 1 Part A failed (%s): wanted %d, got %d", f, resultA, targetA)
	}

	if targetB != resultB {
		t.Fatalf("Day 1 Part B failed (%s): wanted %d, got %d", f, resultB, targetB)
	}
}
