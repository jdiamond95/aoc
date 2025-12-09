package internal

import "fmt"

type Bank struct {
	batteries []int
}

func (b *Bank) HighestJoltage() int {
	highestIndex := 0
	for i := 0; i < len(b.batteries)-1; i++ {
		if b.batteries[i] > b.batteries[highestIndex] {
			highestIndex = i
		}

	}

	// Second last number was the highest
	if highestIndex == len(b.batteries)-1 {
		return b.batteries[highestIndex]*10 + b.batteries[len(b.batteries)-1]
	}

	secondHighestIndex := 0
	for i := highestIndex + 1; i < len(b.batteries); i++ {
		if b.batteries[i] > b.batteries[secondHighestIndex] || secondHighestIndex == 0 {
			secondHighestIndex = i
		}
	}
	return b.batteries[highestIndex]*10 + b.batteries[secondHighestIndex]
}

func Day3Part01(line []string) {
	sum := 0
	for _, bank := range line {
		var i []int
		for _, v := range bank {
			i = append(i, int(v-'0'))
		}
		b := Bank{batteries: i}
		sum += b.HighestJoltage()
	}
	fmt.Println(sum)
}
