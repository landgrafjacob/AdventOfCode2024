package day04

import (
	"github.com/landgrafjacob/AdventOfCode2024/helpers"
	"testing"
)

var testAnswers helpers.TestAnswers

func init() {
	testAnswers.Part1 = 18
	testAnswers.Part2 = 9
}

func TestPart1(t *testing.T) {
	d := &Day4{}
	got := d.Part1("test.txt")
	if got != testAnswers.Part1 {
		t.Errorf("Expected: %d, Got: %d", testAnswers.Part1, got)
	}

}

func TestPart2(t *testing.T) {
	d := &Day4{}
	got := d.Part2("test.txt")
	if got != testAnswers.Part2 {
		t.Errorf("Expected: %d, Got: %d", testAnswers.Part2, got)
	}

}