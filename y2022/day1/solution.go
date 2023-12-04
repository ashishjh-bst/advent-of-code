package day1

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	elfCals := calSum(input)
	return strconv.Itoa((*elfCals)[0])
}

func Part2(input *string) string {
	elfCals := calSum(input)
	return strconv.Itoa((*elfCals)[0] + (*elfCals)[1] + (*elfCals)[2])
}

func calSum(input *string) *[]int {
	foodItems := strings.Split(*input, "\n")
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
	return &elfCals
}
