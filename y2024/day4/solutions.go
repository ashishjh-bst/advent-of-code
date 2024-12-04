package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	grid := makeGrid(input)
	//all eight directions
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
	word := "XMAS"
	wordCount := 0
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len((*grid)[i]); j++ {
			startPos := [2]int{i, j}
			fmt.Printf("StartPOS: %v \n", startPos)
			for _, dir := range directions {
				fmt.Printf("Direction: %v", dir)
				if searchDirectionForWord(grid, word, startPos, dir) {
					wordCount++
				}
				fmt.Print("\n")
			}
			fmt.Print("\n")
		}
	}
	return strconv.Itoa(wordCount)
}

func Part2(input *string) string {
	grid := makeGrid(input)
	crossCount := 0
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len((*grid)[i]); j++ {
			gridPoint := (*grid)
			if gridPoint[i][j] != "A" {
				continue
			}
			if checkDiagonals(grid, [2]int{i, j}) {
				crossCount++
			}
		}
	}
	return strconv.Itoa(crossCount)
}

func makeGrid(input *string) *[][]string {
	grid := make([][]string, 0)
	lines := strings.Split(*input, "\n")
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}
	return &grid
}

func searchDirectionForWord(grid *[][]string, word string, start [2]int, dir [2]int) bool {
	gridPoint := *grid
	for i := 0; i < len(word); i++ {
		checkPos := [2]int{start[0] + i*dir[0], start[1] + i*dir[1]}
		// fmt.Printf("Checking pos %v == %s, rowLen: %d, colLen: %d \n", checkPos, string(word[i-1]), len(gridPoint), len(gridPoint[0]))
		if checkPos[0] == -1 || checkPos[1] == -1 || checkPos[0] >= len(gridPoint) || checkPos[1] >= len(gridPoint[0]) {
			return false
		}
		if gridPoint[checkPos[0]][checkPos[1]] != string(word[i]) {
			return false
		}
		fmt.Printf("%s", string(word[i]))
	}
	return true
}

func checkPosForChar(grid *[][]string, char string, checkPos [2]int) bool {
	gridPoint := *grid
	if checkPos[0] == -1 || checkPos[1] == -1 || checkPos[0] >= len(gridPoint) || checkPos[1] >= len(gridPoint[0]) {
		return false
	}
	return gridPoint[checkPos[0]][checkPos[1]] == char
}

func checkDiagonals(grid *[][]string, start [2]int) bool {
	i := start[0]
	j := start[1]
	topLeft := [2]int{i - 1, j - 1}
	topRight := [2]int{i - 1, j + 1}
	bottomLeft := [2]int{i + 1, j - 1}
	bottomRight := [2]int{i + 1, j + 1}
	rightDiagonal := (checkPosForChar(grid, "S", topLeft) && checkPosForChar(grid, "M", bottomRight)) || (checkPosForChar(grid, "M", topLeft) && checkPosForChar(grid, "S", bottomRight))
	leftDiagonal := (checkPosForChar(grid, "S", topRight) && checkPosForChar(grid, "M", bottomLeft)) || (checkPosForChar(grid, "M", topRight) && checkPosForChar(grid, "S", bottomLeft))
	return rightDiagonal && leftDiagonal
}
