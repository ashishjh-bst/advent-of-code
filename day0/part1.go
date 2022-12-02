package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ashishjh-bst/aoc2022/common"
)

func main() {
	args := os.Args[1:]
	inputFilePath := args[0]
	input := common.ReadFileInput(inputFilePath)
	foodItems := strings.Split(input, "\n")
	maxCal := 0
	currElfItemSum := 0
	for _, item := range foodItems {
		if item == "" {
			currElfItemSum = 0
		} else {
			cals, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal("Couldn't convert string to int for %s, check input")
			}
			currElfItemSum += cals
		}
		if maxCal < currElfItemSum {
			maxCal = currElfItemSum
		}
	}
	fmt.Printf("Max Calories %d \n", maxCal)
}
