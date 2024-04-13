package days

import "testing"

func TestDay2(t *testing.T) {
    f := "day02.txt"
    targetA := 2207
    targetB := 62241
    resultA, resultB := Day2(f)

    if targetA != resultA {
        t.Fatalf("Day 1 Part A failed (%s): wanted %d, got %d", f, resultA, targetA)
    }

    if targetB != resultB {
        t.Fatalf("Day 1 Part B failed (%s): wanted %d, got %d", f, resultB, targetB)
    }
}
