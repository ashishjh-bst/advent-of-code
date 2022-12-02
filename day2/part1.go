package day2

import (
	"log"
	"strings"
)

func Part1(input string) {
	rounds := strings.Split(input, "\n")
	log.Printf("rounds %#v", rounds)
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
	log.Printf("Total Score for strategy is %d", score)
}
