package days

import (
	"fmt"

	"github.com/landgrafjacob/AdventOfCode2024/days/day01"
	"github.com/landgrafjacob/AdventOfCode2024/days/day02"
	"github.com/landgrafjacob/AdventOfCode2024/days/day03"
	"github.com/landgrafjacob/AdventOfCode2024/days/day04"
	"github.com/landgrafjacob/AdventOfCode2024/days/day05"
	"github.com/landgrafjacob/AdventOfCode2024/days/day06"
	"github.com/landgrafjacob/AdventOfCode2024/days/day07"
	"github.com/landgrafjacob/AdventOfCode2024/days/day08"
	"github.com/landgrafjacob/AdventOfCode2024/days/day09"
)

type DayInterface interface {
	Part1(string) int
	Part2(string) int
}

var dayMap map[string]DayInterface

func init() {
	dayMap = map[string]DayInterface{
		"1":  &day01.Day{},
		"2":  &day02.Day{},
		"3":  &day03.Day{},
		"4":  &day04.Day{},
		"5":  &day05.Day{},
		"6":  &day06.Day{},
		"7":  &day07.Day{},
		"8":  &day08.Day{},
		"9":  &day09.Day{},
		"10": &day01.Day{},
		"11": &day01.Day{},
		"12": &day01.Day{},
		"13": &day01.Day{},
		"14": &day01.Day{},
		"15": &day01.Day{},
		"16": &day01.Day{},
		"17": &day01.Day{},
		"18": &day01.Day{},
		"19": &day01.Day{},
		"20": &day01.Day{},
		"21": &day01.Day{},
		"22": &day01.Day{},
		"23": &day01.Day{},
		"24": &day01.Day{},
		"25": &day01.Day{},
	}
}

func GetDay(day string) DayInterface {
	return dayMap[day]
}

func Execute(d DayInterface) {
	fmt.Printf("Part 1: %d\n", d.Part1("input.txt"))
	fmt.Printf("Part 2: %d\n", d.Part2("input.txt"))
}
