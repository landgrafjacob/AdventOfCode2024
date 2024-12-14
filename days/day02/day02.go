package day02

import (
	"strconv"
	"strings"

	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day struct{}

// Given a report (as slice of strings), returns whether or not the report is safe
func isSafe(lineList []string) bool {
	if len(lineList) < 2 {
		return true
	}

	// Initialize pointers to step through slice
	left := 0
	right := 1
	leftVal, _ := strconv.Atoi(lineList[left])
	rightVal, _ := strconv.Atoi(lineList[right])

	// Determine from the first two entries if the slice should be increasing or decreasing
	var increasing bool
	if leftVal == rightVal {
		return false
	} else {
		increasing = (leftVal < rightVal)
	}

	for right < len(lineList) {
		leftVal, _ = strconv.Atoi(lineList[left])
		rightVal, _ = strconv.Atoi(lineList[right])

		// Bad cases:
		// - Have been increasing, but found decreasing pair
		// - Have been decreasing, but found increasing pair
		// - Found pair that differs by more than 3
		// Note: the case of an equal pair is handled in the first two cases
		if increasing && (leftVal >= rightVal) {
			return false
		} else if !increasing && (rightVal >= leftVal) {
			return false
		} else if helpers.AbsInt(leftVal, rightVal) > 3 {
			return false
		}

		left += 1
		right += 1
	}
	return true
}

// Given a report (as slice of strings), returns true if the report is safe,
// or can be made safe by removing a single entry
func isSafeWithDampener(lineList []string) bool {
	// Check if the report is already safe
	if isSafe(lineList) {
		return true
	}

	// Remove entries one at a time, and check if the resulting report is safe
	for i, _ := range lineList {
		if isSafe(sliceRemoveIndex(lineList, i)) {
			return true
		}
	}

	// Otherwise, return false
	return false
}

// Given a slice of strings, returns a copy of the slice with the given index removed
func sliceRemoveIndex(lineList []string, index int) []string {
	newSlice := []string{}

	for i, val := range lineList {
		if i == index {
			continue
		}
		newSlice = append(newSlice, val)
	}
	return newSlice
}

func (d *Day) Part1(fileName string) int {
	lines := helpers.GetLines("days/day02", fileName)

	answer := 0
	for _, line := range lines {
		lineList := strings.Fields(line)

		if isSafe(lineList) {
			answer += 1
		}
	}

	return answer
}

func (d *Day) Part2(fileName string) int {
	lines := helpers.GetLines("days/day02", fileName)

	answer := 0
	for _, line := range lines {
		lineList := strings.Fields(line)

		if isSafeWithDampener(lineList) {
			answer += 1
		}
	}

	return answer
}
