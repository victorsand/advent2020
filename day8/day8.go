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

func resetRunFlags(instructions []instruction) {
	for i := range instructions {
		instructions[i].hasRun = false
	}
}

func switchOperation(instructions []instruction, index int) {
	if instructions[index].operation == jump {
		instructions[index].operation = noOperation
	} else if instructions[index].operation == noOperation {
		instructions[index].operation = jump
	}
}

func runAccumulator(instructions []instruction) (int, bool) {
	loopDetected := false
	accumulator := int64(0)
	next := 0
	for next < len(instructions) {
		current := next
		instruction := instructions[next]
		if instruction.hasRun {
			loopDetected = true
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
	return int(accumulator), loopDetected
}

func switchAndRun(instructions []instruction) {
	operationToSwitch := 0
	isBroken := true
	for operationToSwitch < len(instructions) && isBroken {
		resetRunFlags(instructions)
		switchOperation(instructions, operationToSwitch)
		value, loopDetected := runAccumulator(instructions)
		if !loopDetected {
			fmt.Println("FIXED! Accumulator value:", value)
			break
		}
		switchOperation(instructions, operationToSwitch)
		operationToSwitch++
	}
}

func main() {
	inputLines := util.ReadFile("day8/input.txt")
	instructions := parseInstructions(inputLines)
	value, _ := runAccumulator(instructions)
	fmt.Println("Accumulator value:", value)
	switchAndRun(instructions)

}
