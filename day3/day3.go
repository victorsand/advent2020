package main

import (
	"advent2020/util"
	"fmt"
)

func countSlopeTrees(grid []string, xStep int, yStep int) int {
	numTrees := 0
	xPos, yPos := 0, 0
	numRows := len(grid)
	for {
		xPos = (xPos + xStep) % len(grid[yPos])
		yPos = yPos + yStep
		if yPos >= numRows {
			break
		}
		if grid[yPos][xPos] == '#' {
			numTrees++
		}
	}
	return numTrees
}

func main() {
	var areaGrid = util.ReadFile("day3/input.txt")

	// part 1
	// 274
	fmt.Println(countSlopeTrees(areaGrid, 3, 1))

	// part 2
	// 6050183040
	slopeProduct := countSlopeTrees(areaGrid, 1, 1)
	slopeProduct *= countSlopeTrees(areaGrid, 3, 1)
	slopeProduct *= countSlopeTrees(areaGrid, 5, 1)
	slopeProduct *= countSlopeTrees(areaGrid, 7, 1)
	slopeProduct *= countSlopeTrees(areaGrid, 1, 2)
	fmt.Println(slopeProduct)
}
