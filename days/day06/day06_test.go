package day06

import (
	"testing"

	"github.com/landgrafjacob/AdventOfCode2024/helpers"
)

var testAnswers helpers.TestAnswers

func init() {
	testAnswers.Part1 = 41
	testAnswers.Part2 = 6
}

func TestPart1(t *testing.T) {
	d := &Day6{}
	got := d.Part1("test.txt")
	if got != testAnswers.Part1 {
		t.Errorf("Expected: %d, Got: %d", testAnswers.Part1, got)
	}

}

func TestPart2(t *testing.T) {
	d := &Day6{}
	got := d.Part2("test.txt")
	if got != testAnswers.Part2 {
		t.Errorf("Expected: %d, Got: %d", testAnswers.Part2, got)
	}

}
