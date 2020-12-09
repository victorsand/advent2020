package main

import (
	"advent2020/util"
	"fmt"
	"strconv"
)

func findPair(preambleStart int, preambleSize int, index int, values []int) (int, int) {
	for i := preambleStart; i < preambleStart+preambleSize; i++ {
		for j := i; j < preambleStart+preambleSize; j++ {
			if i == j {
				continue
			}
			if values[i]+values[j] == values[index] {
				return values[i], values[j]
			}
		}
	}
	return -1, -1
}

func main() {
	inputLines := util.ReadFile("day9/input.txt")
	var values []int
	for _, line := range inputLines {
		value, _ := strconv.ParseInt(line, 10, 32)
		values = append(values, int(value))
	}

	preambleSize := 25
	fmt.Println(len(values))
	for i := preambleSize; i < len(values); i++ {
		x, _ := findPair(i-preambleSize, preambleSize, i, values)
		if x == -1 {
			fmt.Println("Weakness detected!", values[i])
			break
		}
	}
}
