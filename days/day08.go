package days

import (
	"fmt"
	"github.com/jdiamond95/aoc2023/utils"
	"regexp"
)

type node struct {
	left  string
	right string
}

func Day8(p string) (int, int) {
	lines := utils.ReadLines(p)

	a := day8PartA(lines)
	b := day8PartB(lines)

	fmt.Printf("Part a: %d\n", a)
	fmt.Printf("Part b: %d\n", b)

	return a, b
}

func day8PartB(lines []string) int {
	return 5
}

func day8PartA(lines []string) int {
	path := lines[0]
	tree := make(map[string]node)
	for _, line := range lines[2:] {
		key, node := createNode(line)
		tree[key] = node
	}

	index := 0
	steps := 0
	currStr := "AAA"
	for {
		turn := path[index]
		currNode := tree[currStr]
		if turn == 'L' {
			currStr = currNode.left
		} else {
			currStr = currNode.right
		}

		steps += 1
		if currStr == "ZZZ" {
			break
		}

		if index == len(path)-1 {
			index = 0
		} else {
			index += 1
		}
	}
	return steps
}

func createNode(line string) (string, node) {
	r, _ := regexp.Compile("([A-Z]+)")
	m := r.FindAllString(line, -1)
	return m[0], node{left: m[1], right: m[2]}
}
