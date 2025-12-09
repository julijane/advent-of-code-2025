package main

import (
	"github.com/julijane/advent-of-code-2025/aoc"
)

type edge struct {
	a, b aoc.Coordinate
}

func calc(input *aoc.Input, _, _ bool, params ...any) (any, any) {
	nodes := aoc.Coordinates{}
	edges := []edge{}

	for i, line := range input.PlainLines() {
		val := aoc.ExtractNumbers(line)
		coord := aoc.Coordinate{X: val[0], Y: val[1]}

		nodes = append(nodes, coord)
		if i > 0 {
			edges = append(edges, edge{a: nodes[i-1], b: coord})
		}
	}
	edges = append(edges, edge{a: nodes[len(nodes)-1], b: nodes[0]})

	part1 := 0
	part2 := 0

	for i, a := range nodes {
	rectloop:
		for _, b := range nodes[i+1:] {
			nodeXMin := min(a.X, b.X)
			nodeXMax := max(a.X, b.X)
			nodeYMin := min(a.Y, b.Y)
			nodeYMax := max(a.Y, b.Y)

			area := (nodeXMax - nodeXMin + 1) * (nodeYMax - nodeYMin + 1)
			if area > part1 {
				part1 = area
			}

		lineloop:
			for _, edge := range edges {
				if edge.a.X == edge.b.X {
					if edge.a.X <= nodeXMin || edge.a.X >= nodeXMax {
						continue lineloop
					}

					edgeYMin := min(edge.a.Y, edge.b.Y)
					edgeYMax := max(edge.a.Y, edge.b.Y)

					if (edgeYMin <= nodeYMin && edgeYMax > nodeYMin) ||
						(edgeYMin < nodeYMax && edgeYMax >= nodeYMax) {
						continue rectloop
					}
				} else {
					if edge.a.Y <= nodeYMin || edge.a.Y >= nodeYMax {
						continue lineloop
					}

					edgeXMin := min(edge.a.X, edge.b.X)
					edgeXMax := max(edge.a.X, edge.b.X)

					if (edgeXMin <= nodeXMin && edgeXMax > nodeXMin) ||
						(edgeXMin < nodeXMax && edgeXMax >= nodeXMax) {
						continue rectloop
					}
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
