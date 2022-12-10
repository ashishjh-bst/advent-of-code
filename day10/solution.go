package day10

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	instructions := strings.Split(*input, "\n")
	signaStrengthSum := 0
	xValue := 1
	cycleCounter := 0
	for _, instruction := range instructions {
		command := strings.Split(instruction, " ")
		if command[0] == "noop" {
			cycleCounter++
			fmt.Println(cycleCounter, command[0], xValue)
			signaStrengthSum += logStrength(cycleCounter, xValue)
			continue
		}
		if command[0] == "addx" {
			for i := 0; i < 2; i++ {
				cycleCounter++
				signaStrengthSum += logStrength(cycleCounter, xValue)
				fmt.Println(cycleCounter, command[0], xValue)
			}
			xIncr, _ := strconv.Atoi(command[1])
			xValue += xIncr
			fmt.Println(cycleCounter, command[0], xValue)
		}
	}
	return strconv.Itoa(signaStrengthSum)
}

func logStrength(cycle int, x int) int {
	if cycle == 20 {
		return x * cycle
	}
	if cycle > 220 {
		return 0
	}
	if (cycle-20)%40 == 0 {
		return x * cycle
	}
	return 0
}

func Part2(input *string) string {
	instructions := strings.Split(*input, "\n")
	cycle := 0
	display := initDisplay()
	xValue := 1
	currentRow := 0
	currentColumn := 0
	for _, instruction := range instructions {
		command := strings.Split(instruction, " ")
		if command[0] == "noop" {
			updateDisplay(currentRow, currentColumn, xValue, &display)
			currentRow, currentColumn = getNextPixel(currentRow, currentColumn)
			cycle++
			continue
		} else if command[0] == "addx" {
			for i := 0; i < 2; i++ {
				updateDisplay(currentRow, currentColumn, xValue, &display)
				currentRow, currentColumn = getNextPixel(currentRow, currentColumn)
				cycle++
			}
			xIncr, _ := strconv.Atoi(command[1])
			xValue += xIncr
		}
	}
	drawDisplay(&display)
	return "use your eyes"
}

func updateDisplay(row int, column int, x int, display *[][]string) {
	if column-1 == x || column == x || column+1 == x {
		(*display)[row][column] = "#"
	} else {
		(*display)[row][column] = "."
	}
}

func drawDisplay(display *[][]string) {
	for i := 0; i < 6; i++ {
		fmt.Println(strings.Join((*display)[i], ""))
	}
}

func getNextPixel(row, column int) (int, int) {
	column++
	if column > 39 {
		column = 0
		row++
	}
	if row > 5 {
		row = 0
	}
	return row, column
}

func initDisplay() [][]string {
	display := make([][]string, 6)
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			display[i] = append(display[i], " ")
		}
	}
	return display
}
