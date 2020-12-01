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

func part2() int {
	var inputLines = util.ReadFile("day1/input.txt")
	var i, j, k int

	// just brute force it, as a star
	for i=0; i<len(inputLines); i++ {
		for j=i+1; j<len(inputLines); j++ {
			for k=j+1; k<len(inputLines); k++ {
				var x, y, z int64
				var err error
				x, err = strconv.ParseInt(inputLines[i], 0, 32)
				if (err != nil) {
					fmt.Println("", err)
				}
				y, err = strconv.ParseInt(inputLines[j], 0, 32)
				if (err != nil) {
					fmt.Println("", err)
				}
				z, err = strconv.ParseInt(inputLines[k], 0, 32)
				if (err != nil) {
					fmt.Println("", err)
				}
				if (x + y + z == 2020) {
					return int(x * y * z)
				}
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
