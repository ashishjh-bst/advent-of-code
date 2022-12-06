package day6

import (
	"strconv"
)

func Part1(input *string) string {
	markerSize := 4
	for i := 0; i < len(*input); i++ {
		if i+(markerSize-1) >= len(*input) {
			return "0"
		}
		marker := make(map[rune]int, 0)
		for j := 0; j < markerSize; j++ {
			char := rune((*input)[i+j])
			marker[char]++
			if marker[char] > 1 {
				break
			}
		}
		if len(marker) == markerSize {
			return strconv.Itoa(i + markerSize)
		}
	}
	return "0"
}

func Part2(input *string) string {
	markerSize := 14
	for i := 0; i < len(*input); i++ {
		if i+markerSize-1 >= len(*input) {
			return "0"
		}
		marker := make(map[rune]int, 0)
		for j := 0; j < markerSize; j++ {
			char := rune((*input)[i+j])
			marker[char]++
			if marker[char] > 1 {
				break
			}
		}
		if len(marker) == markerSize {
			return strconv.Itoa(i + markerSize)
		}
	}
	return "0"
}
