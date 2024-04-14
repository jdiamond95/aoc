package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	days "github.com/jdiamond95/2023/internal"
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
		for i := 1; i < 9; i++ {
			challenges = append(challenges, i)
		}
	}

	for _, val := range challenges {
		fmt.Printf("🎅 Day %02d 🎅\n", val)

		start := time.Now()
		switch val {
		case 1:
			path, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(path)
			days.Day1("../internal/inputs/day01.txt")
		case 2:
			days.Day2("../internal/inputs/day02.txt")
		// case 3:
		// 	days.Day3("../internal/inputs/day03.txt")
		case 4:
			days.Day4("../internal/inputs/day04.txt")
		case 5:
			days.Day5("../internal/inputs/day05.txt")
		case 6:
			days.Day6("../internal/inputs/day06.txt")
		case 7:
			days.Day7("../internal/inputs/day07.txt")
		case 8:
			days.Day8("../internal/inputs/day08.txt")
		default:
			fmt.Println("Not implemented")
		}

		elapsed := time.Since(start)
		fmt.Printf("\nTime taken: %s\n\n", elapsed.String())
	}
}
