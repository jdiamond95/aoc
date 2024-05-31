package internal

import (
	"testing"

	days "github.com/jdiamond95/aoc/2023/internal"
)

func TestDay2(t *testing.T) {
	f := "./inputs/day02.txt"
	targetA := 2207
	targetB := 62241
	resultA, resultB := days.Day2(f)

	if targetA != resultA {
		t.Fatalf("Day 1 Part A failed (%s): wanted %d, got %d", f, resultA, targetA)
	}

	if targetB != resultB {
		t.Fatalf("Day 1 Part B failed (%s): wanted %d, got %d", f, resultB, targetB)
	}
}
