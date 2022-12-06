package day6

import (
	"strconv"
)

func Part1(input *string) string {
	for i := 0; i < len(*input); i++ {
		if i+3 >= len(*input) {
			return "0"
		}
		marker := make(map[rune]int, 0)
		for j := 0; j < 4; j++ {
			char := rune((*input)[i+j])
			marker[char]++
			if marker[char] > 1 {
				break
			}
		}
		if len(marker) == 4 {
			return strconv.Itoa(i + 4)
		}
	}
	return "0"
}

func Part2(input *string) string {
	for i := 0; i < len(*input); i++ {
		if i+13 >= len(*input) {
			return "0"
		}
		marker := make(map[rune]int, 0)
		for j := 0; j < 14; j++ {
			char := rune((*input)[i+j])
			marker[char]++
			if marker[char] > 1 {
				break
			}
		}
		if len(marker) == 14 {
			return strconv.Itoa(i + 14)
		}
	}
	return "0"
}
