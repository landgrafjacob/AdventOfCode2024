package day08

import (
	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day struct{}

func (d *Day) Part1(fileName string) int {
	lines := helpers.GetLines("days/day08", fileName)
	a := &Antenna{}
	a.Init(lines)
	a.GetAntiNodes()
	return len(a.Antinodes)
}

func (d *Day) Part2(fileName string) int {
	lines := helpers.GetLines("days/day08", fileName)
	a := &Antenna{}
	a.Init(lines)
	a.GetNewAntiNodes()
	return len(a.Antinodes)
}

type Antenna struct {
	Rows      int
	Columns   int
	Map       map[rune][][2]int
	Antinodes map[[2]int]bool
}

// Go through lines and create a map which sends a character to the list of coordinates
// at which an antenna of that frequency exists
func (a *Antenna) Init(lines []string) {
	a.Rows = len(lines)
	a.Columns = len(lines[0])
	a.Map = make(map[rune][][2]int)
	a.Antinodes = make(map[[2]int]bool)
	for i, line := range lines {
		for j, char := range line {
			if string(char) != "." {
				if _, ok := a.Map[char]; ok {
					a.Map[char] = append(a.Map[char], [2]int{i, j})
				} else {
					a.Map[char] = [][2]int{{i, j}}
				}
			}
		}
	}
}

// Given two antennae of the same frequency, computes the location of the antinode closer to a
func computeAntiNode(a, b [2]int) [2]int {
	return [2]int{
		2*a[0] - b[0],
		2*a[1] - b[1],
	}
}

// Computes the smallest integer vector pointing from one point to the other
func computeVector(a, b [2]int) [2]int {
	vector := [2]int{
		a[0] - b[0],
		a[1] - b[1],
	}

	gcd := helpers.GCD(vector[0], vector[1])
	vector[0] /= gcd
	vector[1] /= gcd

	return vector
}

// Go through every pair of antennae with the same frequency, find their antinodes, and add it to a.Antinodes
func (a *Antenna) GetAntiNodes() {
	for _, val := range a.Map {
		for i := 0; i < len(val)-1; i++ {
			for j := i + 1; j < len(val); j++ {
				antiNode1 := computeAntiNode(val[i], val[j])
				if a.InGrid(antiNode1) {
					a.Antinodes[antiNode1] = true
				}

				antiNode2 := computeAntiNode(val[j], val[i])
				if a.InGrid(antiNode2) {
					a.Antinodes[antiNode2] = true
				}
			}
		}
	}
}

// Go through every pair of antennae with the same frequency, find their antinodes, and add it to a.Antinodes
func (a *Antenna) GetNewAntiNodes() {
	for _, val := range a.Map {
		for i := 0; i < len(val)-1; i++ {
			for j := i + 1; j < len(val); j++ {
				vector := computeVector(val[i], val[j])

				// Find every integer point starting at val[i] in the direction vector (need to consider both ways)
				coord := val[i]
				for a.InGrid(coord) {
					a.Antinodes[coord] = true
					coord[0] += vector[0]
					coord[1] += vector[1]
				}

				coord = val[i]
				for a.InGrid(coord) {
					a.Antinodes[coord] = true
					coord[0] -= vector[0]
					coord[1] -= vector[1]
				}
			}
		}
	}
}

// Given a coordinate, return true if it is in the grid
func (a *Antenna) InGrid(coord [2]int) bool {
	if coord[0] < 0 || coord[0] >= a.Rows {
		return false
	}
	if coord[1] < 0 || coord[1] >= a.Columns {
		return false
	}
	return true
}
