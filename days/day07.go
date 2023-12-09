package days

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/jdiamond95/aoc2023/utils"
)

type hand struct {
	cards  []rune
	score  int
	valMap map[rune]int
}

type orderedHands []hand

func (h orderedHands) Len() int {
	return len(h)
}

func (h orderedHands) Less(i, j int) bool {
	iScore, jScore := scoreHand(h[i]), scoreHand(h[j])
	if iScore != jScore {
		return iScore < jScore
	}

	iCardVals, jCardVals := cardsToValues(h[i].cards), cardsToValues(h[j].cards)
	for k := range iCardVals {
		if iCardVals[k] != jCardVals[k] {
			return iCardVals[k] < jCardVals[k]
		}
	}

	return false
}

func (h orderedHands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

var cardMapping = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
}

func Day7(p string) (int, int) {
	lines := utils.ReadLines(p)

	a := day7PartA(lines)
	b := day7PartB(lines)

	fmt.Printf("Part a: %d\n", a)
	fmt.Printf("Part b: %d\n", b)

	return a, b
}

func day7PartB(lines []string) int {
	return 6
}

func day7PartA(lines []string) int {
	hands := parseCamelCards(lines)
	sort.Sort(orderedHands(hands))

	sum := 0
	for i, hand := range hands {
		score := (i + 1) * hand.score
		if sum == 0 {
			sum = score
		} else {
			sum += score
		}
	}

	return sum
}

func parseCamelCards(lines []string) []hand {
	hands := make([]hand, len(lines))
	for i, line := range lines {
		split := strings.Split(line, " ")
		score, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println(err)
		}
		hand := hand{
			[]rune(split[0]),
			score,
			handToType([]rune(split[0])),
		}
		hands[i] = hand
	}
	return hands
}

func handToType(cards []rune) map[rune]int {
	typeMap := make(map[rune]int)
	for _, val := range cards {
		_, ok := typeMap[val]
		if !ok {
			typeMap[val] = 1
		} else {
			typeMap[val] += 1
		}
	}
	return typeMap
}

func cardsToValues(cards []rune) []int {
	cardVals := make([]int, 5)
	for i, card := range cards {
		str := string(card)
		num, err := strconv.Atoi(str)
		if err != nil {
			num = cardMapping[card]
		}
		cardVals[i] = num
	}
	return cardVals
}

func scoreHand(hand hand) int {
	scorecard := make([]int, 0)

	for _, val := range hand.valMap {
		scorecard = append(scorecard, val)
	}
	sort.Slice(scorecard, func(i, j int) bool {
		return scorecard[i] > scorecard[j]
	})

	score := 0
	// Five of a kind

	if reflect.DeepEqual(scorecard, []int{5}) {
		score = 7
		// Four of a kind
	} else if reflect.DeepEqual(scorecard, []int{4, 1}) {
		score = 6
		// Full house
	} else if reflect.DeepEqual(scorecard, []int{3, 2}) {
		score = 5
		// Three of a kind
	} else if reflect.DeepEqual(scorecard, []int{3, 1, 1}) {
		score = 4
		// Two pair
	} else if reflect.DeepEqual(scorecard, []int{2, 2, 1}) {
		score = 3
		// One pair
	} else if reflect.DeepEqual(scorecard, []int{2, 1, 1, 1}) {
		score = 2
	} else if reflect.DeepEqual(scorecard, []int{1, 1, 1, 1, 1}) {
		score = 1
	}
	return score
}
