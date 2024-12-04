package day01

import (
	"strings"
	"strconv"
	"slices"

	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day1 struct {}

// Given a slice representing lines of the file, returns the two columns (as slices of ints)
func getColumns(fileSlice []string) ([]int, []int) {
	var firstSlice, secondSlice []int

	for _, line := range fileSlice {
		lineList := strings.Fields(line)
		first, _ := strconv.Atoi(lineList[0])
		second, _ := strconv.Atoi(lineList[1])

		firstSlice = append(firstSlice, first)
		secondSlice = append(secondSlice, second)
	}

	return firstSlice, secondSlice
}

func (d *Day1) Part1(fileName string) int {
	lines := helpers.GetLines("days/day01", fileName)
	first, second := getColumns(lines)

	slices.Sort(first)
	slices.Sort(second)

	var answer int
	var lineDiff int
	// Compute differences of sorted slices, line by line
	for i, val := range first {
		if val > second[i] {
			lineDiff = val - second[i]
		} else {
			lineDiff = second[i] - val
		}

		answer += lineDiff
	}

	return answer
}

func (d *Day1) Part2(fileName string) int {
	lines := helpers.GetLines("days/day01", fileName)
	first, second := getColumns(lines)
	
	// Count the occurences of each number in the second column
	countMap := make(map[int]int)

	for _, val := range second {
		countMap[val] += 1
	}

	// Compute the similarity score
	var answer int
	for _, val := range first {
		answer += val * countMap[val]
	}
	return answer
}
