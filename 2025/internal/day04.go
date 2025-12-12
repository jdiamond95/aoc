package internal

import (
	"fmt"
)

const MAX_ADJACENT_PAPER = 4

type PaperMap struct {
	positions [][]string
}

func (pm *PaperMap) ToString() string {
	result := ""
	for _, row := range pm.positions {
		for _, pos := range row {
			result += pos
		}
		result += "\n"
	}
	return result
}

// Returns the number of papers removed and a new Papermap
func (pm PaperMap) RemovePaper() (PaperMap, int) {
	paperRemoved := 0
	// Make a copy of the map
	newMap := PaperMap{positions: make([][]string, len(pm.positions))}
	for i := range pm.positions {
		newMap.positions[i] = make([]string, len(pm.positions[i]))
		for j := range pm.positions[i] {
			newMap.positions[i][j] = pm.positions[i][j]
		}
	}

	for i := range pm.positions {
		for j := range pm.positions[i] {
			if pm.positions[i][j] == "@" && pm.CanPaperBeRemoved(i, j) {
				newMap.positions[i][j] = "X"
				paperRemoved++
			}
		}
	}
	return newMap, paperRemoved
}

func (pm *PaperMap) CanPaperBeRemoved(x, y int) bool {
	totalPaper := 0
	for i := -1; i <= 1; i++ {
		if x+i < 0 || x+i >= len(pm.positions) {
			continue
		}
		for j := -1; j <= 1; j++ {
			if y+j < 0 || y+j >= len(pm.positions[x+i]) {
				continue
			}
			if pm.positions[x+i][y+j] == "@" {
				totalPaper++
			}
		}
	}

	return totalPaper <= MAX_ADJACENT_PAPER
}

func Day4Part1(line []string) {
	pm := PaperMap{}
	for _, v := range line {
		row := []string{}
		for _, c := range v {
			row = append(row, string(c))
		}
		pm.positions = append(pm.positions, row)
	}

	pm, removed := pm.RemovePaper()
	fmt.Println(removed)
	fmt.Println(pm.ToString())
}

func Day4Part2(line []string) {
	pm := PaperMap{}
	for _, v := range line {
		row := []string{}
		for _, c := range v {
			row = append(row, string(c))
		}
		pm.positions = append(pm.positions, row)
	}

	var removed int
	totalPaperRemoved := 0
	for {
		fmt.Println(pm.ToString())
		pm, removed = pm.RemovePaper()
		totalPaperRemoved += removed
		if removed == 0 {
			break
		}
	}
	fmt.Println(totalPaperRemoved)
}
