package day6

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	raceData := strings.Split(*input, "\n")
	numRegex := regexp.MustCompile("[0-9]+")
	times := numRegex.FindAllString(raceData[0], -1)
	distance := numRegex.FindAllString(raceData[1], -1)
	prod := 1
	for i, time := range times {
		intTime, _ := strconv.Atoi(time)
		intDist, _ := strconv.Atoi(distance[i])
		prod *= calc_possibility(intTime, intDist)
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
	return strconv.Itoa(calc_possibility(intTime, intDist))
}

func calc_possibility(time, dist int) int {
	speed := math.Sqrt(float64(time*time - 4*dist))
	holdMin := int(math.Ceil((float64(time) - speed) / 2))
	holdMax := int(math.Floor((float64(time) + speed) / 2))
	return holdMax - holdMin + 1
}
