package day8

import (
	"strconv"
	"strings"
)

func Part1(input *string) string {
	matrix := makeMatrix(input)
	visibleTrees := 0
	for i := 0; i < len(matrix); i++ {
		rows := len(matrix)
		for j := 0; j < len(matrix[i]); j++ {
			cols := len(matrix[i])

			//if edge tree, don't scan, they are all visible
			if i == 0 || j == 0 || i == rows-1 || j == cols-1 {
				visibleTrees++
				continue
			}
			currentTree := matrix[i][j]

			//scan from left
			isTreeInvisible := false
			for k := 0; k < j; k++ {
				if matrix[i][k] >= currentTree {
					isTreeInvisible = true
				}
			}
			if !isTreeInvisible {
				visibleTrees++
				continue
			}

			//scan from right
			isTreeInvisible = false
			for k := j + 1; k < cols; k++ {
				if matrix[i][k] >= currentTree {
					isTreeInvisible = true
				}
			}
			if !isTreeInvisible {
				visibleTrees++
				continue
			}

			//scan from top
			isTreeInvisible = false
			for k := 0; k < i; k++ {
				if matrix[k][j] >= currentTree {
					isTreeInvisible = true
				}
			}
			if !isTreeInvisible {
				visibleTrees++
				continue
			}

			//scan from bottom
			isTreeInvisible = false
			for k := i + 1; k < rows; k++ {
				if matrix[k][j] >= currentTree {
					isTreeInvisible = true
				}
			}
			if !isTreeInvisible {
				visibleTrees++
				continue
			}
		}
	}
	return strconv.Itoa(visibleTrees)
}

func Part2(input *string) string {
	matrix := makeMatrix(input)
	maxScore := 0
	for i := 0; i < len(matrix); i++ {
		rows := len(matrix)
		for j := 0; j < len(matrix[i]); j++ {
			cols := len(matrix[i])
			//if edge tree, don't scan, they are all visible
			if i == 0 || j == 0 || i == rows-1 || j == cols-1 {
				continue
			}
			currentTree := matrix[i][j]

			scoreRight := 1
			for k := j + 1; k < rows-1; k++ {
				if matrix[i][k] >= currentTree {
					break
				}
				scoreRight++
			}

			scoreLeft := 1
			for k := j - 1; k > 0; k-- {
				if matrix[i][k] >= currentTree {
					break
				}
				scoreLeft++
			}

			scoreBottom := 1
			for k := i + 1; k < cols-1; k++ {
				if matrix[k][j] >= currentTree {
					break
				}
				scoreBottom++
			}

			scoreTop := 1
			for k := i - 1; k > 0; k-- {
				if matrix[k][j] >= currentTree {
					break
				}
				scoreTop++
			}

			score := scoreBottom * scoreLeft * scoreRight * scoreTop
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return strconv.Itoa(maxScore)
}

func makeMatrix(input *string) [][]int {
	matrix := make([][]int, 0)
	lines := strings.Split(*input, "\n")
	for _, line := range lines {
		row := make([]int, 0)
		rowString := strings.Split(line, "")
		for _, ele := range rowString {
			eleInt, _ := strconv.Atoi(ele)
			row = append(row, eleInt)
		}
		matrix = append(matrix, row)
	}
	return matrix
}
