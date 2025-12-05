package main

import (
	"slices"
	"strings"

	"github.com/julijane/advent-of-code-2025/aoc"
)

type numrange struct {
	min, max int
}

func calc(input *aoc.Input, _, _ bool, _ ...any) (any, any) {
	part1 := 0

	blocks := input.TextBlocks()

	numRanges := []numrange{}
	for _, line := range blocks[0] {
		r := strings.Split(line, "-")
		nr := numrange{min: aoc.Atoi(r[0]), max: aoc.Atoi(r[1])}

		for i := 0; i < len(numRanges); i++ {
			cr := numRanges[i]
			if nr.min <= cr.max+1 && nr.max >= cr.min-1 {
				cr.min = min(nr.min, cr.min)
				cr.max = max(nr.max, cr.max)
				nr = cr
				numRanges = slices.Delete(numRanges, i, i+1)
				i--
			}
		}
		numRanges = append(numRanges, nr)
	}

	for _, line := range blocks[1] {
		num := aoc.Atoi(line)

		for _, nr := range numRanges {
			if num >= nr.min && num <= nr.max {
				part1++
			}
		}
	}

	part2 := 0
	for _, cr := range numRanges {
		part2 += cr.max - cr.min + 1
	}

	return part1, part2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
