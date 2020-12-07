package main

import (
	"advent2020/util"
	"fmt"
	"strconv"
	"strings"
)

type edge struct {
	weight int64
	target string
}

func buildAdjListBackwards(inputLines []string) map[string][]edge {
	result := make(map[string][]edge)
	for _, line := range inputLines {
		if strings.Contains(line, "no other bags") {
			continue
		}
		split := strings.Split(line, " ")
		rootColor := split[0] + split[1]
		bagsDef := strings.Join(split[4:len(split)-1], " ")
		bagsDefStrings := strings.Split(bagsDef, ",")
		for _, bag := range bagsDefStrings {
			bagDefSplit := strings.Split(strings.Trim(bag, " "), " ")
			number, _ := strconv.ParseInt(bagDefSplit[0], 10, 64)
			color := bagDefSplit[1] + bagDefSplit[2]
			bagEdge := edge{weight: number, target: rootColor}
			if result[color] != nil {
				result[color] = append(result[color], bagEdge)
			} else {
				result[color] = []edge{bagEdge}
			}
		}

	}
	return result
}

func main() {
	var inputLines = util.ReadFile("day7/input.txt")

	// build the adjacency list backwards, for traversing from the innermost bag and outwards
	adjListBackwards := buildAdjListBackwards(inputLines)
	fmt.Println("Part 1", countAdjacentNodes(adjListBackwards, "shinygold"))

}

func countAdjacentNodes(list map[string][]edge, startNode string) int {
	var queue []string
	var visited []string
	count := 0
	var start = list[startNode]
	queue = append(queue, edgeSliceToTargetStringSlice(start)...)
	for len(queue) > 0 {
		node := queue[0]
		visited = append(visited, node)
		count++
		toVisit := edgeSliceToTargetStringSlice(list[node])
		for _, edge := range toVisit {
			if !containsString(queue, edge) && !containsString(visited, edge) {
				queue = append(queue, edge)
			}
		}
		queue = queue[1:]
	}
	return count
}

func edgeSliceToTargetStringSlice(source []edge) []string {
	var result []string
	for _, e := range source {
		result = append(result, e.target)
	}
	return result
}

func containsString(target []string, item string) bool {
	existsInQueue := false
	for _, itemInTarget := range target {
		if itemInTarget == item {
			existsInQueue = true
			break
		}
	}
	return existsInQueue
}
