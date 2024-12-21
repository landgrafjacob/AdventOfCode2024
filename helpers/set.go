package helpers

type Set struct {
	setMap map[any]bool
}

func GetEmptySet() *Set {
	return &Set{
		setMap: make(map[any]bool),
	}
}

func (s *Set) Add(element any) {
	s.setMap[element] = true
}

func (s *Set) Union(s2 *Set) *Set {
	newMap := make(map[any]bool)

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

func (s *Set) IsIn(element any) bool {
	_, ok := s.setMap[element]
	return ok
}

func (s *Set) Length() int {
	return len(s.setMap)
}
