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
    fmt.Println(timesInt)

    distances := strings.Split(lines[1], ":")[1]
    distances = strings.Replace(distances, " ", "", -1)
    distancesInt, _ := strconv.Atoi(distances)
    fmt.Println(distancesInt)

    raceSum := 0
    // For each range of durations button held
    for j := 0; j < timesInt; j++ {
        travelTime := timesInt - j
        distance := j * travelTime
        if distance > distancesInt {
            raceSum += 1
        }
    }

    return raceSum
}

func day6PartA(lines []string) int {
    times := strings.Split(lines[0], ":")[1]
    timesArr := strings.Fields(times)
    timesIntArr := []int{}

    for i := range timesArr {
        num, _ := strconv.Atoi(timesArr[i])
        timesIntArr = append(timesIntArr, num)
    }

    distances := strings.Split(lines[1], ":")[1]
    distancesArr := strings.Fields(distances)
    distancesIntArr := []int{}

    for i := range distancesArr {
        num, _ := strconv.Atoi(distancesArr[i])
        distancesIntArr = append(distancesIntArr, num)
    }

    // For each race

    total := 0
    for i := 0; i < len(timesArr); i++ {
        fmt.Printf("Race number: %d\n", i)
        fmt.Printf("Duration: %d - Record: %d\n", timesIntArr[i], distancesIntArr[i])
        raceSum := 0
        // For each range of durations button held
        for j := 0; j < timesIntArr[i]; j++ {
            travelTime := timesIntArr[i] - j
            distance := j * travelTime
            fmt.Printf("Seconds Pressed: %d - Distance: %d\n", j, distance)
            if distance > distancesIntArr[i] {
                raceSum += 1
            }
        }
        if total == 0 {
            total = raceSum
        } else {
            total = total * raceSum
        }
    }

    return total
}


