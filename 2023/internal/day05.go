package internal

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/jdiamond95/2023/pkg"
)

func Day5(p string) (int, int) {
	input := utils.ReadBlock(p)

	a := day5PartA(input)
	b := day5PartB(input)

	fmt.Printf("Part a: %d\n", a)
	fmt.Printf("Part b: %d\n", b)

	return a, b
}

func day5PartB(input string) int {
	almanac := parseAlmanac(input)
	seedList := strings.Split(almanac[0], " ")

	seedsArr := make([]int, 0)
	for i := 0; i < len(seedList); i += 2 {
		seedStart, _ := strconv.Atoi(seedList[i])
		seedLen, _ := strconv.Atoi(seedList[i+1])
		for j := seedStart; j < (seedStart + seedLen); j++ {
			seedsArr = append(seedsArr, j)
		}
	}

	// Dedup seedsArr
	lowest := calculateLowestSeed(seedsArr, almanac)
	return lowest
}

func day5PartA(input string) int {
	almanac := parseAlmanac(input)
	seeds := strings.Split(almanac[0], " ")

	seedsIntArr := make([]int, 0)
	for _, val := range seeds {
		num, _ := strconv.Atoi(val)
		seedsIntArr = append(seedsIntArr, num)
	}
	lowest := calculateLowestSeed(seedsIntArr, almanac)

	return lowest
}

func parseAlmanac(input string) []string {
	almanac := make([]string, 0)

	str := strings.Split(input, "\n\n")
	for _, val := range str {
		nums := strings.Split(val, ":")[1]
		nums = strings.TrimSpace(nums)
		almanac = append(almanac, nums)
	}
	return almanac
}

func calculateLowestSeed(seeds []int, almanac []string) int {
	lowest := -1
	for _, seed := range seeds {
		currVal := seed

	almanac:
		for i := 1; i < len(almanac); i++ {
			for _, line := range strings.Split(almanac[i], "\n") {
				splitLine := strings.Split(line, " ")
				destStart, _ := strconv.Atoi(splitLine[0])
				sourceStart, _ := strconv.Atoi(splitLine[1])
				length, _ := strconv.Atoi(splitLine[2])

				// Input number needs to be updated
				if currVal >= sourceStart && currVal <= sourceStart+length {
					currVal = currVal + (destStart - sourceStart)
					continue almanac
				}
			}
		}
		if lowest < 0 || currVal < lowest {
			lowest = currVal
		}
	}
	return lowest
}
