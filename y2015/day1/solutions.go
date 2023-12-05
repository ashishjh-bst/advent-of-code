package day1

import "strconv"

func Part1(input *string) string {
	floor := 0
	for _, dir := range *input {
		if dir == '(' {
			floor++
		} else {
			floor--
		}
	}
	return strconv.Itoa(floor)
}

func Part2(input *string) string {
	floor := 0
	for pos, dir := range *input {
		if dir == '(' {
			floor++
		} else {
			floor--
		}
		if floor < 0 {
			floor = pos + 1
			break
		}
	}
	return strconv.Itoa(floor)
}
