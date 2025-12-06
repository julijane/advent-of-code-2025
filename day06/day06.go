package main

import (
	"strings"

	"github.com/julijane/advent-of-code-2025/aoc"
)

type column struct {
	numbers []string
	width   int
	op      byte
}

func calc(input *aoc.Input, _, _ bool, _ ...any) (any, any) {
	part1 := 0
	part2 := 0

	lines := input.PlainLines()

	operators := aoc.ExtractRegexps(lines[len(lines)-1], `[+*]\s+`)

	columns := make([]column, 0, len(operators))
	for _, op := range operators {
		columns = append(columns, column{op: op[0], width: len(op) - 1})
	}
	columns[len(columns)-1].width++

	for _, line := range lines[:len(lines)-1] {
		for i, col := range columns {
			n := line[:col.width]
			if len(line) > col.width {
				line = line[col.width+1:]
			}
			col.numbers = append(col.numbers, n)

			columns[i] = col
		}
	}

	for _, col := range columns {
		colSumPart1 := 0
		colSumPart2 := 0

		switch col.op {
		case '+':
			for _, n := range col.numbers {
				colSumPart1 += aoc.Atoi(strings.TrimSpace(n))
			}
		case '*':
			colSumPart1 = 1
			colSumPart2 = 1

			for _, n := range col.numbers {
				colSumPart1 *= aoc.Atoi(strings.TrimSpace(n))
			}
		}

		part1 += colSumPart1

		for colcol := range col.width {
			numberString := ""
			for _, n := range col.numbers {
				numberString += string(n[colcol])
			}
			number := aoc.Atoi(strings.TrimSpace(numberString))

			switch col.op {
			case '+':
				colSumPart2 += number
			case '*':
				colSumPart2 *= number
			}
		}

		part2 += colSumPart2
	}

	return part1, part2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
