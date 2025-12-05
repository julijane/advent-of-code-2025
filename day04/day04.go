package main

import (
	"github.com/julijane/advent-of-code-2025/aoc"
)

func calc(input *aoc.Input, _, _ bool, _ ...any) (any, any) {
	part1 := 0
	part2 := 0

	g := input.Grid()

	for {
		canRemove := aoc.Coordinates{}

		for _, c := range g.FindAll('@', g.AllCoordinates()) {
			rolls := g.FindAll('@', c.Adjacent(false))
			if len(rolls) < 4 {
				canRemove = append(canRemove, c)
			}
		}
		if len(canRemove) == 0 {
			break
		}

		if part1 == 0 {
			part1 = len(canRemove)
		}

		part2 += len(canRemove)

		g.SetAll(canRemove, '.')
	}

	return part1, part2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
