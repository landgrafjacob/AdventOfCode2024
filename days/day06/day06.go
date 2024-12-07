package day06

import (
	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day6 struct {}

func (d *Day6) Part1(fileName string) int {
	lines := helpers.GetLines("days/day06", fileName)
	m := &helpers.Map{}
	m.InitMap(lines)

	for m.InsideGrid {
		m.Step()
	}
	return len(m.Visited)
}

func (d *Day6) Part2(fileName string) int {
	_ = helpers.GetLineSections("days/day05", fileName)

	return 0
}
