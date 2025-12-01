package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

func calc(input *aoc.Input, _, _ bool, _ ...any) (any, any) {
	dialPos := 50
	numZeroPart1 := 0
	numZeroPart2 := 0

	for _, line := range input.PlainLines() {
		dir := line[0]
		steps := aoc.Atoi(line[1:])

		fullRotations := steps / 100
		steps = steps % 100

		switch dir {
		case 'L':
			if dialPos != 0 && steps > dialPos {
				numZeroPart2 += 1
			}
			dialPos = (dialPos - steps + 100) % 100
		case 'R':
			if dialPos != 0 && steps > (100-dialPos) {
				numZeroPart2 += 1
			}
			dialPos = (dialPos + steps + 100) % 100
		}

		numZeroPart2 += fullRotations

		if dialPos == 0 {
			numZeroPart1++
			numZeroPart2++
		}

		// fmt.Printf("After %c%02d: dial at %02d; numZeroPart1=%d, numZeroPart2=%d\n", dir, steps, dialPos, numZeroPart1, numZeroPart2)
	}

	return numZeroPart1, numZeroPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
