package day13

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/landgrafjacob/AdventOfCode2024/helpers"
	"gonum.org/v1/gonum/mat"
)

// Error for float to be considered an integer
const epsilon = 1e-3

// Cost for button pushes
const costA = 3
const costB = 1

type Day struct{}

func (d *Day) Part1(fileName string) int {
	lines := helpers.GetLineSections("days/day13", fileName)
	ms := GetMachines(lines, false)
	return int(ms.Solve())
}

func (d *Day) Part2(fileName string) int {
	lines := helpers.GetLineSections("days/day13", fileName)
	ms := GetMachines(lines, true)
	return int(ms.Solve())
}

type Machines struct {
	MachineSlice []Machine
}

type Machine struct {
	ButtonA []float64
	ButtonB []float64
	Prize   []float64
}

func (m *Machine) Solve() float64 {
	costSlice := []float64{costA, costB}

	// Set up matrix equation to solve system:
	// (# A pushes) * (A x-shift) + (# B pushes) * (B x-shift) = x-val of prize
	// (# A pushes) * (A y-shift) + (# B pushes) * (B y-shift) = y-val of prize
	A := mat.NewDense(2, 2, []float64{
		m.ButtonA[0],
		m.ButtonB[0],
		m.ButtonA[1],
		m.ButtonB[1],
	})
	b := mat.NewVecDense(2, []float64{m.Prize[0], m.Prize[1]})

	// Solve the equation for the number of A and B pushes required
	// Note: the problem is tricky by insinuating that there are multiple
	// combinations of button pushes to get to the prize
	// However, there's only one (2 equations with 2 unknowns)
	var x mat.VecDense
	if err := x.SolveVec(A, b); err != nil {
		fmt.Println(err)
	}
	answer := float64(0)
	for i, val := range x.RawVector().Data {
		// Check that the solutions are both positive (negative button pushes doesn't make sense)
		if val < 0 {
			return 0
		}
		// Check that the solutions are both integers
		if _, frac := math.Modf(math.Abs(val)); frac < epsilon || frac > 1.0-epsilon {
			answer += costSlice[i] * math.Round(val)
		} else {
			return 0
		}
	}
	return answer
}

func (ms *Machines) Solve() float64 {
	answer := float64(0)
	for _, machine := range ms.MachineSlice {
		answer += machine.Solve()
	}
	return answer
}

func GetMachines(lines [][]string, scalePrize bool) Machines {
	r, _ := regexp.Compile("[0-9]+")
	ms := Machines{[]Machine{}}
	for _, block := range lines {
		m := Machine{}
		m.ButtonA = ToFloat(r.FindAllString(block[0], -1))
		m.ButtonB = ToFloat(r.FindAllString(block[1], -1))
		m.Prize = ToFloat(r.FindAllString(block[2], -1))

		if scalePrize {
			for i := range m.Prize {
				m.Prize[i] += 10000000000000
			}
		}
		ms.MachineSlice = append(ms.MachineSlice, m)
	}
	return ms
}

func ToFloat(s []string) []float64 {
	a := []float64{}
	for _, val := range s {
		valInt, _ := strconv.Atoi(val)
		a = append(a, float64(valInt))
	}
	return a
}
