package days

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/jdiamond95/aoc2023/utils"
)

func Day3(p string) {
    lines := utils.ReadLines(p)

    partA(lines)
}

func partA(lines []string) {
    runeMap := [][]rune{}
    for _, line := range lines {
        runeLine := []rune(line)
        runeMap = append(runeMap, runeLine)
    }

    touchMap := make([][]bool, 0)
    for i := 0; i < len(runeMap[0]); i++ {
        tmp := make([]bool, 0)
        for j := 0; j < len(lines); j++ {
            tmp = append(tmp, checkTouching(i, j, runeMap))
        }
        touchMap = append(touchMap, tmp)
    }

    sum := 0
    for i := 0; i < len(lines); i++ {
        numStartIndex := -1
        for j := 0; j < len(runeMap[0]); j++ {
            // Start of a number
            if numStartIndex == -1 && unicode.IsDigit(runeMap[i][j]) {
                numStartIndex = j
            }

            // Come to the end of a number
            if (!unicode.IsDigit(runeMap[i][j]) || j == len(runeMap[0]) - 1) && numStartIndex != -1 {
                // If theres a true in the touchMap for the range, append the number
                for k := numStartIndex; k <= j; k++ {
                    if touchMap[i][k] == true {
                        numSlice := string(runeMap[i][numStartIndex:j])
                        // Slicing edge case
                        if j == len(runeMap[0]) - 1 && unicode.IsDigit(runeMap[i][j]) {
                            numSlice = string(runeMap[i][numStartIndex:])
                        }
                        numSliceInt, err := strconv.Atoi(numSlice)
                        if err != nil {
                            fmt.Println(err)
                        }
                        sum += numSliceInt
                        break
                    }
                }
                numStartIndex = -1
            }
        }
    }
    fmt.Println(sum)
}

func checkTouching(x int, y int, schematic [][]rune) bool {
    if schematic[x][y] == '.' {
        return false
    }

    width := len(schematic[0])
    height := len(schematic)

    coordinates := [][]int{
        {x-1, y-1},
        {x, y-1},
        {x+1, y-1},
        {x-1, y},
        {x+1, y},
        {x-1, y+1},
        {x, y+1},
        {x+1, y+1},
    }

    // For each possible surrounding coordinate
    for i := 0; i < len(coordinates); i++ {
        // Skip coordinates outside of the schematic
        if coordinates[i][0] == -1 || coordinates[i][1] == -1 || coordinates[i][0] >= width || coordinates[i][1] >= height {
            continue
        }
        // If neighbour is not a number or ., its a symbol - return true
        if !unicode.IsDigit(schematic[coordinates[i][0]][coordinates[i][1]]) && schematic[coordinates[i][0]][coordinates[i][1]] != '.' {
            return true
        }
    }
    return false
}
