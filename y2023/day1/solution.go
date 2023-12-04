package day1

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	lines := strings.Split(*input, "\n")
	numRegex := regexp.MustCompile("[0-9]")
	sum := 0
	for _, line := range lines {
		lineNums := numRegex.FindAllString(line, -1)
		lineValue, err := strconv.Atoi(fmt.Sprintf("%s%s", lineNums[0], lineNums[len(lineNums)-1]))
		if err != nil {
			log.Fatalf(err.Error())
		}
		sum += lineValue
	}
	return strconv.Itoa(sum)
}

func Part2(input *string) string {
	lines := strings.Split(*input, "\n")

	sum := 0
	for _, line := range lines {
		lineNums := getDigits(line)
		first := lineNums[0]
		last := lineNums[len(lineNums)-1]
		lineValue := (first * 10) + last
		sum += lineValue
	}
	return strconv.Itoa(sum)
}

func getDigits(line string) []int {
	word2Num := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	numRegex := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|[0-9])")
	var nums []int
	for i := 0; i < len(line); i++ {
		part := line[i:]
		indexes := numRegex.FindStringIndex(part)
		if indexes == nil {
			return nums
		}
		num := part[indexes[0]:indexes[1]]
		if val, ok := word2Num[num]; ok {
			num = val
		}
		iNum, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, iNum)
		i += indexes[0]
	}
	return nums
}
