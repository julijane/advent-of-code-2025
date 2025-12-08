package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/julijane/advent-of-code-2025/aoc"
)

type loc struct {
	x, y, z int
}

func (l *loc) distance(o *loc) float64 {
	dx := float64(l.x - o.x)
	dy := float64(l.y - o.y)
	dz := float64(l.z - o.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

type distance struct {
	a, b loc
	d    float64
}

type circuit []loc

func calc(input *aoc.Input, _, _ bool, params ...any) (any, any) {

	locations := []loc{}

	for _, line := range input.PlainLines() {
		var l loc
		_, err := fmt.Sscanf(line, "%d,%d,%d", &l.x, &l.y, &l.z)
		if err != nil {
			continue
		}
		locations = append(locations, l)
	}

	distances := []distance{}
	for i, a := range locations {
		for _, b := range locations[i+1:] {
			distances = append(distances, distance{a: a, b: b, d: a.distance(&b)})
		}
	}

	slices.SortFunc(distances, func(i, j distance) int {
		if i.d < j.d {
			return -1
		} else if i.d > j.d {
			return 1
		}
		return 0
	})

	param := params[0].([]interface{})
	numConnections := param[0].(int)

	circuits := []circuit{}

	part1 := 0
	part2 := 0

	for i, dist := range distances {

		if i == numConnections {
			slices.SortFunc(circuits, func(i, j circuit) int {
				if len(i) > len(j) {
					return -1
				} else if len(i) < len(j) {
					return 1
				}
				return 0
			})

			part1 = len(circuits[0]) * len(circuits[1]) * len(circuits[2])
		}

		circuitContainsA := -1
		circuitContainsB := -1

		for i, c := range circuits {
			if slices.Contains(c, dist.a) {
				circuitContainsA = i
			}

			if slices.Contains(c, dist.b) {
				circuitContainsB = i
			}

		}

		if circuitContainsA == circuitContainsB {
			if circuitContainsA == -1 {
				circuits = append(circuits, circuit{dist.a, dist.b})
			}
		} else if circuitContainsA == -1 {
			circuits[circuitContainsB] = append(circuits[circuitContainsB], dist.a)
		} else if circuitContainsB == -1 {
			circuits[circuitContainsA] = append(circuits[circuitContainsA], dist.b)
		} else {
			circuits[circuitContainsA] = append(circuits[circuitContainsA], circuits[circuitContainsB]...)
			circuits = append(circuits[:circuitContainsB], circuits[circuitContainsB+1:]...)
		}

		if len(circuits[0]) == len(locations) {
			part2 = dist.a.x * dist.b.x
			break
		}
	}

	return part1, part2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true, 10)
	aoc.Run("input.txt", calc, true, true, 1000)
}
