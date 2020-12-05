package main

import (
	"advent2020/util"
	"fmt"
	"math"
)

func test() {
	util.AssertTrue(calculateRow("FBFBBFFRLR") == 44)
	util.AssertTrue(calculateColumn("FBFBBFFRLR") == 5)
}

func lowerHalf(lower, upper, diff int) (int, int, int) {
	upper = upper - int(math.Ceil(float64(diff)/2.0))
	diff = upper - lower
	return lower, upper, diff
}

func upperHalf(lower, upper, diff int) (int, int, int) {
	lower = lower + int(math.Ceil(float64(diff)/2.0))
	diff = upper - lower
	return lower, upper, diff
}

func calculateRow(input string) int {
	lower, upper, diff := 0, 127, 127
	for i := 0; i < 7; i++ {
		char := input[i]
		if char == 'F' {
			lower, upper, diff = lowerHalf(lower, upper, diff)
		} else if char == 'B' {
			lower, upper, diff = upperHalf(lower, upper, diff)
		}
	}
	return lower
}

func calculateColumn(input string) int {
	lower, upper, diff := 0, 7, 7
	for i := 7; i < 10; i++ {
		char := input[i]
		if char == 'L' {
			lower, upper, diff = lowerHalf(lower, upper, diff)
		} else if char == 'R' {
			lower, upper, diff = upperHalf(lower, upper, diff)
		}
	}
	return lower
}

func main() {
	//test()
	var inputLines = util.ReadFile("day5/input.txt")

	maxID := 0
	var IDs map[int]bool
	IDs = make(map[int]bool)
	for _, line := range inputLines {
		row := calculateRow(line)
		column := calculateColumn(line)
		id := row*8 + column
		IDs[id] = true
		if id > maxID {
			maxID = id
		}
	}

	fmt.Println("Max ID:", maxID)

	for i := 0; i < maxID; i++ {
		if !IDs[i] && IDs[i-1] && IDs[i+1] {
			fmt.Println("My ID:", i)
			break
		}
	}
}
