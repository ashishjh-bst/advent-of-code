package y2022

import (
	"fmt"
	"log"
	"os"

	"github.com/ashishjh-bst/aoc/y2022/day1"
	"github.com/ashishjh-bst/aoc/y2022/day10"
	"github.com/ashishjh-bst/aoc/y2022/day11"
	"github.com/ashishjh-bst/aoc/y2022/day12"
	"github.com/ashishjh-bst/aoc/y2022/day13"
	"github.com/ashishjh-bst/aoc/y2022/day2"
	"github.com/ashishjh-bst/aoc/y2022/day3"
	"github.com/ashishjh-bst/aoc/y2022/day4"
	"github.com/ashishjh-bst/aoc/y2022/day5"
	"github.com/ashishjh-bst/aoc/y2022/day6"
	"github.com/ashishjh-bst/aoc/y2022/day7"
	"github.com/ashishjh-bst/aoc/y2022/day8"
	"github.com/ashishjh-bst/aoc/y2022/day9"
)

func Calculate(day int, part int, input *string) (string, error) {
	args := os.Args[1:]
	type Part func(input *string) string
	type Day map[int]Part

	// map of solutions
	Days := map[int]Day{
		1:  {1: day1.Part1, 2: day1.Part2},
		2:  {1: day2.Part1, 2: day2.Part2},
		3:  {1: day3.Part1, 2: day3.Part2},
		4:  {1: day4.Part1, 2: day4.Part2},
		5:  {1: day5.Part1, 2: day5.Part2},
		6:  {1: day6.Part1, 2: day6.Part2},
		7:  {1: day7.Part1, 2: day7.Part2},
		8:  {1: day8.Part1, 2: day8.Part2},
		9:  {1: day9.Part1, 2: day9.Part2},
		10: {1: day10.Part1, 2: day10.Part2},
		11: {1: day11.Part1, 2: day11.Part2},
		12: {1: day12.Part1, 2: day12.Part2},
		13: {1: day13.Part1, 2: day13.Part2},
	}

	if Days[day] == nil {
		return "", fmt.Errorf("invalid Day %s, Either not created or doesn't exist", args[0])
	}

	if part < 1 || part > 2 {
		return "", fmt.Errorf("invalid Part %s, should be 1 or 2", args[1])
	}

	log.Printf("\nExecuting y2022 Day %d Part %d", day, part)
	answer := Days[day][part](input)
	return answer, nil
}
