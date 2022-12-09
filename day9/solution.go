package day9

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	moves := strings.Split(*input, "\n")
	headPos := [2]int{0, 0}
	tailPos := [2]int{0, 0}
	moveMap := make(map[string]int, 0)
	moveMap["0,0"] = 1
	for _, move := range moves {
		moveParser := strings.Split(move, " ")
		dir := moveParser[0]
		steps, _ := strconv.Atoi(moveParser[1])
		for i := 0; i < steps; i++ {
			switch dir {
			case "R":
				headPos[0]++
				if abs(headPos[0]-tailPos[0]) == 2 {
					tailPos[0]++
					tailPos[1] = headPos[1]
				}
			case "L":
				headPos[0]--
				if abs(headPos[0]-tailPos[0]) == 2 {
					tailPos[0]--
					tailPos[1] = headPos[1]
				}
			case "U":
				headPos[1]++
				if abs(headPos[1]-tailPos[1]) == 2 {
					tailPos[1]++
					tailPos[0] = headPos[0]
				}
			case "D":
				headPos[1]--
				if abs(headPos[1]-tailPos[1]) == 2 {
					tailPos[1]--
					tailPos[0] = headPos[0]
				}
			}
			moveMap[fmt.Sprintf("%d,%d", tailPos[0], tailPos[1])]++
		}
	}
	return fmt.Sprintf("%d", len(moveMap))
}
func Part2(input *string) string {

	return ""
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -1 * x
}
