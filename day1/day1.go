package day1

import (
	"fmt"

	"github.com/wang429/aoc2023/util"
)

func Run() {
	fmt.Println("Running Part 1")
	Part1()
	fmt.Println("Running Part 2")
}

func Part1() {
	content := util.ReadFile("day1/part1.txt")
	fmt.Printf("File contents: \n%s", content)
}
