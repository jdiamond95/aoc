package internal

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	utils "github.com/jdiamond95/2023/pkg"
)

type hand struct {
	cards  []rune
	score  int
	valMap map[rune]int
}

// TODO - make this not disgusting
type part1Order []hand

func (h part1Order) Len() int {
	return len(h)
}

func (h part1Order) Less(i, j int) bool {
	iScore, jScore := scoreHand(h[i]), scoreHand(h[j])
	if iScore != jScore {
		return iScore < jScore
	}

	iCardVals, jCardVals := cardsToValues(h[i].cards, false), cardsToValues(h[j].cards, false)
	for k := range iCardVals {
		if iCardVals[k] != jCardVals[k] {
			return iCardVals[k] < jCardVals[k]
		}
	}

	return false
}

func (h part1Order) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

type part2Order []hand

func (h part2Order) Len() int {
	return len(h)
}

func (h part2Order) Less(i, j int) bool {
	iScore, jScore := substitute(h[i]), substitute(h[j])
	if iScore != jScore {
		return iScore < jScore
	}

	iCardVals, jCardVals := cardsToValues(h[i].cards, true), cardsToValues(h[j].cards, true)
	for k := range iCardVals {
		if iCardVals[k] != jCardVals[k] {
			return iCardVals[k] < jCardVals[k]
		}
	}

	return false
}

func (h part2Order) Swap(i, j int) {
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
	hands := parseCamelCards(lines)
	sort.Sort(part2Order(hands))

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

func day7PartA(lines []string) int {
	hands := parseCamelCards(lines)
	sort.Sort(part1Order(hands))

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

func substitute(cards hand) int {
	maxScore := 0
	for card, count := range cards.valMap {
		if card == 'J' && count == 5 {
			return 7
		} else if card == 'J' {
			continue
		} else {
			temp := make([]rune, 0)
			for i := range cards.cards {
				if cards.cards[i] == 'J' {
					temp = append(temp, card)
				} else {
					temp = append(temp, cards.cards[i])
				}
			}
			tryHand := hand{
				temp,
				0,
				handToType(temp),
			}
			score := scoreHand(tryHand)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
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

func cardsToValues(cards []rune, part2 bool) []int {
	cardVals := make([]int, 5)
	for i, card := range cards {
		str := string(card)
		num, err := strconv.Atoi(str)
		if err != nil {
			if part2 && card == 'J' {
				num = 0
			} else {
				num = cardMapping[card]
			}
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
