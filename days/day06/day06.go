package day06

import (
	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day6 struct{}

func (d *Day6) Part1(fileName string) int {
	lines := helpers.GetLines("days/day06", fileName)

	// Initialize map, run, and return the number of coordinates visited
	m := &helpers.Map{}
	m.InitMap(lines)
	m.Run()
	return m.NumVisited()
}

func (d *Day6) Part2(fileName string) int {
	lines := helpers.GetLines("days/day06", fileName)

	// Run the trial without adding any additional obstacles
	m := &helpers.Map{}
	m.InitMap(lines)
	m.Run()

	// Observation: an obstacle catching the guard in a loop must be at one of the sites
	// visited in the previous run, otherwise he would just follow the exact same path
	// Idea: add obstacle in each of these spots, and run, checking for loops
	answer := 0
	for visited := range m.Visited.Coords {
		mNew := helpers.Map{}
		mNew.InitMap(lines)
		mNew.Obstacles[visited] = true
		if (visited != mNew.Position) && !mNew.Run() {
			answer += 1
		}
	}
	return answer
}
