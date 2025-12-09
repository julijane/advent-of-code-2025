package main

import (
	"github.com/julijane/advent-of-code-2025/aoc"
)

type twoCoord struct {
	a, b aoc.Coordinate
}

func convertCoords(a, b aoc.Coordinate) twoCoord {
	return twoCoord{
		a: aoc.Coordinate{
			X: min(a.X, b.X),
			Y: min(a.Y, b.Y),
		},
		b: aoc.Coordinate{
			X: max(a.X, b.X),
			Y: max(a.Y, b.Y),
		},
	}
}

func edgeOverlapsRect(edge, rect twoCoord) bool {
	if edge.a.X == edge.b.X {
		if edge.a.X <= rect.a.X || edge.a.X >= rect.b.X {
			return false
		}

		return (edge.a.Y <= rect.a.Y && edge.b.Y > rect.a.Y) ||
			(edge.a.Y < rect.b.Y && edge.b.Y >= rect.b.Y)
	} else {
		if edge.a.Y <= rect.a.Y || edge.a.Y >= rect.b.Y {
			return false
		}

		return (edge.a.X <= rect.a.X && edge.b.X > rect.a.X) ||
			(edge.a.X < rect.b.X && edge.b.X >= rect.b.X)
	}
}

func calc(input *aoc.Input, _, _ bool, params ...any) (any, any) {
	part1 := 0
	part2 := 0

	corners := aoc.Coordinates{}
	for _, line := range input.PlainLines() {
		val := aoc.ExtractNumbers(line)
		corners = append(corners, aoc.Coordinate{X: val[0], Y: val[1]})
	}

	edges := []twoCoord{}
	for i, cornerA := range corners {
		cornerB := corners[(i+1)%len(corners)]
		edges = append(edges, convertCoords(cornerA, cornerB))
	}

	for i, rectCornerA := range corners[1:] {
	rectloop:
		for _, rectCornerB := range corners[i+1:] {
			rect := convertCoords(rectCornerA, rectCornerB)

			area := (rect.b.X - rect.a.X + 1) * (rect.b.Y - rect.a.Y + 1)
			if area > part1 {
				part1 = area
			}

			for _, edge := range edges {
				if edgeOverlapsRect(edge, rect) {
					continue rectloop
				}
			}

			if area > part2 {
				part2 = area
			}
		}
	}

	return part1, part2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
