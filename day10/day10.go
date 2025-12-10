package main

import (
	"github.com/julijane/advent-of-code-2025/aoc"
)

func calcPart1(target int, buttons []int) int {
	numPressesFor := make(map[int]int)
	numPressesFor[0] = 0

	queue := []int{0}

	for len(queue) > 0 {
		val := queue[0]
		queue = queue[1:]
		numPresses := numPressesFor[val]

		for _, button := range buttons {
			newVal := val ^ button

			if newVal == target {
				return numPresses + 1
			}

			if _, ok := numPressesFor[newVal]; !ok {
				numPressesFor[newVal] = numPresses + 1
				queue = append(queue, newVal)
			}
		}
	}

	return -1 // this should not happen
}

func calc(input *aoc.Input, _, _ bool, params ...any) (any, any) {
	part1 := 0
	part2 := 0

	for _, line := range input.PlainLines() {
		targetStr := aoc.ExtractRegexps(line, `\[.*?\]`)[0]
		target := 0
		for i, ch := range targetStr[1 : len(targetStr)-1] {
			if ch == '#' {
				target += 1 << i
			}
		}

		buttonsStr := aoc.ExtractRegexps(line, `\(.*?\)`)
		buttons := []int{}
		for _, buttonStr := range buttonsStr {
			buttonBits := aoc.ExtractNumbers(buttonStr)
			buttonVal := 0
			for _, bit := range buttonBits {
				buttonVal += 1 << bit
			}
			buttons = append(buttons, buttonVal)
		}

		part1 += calcPart1(target, buttons)
	}

	return part1, part2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
