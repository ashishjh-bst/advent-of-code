package day3

import (
	"strconv"
)

func Part1(input *string) string {
	locations := make(map[[2]int]int, 0)
	pos := [2]int{0, 0}
	locations[pos]++
	for _, c := range *input {
		switch c {
		case '^':
			pos[1]++
		case 'v':
			pos[1]--
		case '>':
			pos[0]++
		case '<':
			pos[0]--
		}
		locations[pos]++
	}

	return strconv.Itoa(len(locations))
}

func Part2(input *string) string {
	locations := make(map[[2]int]int, 0)
	santaPos := [2]int{0, 0}
	roboSantaPos := [2]int{0, 0}
	locations[santaPos]++
	locations[roboSantaPos]++
	for index, c := range *input {
		var mover *[2]int
		if index%2 == 0 {
			mover = &santaPos
		} else {
			mover = &roboSantaPos
		}
		switch c {
		case '^':
			mover[1]++
		case 'v':
			mover[1]--
		case '>':
			mover[0]++
		case '<':
			mover[0]--
		}
		locations[*mover]++
	}
	return strconv.Itoa(len(locations))
}
