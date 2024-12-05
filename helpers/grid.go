package helpers

import "fmt"

type LetterGrid struct {
	Rows int
	Columns int
	Grid []string
}

type Cross struct {
	Center [2]int
	UpLeft [2]int
	UpRight [2]int
	DownLeft [2]int
	DownRight [2]int
	isValid bool
}

func BuildLine(start [2]int, direction [2]int, length int) [][2]int {
	line := [][2]int{start}

	newCoords := start
	for i := 1; i < length; i++ {
		newCoords[0] += direction[0]
		newCoords[1] += direction[1]

		line = append(line, newCoords)
	}
	return line
}

func (l *LetterGrid) GetDirections(start [2]int, length int)  [][][2]int {
	lines := [][][2]int{}

	// Check if there is enough room in the Grid to fit a line starting at the given coordinates 
	// and of the given length in the directions down, down-right, right, up-right
	// Note: the other four directions will be reverses of these, so we don't need to include them
	if start[0] + length - 1 < l.Rows {
		direction := [2]int{1,0}
		lines = append(lines, BuildLine(start, direction, length))
	}
	if start[0] + length - 1 < l.Rows && start[1] + length - 1 < l.Columns {
		direction := [2]int{1,1}
		lines = append(lines, BuildLine(start, direction, length))
	}
	if start[1] + length - 1 < l.Columns {
		direction := [2]int{0,1}
		lines = append(lines, BuildLine(start, direction, length))
	}
	if start[0] + 1 >= length  && start[1] + length - 1 < l.Columns {
		direction := [2]int{-1,1}
		lines = append(lines, BuildLine(start, direction, length))
	}

	return lines
}

// Build a cross at the given Center
func (l *LetterGrid) GetCross(center [2]int) Cross {
	if center[0] <= 0 || center[0] >= l.Rows-1 {
		return Cross{isValid: false}
	} else if center[1] <= 0 || center[1] >= l.Columns-1 {
		return Cross{isValid: false}
	}

	cross := Cross{
		Center: center,
		UpLeft: [2]int{center[0]-1,center[1]-1},
		UpRight: [2]int{center[0]-1,center[1]+1},
		DownLeft: [2]int{center[0]+1,center[1]-1},
		DownRight: [2]int{center[0]+1,center[1]+1},
		isValid: true,
	}
	return cross
}

func (l *LetterGrid) GetLetter(coords [2]int) string {
	return string(l.Grid[coords[0]][coords[1]])
}

func (l *LetterGrid) GetStrings(start [2]int, length int) []string {
	strings := []string{}

	lines := l.GetDirections(start, length)
	for _, line := range lines {
		s := ""
		for _, coord := range line {
			s += l.GetLetter(coord)
		}
		strings = append(strings, s)
	}
	return strings
}

func (l *LetterGrid) CountStrings(search string) int {
	length := len(search)
	revSearch := Reverse(search)
	answer := 0
	for i := 0; i < l.Rows; i++ {
		for j := 0; j < l.Columns; j++ {
			strings := l.GetStrings([2]int{i,j}, length)
			for _, s := range strings {
				if s == search || s == revSearch {
					answer += 1
				}
			}
		}
	}
	return answer
}

func (l *LetterGrid) CountCrosses(search string) int {
	revSearch := Reverse(search)
	answer := 0
	for i := 1; i < l.Rows-1; i++ {
		for j := 1; j < l.Columns-1; j++ {
			cross := l.GetCross([2]int{i,j})

			firstDiag := fmt.Sprintf(
				"%s%s%s",
				l.GetLetter(cross.UpLeft),
				l.GetLetter(cross.Center),
				l.GetLetter(cross.DownRight))

			secondDiag := fmt.Sprintf(
				"%s%s%s",
				l.GetLetter(cross.DownLeft),
				l.GetLetter(cross.Center),
				l.GetLetter(cross.UpRight))
			
			if (firstDiag == search || firstDiag == revSearch) &&
					(secondDiag == search || secondDiag == revSearch) {

				answer += 1
			}
		}
	}
	return answer
}