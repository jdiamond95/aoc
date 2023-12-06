package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jdiamond95/aoc2023/utils"
)

func Day6(p string) (int, int) {
    lines := utils.ReadLines(p)

    a := day6PartA(lines)
    b := day6PartB(lines)

    fmt.Printf("Part a: %d\n", a)
    fmt.Printf("Part b: %d\n", b)
    
    return a, b
}

func day6PartB(lines []string) int {
    times := strings.Split(lines[0], ":")[1]
    times = strings.Replace(times, " ", "", -1)
    timesInt, _ := strconv.Atoi(times)

    distances := strings.Split(lines[1], ":")[1]
    distances = strings.Replace(distances, " ", "", -1)
    distancesInt, _ := strconv.Atoi(distances)

    raceSum := calculateRecordBreaks(timesInt, distancesInt)
    return raceSum
}

func day6PartA(lines []string) int {
    times := utils.ReadIntArray(lines[0])
    distances := utils.ReadIntArray(lines[1])

    // For each race
    total := 0
    for i := 0; i < len(times); i++ {
        raceSum := calculateRecordBreaks(times[i], distances[i])
        if total == 0 {
            total = raceSum
        } else {
            total = total * raceSum
        }
    }
    return total
}

func calculateRecordBreaks(time int, record int) int {
    recordBreaks := 0
    for j := 0; j < time; j++ {
        travelTime := time - j
        distance := j * travelTime
        if distance > record {
            recordBreaks += 1
        }
    }
    return recordBreaks
}



