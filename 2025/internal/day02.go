package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func ExtractRanges(line string) []Range {
	ranges := []Range{}
	rangeStrings := strings.Split(line, ",")
	for _, r := range rangeStrings {
		rangeParts := strings.Split(r, "-")
		start, _ := strconv.Atoi(rangeParts[0])
		end, _ := strconv.Atoi(rangeParts[1])
		ranges = append(ranges, Range{Start: start, End: end})
	}
	return ranges
}

func FindInvalidIdsInRange(r Range) []int {
	invalidIds := []int{}
	for i := r.Start; i <= r.End; i++ {
		iStr := strconv.Itoa(i)
		if len(iStr)%2 != 0 {
			continue
		}

		if iStr[0:len(iStr)/2] == iStr[len(iStr)/2:] {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}

func IsInvalidIdPart2(id int) bool {
	idStr := strconv.Itoa(id)
	idLen := len(idStr)
	for patternLen := 1; patternLen <= idLen/2; patternLen++ {
		if idLen%patternLen != 0 {
			continue
		}

		pattern := idStr[:patternLen]
		theRest := idStr[patternLen:]
		if strings.Replace(theRest, pattern, "", -1) == "" {
			return true
		}

	}
	return false
}

func Day2Part01(line string) {
	ranges := ExtractRanges(line)
	invalidIdSum := 0
	for _, r := range ranges {
		invalidIds := FindInvalidIdsInRange(r)
		for _, i := range invalidIds {
			invalidIdSum += i
		}
	}

	fmt.Println(invalidIdSum)
}

func Day2Part02(line string) {
	ranges := ExtractRanges(line)
	sum := 0
	for _, r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			if IsInvalidIdPart2(i) {
				sum += i
			}
		}
	}
	fmt.Println(sum)
}
