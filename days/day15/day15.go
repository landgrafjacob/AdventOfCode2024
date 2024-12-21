package day15

import (
	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day struct{}

func (d *Day) Part1(fileName string) int {
	lines := helpers.GetLineSections("days/day15", fileName)
	w := BuildWarehouse(lines[0], true)
	for _, line := range lines[1] {
		for _, r := range line {
			w.Move(r)
		}
	}
	return w.GetGPSSum()
}

func (d *Day) Part2(fileName string) int {
	lines := helpers.GetLineSections("days/day15", fileName)
	w := BuildWarehouse(lines[0], false)
	for _, line := range lines[1] {
		for _, r := range line {
			w.MovePart2(r)
		}
	}
	return w.GetGPSSum()
}

// Convert the arrow runes to a direction vector
func getDirection(r rune) [2]int {
	if string(r) == "^" {
		return [2]int{-1, 0}
	} else if string(r) == "<" {
		return [2]int{0, -1}
	} else if string(r) == ">" {
		return [2]int{0, 1}
	} else if string(r) == "v" {
		return [2]int{1, 0}
	}
	return [2]int{}
}

type Warehouse struct {
	Robot   [2]int
	Boxes   helpers.Set
	Walls   helpers.Set
	Rows    int
	Columns int
}

// Initialize a warehouse from the input
func BuildWarehouse(input []string, part1 bool) *Warehouse {
	// Stretch width by factor of 2 for part 2
	var rows, columns int
	if part1 {
		rows = len(input)
		columns = len(input[0])
	} else {
		rows = len(input)
		columns = 2 * len(input[0])
	}
	w := &Warehouse{
		Boxes:   *helpers.GetEmptySet(),
		Walls:   *helpers.GetEmptySet(),
		Rows:    rows,
		Columns: columns,
	}

	// Insert Robot/Walls/Boxes
	// Note: for part 2, we are only storing the left side of each box
	for i, line := range input {
		for j, c := range line {
			if string(c) == "O" {
				if part1 {
					w.Boxes.Add([2]int{i, j})
				} else {
					w.Boxes.Add([2]int{i, 2 * j})
				}
			} else if string(c) == "#" {
				if part1 {
					w.Walls.Add([2]int{i, j})
				} else {
					w.Walls.Add([2]int{i, 2 * j})
					w.Walls.Add([2]int{i, 2*j + 1})
				}
			} else if string(c) == "@" {
				if part1 {
					w.Robot = [2]int{i, j}
				} else {
					w.Robot = [2]int{i, 2 * j}
				}
			}
		}
	}
	return w
}

// Part 1 move
func (w *Warehouse) Move(r rune) {
	d := getDirection(r)

	// Calculate the new position
	pos := [2]int{
		w.Robot[0] + d[0],
		w.Robot[1] + d[1],
	}

	// If we land on a box, keep stepping to find out whether there
	// is a wall or empty space on the other side of the row of boxes
	for w.Boxes.IsIn(pos) {
		pos[0] += d[0]
		pos[1] += d[1]
	}

	// If a wall lies on at the end, do nothing
	if w.Walls.IsIn(pos) {
		return
	}

	// If an empty spaces is at the end, make the step
	w.Robot[0] += d[0]
	w.Robot[1] += d[1]

	// Remove box from new position, and place at the end of the line
	if w.Robot != pos {
		w.Boxes.Add(pos)
		w.Boxes.Remove(w.Robot)
	}
}

// Calculate the sum of the GPS coordinates of the boxes
func (w *Warehouse) GetGPSSum() int {
	answer := 0
	for _, coord := range w.Boxes.GetElements() {
		answer += 100*coord[0] + coord[1]
	}
	return answer
}

// Given the coordinates of a box and direction, recursively decide if the box can be moved in that direction
// If the box can be moved, returns a list of the boxes that have to be moved also.
func (w *Warehouse) CanMove(coord, d [2]int) (bool, [][2]int) {
	// Check if hitting wall
	newCoords := [2]int{
		coord[0] + d[0],
		coord[1] + d[1],
	}
	adjCoords := [2]int{
		newCoords[0],
		newCoords[1] + 1,
	}

	if w.Walls.IsIn(newCoords) || w.Walls.IsIn(adjCoords) {
		return false, [][2]int{}
	}

	var toCheck [][2]int
	if d == [2]int{0, 1} {
		toCheck = [][2]int{{coord[0], coord[1] + 2}}
	} else if d == [2]int{0, -1} {
		toCheck = [][2]int{{coord[0], coord[1] - 2}}
	} else {
		toCheck = [][2]int{
			{coord[0] + d[0], coord[1]},
			{coord[0] + d[0], coord[1] + 1},
			{coord[0] + d[0], coord[1] - 1},
		}
	}

	answer := [][2]int{coord}

	for _, check := range toCheck {
		if w.Boxes.IsIn(check) {
			canMoveBox, boxList := w.CanMove(check, d)
			if !canMoveBox {
				return false, [][2]int{}
			}
			answer = append(answer, boxList...)
		}
	}

	return true, answer
}

// Takes the output of CanMove and moves the given boxes
func (w *Warehouse) MoveBoxes(boxes [][2]int, d [2]int) {
	newPositions := [][2]int{}
	for _, box := range boxes {
		newPositions = append(newPositions, [2]int{
			box[0] + d[0],
			box[1] + d[1],
		})
		w.Boxes.Remove(box)
	}
	for _, box := range newPositions {
		w.Boxes.Add(box)
	}
}

// Part 2 move
func (w *Warehouse) MovePart2(r rune) {
	d := getDirection(r)
	newCoord := [2]int{
		w.Robot[0] + d[0],
		w.Robot[1] + d[1],
	}

	adjCoord := [2]int{
		newCoord[0],
		newCoord[1] - 1,
	}

	// If our new position is on a wall, do nothing
	if w.Walls.IsIn(newCoord) {
		return
	}

	// If our new position is in a box, try to move the box and all boxes behind it
	if w.Boxes.IsIn(newCoord) {
		ok, boxes := w.CanMove(newCoord, d)
		if ok {
			w.MoveBoxes(boxes, d)
			w.Robot = newCoord
		}
	} else if w.Boxes.IsIn(adjCoord) {
		ok, boxes := w.CanMove(adjCoord, d)
		if ok {
			w.MoveBoxes(boxes, d)
			w.Robot = newCoord
		}
	} else {
		// If the new position is not on a box or wall, just move
		w.Robot = newCoord
	}
}
