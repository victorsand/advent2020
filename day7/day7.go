package main

import (
	"advent2020/util"
	"fmt"
	"strconv"
	"strings"
)

type edge struct {
	weight int
	target string
}

const directionForward = 0
const directionBackward = 1

func buildAdjacencyList(inputLines []string, direction int) map[string][]edge {
	result := make(map[string][]edge)
	for _, line := range inputLines {
		split := strings.Split(line, " ")
		rootColor := split[0] + split[1]
		if strings.Contains(line, "no other bags") {
			if direction == directionForward {
				result[rootColor] = []edge{}
			}
			continue
		}
		bagsDef := strings.Join(split[4:len(split)-1], " ")
		bagsDefStrings := strings.Split(bagsDef, ",")
		for _, bag := range bagsDefStrings {
			bagDefSplit := strings.Split(strings.Trim(bag, " "), " ")
			number, _ := strconv.ParseInt(bagDefSplit[0], 10, 64)
			color := bagDefSplit[1] + bagDefSplit[2]
			var from, to string
			if direction == directionForward {
				from = rootColor
				to = color
			} else if direction == directionBackward {
				from = color
				to = rootColor
			}
			bagEdge := edge{weight: int(number), target: to}
			if result[from] != nil {
				result[from] = append(result[from], bagEdge)
			} else {
				result[from] = []edge{bagEdge}
			}
		}
	}
	return result
}

func main() {
	var inputLines = util.ReadFile("day7/testinput.txt")

	// build the adjacency list backwards, for traversing from the innermost bag and outwards
	adjListBackwards := buildAdjacencyList(inputLines, directionBackward)
	fmt.Println("Part 1", countAdjacentNodes(adjListBackwards, "shinygold"))

	// build the adjacency list forwards, for traversing from the outermost bag and inwards
	adjListForward := buildAdjacencyList(inputLines, directionForward)
	fmt.Println("Part 2", adjListForward)
	//fmt.Println("Part 2", multiplyAdjacentNodeWeights(adjListForward, "shinygold"))
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

// gotta do depth first

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
