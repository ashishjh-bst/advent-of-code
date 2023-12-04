package day3

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func Part1(input *string) string {
	lines := strings.Split(*input, "\n")
	schematic := make([][]string, 0)
	var parts []int
	for _, line := range lines {
		splitLine := strings.Split(line, "")
		schematic = append(schematic, splitLine)
	}

	for row := range schematic {
		for col := range schematic[row] {
			if isSymbol(schematic[row][col]) {
				adjacent := checkAdjacent(schematic, row, col)
				parts = append(parts, adjacent...)
			}
		}
	}
	sum := 0
	for _, num := range parts {
		sum += num
	}
	return strconv.Itoa(sum)
}

func Part2(input *string) string {
	lines := strings.Split(*input, "\n")
	schematic := make([][]string, 0)
	var parts []int
	for _, line := range lines {
		splitLine := strings.Split(line, "")
		schematic = append(schematic, splitLine)
	}

	for row := range schematic {
		for col := range schematic[row] {
			if schematic[row][col] == "*" {
				adjacent := checkAdjacent(schematic, row, col)
				if len(adjacent) == 2 {
					parts = append(parts, adjacent[0]*adjacent[1])
				}
			}
		}
	}
	sum := 0
	for _, num := range parts {
		sum += num
	}
	return strconv.Itoa(sum)
}

func isDigit(c string) bool {
	if len(c) > 1 {
		return false
	}
	for _, r := range c {
		return unicode.IsDigit(r)
	}
	return false
}

func isSymbol(c string) bool {
	return c != "." && !isDigit(c)
}

func checkAdjacent(schematic [][]string, x int, y int) []int {
	var adjacentNumbers []int

	//top left
	if x-1 >= 0 && y-1 >= 0 && isDigit(schematic[x-1][y-1]) {
		num := parseNumber(schematic, x-1, y-1)
		adjacentNumbers = append(adjacentNumbers, num)
	}
	//top right
	if x-1 >= 0 && y+1 < len(schematic[x]) && isDigit(schematic[x-1][y+1]) {
		num := parseNumber(schematic, x-1, y+1)
		adjacentNumbers = append(adjacentNumbers, num)
	}
	//left
	if y-1 >= 0 && isDigit(schematic[x][y-1]) {
		num := parseNumber(schematic, x, y-1)
		adjacentNumbers = append(adjacentNumbers, num)
	}
	//right
	if y+1 < len(schematic[x]) && isDigit(schematic[x][y+1]) {
		num := parseNumber(schematic, x, y+1)
		adjacentNumbers = append(adjacentNumbers, num)
	}
	//bottom
	if x+1 < len(schematic) && isDigit(schematic[x+1][y]) {
		num := parseNumber(schematic, x+1, y)
		adjacentNumbers = append(adjacentNumbers, num)
	}
	//top
	if x-1 >= 0 && isDigit(schematic[x-1][y]) {
		num := parseNumber(schematic, x-1, y)
		adjacentNumbers = append(adjacentNumbers, num)
	}
	//bottom-left
	if x+1 < len(schematic) && y-1 >= 0 && isDigit(schematic[x+1][y-1]) {
		num := parseNumber(schematic, x+1, y-1)
		adjacentNumbers = append(adjacentNumbers, num)
	}
	//bottom-right
	if x+1 < len(schematic) && y+1 < len(schematic[x]) && isDigit(schematic[x+1][y+1]) {
		num := parseNumber(schematic, x+1, y+1)
		adjacentNumbers = append(adjacentNumbers, num)
	}
	return adjacentNumbers
}

func parseNumber(schematic [][]string, x int, y int) int {
	for y > 0 && isDigit(schematic[x][y-1]) {
		y--
	}

	num := ""
	for y+1 <= len(schematic[x]) && isDigit(schematic[x][y]) {
		num = fmt.Sprintf("%s%s", num, schematic[x][y])
		schematic[x][y] = "."
		y++
	}

	parsedNum, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return parsedNum
}
