package day07

import (
	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day7 struct{}

func (d *Day7) Part1(fileName string) int {
	lines := helpers.GetLines("days/day07", fileName)

	answer := 0
	for _, line := range lines {
		lineExp := &helpers.Expression{
			Part: 1,
		}
		lineExp.Load(line)
		if lineExp.Evaluate() {
			answer += lineExp.Target
		}
	}
	return answer
}

func (d *Day7) Part2(fileName string) int {
	lines := helpers.GetLines("days/day07", fileName)

	answer := 0
	for _, line := range lines {
		lineExp := &helpers.Expression{
			Part: 2,
		}
		lineExp.Load(line)
		if lineExp.Evaluate() {
			answer += lineExp.Target
		}
	}
	return answer
}
