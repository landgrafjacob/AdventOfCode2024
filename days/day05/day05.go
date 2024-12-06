package day05

import (
	"strings"
	"strconv"
	"slices"
	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

type Day5 struct {}

func (d *Day5) Part1(fileName string) int {
	sections := helpers.GetLineSections("days/day05", fileName)
	rules := generateRules(sections[0])

	answer := 0
	for _, update := range sections[1] {
		pages := strings.Split(update, ",")
		if isOrdered(pages, rules) {
			answer += getCenter(pages)
		}
	}
	return answer
}

func (d *Day5) Part2(fileName string) int {
	sections := helpers.GetLineSections("days/day05", fileName)
	rules := generateRules(sections[0])
	o := helpers.Ordering{}
	o.InitOrdering(rules)

	answer := 0
	for _, update := range sections[1] {
		pages := strings.Split(update, ",")
		if !isOrdered(pages, rules) {
			slices.SortFunc(pages, o.GetFunction())
			answer += getCenter(pages)
		}
	}
	return answer
}

func generateRules(ruleList []string) [][]string {
	answer := [][]string{}

	for _, ruleString := range ruleList {
		splitRule := strings.Split(ruleString, "|")
		rule := []string{}
		for _, page := range splitRule {
			rule = append(rule, page)
		}
		answer = append(answer, rule)
	}
	return answer
}

func buildMap(pageList []string) map[string]int {
	pageMap := make(map[string]int)

	for index, page := range pageList {
		pageMap[page] = index
	}
	return pageMap
}

func followsRule(pageMap map[string]int, rule []string) bool {
	for i:=0; i<len(rule)-1; i++ {
		if _, ok := pageMap[rule[i]]; !ok {
			continue
		}
		if _, ok := pageMap[rule[i+1]]; !ok {
			continue
		}
		if pageMap[rule[i]] > pageMap[rule[i+1]] {
			return false
		}
	}
	return true
}

func isOrdered(pages []string, rules [][]string) bool {
	pageMap := buildMap(pages)

	for _, rule := range rules {
		if !followsRule(pageMap,rule) {
			return false
		}
	}
	return true
}

func getCenter(update []string) int {
	centerInd := len(update) / 2
	center, _ := strconv.Atoi(update[centerInd])
	return center
}
