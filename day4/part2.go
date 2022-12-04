package day4

import (
	"log"
	"strings"
)

func Part2(input string) {
	elfPairs := strings.Split(input, "\n")
	overlaps := 0
	for _, pair := range elfPairs {
		elfs := strings.Split(pair, ",")
		elf1 := strToElf(elfs[0])
		elf2 := strToElf(elfs[1])
		// check if the elfs don't overlap, if they don't not overlap, it's an overlap??
		if !(elf1.end < elf2.start || elf1.start > elf2.end) {
			overlaps++
		}
	}
	log.Printf("Total Overlaps %d", overlaps)
}
