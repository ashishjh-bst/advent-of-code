package day2

import (
	"strconv"
	"strings"
)

func Part1(input *string) string {
	gifts := strings.Split(*input, "\n")
	sum := 0
	for _, gift := range gifts {
		sides := strings.Split(gift, "x")
		l, _ := strconv.Atoi(sides[0])
		b, _ := strconv.Atoi(sides[1])
		h, _ := strconv.Atoi(sides[2])
		sum += 2*(l*b+b*h+h*l) + min(l*b, b*h, h*l)
	}
	return strconv.Itoa(sum)
}

func Part2(input *string) string {
	gifts := strings.Split(*input, "\n")
	sum := 0
	for _, gift := range gifts {
		sides := strings.Split(gift, "x")
		l, _ := strconv.Atoi(sides[0])
		b, _ := strconv.Atoi(sides[1])
		h, _ := strconv.Atoi(sides[2])
		sum += 2*min(l+b, l+h, b+h) + (l * b * h)
	}
	return strconv.Itoa(sum)
}
