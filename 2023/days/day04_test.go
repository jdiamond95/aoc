package days

import "testing"

func TestDay4(t *testing.T) {
    f := "day04.txt"
    targetA := 24175
    targetB := 18846301
    resultA, resultB := Day4(f)

    if targetA != resultA {
        t.Fatalf("Day 4 Part A failed (%s): wanted %d, got %d", f, resultA, targetA)
    }

    if targetB != resultB {
        t.Fatalf("Day 4 Part B failed (%s): wanted %d, got %d", f, resultB, targetB)
    }
}
