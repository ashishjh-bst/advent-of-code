package day1

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) string {
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
	return strconv.Itoa(maxCal)
}

func Part2(input string) string {
	foodItems := strings.Split(input, "\n")
	elfCals := []int{}
	currElfItemSum := 0
	for _, item := range foodItems {
		if item == "" {
			elfCals = append(elfCals, currElfItemSum)
			currElfItemSum = 0
		} else {
			cals, err := strconv.Atoi(item)
			if err != nil {
				log.Fatalf("Couldn't convert string to int for %s, check input", item)
			}
			currElfItemSum += cals
		}
	}
	elfCals = append(elfCals, currElfItemSum)
	sort.Slice(elfCals, func(i, j int) bool {
		return elfCals[i] > elfCals[j]
	})
	return strconv.Itoa(elfCals[0] + elfCals[1] + elfCals[2])
}
