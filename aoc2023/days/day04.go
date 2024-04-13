package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jdiamond95/aoc2023/utils"
)

func Day4(p string) (int, int) {
    lines := utils.ReadLines(p)

    a := day4PartA(lines)
    b := day4PartB(lines)

    fmt.Printf("Part a: %d\n", a)
    fmt.Printf("Part b: %d\n", b)
    
    return a, b
}

func day4PartB(lines []string) int {
    scorecardsMap := make(map[int]int)
    // Initialise scorecardsMap. We start with one scorecard for each line
    for i := range lines {
        scorecardsMap[i] = 1
    }

    for i, line := range lines {
        winningMatches := processScorecardLine(line)
        for j := i + 1; j <= i + winningMatches; j++ {
            if j >= len(scorecardsMap) {
                continue
            }
            scorecardsMap[j] += scorecardsMap[i]
        }
    }

    sum := 0
    for _, val := range scorecardsMap {
        sum += val
    }
    return sum
}

func day4PartA(lines []string) int {
    pointSum := 0
    for _, line := range lines {
        winningMatches := processScorecardLine(line)
        if winningMatches > 0 {
            pointSum = pointSum + int(math.Pow(2, float64(winningMatches)- 1))
        }
    }
    return pointSum
}

func processScorecardLine(line string) int {
    count := 0
    scratchie := strings.Split(line, ":")[1]
    winningNums := parseLine(strings.Split(scratchie, "|")[0])
    winnersMap := make(map[int]int)

    // Put winning numbers into a hashmap
    for _, val := range winningNums {
        num, err := strconv.Atoi(val)
        if err != nil {
            fmt.Println(err)
        }
        winnersMap[num] = 1
    }

    // If our nums are in the hashmap, count them. Essentially finding an intersection here
    ourNums := parseLine(strings.Split(scratchie, "|")[1])
    for _, val := range ourNums {
        num, err := strconv.Atoi(val)
        if err != nil {
            fmt.Println(err)
        }

        _, ok := winnersMap[num]
        if ok {
            count += 1
        }
    }
    return count
}

func parseLine(line string) []string {
    line = strings.TrimSpace(line)
    lineArr := strings.Fields(line)
    return lineArr
}


