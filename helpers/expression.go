package helpers

import (
	"strconv"
	"strings"
)

type Expression struct {
	Nums   []int
	Target int
	Part   int
}

// Load line into Expression
// Input:
// {target}: fields[0] fields[1] fields[2] ...
func (e *Expression) Load(s string) {
	fields := strings.Fields(s)
	e.Target, _ = strconv.Atoi(strings.Trim(fields[0], ":"))
	e.Nums = []int{}
	for _, numStr := range fields[1:] {
		numInt, _ := strconv.Atoi(numStr)
		e.Nums = append(e.Nums, numInt)
	}
}

// Return true if the target can be achieved by inserting operations
// addition/multiplication/concatenation(for part 2) into the list of nums
func (e *Expression) Evaluate() bool {
	if len(e.Nums) == 1 {
		return e.Nums[0] == e.Target
	}

	// Idea: pop off last number and build a new target for the expression to evaluate to
	last := e.Nums[len(e.Nums)-1]
	newNums := e.Nums[:len(e.Nums)-1]

	// The new expression if the last operation is addition
	addExp := &Expression{
		Nums:   newNums,
		Target: e.Target - last,
		Part:   e.Part,
	}

	// The new expression if the last operation is multiplication
	multExp := &Expression{
		Nums:   newNums,
		Target: e.Target / last,
		Part:   e.Part,
	}

	// The new expression if the last operation is concatenation
	powTen := 10
	for powTen < last {
		powTen *= 10
	}
	concatExp := &Expression{
		Nums:   newNums,
		Target: e.Target / powTen,
		Part:   e.Part,
	}

	// Evaluate these new expressions if valid
	if e.Target%last == 0 && multExp.Evaluate() {
		// Target must be divisible by the last entry for multExp to be valid
		return true
	} else if e.Part == 2 && e.Target%powTen == last && concatExp.Evaluate() {
		// Target must end with the last entry for concatExp to be valid
		return true
	} else {
		// addExp is always valid
		return addExp.Evaluate()
	}
}
