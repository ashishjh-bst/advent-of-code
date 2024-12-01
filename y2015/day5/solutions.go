package day5

import (
	"strconv"
	"strings"
)

func Part1(input *string) string {
	words := strings.Split(*input, "\n")
	count := 0
	for _, word := range words {
		if hasVowels(word) && hasDoubleLetter(word) && !hasForbiddenStrings(word) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func Part2(input *string) string {
	words := strings.Split(*input, "\n")
	count := 0
	for _, word := range words {
		if hasRepeatingLetterWithGap(word) && hasRepeatingNonOverlappingPairs(word) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func hasRepeatingNonOverlappingPairs(word string) bool {
	for i := 0; i < len(word)-3; i++ {
		if strings.Contains(word[i+2:], word[i:i+2]) {
			return true
		}
	}
	return false
}

func hasRepeatingLetterWithGap(word string) bool {
	for i := 0; i < len(word)-2; i++ {
		if word[i] == word[i+2] {
			return true
		}
	}
	return false
}

func hasVowels(word string) bool {
	vowels := "aeiou"
	count := 0
	for _, c := range word {
		if strings.Contains(vowels, string(c)) {
			count++
		}
	}
	return count >= 3
}

func hasDoubleLetter(word string) bool {
	for i := 0; i < len(word)-1; i++ {
		if word[i] == word[i+1] {
			return true
		}
	}
	return false
}

func hasForbiddenStrings(word string) bool {
	for _, s := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(word, s) {
			return true
		}
	}
	return false
}
