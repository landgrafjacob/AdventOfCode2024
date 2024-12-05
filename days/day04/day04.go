package day04

import (
	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day4 struct {}

func (d *Day4) Part1(fileName string) int {
	lines := helpers.GetLines("days/day04", fileName)

	rows := len(lines)
	columns := len(lines[0])
	lg := &helpers.LetterGrid{
		Rows: rows,
		Columns: columns,
		Grid: lines,
	}

	return lg.CountStrings("XMAS")
}

func (d *Day4) Part2(fileName string) int {
	lines := helpers.GetLines("days/day04", fileName)

	rows := len(lines)
	columns := len(lines[0])
	lg := &helpers.LetterGrid{
		Rows: rows,
		Columns: columns,
		Grid: lines,
	}

	return lg.CountCrosses("MAS")
}
