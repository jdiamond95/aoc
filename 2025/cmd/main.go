package main

import (
	"fmt"
	"os"

	"github.com/jdiamond95/aoc/2025/internal"
	"github.com/jdiamond95/aoc/2025/pkg"
)

type Turn struct {
	Direction string
	Rotation  int
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go day <day_num>")
		os.Exit(1)
	}

	args := os.Args[1:]
	day := args[1]

	switch day {
	case "1":
		lines := pkg.ReadLines("inputs/01.txt")
		internal.Day1Part1(lines)
		internal.Day1Part2(lines)
	case "2":
		lines := pkg.ReadLines("inputs/02.txt")
		internal.Day2Part01(lines[0])
		internal.Day2Part02(lines[0])
	case "3":
		lines := pkg.ReadLines("inputs/03.txt")
		internal.Day3Part01(lines)
		internal.Day3Part02(lines)
	case "4":
		lines := pkg.ReadLines("inputs/04.txt")
		internal.Day4Part1(lines)
		internal.Day4Part2(lines)
	default:
		fmt.Println("Invalid day")
	}
}
