package main

import (
	"advent2020/util"
	"fmt"
)

type grid struct {
	Rows []string
}

func buildGrid(inputLines []string) grid {
	gridSizeX := 0
	gridSizeY := 0
	var areaGrid grid
	for _, line := range inputLines {
		var areaRow string
		for pos, char := range line {
			areaRow += string(char)
			if pos+1 > gridSizeX {
				gridSizeX = pos + 1
			}
		}
		areaGrid.Rows = append(areaGrid.Rows, areaRow)
		gridSizeY++
	}
	fmt.Printf("Grid size: %d x %d\n", gridSizeX, gridSizeY)
	return areaGrid
}

func countSlopeTrees(g grid, xStep int, yStep int) int {
	numTrees := 0
	xPos, yPos := 0, 0
	numRows := len(g.Rows)
	for {
		currentRow := g.Rows[yPos]
		rowWidth := len(currentRow)
		xPos = (xPos + xStep) % rowWidth
		yPos = yPos + yStep
		if yPos >= numRows {
			break
		}
		if g.Rows[yPos][xPos] == '#' {
			numTrees++
		}
	}
	return numTrees
}

func main() {
	var inputLines = util.ReadFile("day3/input.txt")
	areaGrid := buildGrid(inputLines)

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
