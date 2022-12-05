package day5

import (
	"regexp"
	"sort"
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
	stackSortedKeys := make([]int, 0, len(stacks))
	for key := range stacks {
		stackSortedKeys = append(stackSortedKeys, key)
	}
	sort.Slice(stackSortedKeys, func(i, j int) bool {
		return stackSortedKeys[i] <= stackSortedKeys[j]
	})
	for _, key := range stackSortedKeys {
		stack := stacks[key]
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
	stackSortedKeys := make([]int, 0, len(stacks))
	for key := range stacks {
		stackSortedKeys = append(stackSortedKeys, key)
	}
	sort.Slice(stackSortedKeys, func(i, j int) bool {
		return stackSortedKeys[i] <= stackSortedKeys[j]
	})
	for _, key := range stackSortedKeys {
		stack := stacks[key]
		topCrates = append(topCrates, stack[len(stack)-1])
	}
	return strings.Join(topCrates, "")
}

func parseInput(input string) (map[int][]string, []string) {
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
	return stacks, moves
}

func parseMove(move string) (int, int, int) {
	moveSplit := strings.Split(move, " ")
	count, _ := strconv.Atoi(moveSplit[1])
	from, _ := strconv.Atoi(moveSplit[3])
	to, _ := strconv.Atoi(moveSplit[5])
	return count, from, to
}
