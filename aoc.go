package main

import (
	"flag"
	"fmt"

	"github.com/wang429/aoc2023/day1"
	"github.com/wang429/aoc2023/util"
)

func main() {
	dayNumPtr := flag.Int("day", 0, "which day to run")
	partPtr := flag.Int("part", 0, "which part for the day to run. 0 for both parts")
	quietPtr := flag.Bool("q", true, "Toggle running quietly")
	inputPathPtr := flag.String("path", "", "relative path of the input file")

	flag.Parse()
	util.SetQuiet(*quietPtr)
	if *quietPtr {
		fmt.Println("Running day", *dayNumPtr, ", quietly")
	} else {
		fmt.Println("Running day", *dayNumPtr, ", loudly")
	}

	switch *dayNumPtr {
	case 1:
		day1.Run(*partPtr, *inputPathPtr)
	}
}
