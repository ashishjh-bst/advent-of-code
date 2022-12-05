package day2

import (
	"strconv"
	"strings"
)

func Part1(input *string) string {
	rounds := strings.Split(*input, "\n")
	score := 0
	for _, round := range rounds {
		switch round {
		case "A X": // rock rock draw
			score += 4
		case "A Y": // rock paper win
			score += 8
		case "A Z": // rock scissor loss
			score += 3
		case "B X": // paper rock loss
			score += 1
		case "B Y": // paper paper draw
			score += 5
		case "B Z": // paper scissor win
			score += 9
		case "C X": // scissor rock win
			score += 7
		case "C Y": // scissor paper loss
			score += 2
		case "C Z": // scissor scissor draw
			score += 6
		}
	}
	return strconv.Itoa(score)
}

func Part2(input *string) string {
	rounds := strings.Split(*input, "\n")
	score := 0
	for _, round := range rounds {
		switch round {
		case "A X": // rock lose scissor
			score += 3
		case "A Y": // rock draw rock
			score += 4
		case "A Z": // rock win paper
			score += 8
		case "B X": // paper lose rock
			score += 1
		case "B Y": // paper draw paper
			score += 5
		case "B Z": // paper win scissor
			score += 9
		case "C X": // scissor lose paper
			score += 2
		case "C Y": // scissor draw scissor
			score += 6
		case "C Z": // scissor win rock
			score += 7
		}
	}
	return strconv.Itoa(score)
}
