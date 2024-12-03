package day3

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	mulFinderRegex := regexp.MustCompile(`mul\([\d]{1,3},[\d]{1,3}\)`)
	mulInstructions := mulFinderRegex.FindAllString(*input, -1)
	var multSum int
	for _, mul := range mulInstructions {
		multSum += parseMul(mul)
	}
	return strconv.Itoa(multSum)
}

func Part2(input *string) string {
	instructionRegex := regexp.MustCompile(`(do\(\)|don't\(\)|mul\([\d]{1,3},[\d]{1,3}\))`)
	instructions := instructionRegex.FindAllString(*input, -1)
	var multSum int
	allowSum := true
	for _, inst := range instructions {
		if inst == "do()" {
			allowSum = true
		} else if inst == "don't()" {
			allowSum = false
		} else if strings.HasPrefix(inst, "mul(") && allowSum {
			multSum += parseMul(inst)
		}
	}
	return strconv.Itoa(multSum)
}

func parseMul(mul string) int {
	getNumsRegex := regexp.MustCompile(`[\d]{1,3}`)
	nums := getNumsRegex.FindAllString(mul, -1)
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])
	return num1 * num2
}
