package main

import (
	"flag"
	"github.com/landgrafjacob/AdventOfCode2024/days"
)

func main() {
	dayPtr := flag.String("day", "1", "The number of the day to run")
	flag.Parse()
	day := days.GetDay(*dayPtr)
	days.Execute(day)
	
}
