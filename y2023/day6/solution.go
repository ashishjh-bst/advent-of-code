package day6

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	raceData := strings.Split(*input, "\n")
	numRegex := regexp.MustCompile("[0-9]+")
	times := numRegex.FindAllString(raceData[0], -1)
	distance := numRegex.FindAllString(raceData[1], -1)
	var possibilities []int
	for i, time := range times {
		intTime, _ := strconv.Atoi(time)
		intDist, _ := strconv.Atoi(distance[i])
		possible := 0
		for hold := 1; hold <= intTime; hold++ {
			if hold*(intTime-hold) > intDist {
				possible++
			}
		}
		possibilities = append(possibilities, possible)
	}

	prod := 1
	for _, possible := range possibilities {
		prod *= possible
	}
	return strconv.Itoa(prod)
}

func Part2(input *string) string {
	raceData := strings.Split(*input, "\n")
	numRegex := regexp.MustCompile("[0-9]+")
	times := numRegex.FindAllString(raceData[0], -1)
	distance := numRegex.FindAllString(raceData[1], -1)
	intTime, _ := strconv.Atoi(strings.Join(times, ""))
	intDist, _ := strconv.Atoi(strings.Join(distance, ""))
	possible := 0
	for hold := 1; hold <= intTime; hold++ {
		if hold*(intTime-hold) > intDist {
			possible++
		}
	}

	return strconv.Itoa(possible)
}
