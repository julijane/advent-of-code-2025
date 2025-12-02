package main

import (
	"fmt"
	"strings"

	"github.com/julijane/advent-of-code-2025/aoc"
)

func calc(input *aoc.Input, _, _ bool, _ ...any) (any, any) {
	part1 := 0
	part2 := 0

	numRanges := input.FindObjects(`(\d+)-(\d+)`)

	for _, numRange := range numRanges {
		numbers := strings.Split(numRange.String(), "-")

		fromNumber := aoc.Atoi(numbers[0])
		toNumber := aoc.Atoi(numbers[1])

		for testNumber := fromNumber; testNumber <= toNumber; testNumber++ {
			numberString := fmt.Sprintf("%d", testNumber)

			halfLen := len(numberString) / 2
			if numberString[0:halfLen] == numberString[halfLen:] {
				part1 += testNumber
			}

		testloop:
			for testlen := 1; testlen <= halfLen; testlen++ {
				if len(numberString)%testlen != 0 {
					continue
				}

				substr := numberString[0:testlen]

				for pos := testlen; pos < len(numberString); pos += testlen {
					if numberString[pos:pos+testlen] != substr {
						continue testloop
					}
				}

				part2 += testNumber
				break
			}
		}
	}

	return part1, part2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
