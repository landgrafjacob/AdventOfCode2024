package day02

import (
	"github.com/landgrafjacob/AdventOfCode2024/helpers"
	"testing"
)

var testAnswers helpers.TestAnswers

func init() {
	testAnswers.Part1 = 2
	testAnswers.Part2 = 4
}

func TestPart1(t *testing.T) {
	d := &Day2{}
	got := d.Part1("test.txt")
	if got != testAnswers.Part1 {
		t.Errorf("Expected: %d, Got: %d", testAnswers.Part1, got)
	}

}

func TestPart2(t *testing.T) {
	d := &Day2{}
	got := d.Part2("test.txt")
	if got != testAnswers.Part2 {
		t.Errorf("Expected: %d, Got: %d", testAnswers.Part2, got)
	}

}

func TestIsSafe(t *testing.T) {
	tests := []struct{
		list []string
		expectedValue bool
	}{
		{
			list: []string{"7","6","4","2","1"},
			expectedValue: true,
		},
		{
			list: []string{"1","2","7","8","9"},
			expectedValue: false,
		},
		{
			list: []string{"1","3","6","7","9"},
			expectedValue: true,
		},
	}

	for _, test := range tests {
		value := isSafe(test.list)

		if value != test.expectedValue {
			t.Errorf("Expected: %t, Got: %t\n", test.expectedValue, value)
		}
	}
}