package main

import (
	"fmt" 
	"advent2020/util"
	"strconv"
)

func part1() int {
	var inputLines = util.ReadFile("day1/input.txt")
	var i, j int
	for i=0; i<len(inputLines); i++ {
		for j=i+1; j<len(inputLines); j++ {
			x, err := strconv.ParseInt(inputLines[i], 0, 32)
			if (err != nil) {
				fmt.Println("", err)
			}
			y, err := strconv.ParseInt(inputLines[j], 0, 32)
			if x + y == 2020 {
				return int(x * y)
			}
			if (err != nil) {
				fmt.Println("", err)
			}
		}
	}
	return 0
}

type sumOfTwo struct {
	a int
	b int
	sum int
}

func part2() int {
	var inputLines = util.ReadFile("day1/input.txt")
	var sumsOfTwo []sumOfTwo
	var i, j int
	for i=0; i<len(inputLines); i++ {
		for j=i+1; j<len(inputLines); j++ {
			var x, y int64
			var err error
			x, err = strconv.ParseInt(inputLines[i], 0, 32)
			y, err = strconv.ParseInt(inputLines[j], 0, 32)
			if (err != nil) {
				fmt.Println("", err)
			}
			sumsOfTwo = append(sumsOfTwo, sumOfTwo{int(x), int(y), int(x + y)})
		}
	}
	for i=0; i<len(sumsOfTwo); i++ {
		for j=0; j<len(inputLines); j++ {
			x, err := strconv.ParseInt(inputLines[j], 0, 32)
			if (err != nil) {
				fmt.Println("", err)
			}
			if (sumsOfTwo[i].sum + int(x) == 2020) {
				return sumsOfTwo[i].a * sumsOfTwo[i].b * int(x)
			}
		}
	}
	return 0
}

// 927684
// 292093004
func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
