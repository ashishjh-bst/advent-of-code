package day7

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	lines := strings.Split(*input, "\n")
	sum := 0
	for _, line := range lines {
		lineResult, result := checkLine(line, []string{"+", "*"})
		if lineResult {
			sum += result
		}
	}
	return strconv.Itoa(sum)
}

func Part2(input *string) string {
	lines := strings.Split(*input, "\n")
	sum := 0
	for _, line := range lines {
		lineResult, result := checkLine(line, []string{"+", "*", "|"})
		if lineResult {
			sum += result
		}
	}
	return strconv.Itoa(sum)
}

func parseLine(line string) (int, []int) {
	lineSplit := strings.Split(line, ": ")
	equationParts := strings.Split(lineSplit[1], " ")
	equationNums := make([]int, 0)
	for _, ep := range equationParts {
		e, _ := strconv.Atoi(ep)
		equationNums = append(equationNums, e)
	}
	rNum, _ := strconv.Atoi(lineSplit[0])
	return rNum, equationNums
}

func checkLine(line string, ops []string) (bool, int) {
	result, eqNums := parseLine(line)
	shouldAdd := isValidPath(result, eqNums, eqNums[0], 1, ops)
	return shouldAdd, result
}

func doOp(op string, a, b int) int {
	switch op {
	case "+":
		return a + b
	case "*":
		return a * b
	case "|":
		concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
		return concat
	}
	fmt.Printf("THIS SHOULD NEVER HAPPEN")
	return 0
}

func isValidPath(expected int, list []int, lastSum int, nextPos int, ops []string) bool {
	if lastSum > expected {
		return false
	}
	if nextPos == len(list) {
		return expected == lastSum
	}
	for _, op := range ops {
		if isValidPath(expected, list, doOp(op, lastSum, list[nextPos]), nextPos+1, ops) {
			return true
		}
	}
	return false
}
