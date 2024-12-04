package day03

import (
	"strconv"
	"regexp"
	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day3 struct {}

// Given a string of the form `mul({first_num},{second_sum})`,
// returns the product first_num * second_num as an int
func parse(s string) int {
	r, _ := regexp.Compile(`\d+`)
	nums := r.FindAllString(s, -1)
	product := 1
	for _, n := range nums {
		nInt, _ := strconv.Atoi(n)
		product *= nInt
	}
	return product
}

func (d *Day3) Part1(fileName string) int {
	lines := helpers.GetLines("days/day03", fileName)

	// Find all substrings of the form `mul({first_num},{second_sum})`
	r, _ := regexp.Compile(`mul\(\d+,\d+\)`)

	answer := 0
	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		for _, m := range matches {
			answer += parse(m)
		}
	}
	return answer
}

func (d *Day3) Part2(fileName string) int {
	lines := helpers.GetLines("days/day03", fileName)

	// Find all substrings of the form `mul({first_num},{second_sum})`,
	// as well as all do()'s and don't()'s
	r, _ := regexp.Compile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)

	// The `on` variable will keep track of when the last seen command
	// was `do()` (true) or `don't()` (false)
	on := true
	answer := 0
	for _, line := range lines {
		matches := r.FindAllStringIndex(line, -1)
		for _, m := range matches {
			if line[m[0]:m[1]] == "do()" {
				on = true
			} else if line[m[0]:m[1]] == "don't()" {
				on = false
			} else if on {
				answer += parse(line[m[0]:m[1]])
			}
		}
	}

	return answer
}
