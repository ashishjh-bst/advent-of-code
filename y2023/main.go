package y2023

import (
	"fmt"
	"log"
	"os"

	"github.com/ashishjh-bst/aoc/y2023/day1"
)

func Calculate(day int, part int, input *string) (string, error) {
	args := os.Args[1:]
	type Part func(input *string) string
	type Day map[int]Part

	// map of solutions
	Days := map[int]Day{
		1: {1: day1.Part1, 2: day1.Part2},
	}

	if Days[day] == nil {
		return "", fmt.Errorf("invalid Day %s, Either not created or doesn't exist", args[0])
	}

	if part < 1 || part > 2 {
		return "", fmt.Errorf("invalid Part %s, should be 1 or 2", args[1])
	}

	log.Printf("\nExecuting y2023 Day %d Part %d", day, part)
	answer := Days[day][part](input)
	return answer, nil
}