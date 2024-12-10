package day8

import (
	"fmt"
	"strings"
)

func Part1(input *string) string {
	lines := strings.Split(*input, "\n")
	storageSum := 0
	memorySum := 0
	for _, line := range lines {
		storageSum += len(line)
		memorySum += getLineCount(line)
	}
	return fmt.Sprintf("%d", storageSum-memorySum)
}

func Part2(input *string) string {
	lines := strings.Split(*input, "\n")
	storageSum := 0
	memorySum := 0
	for _, line := range lines {
		storageSum += len(line)
		memorySum += getLineCountAfterEncoding(line)
	}
	return fmt.Sprintf("%d", memorySum-storageSum)
}

func getLineCount(line string) (count int) {
	count = 0
	for i := 1; i < len(line)-1; i++ {
		count++
		if line[i] == '\\' {
			switch line[i+1] {
			case '\\', '"':
				i++
			case 'x':
				i = i + 3
			}
		}
	}
	return
}

func getLineCountAfterEncoding(line string) (count int) {
	count = 2
	for i := 0; i < len(line); i++ {
		count++
		switch line[i] {
		case '"', '\\':
			count++
		}
	}
	return
}
