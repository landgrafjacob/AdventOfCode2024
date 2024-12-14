package day10

import (
	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day struct{}

func (d *Day) Part1(fileName string) int {
	lines := helpers.GetLines("days/day10", fileName)
	m := BuildTopMap(lines)
	return m.SumTrailheads()
}

func (d *Day) Part2(fileName string) int {
	lines := helpers.GetLines("days/day10", fileName)
	m := BuildTopMap(lines)
	return m.SumTrails()
}

func BuildTopMap(lines []string) *TopMap {
	m := &TopMap{
		Rows:    len(lines),
		Columns: len(lines[0]),
		Grid:    [][]Node{},
	}
	for _, line := range lines {
		nodeLine := []Node{}
		for _, char := range line {
			nodeLine = append(nodeLine, Node{
				Height: int(char - '0'),
			})
		}
		m.Grid = append(m.Grid, nodeLine)
	}
	return m
}

type TopMap struct {
	Rows    int
	Columns int
	Grid    [][]Node
}

type Node struct {
	Height int
}

// Use dfs to find the set of trails ends reachable from given coordinates i,j
func (tm *TopMap) FindTrailEnds(i, j int) *helpers.Set {
	if i < 0 || i >= tm.Rows {
		return helpers.GetEmptySet()
	} else if j < 0 || j >= tm.Columns {
		return helpers.GetEmptySet()
	}
	trailEnds := helpers.GetEmptySet()
	currentHeight := tm.Grid[i][j].Height
	if currentHeight == 9 {
		trailEnds.Add([2]int{i, j})
		return trailEnds
	}

	if i > 0 && tm.Grid[i-1][j].Height == currentHeight+1 {
		trailEnds = trailEnds.Union(tm.FindTrailEnds(i-1, j))
	}
	if i < tm.Rows-1 && tm.Grid[i+1][j].Height == currentHeight+1 {
		trailEnds = trailEnds.Union(tm.FindTrailEnds(i+1, j))
	}
	if j > 0 && tm.Grid[i][j-1].Height == currentHeight+1 {
		trailEnds = trailEnds.Union(tm.FindTrailEnds(i, j-1))
	}
	if j < tm.Columns-1 && tm.Grid[i][j+1].Height == currentHeight+1 {
		trailEnds = trailEnds.Union(tm.FindTrailEnds(i, j+1))
	}

	return trailEnds
}

// Find all trail heads and sum the number of trail ends reachable from them
func (tm *TopMap) SumTrailheads() int {
	answer := 0
	for i, nodeList := range tm.Grid {
		for j, node := range nodeList {
			if node.Height == 0 {
				answer += tm.FindTrailEnds(i, j).Length()
			}
		}
	}
	return answer
}

// Use dfs to find the total number of trails to trail ends
func (tm *TopMap) TotalTrails(i, j int) int {
	if i < 0 || i >= tm.Rows {
		return 0
	} else if j < 0 || j >= tm.Columns {
		return 0
	}
	totalTrails := 0
	currentHeight := tm.Grid[i][j].Height
	if currentHeight == 9 {
		return 1
	}

	if i > 0 && tm.Grid[i-1][j].Height == currentHeight+1 {
		totalTrails += tm.TotalTrails(i-1, j)
	}
	if i < tm.Rows-1 && tm.Grid[i+1][j].Height == currentHeight+1 {
		totalTrails += tm.TotalTrails(i+1, j)
	}
	if j > 0 && tm.Grid[i][j-1].Height == currentHeight+1 {
		totalTrails += tm.TotalTrails(i, j-1)
	}
	if j < tm.Columns-1 && tm.Grid[i][j+1].Height == currentHeight+1 {
		totalTrails += tm.TotalTrails(i, j+1)
	}

	return totalTrails
}

// Count all possible trails from 0's to 9's
func (tm *TopMap) SumTrails() int {
	answer := 0
	for i, nodeList := range tm.Grid {
		for j, node := range nodeList {
			if node.Height == 0 {
				answer += tm.TotalTrails(i, j)
			}
		}
	}
	return answer
}
