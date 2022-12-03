package day3

import (
	"log"
	"reflect"
	"strings"
)

func Part2(input string) {
	rucksacks := strings.Split(input, "\n")
	type GiftMap map[rune]int
	prioritySum := 0
	for g := 0; g < len(rucksacks); g = g + 3 {
		elf1 := rucksacks[g]
		elf2 := rucksacks[g+1]
		elf3 := rucksacks[g+2]
		elfCommon := make(GiftMap, 0)
		for _, item := range elf1 {
			elfCommon[item] = 1
		}
		for _, item := range elf2 {
			_, ok := elfCommon[item]
			if ok {
				elfCommon[item] = 2
			}
		}
		for _, item := range elf3 {
			val, ok := elfCommon[item]
			if ok && val == 2 {
				elfCommon[item] = 3
			}
		}
		iter := reflect.ValueOf(elfCommon).MapRange()
		for iter.Next() {
			if iter.Value().Int() < 3 {
				continue
			}
			gift := iter.Key().Int()
			if gift >= 'a' && gift <= 'z' {
				prioritySum += int(gift - 'a' + 1)
			} else if gift >= 'A' && gift <= 'Z' {
				prioritySum += int(gift - 'A' + 27)
			}
		}
	}
	log.Printf("Priority Sum is %d", prioritySum)
}
