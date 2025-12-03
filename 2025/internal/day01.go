package internal

import (
	"fmt"
	"math"
	"strconv"
)

const INITIAL_POSITION = 50
const MIN_SAFE_POSITION = 0
const MAX_SAFE_POSITION = 100

type Safe struct {
	position int
}

type SafeTurn struct {
	direction string
	rotation  int
}

func (s *Safe) Turn(st SafeTurn) int {
	c := 0
	st.rotation %= MAX_SAFE_POSITION
	if st.direction == "L" {
		st.rotation *= -1
	}
	s.position = s.position + st.rotation
	if s.position >= MAX_SAFE_POSITION {
		s.position -= MAX_SAFE_POSITION
	} else if s.position < MIN_SAFE_POSITION {
		s.position += MAX_SAFE_POSITION
	}
	return c
}

func Day1Part1(l []string) {
	count := 0
	safe := Safe{position: INITIAL_POSITION}
	for _, v := range l {
		dir := string(v[0])
		rot, _ := strconv.Atoi(string(v[1:]))
		st := SafeTurn{direction: dir, rotation: rot}
		count += safe.Turn(st)

		// Count if the turn lands at MIN_SAFE_POSITION
		if safe.position == MIN_SAFE_POSITION {
			count++
		}
	}
	fmt.Printf("Password: %d\n", count)
}

func Day1Part2(l []string) {
	count := 0
	s := Safe{position: INITIAL_POSITION}
	for _, v := range l {
		dir := string(v[0])
		rot, _ := strconv.Atoi(string(v[1:]))
		st := SafeTurn{direction: dir, rotation: rot}

		// Count full turns
		fullTurns := math.Floor(float64(st.rotation) / float64(MAX_SAFE_POSITION))
		count += int(fullTurns)

		// Count wrapping turns
		localrot := st.rotation % MAX_SAFE_POSITION
		if ((st.direction == "L" && s.position-localrot <= MIN_SAFE_POSITION) || (st.direction == "R" && s.position+localrot >= MAX_SAFE_POSITION)) && s.position != 0 {
			count++
		}
		s.Turn(st)
	}
	fmt.Printf("Password: %d\n", count)
}
