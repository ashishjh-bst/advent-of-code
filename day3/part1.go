package day3

import (
	"log"
	"strings"
)

func Part1(input string) {
	rucksacks := strings.Split(input, "\n")
	type GiftMap map[rune]int
	prioritySum := 0
	for _, rucksack := range rucksacks {
		size := len(rucksack)
		compartment1, compartment2 := getCompartments(rucksack)
		giftsComp1 := make(GiftMap, 0)
		giftsComp2 := make(GiftMap, 0)
		for i := 0; i < size/2; i++ {
			giftsComp1[compartment1[i]]++
			giftsComp2[compartment2[i]]++
		}
		for gift := range giftsComp1 {
			_, ok := giftsComp2[gift]
			if ok {
				if gift >= 'a' && gift <= 'z' {
					prioritySum += int(gift - 'a' + 1)
				} else if gift >= 'A' && gift <= 'Z' {
					prioritySum += int(gift - 'A' + 27)
				}
			}
		}
	}
	log.Printf("Priority Sum is %d", prioritySum)
}

func getCompartments(rucksack string) ([]rune, []rune) {
	mid := (len(rucksack) / 2)
	//end := (len(rucksack) / 2)
	compartment1 := []rune(rucksack[:mid])
	compartment2 := []rune(rucksack[mid:])
	return compartment1, compartment2
}
