package day6

import (
	"strconv"
	"strings"
)

func Part1(input *string) string {
	lights := make(lightsMap, 0)
	instructions := strings.Split(*input, "\n")
	for _, i := range instructions {
		instruction := parseInstruction(i)
		followInstruction(&instruction, &lights)
	}

	var totalOn int
	for _, v := range lights {
		if v {
			totalOn++
		}
	}

	return strconv.Itoa(totalOn)
}

func Part2(input *string) string {
	lights := make(lightsBrightness, 0)
	instructions := strings.Split(*input, "\n")
	for _, i := range instructions {
		instruction := parseInstruction(i)
		followInstruction2(&instruction, &lights)
	}

	var totalOn int
	for _, v := range lights {
		totalOn += v
	}

	return strconv.Itoa(totalOn)
}

type instruction struct {
	task  string
	start [2]int
	end   [2]int
}

type lightsMap map[[2]int]bool

func followInstruction(ins *instruction, light *lightsMap) {
	for i := ins.start[0]; i <= ins.end[0]; i++ {
		for j := ins.start[1]; j <= ins.end[1]; j++ {
			switch ins.task {
			case "toggle":
				(*light)[[2]int{i, j}] = !(*light)[[2]int{i, j}]
			case "on":
				(*light)[[2]int{i, j}] = true
			case "off":
				(*light)[[2]int{i, j}] = false
			}
		}
	}
}

func parseInstruction(word string) (ins instruction) {
	splitWord := strings.Split(word, " ")
	var start []string
	var end []string
	if splitWord[0] == "turn" {
		ins.task = splitWord[1]
		start = strings.Split(splitWord[2], ",")
		end = strings.Split(splitWord[4], ",")
	} else {
		ins.task = splitWord[0]
		start = strings.Split(splitWord[1], ",")
		end = strings.Split(splitWord[3], ",")
	}
	ins.start[0], _ = strconv.Atoi(start[0])
	ins.start[1], _ = strconv.Atoi(start[1])
	ins.end[0], _ = strconv.Atoi(end[0])
	ins.end[1], _ = strconv.Atoi(end[1])
	return
}

type lightsBrightness map[[2]int]int

func followInstruction2(ins *instruction, light *lightsBrightness) {
	for i := ins.start[0]; i <= ins.end[0]; i++ {
		for j := ins.start[1]; j <= ins.end[1]; j++ {
			switch ins.task {
			case "toggle":
				(*light)[[2]int{i, j}] += 2
			case "on":
				(*light)[[2]int{i, j}]++
			case "off":
				if (*light)[[2]int{i, j}] > 0 {
					(*light)[[2]int{i, j}]--
				}
			}
		}
	}
}
