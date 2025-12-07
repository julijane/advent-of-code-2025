package main

import (
	"slices"
	"strings"

	"github.com/julijane/advent-of-code-2025/aoc"
)

type pos struct {
	x          int
	numReached int
}

type positions []pos

func (s *positions) Add(x, numReached int) {
	for i, p := range *s {
		if p.x == x {
			(*s)[i].numReached += numReached
			return
		}
	}
	*s = append(*s, pos{x: x, numReached: numReached})
}

func calc(input *aoc.Input, _, _ bool, _ ...any) (any, any) {
	part1 := 0

	state := positions{{x: strings.Index(input.Lines[0].Data, "S"), numReached: 1}}

	for _, line := range input.Lines[1:] {
		splitterPosO := line.FindObjects("\\^")
		splitterPos := splitterPosO.StartPostions()

		if len(splitterPos) == 0 {
			continue
		}

		newState := positions{}

		for _, s := range state {
			if slices.Contains(splitterPos, s.x) {
				part1++
				newState.Add(s.x-1, s.numReached)
				newState.Add(s.x+1, s.numReached)
			} else {
				newState.Add(s.x, s.numReached)
			}

		}
		state = newState
	}

	part2 := 0
	for _, s := range state {
		part2 += s.numReached
	}
	return part1, part2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
