package day9

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	lines := strings.Split(*input, "\n")
	sum := 0
	for _, sStr := range lines {
		s := IntSlice(strings.Split(sStr, " "))
		v := diffSlice(s)
		fmt.Printf("\n %d", v)
		sum += v
	}
	return strconv.Itoa(sum)
}

func Part2(input *string) string {
	lines := strings.Split(*input, "\n")
	sum := 0
	for _, sStr := range lines {
		s := IntSlice(strings.Split(sStr, " "))
		v := diffSlice2(s)
		sum += v
	}
	return strconv.Itoa(sum)
}

func IntSlice(s []string) []int {
	var r []int
	for _, i := range s {
		c, _ := strconv.Atoi(i)
		r = append(r, c)
	}
	return r
}

func diffSlice(s []int) int {
	var r []int
	isAllZero := true
	for i := 0; i < len(s)-1; i++ {
		d := s[i+1] - s[i]
		if d != 0 {
			isAllZero = false
		}
		r = append(r, d)
	}
	if isAllZero {
		return s[len(s)-1]
	}

	return s[len(s)-1] + diffSlice(r)
}

func diffSlice2(s []int) int {
	var r []int
	isAllZero := true
	for i := 0; i < len(s)-1; i++ {
		d := s[i+1] - s[i]
		if d != 0 {
			isAllZero = false
		}
		r = append(r, d)
	}
	if isAllZero {
		return s[0]
	}

	return s[0] - diffSlice2(r)
}
