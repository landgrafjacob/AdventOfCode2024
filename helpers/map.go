package helpers

import "fmt"

type Map struct {
	Rows int
	Columns int
	Position [2]int
	Facing [2]int
	Obstacles map[[2]int]bool
	Visited map[[2]int]bool
	InsideGrid bool
}

func (m *Map) InitMap(lines []string) {
	m.Rows = len(lines)
	m.Columns = len(lines[0])
	m.Obstacles = make(map[[2]int]bool)
	m.Visited = make(map[[2]int]bool)

	for i, line := range lines {
		for j, char := range line {
			if string(char) == "#" {
				m.Obstacles[[2]int{i,j}] = true
			} else if string(char) == "^" {
				m.Position = [2]int{i,j}
				fmt.Println("Starting at:", m.Position)
				m.Facing = [2]int{-1,0}
				m.Visited[[2]int{i,j}] = true
				m.InsideGrid = true
			}
		}
	}
	fmt.Println("Obstacles:", m.Obstacles)
}

func (m *Map) Step() {
	newPos := [2]int{
		m.Position[0] + m.Facing[0],
		m.Position[1] + m.Facing[1],
	}

	if m.Obstacles[newPos] {
		fmt.Println("Hit obstacle at", newPos)
		m.Rotate()
		fmt.Println("Now facing", m.Facing)
		return
	} else {
		fmt.Println("Stepping to", newPos)
		m.Position = newPos
	}

	m.UpdateInsideGrid()

	if m.InsideGrid {
		m.Visited[newPos] = true
	}
}

func (m *Map) Rotate() {
	if m.Facing == [2]int{-1,0} {
		m.Facing = [2]int{0,1}
	} else if m.Facing == [2]int{0,1} {
		m.Facing = [2]int{1,0}
	} else if m.Facing == [2]int{1,0} {
		m.Facing = [2]int{0,-1}
	} else if m.Facing == [2]int{0,-1} {
		m.Facing = [2]int{-1,0}
	}
}

func (m *Map) UpdateInsideGrid() {
	row := m.Position[0]
	column := m.Position[1]

	if row < 0 || row >= m.Rows {
		m.InsideGrid = false
	} else if column < 0 || column >= m.Columns {
		m.InsideGrid = false
	} else{
		m.InsideGrid = true
	}
}