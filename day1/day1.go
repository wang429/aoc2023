package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wang429/aoc2023/util"
)

func Run(part int, inputPath string) {
	if part == 0 || part == 1 {
		fmt.Println("Running Part 1")
		Part1(inputPath)
	}
	if part == 0 || part == 2 {
		fmt.Println("Running Part 2")
		Part2(inputPath)
	}
}

func Part1(inputPath string) {
	if inputPath == "" {
		inputPath = "day1/part1.txt"
	}
	lines := util.ReadFile(inputPath)
	util.PrintLines(lines)
	//code below

	sum := 0
	for _, v := range lines {
		var num []int
		for i := 0; i < len(v); i++ {
			val, err := strconv.Atoi(string(v[i]))
			if err == nil {
				num = append(num, val)
			}
		}
		util.Println(num)
		if len(num) > 0 {
			sum += num[0]*10 + num[len(num)-1]
		}
	}

	fmt.Printf("Output sum for Part1: %d\n", sum)
}

func Part2(inputPath string) {
	if inputPath == "" {
		inputPath = "day1/part2.txt"
	}
	lines := util.ReadFile(inputPath)
	util.PrintLines(lines)
	// code below

	sum := 0
	for _, v := range lines {
		first := FindFirst(v)
		last := FindLast(v)
		fmt.Printf("first: %d, last: %d\n", first, last)
		sum += first*10 + last
	}

	fmt.Printf("Output sum for Part2: %d\n", sum)
}

func FindFirst(line string) int {
	for i := 0; i < len(line); i++ {
		subLine := line[i:]
		if strings.HasPrefix(subLine, "one") {
			return 1
		} else if strings.HasPrefix(subLine, "two") {
			return 2
		} else if strings.HasPrefix(subLine, "three") {
			return 3
		} else if strings.HasPrefix(subLine, "four") {
			return 4
		} else if strings.HasPrefix(subLine, "five") {
			return 5
		} else if strings.HasPrefix(subLine, "six") {
			return 6
		} else if strings.HasPrefix(subLine, "seven") {
			return 7
		} else if strings.HasPrefix(subLine, "eight") {
			return 8
		} else if strings.HasPrefix(subLine, "nine") {
			return 9
		} else {
			v, err := strconv.Atoi(string(line[i]))
			if err == nil && v > 0 {
				return v
			}
		}
	}
	return 0
}

func FindLast(line string) int {
	for i := 0; i < len(line); i++ {
		subLine := line[:len(line)-i]
		if strings.HasSuffix(subLine, "one") {
			return 1
		} else if strings.HasSuffix(subLine, "two") {
			return 2
		} else if strings.HasSuffix(subLine, "three") {
			return 3
		} else if strings.HasSuffix(subLine, "four") {
			return 4
		} else if strings.HasSuffix(subLine, "five") {
			return 5
		} else if strings.HasSuffix(subLine, "six") {
			return 6
		} else if strings.HasSuffix(subLine, "seven") {
			return 7
		} else if strings.HasSuffix(subLine, "eight") {
			return 8
		} else if strings.HasSuffix(subLine, "nine") {
			return 9
		} else {
			v, err := strconv.Atoi(string(line[len(line)-i-1]))
			if err == nil && v > 0 {
				return v
			}
		}
	}
	return 0
}
