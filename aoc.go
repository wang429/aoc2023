package main

import (
	"flag"
	"fmt"

	"github.com/wang429/aoc2023/day1"
)

func main() {
	dayNumPtr := flag.Int("day", 0, "which day to run")
	flag.Parse()
	fmt.Println("Running day", *dayNumPtr)

	switch *dayNumPtr {
	case 1:
		day1.Run()
	}
}
