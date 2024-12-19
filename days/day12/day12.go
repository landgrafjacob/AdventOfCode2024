package day12

import (
	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day struct{}

func (d *Day) Part1(fileName string) int {
	lines := helpers.GetLines("days/day12", fileName)
	g := BuildGarden(lines)

	return g.ComputePerimArea()
}

func (d *Day) Part2(fileName string) int {
	_ = helpers.GetLines("days/day12", fileName)
	return 0
}

type Garden struct {
	Rows    int
	Columns int
	Plots   [][]rune
}

func BuildGarden(input []string) Garden {
	plots := [][]rune{}
	for _, line := range input {
		plots = append(plots, []rune(line))
	}

	g := Garden{
		Rows:    len(input),
		Columns: len(input[0]),
		Plots:   plots,
	}

	return g
}

func (g *Garden) ComputePerimArea() int {
	answer := 0

	for i, row := range g.Plots {
		for j, plot := range row {
			if plot > 0 {
				result := g.dfs(i, j, plot)
				answer += result[0] * result[1]
			}
		}
	}
	return answer
}

func (g *Garden) dfs(i, j int, label rune) [2]int {
	if g.Plots[i][j] != label {
		return [2]int{0, 0}
	}

	perim := 0
	area := 1

	g.Plots[i][j] = -1 * g.Plots[i][j]

	if i == 0 || helpers.AbsRune(g.Plots[i-1][j]) != label {
		// fmt.Println("Boundary between", i, j, "and", i-1, j)
		perim += 1
	} else if g.Plots[i-1][j] == label {
		result := g.dfs(i-1, j, label)
		perim += result[0]
		area += result[1]
	}

	if i == g.Rows-1 || helpers.AbsRune(g.Plots[i+1][j]) != label {
		// fmt.Println("Boundary between", i, j, "and", i+1, j)
		perim += 1
	} else if g.Plots[i+1][j] == label {
		result := g.dfs(i+1, j, label)
		perim += result[0]
		area += result[1]
	}

	if j == 0 || helpers.AbsRune(g.Plots[i][j-1]) != label {
		// fmt.Println("Boundary between", i, j, "and", i, j-1)
		perim += 1
	} else if g.Plots[i][j-1] == label {
		result := g.dfs(i, j-1, label)
		perim += result[0]
		area += result[1]
	}

	if j == g.Columns-1 || helpers.AbsRune(g.Plots[i][j+1]) != label {
		// fmt.Println("Boundary between", i, j, "and", i, j+1)
		perim += 1
	} else if g.Plots[i][j+1] == label {
		result := g.dfs(i, j+1, label)
		perim += result[0]
		area += result[1]
	}

	return [2]int{perim, area}
}
