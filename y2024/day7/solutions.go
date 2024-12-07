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
	combinations := generateCombinations(ops, len(eqNums)-1)
	gotResult := false
	for _, combination := range combinations {
		gotResult = checkCombination(result, eqNums, combination)
		if gotResult {
			return true, result
		}
	}
	return gotResult, result
}

func checkCombination(result int, eqNums []int, combination string) bool {
	sum := 0
	for i := 0; i < len(eqNums)-1; i++ {
		if sum > result {
			return false
		}
		op := string(combination[i])
		a := eqNums[i]
		if sum > 0 {
			a = sum
		}
		sum = doOp(op, a, eqNums[i+1])
	}
	return sum == result
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

func generateCombinations(operators []string, slots int) []string {
	if slots == 0 {
		return []string{""}
	}
	previousCombinations := generateCombinations(operators, slots-1)
	var combinations []string
	for _, comb := range previousCombinations {
		for _, op := range operators {
			combinations = append(combinations, comb+op)
		}
	}
	return combinations
}
