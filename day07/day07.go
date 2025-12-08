package main

import (
	"slices"
	"strings"

	"github.com/julijane/advent-of-code-2025/aoc"
)

// Store all current beam positions and how many ways they each were reached
type positions map[int]int

// helper function, just for more clarity/intent in the code
func (s positions) Add(x, numReached int) {
	s[x] += numReached
}

func calc(input *aoc.Input, _, _ bool, _ ...any) (any, any) {
	part1 := 0

	state := positions{}
	state.Add(strings.Index(input.Lines[0].Data, "S"), 1)

	for _, line := range input.Lines[1:] {
		splitterPosO := line.FindObjects("\\^")
		splitterPos := splitterPosO.StartPostions()

		if len(splitterPos) == 0 {
			continue
		}

		newState := positions{}

		for x, numReached := range state {
			if slices.Contains(splitterPos, x) {
				part1++
				newState.Add(x-1, numReached)
				newState.Add(x+1, numReached)
			} else {
				newState.Add(x, numReached)
			}

		}
		state = newState
	}

	part2 := 0
	for _, numReached := range state {
		part2 += numReached
	}
	return part1, part2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
