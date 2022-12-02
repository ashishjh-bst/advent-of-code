package day2

import (
	"log"
	"strings"
)

func Part2(input string) {
	rounds := strings.Split(input, "\n")
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
	log.Printf("Total Score for strategy is %d", score)
}
