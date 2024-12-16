package day13

import (
	"fmt"
	"strconv"
	"strings"
)

type ClawMachine struct {
	A     [2]int
	B     [2]int
	Prize [2]int
}

func Part1(input *string) string {
	machines := parseInput(input)
	costA := 3
	costB := 1
	totalCost := 0
	for _, machine := range *machines {
		solution, isWinnable := machine.solve(0)
		if !isWinnable {
			continue
		}
		totalCost += (*solution)[0]*costA + (*solution)[1]*costB
	}
	return fmt.Sprintf("%d", totalCost)
}

func Part2(input *string) string {
	machines := parseInput(input)
	costA := 3
	costB := 1
	totalCost := 0
	for _, machine := range *machines {
		solution, isWinnable := machine.solve(10000000000000)
		if !isWinnable {
			continue
		}
		totalCost += (*solution)[0]*costA + (*solution)[1]*costB
	}
	return fmt.Sprintf("%d", totalCost)
}

func parseInput(input *string) *[]ClawMachine {
	machines := strings.Split(*input, "\n\n")
	parsedMachines := make([]ClawMachine, 0)
	for _, machine := range machines {
		parsedMachines = append(parsedMachines, *createMachine(machine))
	}
	return &parsedMachines
}

func createMachine(str string) *ClawMachine {
	lines := strings.Split(str, "\n")
	A := parseButton(lines[0])
	B := parseButton(lines[1])
	Prize := parsePrize(lines[2])
	return &ClawMachine{A, B, Prize}
}

func parsePrize(str string) [2]int {
	inst := strings.Split(str, ": ")
	values := strings.Split(inst[1], ", ")
	xStr, _ := strings.CutPrefix(values[0], "X=")
	xVal, _ := strconv.Atoi(xStr)
	yStr, _ := strings.CutPrefix(values[1], "Y=")
	yVal, _ := strconv.Atoi(yStr)
	return [2]int{xVal, yVal}
}

func parseButton(str string) [2]int {
	inst := strings.Split(str, ": ")
	values := strings.Split(inst[1], ", ")
	xStr, _ := strings.CutPrefix(values[0], "X+")
	xVal, _ := strconv.Atoi(xStr)
	yStr, _ := strings.CutPrefix(values[1], "Y+")
	yVal, _ := strconv.Atoi(yStr)
	return [2]int{xVal, yVal}
}

func (c *ClawMachine) solve(buffer int) (solution *[2]int, isWinnable bool) {
	Xa, Xb := c.A[0], c.B[0]
	Ya, Yb := c.A[1], c.B[1]
	Px, Py := c.Prize[0]+buffer, c.Prize[1]+buffer
	BDividend := (Xa * Py) - (Px * Ya)
	BDivsor := (Xa * Yb) - (Xb * Ya)
	if BDivsor == 0 {
		return nil, false
	}

	if BDividend%BDivsor != 0 {
		return nil, false
	}
	B := BDividend / BDivsor

	ADividend := Px - (B * Xb)
	if Xa == 0 {
		return nil, false
	}

	if ADividend%Xa != 0 {
		return nil, false
	}
	A := ADividend / Xa
	return &[2]int{A, B}, true
}
