package day11

import (
	"fmt"
	"strconv"
	"strings"
)

var SplitTimeLengthCache = make(map[[2]int]int)

func Part1(input *string) string {
	stones := parseInput(input)
	splits := 25
	length := 0

	for _, stone := range *stones {
		stoneLength := getSplitLength(stone, splits)
		length += stoneLength
		fmt.Printf("\n length for stone %d after splitting: %d ", stone, stoneLength)
	}

	return fmt.Sprintf("%d", length)
}

func Part2(input *string) string {
	stones := parseInput(input)
	splits := 75
	length := 0

	for _, stone := range *stones {
		stoneLength := getSplitLength(stone, splits)
		length += stoneLength
		fmt.Printf("\n length for stone %d after splitting: %d ", stone, stoneLength)
	}
	return fmt.Sprintf("%d", length)
}

func getSplit(stone int) []int {
	if stone == 0 {
		return []int{1}
	}
	stoneStr := strconv.Itoa(stone)
	stoneLen := len(stoneStr)
	if len(stoneStr)%2 == 0 {
		a, _ := strconv.Atoi(stoneStr[:stoneLen/2])
		b, _ := strconv.Atoi(stoneStr[stoneLen/2:])
		return []int{a, b}
	}
	return []int{stone * 2024}
}

func getSplitLength(stone, times int) int {
	cachedLength, ok := SplitTimeLengthCache[[2]int{stone, times}]
	if ok {
		return cachedLength
	}
	if times == 1 {
		splitLen := len(getSplit(stone))
		SplitTimeLengthCache[[2]int{stone, times}] = splitLen
		return splitLen
	}
	length := 0
	for _, splittedStone := range getSplit(stone) {
		splitLen := getSplitLength(splittedStone, times-1)
		SplitTimeLengthCache[[2]int{splittedStone, times - 1}] = splitLen
		length += splitLen
	}
	return length
}

func parseInput(input *string) *[]int {
	stones := make([]int, 0)
	list := strings.Split(*input, " ")
	for _, item := range list {
		stone, _ := strconv.Atoi(item)
		stones = append(stones, stone)
	}
	return &stones
}
