package main

import (
	"github.com/julijane/advent-of-code-2025/aoc"
)

func findHighestSequence(inputDigits []int, numDigits int) int {
	startPos := 0
	result := 0

	for digit := range numDigits {
		max := -1
		maxpos := -1

		for x := startPos; x <= len(inputDigits)-numDigits+digit; x++ {
			if inputDigits[x] > max {
				max = inputDigits[x]
				maxpos = x
			}
		}

		startPos = maxpos + 1
		result = result*10 + max
	}

	return result
}

func calc(input *aoc.Input, _, _ bool, _ ...any) (any, any) {
	part1 := 0
	part2 := 0

	for _, line := range input.PlainLines() {
		digits := aoc.ExtractDigits(line)

		part1 += findHighestSequence(digits, 2)
		part2 += findHighestSequence(digits, 12)
	}

	return part1, part2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
