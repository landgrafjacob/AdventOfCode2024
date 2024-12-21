package helpers

type Set struct {
	setMap map[[2]int]bool
}

func GetEmptySet() *Set {
	return &Set{
		setMap: make(map[[2]int]bool),
	}
}

func (s *Set) Add(element [2]int) {
	s.setMap[element] = true
}

func (s *Set) Remove(element [2]int) {
	delete(s.setMap, element)
}

func (s *Set) Union(s2 *Set) *Set {
	newMap := make(map[[2]int]bool)

	for key, val := range s.setMap {
		newMap[key] = val
	}

	for key, val := range s2.setMap {
		newMap[key] = val
	}

	return &Set{
		setMap: newMap,
	}
}

func (s *Set) IsIn(element [2]int) bool {
	_, ok := s.setMap[element]
	return ok
}

func (s *Set) Length() int {
	return len(s.setMap)
}

func (s *Set) GetElements() [][2]int {
	a := [][2]int{}
	for key := range s.setMap {
		a = append(a, key)
	}
	return a
}
