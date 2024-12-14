package day11

import (
	"strconv"
	"strings"

	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day struct{}

func (d *Day) Part1(fileName string) int {
	lines := helpers.GetLines("days/day11", fileName)
	sl := GetStoneLine(lines[0])
	for i := 0; i < 25; i++ {
		sl.Blink()
	}
	return sl.GetNumStones()
}

func (d *Day) Part2(fileName string) int {
	lines := helpers.GetLines("days/day11", fileName)
	sl := GetStoneLine(lines[0])
	for i := 0; i < 75; i++ {
		sl.Blink()
	}
	return sl.GetNumStones()
}

// Key observation: the order of the stones doesn't matter
// Only keep track of the number of each stone by value
type StoneLine struct {
	CountMap map[int]int
}

// Initialize stones from input
func GetStoneLine(line string) StoneLine {
	s := StoneLine{
		CountMap: make(map[int]int),
	}

	for _, field := range strings.Fields(line) {
		fieldInt, _ := strconv.Atoi(field)
		s.CountMap[fieldInt] += 1
	}

	return s
}

// Blink each stone in the line
func (sl *StoneLine) Blink() {
	newMap := make(map[int]int)

	for key, count := range sl.CountMap {
		keyString := strconv.Itoa(key)
		keyLength := len(keyString)

		if key == 0 {
			// All the stones marked 0 turn into 1's
			newMap[1] += count
		} else if keyLength%2 == 0 {
			// All the stones with even number of digits get split apart
			leftString := keyString[:keyLength/2]
			leftInt, _ := strconv.Atoi(leftString)
			rightString := keyString[keyLength/2:]
			rightInt, _ := strconv.Atoi(rightString)

			newMap[leftInt] += count
			newMap[rightInt] += count
		} else {
			// All others get multiplied to 2024
			newMap[key*2024] += count
		}
	}
	sl.CountMap = newMap
}

// Count all the stones in the line
func (sl *StoneLine) GetNumStones() int {
	answer := 0
	for _, count := range sl.CountMap {
		answer += count
	}
	return answer
}
