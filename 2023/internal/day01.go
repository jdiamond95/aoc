package internal

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	utils "github.com/jdiamond95/aoc/2023/pkg"
)

func Day1(p string) (int, int) {
	lines := utils.ReadLines(p)

	a := calibrate(lines)
	b := calibrate(wordsToDigits(lines))

	fmt.Printf("Part a: %d\n", a)
	fmt.Printf("Part b: %d\n", b)

	return a, b
}

func calibrate(lines []string) int {
	sum := 0
	for _, line := range lines {
		nums := strings.Map(filterNums, line)
		numsArray := strings.Split(nums, "")
		tens, err := strconv.Atoi(numsArray[0])
		if err != nil {
			fmt.Println(err)
		}
		units, err := strconv.Atoi(numsArray[len(numsArray)-1])
		if err != nil {
			fmt.Println(err)
		}

		lineSum := 10*tens + units
		sum += int(lineSum)
	}
	return sum
}

func wordsToDigits(lines []string) []string {
	mappings := [][]string{
		{"one", "one1one"},
		{"two", "two2two"},
		{"three", "three3three"},
		{"four", "four4four"},
		{"five", "five5five"},
		{"six", "six6six"},
		{"seven", "seven7seven"},
		{"eight", "eight8eight"},
		{"nine", "nine9nine"},
	}

	newLines := []string{}
	for _, l := range lines {
		updatedLine := l
		for _, m := range mappings {
			updatedLine = strings.Replace(updatedLine, m[0], m[1], -1)
		}
		newLines = append(newLines, updatedLine)
	}
	return newLines
}

func filterNums(r rune) rune {
	// If rune is a number, return the number
	// Else return negative number, dropped with map
	if !unicode.IsDigit(r) {
		return -1
	}
	return r
}
