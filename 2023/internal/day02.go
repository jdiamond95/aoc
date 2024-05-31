package internal

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/jdiamond95/aoc/2023/pkg"
)

func Day2(p string) (int, int) {
	lines := utils.ReadLines(p)
	a := PartA(lines)
	b := PartB(lines)

	fmt.Printf("Part a: %d\n", a)
	fmt.Printf("Part b: %d\n", b)

	return a, b
}

func PartB(lines []string) int {
	powerSum := 0

	for _, l := range lines {
		minCubes := map[string]int{}
		lSplit := formatLine(l)

		for j := 0; j < len(lSplit); j += 2 {
			cubeCount, err := strconv.Atoi(lSplit[j])
			if err != nil {
				fmt.Println(err)
			}
			color := lSplit[j+1]
			i, exists := minCubes[color]
			if exists {
				if i < cubeCount {
					minCubes[color] = cubeCount
				}
			} else {
				minCubes[color] = cubeCount
			}
		}
		power := 0
		for _, val := range minCubes {
			if power == 0 {
				power = val
			} else {
				power = power * val
			}
		}
		powerSum += power
	}
	return powerSum
}

func PartA(lines []string) int {
	maxMappings := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0

lines:
	for i, l := range lines {
		lSplit := formatLine(l)
		for j := 0; j < len(lSplit); j += 2 {
			cubeCount, err := strconv.Atoi(lSplit[j])
			if err != nil {
				fmt.Println(err)
			}
			color := lSplit[j+1]

			if !(cubeCount <= maxMappings[color]) {
				continue lines
			}

		}
		sum += i + 1
	}
	return sum
}

func formatLine(l string) []string {
	newLine := strings.Split(l, ":")[1]
	newLine = strings.Replace(newLine, "; ", ",", -1)
	newLine = strings.Replace(newLine, ", ", ",", -1)
	newLine = strings.Replace(newLine, " ", ",", -1)[1:]
	return strings.Split(newLine, ",")
}
