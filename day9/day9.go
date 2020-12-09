package main

import (
	"advent2020/util"
	"fmt"
	"sort"
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

func findSeries(target int, values []int) []int {
	for i := 0; i < len(values); i++ {
		var series []int
		sum := 0
		for j := i; j < len(values); j++ {
			series = append(series, values[j])
			sum += values[j]
			if sum == target && len(series) > 1 {
				return series
			}
			if sum > target {
				continue
			}
		}
	}
	return nil
}

func main() {
	inputLines := util.ReadFile("day9/input.txt")
	var values []int
	for _, line := range inputLines {
		value, _ := strconv.ParseInt(line, 10, 32)
		values = append(values, int(value))
	}

	preambleSize := 25
	var invalidNumber int
	fmt.Println(len(values))
	for i := preambleSize; i < len(values); i++ {
		x, _ := findPair(i-preambleSize, preambleSize, i, values)
		if x == -1 {
			fmt.Println("Weakness detected!", values[i])
			invalidNumber = values[i]
			break
		}
	}

	weaknessSeries := findSeries(invalidNumber, values)
	sort.Ints(weaknessSeries)
	fmt.Println("Weakness", weaknessSeries[0]+weaknessSeries[len(weaknessSeries)-1])
}
