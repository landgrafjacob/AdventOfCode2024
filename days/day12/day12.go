package day12

import (
	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day struct{}

func (d *Day) Part1(fileName string) int {
	lines := helpers.GetLines("days/day12", fileName)
	g := BuildGarden(lines)
	g.FindRegions()
	return g.ComputePrice(false)
}

func (d *Day) Part2(fileName string) int {
	lines := helpers.GetLines("days/day12", fileName)
	g := BuildGarden(lines)
	g.FindRegions()
	return g.ComputePrice(true)
}

type Garden struct {
	Rows    int
	Columns int
	Plots   [][]rune
	Regions map[rune][]*Region
}

// Take input and initialize a Garden object
func BuildGarden(input []string) Garden {
	plots := [][]rune{}
	for _, line := range input {
		plots = append(plots, []rune(line))
	}

	g := Garden{
		Rows:    len(input),
		Columns: len(input[0]),
		Plots:   plots,
		Regions: make(map[rune][]*Region),
	}

	return g
}

// Compute the price of the fences for the garden
// For part 1, use the area and perimeter of regions
// For part 2, use the area and number of sides of each region
func (g *Garden) ComputePrice(part2 bool) int {
	answer := 0
	for _, rSlice := range g.Regions {
		for _, r := range rSlice {
			if part2 {
				answer += r.GetSides() * r.area
			} else {
				answer += r.area * r.perim
			}
		}
	}
	return answer
}

// Look through garden for new plots, and run a dfs on each plot to fill out new regions
func (g *Garden) FindRegions() {
	for i, row := range g.Plots {
		for j, plot := range row {
			// If the current plot hasn't been seen, start a dfs
			// Note: we use negative values to denote plots that have been seen
			if plot > 0 {
				g.Regions[plot] = append(g.Regions[plot], &Region{
					minX:    g.Rows,
					maxX:    -1,
					minY:    g.Columns,
					maxY:    -1,
					tops:    *helpers.GetEmptySet(),
					bottoms: *helpers.GetEmptySet(),
					rights:  *helpers.GetEmptySet(),
					lefts:   *helpers.GetEmptySet(),
					perim:   0,
					area:    0,
				})
				g.dfs(i, j, plot)
			}
		}
	}
}

// Given coordinates and a plant label, create/fill out a region, and add it to the garden
func (g *Garden) dfs(i, j int, label rune) {
	if g.Plots[i][j] != label {
		return
	}

	// Index of the last region in the garden with the given label
	// Note: there can be multiple regions with the same label.
	// Since we are using a dfs, we are guarateed to completely fill out a region
	// before finding another with the same label
	// So we can always add to the last region
	index := len(g.Regions[label]) - 1

	// Count the current plot in the regions area
	g.Regions[label][index].area += 1

	// Mark the current plot as seen
	g.Plots[i][j] = -1 * g.Plots[i][j]

	// If the plot is either on the top edge of the garden,
	// or the plot above has a different label (excluding sign),
	// then we need a fence along the top
	if i == 0 || helpers.AbsRune(g.Plots[i-1][j]) != label {
		g.Regions[label][index].perim += 1
		g.Regions[label][index].AddTop(i, j)
	} else if g.Plots[i-1][j] == label {
		g.dfs(i-1, j, label)
	}

	// Same for bottom
	if i == g.Rows-1 || helpers.AbsRune(g.Plots[i+1][j]) != label {
		g.Regions[label][index].AddBottom(i, j)
		g.Regions[label][index].perim += 1
	} else if g.Plots[i+1][j] == label {
		g.dfs(i+1, j, label)
	}

	// Same for left
	if j == 0 || helpers.AbsRune(g.Plots[i][j-1]) != label {
		g.Regions[label][index].AddLeft(i, j)
		g.Regions[label][index].perim += 1
	} else if g.Plots[i][j-1] == label {
		g.dfs(i, j-1, label)
	}

	// Same for right
	if j == g.Columns-1 || helpers.AbsRune(g.Plots[i][j+1]) != label {
		g.Regions[label][index].AddRight(i, j)
		g.Regions[label][index].perim += 1
	} else if g.Plots[i][j+1] == label {
		g.dfs(i, j+1, label)
	}
}

// Count the total number of sides in the garden
func (g *Garden) GetSides() int {
	sides := 0
	for _, rSlice := range g.Regions {
		for _, r := range rSlice {
			sides += r.GetLeftRightSides() + r.GetTopBottomSides()
		}
	}
	return sides
}

type Region struct {
	minX    int
	maxX    int
	minY    int
	maxY    int
	tops    helpers.Set
	rights  helpers.Set
	lefts   helpers.Set
	bottoms helpers.Set
	perim   int
	area    int
}

// When inserting a plot into the region, resize the dimensions of the region if necessary
// This will significantly pare down the number of computations needed to find the number of sides later on
func (r *Region) Rescale(i, j int) {
	if i > r.maxX {
		r.maxX = i
	}

	if i < r.minX {
		r.minX = i
	}

	if j > r.maxY {
		r.maxY = j
	}

	if j < r.minY {
		r.minY = j
	}
}

// Add top/bottom/left/right plot to region
func (r *Region) AddTop(i, j int) {
	r.Rescale(i, j)
	r.tops.Add([2]int{i, j})
}

func (r *Region) AddBottom(i, j int) {
	r.Rescale(i, j)
	r.bottoms.Add([2]int{i, j})
}

func (r *Region) AddRight(i, j int) {
	r.Rescale(i, j)
	r.rights.Add([2]int{i, j})
}

func (r *Region) AddLeft(i, j int) {
	r.Rescale(i, j)
	r.lefts.Add([2]int{i, j})
}

// Count the total number of top/bottom sides
func (r *Region) GetTopBottomSides() int {
	topSides := 0
	bottomSides := 0

	// Step through each row. topOn/bottomOn keep track of whether we are currently in the middle of a side
	// If we are not in a side, and the current plot is a top/bottom, then it is the start of a side, and we
	// bump the count
	for i := r.minX; i <= r.maxX; i++ {
		topOn := false
		bottomOn := false
		for j := r.minY; j <= r.maxY; j++ {
			if !topOn && r.tops.IsIn([2]int{i, j}) {
				topSides += 1
			}
			topOn = r.tops.IsIn([2]int{i, j})

			if !bottomOn && r.bottoms.IsIn([2]int{i, j}) {
				bottomSides += 1
			}
			bottomOn = r.bottoms.IsIn([2]int{i, j})
		}
	}
	return topSides + bottomSides
}

// Count the total number of left/right sides
func (r *Region) GetLeftRightSides() int {
	rightSides := 0
	leftSides := 0

	// Same as above, but instead of walking through rows,
	// walk through columns
	for j := r.minY; j <= r.maxY; j++ {
		rightOn := false
		leftOn := false
		for i := r.minX; i <= r.maxX; i++ {
			if !rightOn && r.rights.IsIn([2]int{i, j}) {
				rightSides += 1
			}
			rightOn = r.rights.IsIn([2]int{i, j})

			if !leftOn && r.lefts.IsIn([2]int{i, j}) {
				leftSides += 1
			}
			leftOn = r.lefts.IsIn([2]int{i, j})
		}
	}
	return rightSides + leftSides
}

// Return the total number of sides of the region
func (r *Region) GetSides() int {
	return r.GetLeftRightSides() + r.GetTopBottomSides()
}
