package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jdiamond95/aoc2023/utils"
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
    return 5
}

func day5PartA(input string) int {
    almanac := parseAlmanac(input)

    lowest := -1
    seeds := strings.Split(almanac[0], " ")

    for _, seed := range seeds {
        currVal, err := strconv.Atoi(seed)
        if err != nil {
            fmt.Println(err)
        }

        almanac:
        for i := 1; i < len(almanac); i++ {
            for _, line := range strings.Split(almanac[i], "\n") {
                splitLine := strings.Split(line, " ")
                destStart, _ := strconv.Atoi(splitLine[0])
                sourceStart, _ := strconv.Atoi(splitLine[1])
                length, _ := strconv.Atoi(splitLine[2])

                // Input number needs to be updated
                if currVal >= sourceStart && currVal <= sourceStart + length {
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
