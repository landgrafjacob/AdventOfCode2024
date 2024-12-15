package day14

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
	"strconv"

	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day struct{}

func (d *Day) Part1(fileName string) int {
	lines := helpers.GetLineSections("days/day14", fileName)
	g := MakeGrid(lines)
	for i := 0; i < 100; i++ {
		g.RunRobots()
	}
	return g.SafetyFactor()
}

func (d *Day) Part2(fileName string) int {
	lines := helpers.GetLineSections("days/day14", fileName)
	g := MakeGrid(lines)
	seconds := 0

	// Run the robots for 10000 seconds, and generate images for any configuration
	// where the robots are all in unique places
	for {
		allUnique := g.AllUnique()
		if allUnique {
			fmt.Println("Candidate found at", seconds, "seconds")
			g.Draw(seconds)
		}
		g.RunRobots()
		seconds += 1
		if seconds == 10000 {
			break
		}
	}
	return -1
}

type Grid struct {
	Rows    int
	Columns int
	Robots  []*Robot
}

type Robot struct {
	Columns   int
	Rows      int
	X         int
	Y         int
	VelocityX int
	VelocityY int
}

// Initialize the robots on the grid
// Note: we have added a block on the bottom of the input/test file to indicate the width/height of the grid
func MakeGrid(s [][]string) Grid {
	g := Grid{
		Robots: []*Robot{},
	}
	r, _ := regexp.Compile("-?[0-9]+")

	dimSlice := r.FindAllString(s[1][0], -1)
	g.Columns, _ = strconv.Atoi(dimSlice[0])
	g.Rows, _ = strconv.Atoi(dimSlice[1])

	for _, line := range s[0] {
		coords := r.FindAllString(line, -1)
		px, _ := strconv.Atoi(coords[0])
		py, _ := strconv.Atoi(coords[1])
		vx, _ := strconv.Atoi(coords[2])
		vy, _ := strconv.Atoi(coords[3])
		g.Robots = append(g.Robots, &Robot{
			X:         px,
			Y:         py,
			VelocityX: vx,
			VelocityY: vy,
			Rows:      g.Rows,
			Columns:   g.Columns,
		})
	}

	return g
}

// Have robot move for 1 second according to its velocity
func (r *Robot) Run() {
	r.X = helpers.Mod(r.X+r.VelocityX, r.Columns)
	r.Y = helpers.Mod(r.Y+r.VelocityY, r.Rows)
}

// Run all robots on the grid for one second
func (g *Grid) RunRobots() {
	for _, r := range g.Robots {
		r.Run()
	}
}

// Compute safety factor of grid in current state
func (g *Grid) SafetyFactor() int {
	midX := (g.Columns - 1) / 2
	midY := (g.Rows - 1) / 2

	quad1 := 0
	quad2 := 0
	quad3 := 0
	quad4 := 0

	for _, robot := range g.Robots {
		if robot.X < midX && robot.Y < midY {
			quad1 += 1
		} else if robot.X > midX && robot.Y < midY {
			quad2 += 1
		} else if robot.X < midX && robot.Y > midY {
			quad3 += 1
		} else if robot.X > midX && robot.Y > midY {
			quad4 += 1
		}
	}

	return quad1 * quad2 * quad3 * quad4
}

// Create an image of the current grid
func (g *Grid) Draw(fileNum int) {
	upLeft := image.Point{0, 0}
	downRight := image.Point{g.Columns, g.Rows}
	img := image.NewRGBA(image.Rectangle{upLeft, downRight})

	cyan := color.RGBA{100, 200, 200, 0xff}

	for _, robot := range g.Robots {
		img.Set(robot.X, robot.Y, cyan)
	}
	fileName := fmt.Sprintf("images/image%03d.png", fileNum)
	f, _ := os.Create(fileName)
	png.Encode(f, img)
	f.Close()
}

// Return true if no two robots share a space, otherwise return false
func (g *Grid) AllUnique() bool {
	seen := make(map[[2]int]bool)

	for _, robot := range g.Robots {
		if _, ok := seen[[2]int{robot.X, robot.Y}]; ok {
			return false
		}
		seen[[2]int{robot.X, robot.Y}] = true
	}
	return true
}
