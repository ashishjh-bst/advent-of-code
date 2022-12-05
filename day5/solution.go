package day5

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) string {
	stacks, moves := parseInput(input)
	for _, move := range moves {
		count, from, to := parseMove(move)
		fromStack := stacks[from]
		toStack := stacks[to]
		for i := 0; i < count; i++ {
			fromStackSize := len(fromStack)
			pop := fromStack[fromStackSize-1]
			fromStack = fromStack[:fromStackSize-1]
			toStack = append(toStack, pop)
		}
		stacks[from] = fromStack
		stacks[to] = toStack
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
		fromStack := stacks[from]
		toStack := stacks[to]
		fromStackSize := len(fromStack)
		pop := fromStack[(fromStackSize - count):]
		fromStack = fromStack[:(fromStackSize - count)]
		toStack = append(toStack, pop...)
		stacks[from] = fromStack
		stacks[to] = toStack
	}

	topCrates := make([]string, 0)
	for _, stack := range stacks {
		topCrates = append(topCrates, stack[len(stack)-1])
	}
	return strings.Join(topCrates, "")
}

func parseInput(input string) ([][]string, []string) {
	lines := strings.Split(input, "\n")
	stacks := make(map[int][]string, 0)
	moves := make([]string, 0)
	splitter := regexp.MustCompile(`.{1,4}`)
	for i, line := range lines {
		if line[1] == '1' {
			moves = lines[i+2:]
			break
		}
		crates := splitter.FindAllString(line, -1)
		for index, crate := range crates {
			if string(crate[0]) == " " {
				continue
			}
			stackKey := index + 1
			if len(stacks[stackKey]) == 0 {
				stacks[stackKey] = make([]string, 0)
			}
			stacks[stackKey] = append([]string{string(crate[1])}, stacks[stackKey]...)
		}
	}
	stacksSlice := make([][]string, len(stacks), len(stacks))
	for key, val := range stacks {
		stacksSlice[key-1] = val
	}
	return stacksSlice, moves
}

func parseMove(move string) (int, int, int) {
	moveSplit := strings.Split(move, " ")
	count, _ := strconv.Atoi(moveSplit[1])
	from, _ := strconv.Atoi(moveSplit[3])
	to, _ := strconv.Atoi(moveSplit[5])
	return count, from - 1, to - 1
}
