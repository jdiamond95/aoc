package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

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
        for i := 1; i < 7; i++ {
            challenges = append(challenges, i)
        }
    }

    for _, val := range challenges {
        fmt.Printf("🎅 Day %02d 🎅\n", val)

        start := time.Now()
        switch val {
        case 1:
            days.Day1("./days/day01.txt")
        case 2:
            days.Day2("./days/day02.txt")
        case 3:
            days.Day3("./days/day03.txt")
        case 4:
            days.Day4("./days/day04.txt")
        case 5:
            days.Day5("./days/day05.txt")
        case 6:
            days.Day6("./days/day06.txt")
        default:
            fmt.Println("Not implemented")
        }

        elapsed := time.Since(start)
        fmt.Printf("\nTime taken: %s\n\n", elapsed.String())
    }
}
