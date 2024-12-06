package helpers

type Ordering struct {
	OrderingMap map[string]map[string]int
}

func (o *Ordering) InitOrdering (rules [][]string) {
	o.OrderingMap = make(map[string]map[string]int)
	for _, rules := range rules {
		if o.OrderingMap[rules[0]] == nil {
			o.OrderingMap[rules[0]] = map[string]int{rules[1]: -1}
		} else {
			o.OrderingMap[rules[0]][rules[1]] = -1
		}
		if o.OrderingMap[rules[1]] == nil {
			o.OrderingMap[rules[1]] = map[string]int{rules[0]: 1}
		} else {
			o.OrderingMap[rules[1]][rules[0]] = 1
		}
	}
}

func (o *Ordering) GetFunction() func(string,string) int {
	f := func(a string, b string) int {
		return o.OrderingMap[a][b]
	}
	return f
}