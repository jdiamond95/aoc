package days

import "testing"

func TestDay1(t *testing.T) {
    target := 142
    result := Day1("day01_sample.txt")

    if target != result {
        t.Fatalf("Day1 failed, wanted %d, got %d", target, result)
    }

}
