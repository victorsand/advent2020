package main

import (
	"advent2020/util"
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	value     int64
	hasRun    bool
}

const noOperation = "nop"
const accumulate = "acc"
const jump = "jmp"

func parseInstructions(lines []string) []instruction {
	var instructions []instruction
	for _, line := range lines {
		var split = strings.Split(line, " ")
		val, _ := strconv.ParseInt(split[1], 10, 64)
		instr := instruction{operation: split[0], value: val}
		instructions = append(instructions, instr)
	}
	return instructions
}

func main() {
	inputLines := util.ReadFile("day8/input.txt")
	instructions := parseInstructions(inputLines)

	accumulator := int64(0)
	next := 0
	for next < len(instructions) {
		current := next
		instruction := instructions[next]
		if instruction.hasRun {
			fmt.Println("Loop detected! Accumulator value is ", accumulator)
			break
		}
		switch instruction.operation {
		case accumulate:
			accumulator += instruction.value
			next++
		case jump:
			next += int(instruction.value)
		case noOperation:
			next++
		default:
			panic("Unknown operation" + instruction.operation)

		}
		instructions[current].hasRun = true
	}

}
