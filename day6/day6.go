package main

import (
	"advent2020/util"
	"fmt"
	"strings"
)

type fn func(x, y string) string

func mergeStringsUnion(x, y string) string {
	var union = x
	for _, r := range y {
		if !strings.Contains(union, string(r)) {
			union += string(r)
		}
	}
	return union
}

func mergeStringsIntersection(x, y string) string {
	intersection := ""
	for _, r := range y {
		if !strings.Contains(intersection, string(r)) && strings.Contains(x, string(r)) {
			intersection += string(r)
		}
	}
	return intersection
}

func buildGroups(inputLines []string, mergeFunction fn) []string {
	var groups []string
	group := ""
	first := true
	for _, line := range inputLines {
		if first {
			group = line
		}
		if line != "" && line != "/n" {
			group = mergeFunction(group, line)
			first = false
		} else {
			groups = append(groups, group)
			group = ""
			first = true
		}
	}
	groups = append(groups, group)
	return groups
}

func main() {
	//test()

	var inputLines = util.ReadFile("day6/input.txt")

	var groupsUnion = buildGroups(inputLines, mergeStringsUnion)
	sumUnion := 0
	for _, group := range groupsUnion {
		sumUnion += len(group)
	}

	var groupsIntersection = buildGroups(inputLines, mergeStringsIntersection)
	sumIntersection := 0
	for _, group := range groupsIntersection {
		sumIntersection += len(group)
	}

	fmt.Println("Part 1:", sumUnion)
	fmt.Println("Part 2:", sumIntersection)

}

func test() {
	util.AssertTrue(mergeStringsUnion("a", "ab") == "ab")
	util.AssertTrue(mergeStringsUnion("a", "abc") == "abc")
	util.AssertTrue(mergeStringsUnion("ab", "abc") == "abc")
	util.AssertTrue(mergeStringsUnion("a", "b") == "ab")
	util.AssertTrue(mergeStringsUnion("", "a") == "a")

	util.AssertTrue(mergeStringsIntersection("a", "ab") == "a")
	util.AssertTrue(mergeStringsIntersection("a", "abc") == "a")
	util.AssertTrue(mergeStringsIntersection("ab", "abc") == "ab")
	util.AssertTrue(mergeStringsIntersection("a", "b") == "")
	util.AssertTrue(mergeStringsIntersection("", "a") == "")
}
