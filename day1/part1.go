package day1

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Part1(input string) {
	foodItems := strings.Split(input, "\n")
	maxCal := 0
	currElfItemSum := 0
	for _, item := range foodItems {
		if item == "" {
			currElfItemSum = 0
		} else {
			cals, err := strconv.Atoi(item)
			if err != nil {
				log.Fatalf("Couldn't convert string to int for %s, check input", item)
			}
			currElfItemSum += cals
		}
		if maxCal < currElfItemSum {
			maxCal = currElfItemSum
		}
	}
	fmt.Printf("Max Calories: %d \n", maxCal)
}
