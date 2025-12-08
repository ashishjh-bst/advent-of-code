package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	turns := strings.Split(*input, "\n")
	dial := 50
	totalZero := 0
	for _, turn := range turns {
		dir := string(turn[0])
		num := ""
		for i := 1; i < len(turn); i++ {
			num = num + string(turn[i])
		}
		intNum, _ := strconv.Atoi(num)
		if intNum >= 100 {
			intNum = intNum % 100
		}
		switch dir {
		case "L":
			dial -= intNum
		case "R":
			dial += intNum
		}
		if dial == 100 || dial == 0 {
			dial = 0
			totalZero++
		} else if dial < 0 {
			dial = 100 + dial
		} else if dial > 100 {
			dial = dial - 100
		}
	}
	return fmt.Sprintf("%d", totalZero)
}

func Part2(input *string) string {
	turns := strings.Split(*input, "\n")
	dial := 50
	totalZero := 0
	for _, turn := range turns {
		prevDial := dial
		dir := string(turn[0])
		num := ""
		for i := 1; i < len(turn); i++ {
			num = num + string(turn[i])
		}
		intNum, _ := strconv.Atoi(num)
		if intNum >= 100 {
			totalZero += int(intNum / 100)
			intNum = intNum % 100
		}
		switch dir {
		case "L":
			dial -= intNum
		case "R":
			dial += intNum
		}
		if dial == 100 || dial == 0 {
			dial = 0
			totalZero++
		} else if dial < 0 {
			dial = 100 + dial
			if prevDial != 0 {
				totalZero++
			}
		} else if dial > 100 {
			dial = dial - 100
			if prevDial != 0 {
				totalZero++
			}
		}
	}
	return fmt.Sprintf("%d", totalZero)
}
