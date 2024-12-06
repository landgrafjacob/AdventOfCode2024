package days

import (
	"fmt"
	"github.com/landgrafjacob/AdventOfCode2024/days/day01"
	"github.com/landgrafjacob/AdventOfCode2024/days/day02"
	"github.com/landgrafjacob/AdventOfCode2024/days/day03"
	"github.com/landgrafjacob/AdventOfCode2024/days/day04"
	"github.com/landgrafjacob/AdventOfCode2024/days/day05"
)

type Day interface {
	Part1(string) int
	Part2(string) int
}

var dayMap map[string]Day
func init() {
	dayMap = map[string]Day {
		"1": &day01.Day1{},
		"2": &day02.Day2{},
		"3": &day03.Day3{},
		"4": &day04.Day4{},
		"5": &day05.Day5{},
		"6": &day01.Day1{},
		"7": &day01.Day1{},
		"8": &day01.Day1{},
		"9": &day01.Day1{},
		"10": &day01.Day1{},
		"11": &day01.Day1{},
		"12": &day01.Day1{},
		"13": &day01.Day1{},
		"14": &day01.Day1{},
		"15": &day01.Day1{},
		"16": &day01.Day1{},
		"17": &day01.Day1{},
		"18": &day01.Day1{},
		"19": &day01.Day1{},
		"20": &day01.Day1{},
		"21": &day01.Day1{},
		"22": &day01.Day1{},
		"23": &day01.Day1{},
		"24": &day01.Day1{},
		"25": &day01.Day1{},
	}
}

func GetDay(day string) Day {
	return dayMap[day]
}

func Execute(d Day) {
	fmt.Printf("Part 1: %d\n", d.Part1("input.txt"))
	fmt.Printf("Part 2: %d\n", d.Part2("input.txt"))
}
