package day5

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	stacks, moves := parseInput(input)
	for _, move := range moves {
		count, from, to := parseMove(move)
		fromStack := &(stacks)[from]
		toStack := &(stacks)[to]
		for i := 0; i < count; i++ {
			pop := (*fromStack)[len(*fromStack)-1]
			*fromStack = (*fromStack)[:len(*fromStack)-1]
			*toStack = append(*toStack, pop)
		}
	}
	topCrates := make([]string, 0)
	for _, stack := range stacks {
		topCrates = append(topCrates, stack[len(stack)-1])
	}
	return strings.Join(topCrates, "")
}

func Part2(input string) string {
	stacks, moves := parseInput(input)
	for _, move := range moves {
		count, from, to := parseMove(move)
		fromStack := &(stacks)[from]
		toStack := &(stacks)[to]
		pop := (*fromStack)[(len(*fromStack) - count):]
		*fromStack = (*fromStack)[:(len(*fromStack) - count)]
		*toStack = append(*toStack, pop...)
	}
	topCrates := make([]string, 0)
	for _, stack := range stacks {
		topCrates = append(topCrates, stack[len(stack)-1])
	}
	return strings.Join(topCrates, "")
}

func parseInput(input string) ([][]string, []string) {
	lines := strings.Split(input, "\n")
	totalStacks := (len(lines[0]) / 4) + 1
	stacks := make([][]string, totalStacks)
	moves := make([]string, 0)
	for i, line := range lines {
		if line[1] == '1' {
			moves = lines[i+2:]
			break
		}
		// every 4 lines is possibly crate
		for j := 0; j < len(lines[0]); j = j + 4 {
			crate := line[j : j+3]
			index := j / 4
			// if the first char is empty, skip as it is not a crate
			if string(crate[0]) == " " {
				continue
			}
			stacks[index] = append([]string{string(crate[1])}, stacks[index]...)
		}
	}
	return stacks, moves
}

func parseMove(move string) (int, int, int) {
	moveSplit := strings.Split(move, " ")
	count, _ := strconv.Atoi(moveSplit[1])
	from, _ := strconv.Atoi(moveSplit[3])
	to, _ := strconv.Atoi(moveSplit[5])
	return count, from - 1, to - 1
}
