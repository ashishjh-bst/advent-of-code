package day10

import (
	"fmt"
	"strconv"
	"strings"
)

var Directions = [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func Part1(input *string) string {
	grid, starts := makeGrid(input)
	var totalTrailHeads int
	for _, start := range *starts {
		totalTrailHeads += countTrailFromStart(grid, start)
	}
	return fmt.Sprintf("%d", totalTrailHeads)
}

func Part2(input *string) string {
	grid, starts := makeGrid(input)
	tEnds := &map[[2]int]int{}
	for _, start := range *starts {
		for _, dir := range Directions {
			isTrail(grid, start, dir, tEnds)
		}
	}
	var totalRating int
	for _, v := range *tEnds {
		totalRating += v
	}
	return fmt.Sprintf("%d", totalRating)
}

func makeGrid(input *string) (*[][]int, *[][2]int) {
	grid := make([][]int, 0)
	starts := make([][2]int, 0)
	for i, row := range strings.Split(*input, "\n") {
		r := make([]int, 0)
		for j, col := range strings.Split(row, "") {
			val, err := strconv.Atoi(col)
			if err != nil {
				val = -1
			}
			if val == 0 {
				starts = append(starts, [2]int{i, j})
			}
			r = append(r, val)
		}
		grid = append(grid, r)
	}
	return &grid, &starts
}

func countTrailFromStart(grid *[][]int, start [2]int) int {
	tEnds := make(map[[2]int]int)
	for _, dir := range Directions {
		isTrail(grid, start, dir, &tEnds)
	}
	return len(tEnds)
}

func isTrail(grid *[][]int, current [2]int, dir [2]int, tEnds *map[[2]int]int) bool {
	currentValue := (*grid)[current[0]][current[1]]
	next := [2]int{current[0] + dir[0], current[1] + dir[1]}
	if isPosOutside(next, grid) {
		return false
	}
	nextValue := (*grid)[next[0]][next[1]]
	if nextValue != currentValue+1 {
		return false
	}
	if nextValue == 9 {
		(*tEnds)[next]++
		return true
	}
	trail := false
	for _, dir := range Directions {
		isNextTrail := isTrail(grid, next, dir, tEnds)
		trail = trail || isNextTrail
	}
	return trail
}

func isPosOutside(currentPos [2]int, grid *[][]int) bool {
	return currentPos[0] < 0 || currentPos[0] >= len(*grid) || currentPos[1] < 0 || currentPos[1] >= len((*grid)[0])
}
