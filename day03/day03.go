package main

import (
	"strings"

	"github.com/julijane/advent-of-code-2025/aoc"
)

func findMaxFrom(digits []string, startPos, digitsStillNeeded int) string {
	max := "#"
	var maxpos int
	for idx := startPos; idx <= len(digits)-digitsStillNeeded; idx++ {
		if digits[idx] > max {
			max = digits[idx]
			maxpos = idx
		}
	}

	if digitsStillNeeded == 1 {
		return max
	}

	return max + findMaxFrom(digits, maxpos+1, digitsStillNeeded-1)
}

func calc(input *aoc.Input, _, _ bool, _ ...any) (any, any) {
	part1 := 0
	part2 := 0

	for _, line := range input.PlainLines() {
		digits := strings.Split(line, "")

		part1 += aoc.Atoi(findMaxFrom(digits, 0, 2))
		part2 += aoc.Atoi(findMaxFrom(digits, 0, 12))
	}

	return part1, part2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
