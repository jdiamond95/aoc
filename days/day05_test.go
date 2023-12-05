package days

import "testing"

func TestDay5(t *testing.T) {
    f := "day04.txt"
    targetA := 323142486
    targetB := 5
    resultA, resultB := Day4(f)

    if targetA != resultA {
        t.Fatalf("Day 5 Part A failed (%s): wanted %d, got %d", f, resultA, targetA)
    }

    if targetB != resultB {
        t.Fatalf("Day 5 Part B failed (%s): wanted %d, got %d", f, resultB, targetB)
    }
}
