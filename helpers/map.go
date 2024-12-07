package helpers

type Map struct {
	Rows       int
	Columns    int
	Position   [2]int
	Facing     [2]int
	Obstacles  map[[2]int]bool
	Visited    VisitedCoords
	InsideGrid bool
}

// Struct to store the coordinates that have been visited, and the direction faced at the coordinate
type VisitedCoords struct {
	Coords map[[2]int]map[[2]int]bool
}

// Add coord/facing direction to map
func (v *VisitedCoords) Add(coord, facing [2]int) {
	if v.Coords == nil {
		v.Coords = map[[2]int]map[[2]int]bool{
			coord: map[[2]int]bool{
				facing: true,
			},
		}
	} else if v.Coords[coord] == nil {
		v.Coords[coord] = map[[2]int]bool{
			facing: true,
		}
	} else {
		v.Coords[coord][facing] = true
	}
}

// Get number of coords visited
func (v *VisitedCoords) GetLen() int {
	return len(v.Coords)
}

// Check if given coord/direction has been seen already
func (v *VisitedCoords) HaveVisited(coord, facing [2]int) bool {
	if v.Coords[coord] == nil {
		return false
	}
	return v.Coords[coord][facing]
}

// Initialize map from slice of strings (puzzle input)
func (m *Map) InitMap(lines []string) {
	m.Rows = len(lines)
	m.Columns = len(lines[0])
	m.Obstacles = make(map[[2]int]bool)
	m.Visited = VisitedCoords{}

	for i, line := range lines {
		for j, char := range line {
			if string(char) == "#" {
				m.Obstacles[[2]int{i, j}] = true
			} else if string(char) == "^" {
				// If char is "^", starting position/facing direction
				m.Position = [2]int{i, j}
				m.Facing = [2]int{-1, 0}
				m.InsideGrid = true
				m.Visited.Add([2]int{i, j}, m.Facing)
			}
		}
	}
}

// Keep taking steps until either:
// - You step out of the grid (return true)
// - Hit a loop (return false)
func (m *Map) Run() bool {
	for m.InsideGrid {
		if !m.Step() {
			return false
		}
	}
	return true
}

// Take a step if the spot ahead is free of obstacles, or turn if not
// Returns:
// - false if the resulting position/direction has been seen before (resulting in a loop)
// - true otherwise
func (m *Map) Step() bool {
	// Find new position
	newPos := [2]int{
		m.Position[0] + m.Facing[0],
		m.Position[1] + m.Facing[1],
	}

	// If new position contains an obstacle, rotate instead.
	// Otherwise, take step
	if m.Obstacles[newPos] {
		m.Rotate()
		m.Visited.Add(m.Position, m.Facing)
		return true
	} else {
		m.Position = newPos
	}

	// If we land in a spot/direction we have seen before, we are in a loop
	if m.Visited.HaveVisited(m.Position, m.Facing) {
		return false
	}

	// Check if we are still in the grid, and record the new position if so
	m.UpdateInsideGrid()

	if m.InsideGrid {
		m.Visited.Add(newPos, m.Facing)
	}
	return true
}

// Rotate 90 degrees clockwise
func (m *Map) Rotate() {
	if m.Facing == [2]int{-1, 0} {
		m.Facing = [2]int{0, 1}
	} else if m.Facing == [2]int{0, 1} {
		m.Facing = [2]int{1, 0}
	} else if m.Facing == [2]int{1, 0} {
		m.Facing = [2]int{0, -1}
	} else if m.Facing == [2]int{0, -1} {
		m.Facing = [2]int{-1, 0}
	}
}

// Check if current position is inside the grid
func (m *Map) UpdateInsideGrid() {
	row := m.Position[0]
	column := m.Position[1]

	if row < 0 || row >= m.Rows {
		m.InsideGrid = false
	} else if column < 0 || column >= m.Columns {
		m.InsideGrid = false
	} else {
		m.InsideGrid = true
	}
}

// Return the number of positions seen (not counting directions)
func (m *Map) NumVisited() int {
	answer := 0
	for range m.Visited.Coords {
		answer += 1
	}
	return answer
}
