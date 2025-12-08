package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	ranges := strings.Split(*input, ",")
	sum := 0
	for _, r := range ranges {
		ends := strings.Split(r, "-")
		start, _ := strconv.Atoi(ends[0])
		end, _ := strconv.Atoi(ends[1])
		for i := start; i <= end; i++ {
			iStr := fmt.Sprintf("%d", i)
			iLen := len(iStr)
			if iLen%2 != 0 {
				continue
			}
			mid := (iLen / 2)
			first := iStr[0:mid]
			last := iStr[mid:iLen]
			if first == last {
				sum += i
			}
		}
	}
	return fmt.Sprintf("%d", sum)
}

func Part2(input *string) string {
	ranges := strings.Split(*input, ",")
	sum := 0
	for _, r := range ranges {
		ends := strings.Split(r, "-")
		start, _ := strconv.Atoi(ends[0])
		end, _ := strconv.Atoi(ends[1])
		for i := start; i <= end; i++ {
			iStr := fmt.Sprintf("%d", i)
			iLen := len(iStr)
			mid := (iLen / 2)
			isInvalid := false
			for j := 1; j < mid+1; j++ {
				s := iStr[0:j]
				joined := strings.Join(strings.Split(iStr, s), "")
				if joined == "" {
					isInvalid = true
					break
				}
			}
			if isInvalid {
				sum += i
			}
		}
	}
	return fmt.Sprintf("%d", sum)
}
