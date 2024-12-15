package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ashishjh-bst/aoc/common"
)

func Part1(input *string) string {
	visitedMap := make(map[[2]int]bool)
	grid, startPos, startDir, blocks := creategrid(input)
	currentPos := startPos
	currentDir := startDir
	visitedMap[currentPos] = true
	for !common.IsPosOutside(currentPos, &grid) {
		currentPos, currentDir = move(currentPos, currentDir, &blocks)
		visitedMap[currentPos] = true
	}
	delete(visitedMap, startPos)
	return strconv.Itoa(len(visitedMap))
}

func Part2(input *string) string {
	visitedMap := make(map[[2]int]bool)
	grid, startPos, startDir, blocks := creategrid(input)
	currentPos := startPos
	currentDir := startDir
	for !common.IsPosOutside(currentPos, &grid) {
		currentPos, currentDir = move(currentPos, currentDir, &blocks)
		visitedMap[currentPos] = true
	}
	var loops int
	delete(visitedMap, startPos)
	for k := range visitedMap {
		blocks[k] = true
		if IsLoopPossible(&grid, &blocks, startPos, startDir) {
			loops++
		}
		delete(blocks, k)
	}
	return strconv.Itoa(loops)
}

func IsLoopPossible(grid *[][]string, blocks *map[[2]int]bool, currentPos [2]int, currentDir [2]int) bool {
	visitedAndDirMap := make(map[[4]int]bool)
	for !common.IsPosOutside(currentPos, grid) {
		currentPos, currentDir = move(currentPos, currentDir, blocks)
		currentDirPos := [4]int{currentDir[0], currentDir[1], currentPos[0], currentPos[1]}
		_, ok := visitedAndDirMap[currentDirPos]
		if ok {
			return true
		}
		visitedAndDirMap[currentDirPos] = true
	}
	return false
}

func creategrid(input *string) (grid [][]string, startpos [2]int, startDir [2]int, blocks map[[2]int]bool) {
	lines := strings.Split(*input, "\n")
	grid = make([][]string, 0)
	blocks = make(map[[2]int]bool)
	for i, row := range lines {
		col := strings.Split(row, "")
		grid = append(grid, col)
		for j, c := range col {
			if c == "^" || c == "<" || c == ">" || c == "v" {
				startDir = getDir(c)
				startpos = [2]int{i, j}
				continue
			}
			if c == "#" {
				blocks[[2]int{i, j}] = true
			}
		}
	}
	return
}

func getDir(c string) [2]int {
	switch c {
	case "^":
		return [2]int{-1, 0}
	case ">":
		return [2]int{0, 1}
	case "v":
		return [2]int{1, 0}
	case "<":
		return [2]int{0, -1}
	}
	fmt.Printf("THIS SHOULD NEVER HAPPEN")
	return [2]int{}
}

func getNextDir(c [2]int) [2]int {
	switch c {
	case getDir("^"):
		return getDir(">")
	case getDir(">"):
		return getDir("v")
	case getDir("v"):
		return getDir("<")
	case getDir("<"):
		return getDir("^")
	}
	fmt.Printf("THIS SHOULD NEVER HAPPEN")
	return c
}

func move(currentPos [2]int, currentDir [2]int, blocks *map[[2]int]bool) ([2]int, [2]int) {
	nextPos := [2]int{currentPos[0] + currentDir[0], currentPos[1] + currentDir[1]}
	_, ok := (*blocks)[nextPos]
	if ok {
		nextDir := getNextDir(currentDir)
		return currentPos, nextDir
	}
	return nextPos, currentDir
}
