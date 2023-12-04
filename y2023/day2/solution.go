package day2

import (
	"log"
	"strconv"
	"strings"
)

type Cube string

const (
	R Cube = "red"
	G Cube = "green"
	B Cube = "blue"
)

type Round map[Cube]int
type Game []Round

func Part1(input *string) string {
	games := strings.Split(*input, "\n")
	ruleRound := Round{R: 12, G: 13, B: 14}
	indexSum := 0
	for index, game := range games {
		parsedGame := parseGame(game)
		if parsedGame.isValid(&ruleRound) {
			indexSum += index + 1
		}
	}
	return strconv.Itoa(indexSum)
}

func Part2(input *string) string {
	games := strings.Split(*input, "\n")
	sum := 0
	for _, game := range games {
		parsedGame := parseGame(game)
		sum += parsedGame.power()
	}
	return strconv.Itoa(sum)
}

func parseGame(game string) Game {
	round_string := strings.Split(game, ":")[1]
	rounds := strings.Split(strings.Trim(round_string, " "), ";")
	var parsedGame Game
	for _, round := range rounds {
		parsedRound := Round{R: 0, G: 0, B: 0}
		cubes := strings.Split(strings.Trim(round, " "), ",")
		for _, cube := range cubes {
			parsedCube := strings.Split(strings.Trim(cube, " "), " ")
			count, err := strconv.Atoi(parsedCube[0])
			color := parsedCube[1]
			if err != nil {
				log.Fatal(err)
			}
			parsedRound[Cube(color)] = count
		}
		parsedGame = append(parsedGame, parsedRound)
	}
	return parsedGame
}

func (game *Game) isValid(ruleRound *Round) bool {
	for _, round := range *game {
		if round[R] > (*ruleRound)[R] || round[G] > (*ruleRound)[G] || round[B] > (*ruleRound)[B] {
			return false
		}
	}
	return true
}

func (game *Game) power() int {
	minValidRound := Round{R: 0, G: 0, B: 0}
	for _, round := range *game {
		if round[R] > minValidRound[R] {
			minValidRound[R] = round[R]
		}
		if round[G] > minValidRound[G] {
			minValidRound[G] = round[G]
		}
		if round[B] > minValidRound[B] {
			minValidRound[B] = round[B]
		}
	}
	return minValidRound[R] * minValidRound[G] * minValidRound[B]
}
