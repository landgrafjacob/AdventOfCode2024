package days

import (
	"fmt"
)

type Day interface {
	Part1(string) int
	Part2(string) int
}

func Execute(d Day) {
	fmt.Printf("Part 1: %d\n", d.Part1("input.txt"))
	fmt.Printf("Part 2: %d\n", d.Part2("input.txt"))
}
