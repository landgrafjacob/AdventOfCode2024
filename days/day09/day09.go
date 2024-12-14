package day09

import (
	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day9 struct{}

func (d *Day9) Part1(fileName string) int {
	lines := helpers.GetLines("days/day09", fileName)
	diskMap := &DiskMap{
		Input: lines[0],
	}
	diskMap.GenerateFileLayout()
	diskMap.CompressFiles()
	return diskMap.GetCheckSum()
}

func (d *Day9) Part2(fileName string) int {
	_ = helpers.GetLines("days/day09", fileName)
	return 0
}

type DiskMap struct {
	Input      string
	FileLayout []int
}

func addCopies(s *[]int, val int, times int) {
	for i := 0; i < times; i++ {
		*s = append(*s, val)
	}
}

func (d *DiskMap) GenerateFileLayout() {
	start := 0
	file := true
	d.FileLayout = []int{}
	for _, char := range d.Input {
		charInt := int(char - '0')
		if file {
			addCopies(&d.FileLayout, start, charInt)
			start += 1
		} else {
			addCopies(&d.FileLayout, -1, charInt)
		}
		file = !file
	}
}

func (d *DiskMap) CompressFiles() {
	left := 0
	right := len(d.FileLayout) - 1

	for left < right {
		if d.FileLayout[left] != -1 {
			left += 1
		} else if d.FileLayout[right] == -1 {
			right -= 1
		} else {
			d.FileLayout[left], d.FileLayout[right] = d.FileLayout[right], d.FileLayout[left]
			left += 1
			right -= 1
		}
	}
}

func (d *DiskMap) GetCheckSum() int {
	answer := 0
	for i, block := range d.FileLayout {
		if block != -1 {
			answer += i * int(block)
		}
	}
	return answer
}
