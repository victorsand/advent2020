package main

import (
	"advent2020/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func part1() int {
	inputLines := util.ReadFile("day2/input.txt")
	numValid := 0
	for _, line := range inputLines {
		integersRe := regexp.MustCompile("[0-9]+")
		integers := integersRe.FindAllString(line, -1)
		minFreq, _ := strconv.ParseInt(integers[0], 10, 32)
		maxFreq, _ := strconv.ParseInt(integers[1], 10, 32)

		characterRe := regexp.MustCompile("[a-z]")
		character := characterRe.FindString(line)

		password := strings.Split(line, ": ")[1]

		freq := int64(strings.Count(password, character))

		if freq >= minFreq && freq <= maxFreq {
			numValid++
		}
	}
	return numValid
}

func part2() int {
	inputLines := util.ReadFile("day2/input.txt")
	numValid := 0
	for _, line := range inputLines {
		integersRe := regexp.MustCompile("[0-9]+")
		integers := integersRe.FindAllString(line, -1)
		pos1, _ := strconv.ParseInt(integers[0], 10, 32)
		pos2, _ := strconv.ParseInt(integers[1], 10, 32)

		characterRe := regexp.MustCompile("[a-z]")
		character := characterRe.FindString(line)

		password := strings.Split(line, ": ")[1]

		numMatches := 0
		if string(password[pos1-1]) == character {
			numMatches++
		}
		if string(password[pos2-1]) == character {
			numMatches++
		}

		if numMatches == 1 {
			numValid++
		}
	}
	return numValid
}

// 586
// 352
func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
