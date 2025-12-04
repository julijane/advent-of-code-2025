package main

import (
	"github.com/julijane/advent-of-code-2025/aoc"
)

func getRemovable(g *aoc.Grid) []aoc.Coordinate {
	removables := []aoc.Coordinate{}

search:
	for _, c := range g.FindAll('@') {
		countRolls := 0

		for _, dir := range aoc.DirsAll {
			if g.Get(c.Add(dir), '.') == '@' {
				countRolls++

				if countRolls > 3 {
					continue search
				}
			}
		}

		removables = append(removables, c)
	}

	return removables
}

func calc(input *aoc.Input, _, _ bool, _ ...any) (any, any) {
	part1 := 0
	part2 := 0

	g := input.Grid()

	removables := getRemovable(g)
	part1 = len(removables)

	for len(removables) > 0 {
		part2 += len(removables)

		for _, c := range removables {
			g.Set(c, '.')
		}

		removables = getRemovable(g)
	}

	return part1, part2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
