package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/jdiamond95/aoc2023/days"
)

func main() {
    fmt.Printf("🎄 aoc2023 🎄\n\n")

    args := os.Args
    var challenges []int

    // Individual day
    if len(args) == 2 {
        n, err := strconv.Atoi(args[1])

        if err != nil {
            fmt.Println(err)
        }

        challenges = append(challenges, n)

    // All days
    } else {
        for i := 1; i < 2; i++ {
            challenges = append(challenges, i)
        }
    }

    for _, val := range challenges {
        fmt.Printf("🎅 Day %02d 🎅\n", val)
        switch val {
        case 1:
            days.Day1("day01.txt")
        default:
            fmt.Println("Not implemented")
        }
    }
}
