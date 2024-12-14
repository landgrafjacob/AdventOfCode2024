package day09

import (
	"slices"

	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day struct{}

func (d *Day) Part1(fileName string) int {
	lines := helpers.GetLines("days/day09", fileName)
	diskMap := &DiskMap{
		Input: lines[0],
	}
	diskMap.GenerateFileLayout()
	diskMap.CompressFiles()
	return diskMap.GetCheckSum()
}

func (d *Day) Part2(fileName string) int {
	lines := helpers.GetLines("days/day09", fileName)
	fileBlocks := &FileBlocks{
		Input: lines[0],
	}
	fileBlocks.Init()
	fileBlocks.ShiftAll()
	return fileBlocks.ComputeCheckSum()
}

// Struct to store file layout
type DiskMap struct {
	Input      string
	FileLayout []int
}

func addCopies(s *[]int, val int, times int) {
	for i := 0; i < times; i++ {
		*s = append(*s, val)
	}
}

// Given diskmap generate the expanded file layout
// Store empty blocks as -1
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

// Using two pointers, gather the file blocks on the left
func (d *DiskMap) CompressFiles() {
	left := 0
	right := len(d.FileLayout) - 1

	for left < right {
		if d.FileLayout[left] != -1 {
			// If the left pointer is on a file, bump up one
			left += 1
		} else if d.FileLayout[right] == -1 {
			// If the right pointer is on an empty space, bump down one
			right -= 1
		} else {
			// Otherwise, switch the empty block on the left with the file block on the right
			d.FileLayout[left], d.FileLayout[right] = d.FileLayout[right], d.FileLayout[left]
			left += 1
			right -= 1
		}
	}
}

// Compute checksum
func (d *DiskMap) GetCheckSum() int {
	answer := 0
	for i, block := range d.FileLayout {
		if block != -1 {
			answer += i * int(block)
		}
	}
	return answer
}

// For part two, we store files in blocks instead
type FileBlocks struct {
	Input    string
	Blocks   []Block
	BlockMap map[int]Block
	NumFiles int
}

type Block struct {
	FileID int
	Length int
	IsFile bool
}

// Computes the checksum of a block starting at the given index
func (b *Block) CheckSum(startIndex int) int {
	if !b.IsFile {
		return 0
	}
	return b.FileID * (b.Length*startIndex + (b.Length * (b.Length - 1) / 2))
}

// Converts the diskmap to a slice of blocks
func (fb *FileBlocks) Init() {
	fb.Blocks = []Block{}
	fb.BlockMap = make(map[int]Block)
	fb.NumFiles = 0
	file := true
	for _, char := range fb.Input {
		if char == '0' {
			file = !file
			continue
		}
		charInt := int(char - '0')
		if file {
			fb.Blocks = append(fb.Blocks, Block{
				Length: charInt,
				FileID: fb.NumFiles,
				IsFile: true,
			})
			fb.NumFiles += 1

		} else {
			fb.Blocks = append(fb.Blocks, Block{
				Length: charInt,
				FileID: -1,
				IsFile: false,
			})
		}
		file = !file
	}
}

// Look for a block with the given file ID, find the furthest-left empty block that can fit it, and shift it there
func (fb *FileBlocks) Shift(file int) {
	var fileIndex int
	for j := len(fb.Blocks) - 1; j >= 0; j-- {
		if fb.Blocks[j].FileID == file {
			fileIndex = j
			break
		}
	}

	for i := 0; i < fileIndex; i++ {
		// If the block is empty and has the same length as the file being shifted, just switch the blocks
		if !fb.Blocks[i].IsFile && fb.Blocks[i].Length == fb.Blocks[fileIndex].Length {
			fb.Blocks[i], fb.Blocks[fileIndex] = fb.Blocks[fileIndex], fb.Blocks[i]
			return
		}

		// If the block is empty and has length larger than the file being shifted, split the empty block in two and then switch the blocks
		if !fb.Blocks[i].IsFile && fb.Blocks[i].Length > fb.Blocks[fileIndex].Length {
			// Create new block with length the excess of the file block length
			newBlock := Block{
				Length: fb.Blocks[i].Length - fb.Blocks[fileIndex].Length,
				FileID: -1,
				IsFile: false,
			}
			// Shrink the empty block to have same length as the file
			fb.Blocks[i].Length = fb.Blocks[fileIndex].Length

			// Insert the new block
			fb.Blocks = slices.Insert(fb.Blocks, i+1, newBlock)

			// Swap the file block and the same-length empty block
			fb.Blocks[i], fb.Blocks[fileIndex+1] = fb.Blocks[fileIndex+1], fb.Blocks[i]
			return
		}
	}
}

// Starting from the highest file ID, shift all blocks to the left
func (fb *FileBlocks) ShiftAll() {
	for i := fb.NumFiles - 1; i >= 0; i-- {
		fb.Shift(i)
	}
}

func (fb *FileBlocks) ComputeCheckSum() int {
	answer := 0
	index := 0
	for _, b := range fb.Blocks {
		answer += b.CheckSum(index)
		index += b.Length
	}
	return answer
}
