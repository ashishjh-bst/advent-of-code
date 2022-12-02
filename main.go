package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ashishjh-bst/aoc2022/common"
	"github.com/ashishjh-bst/aoc2022/day1"
	"github.com/ashishjh-bst/aoc2022/day2"
)

func main() {
	args := os.Args[1:]
	log.Printf("args %v", args)
	type Part func(input string)
	type Day map[int]Part
	Days := map[int]Day{
		1: {1: day1.Part1, 2: day1.Part2},
		2: {1: day2.Part1, 2: day2.Part2},
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Invalid Day %s, should be an integer", args[0])
	}
	if Days[day] == nil {
		log.Fatalf("Invalid Day %s, Either not created or doesn't exist", args[0])
	}
	part, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("Invalid Part %s, should be an integer", args[1])
	}
	if part < 1 || part > 2 {
		log.Fatalf("Invalid Part %s, should be 1 or 1", args[1])
	}
	input, err := common.ReadFileInput(fmt.Sprintf("./day%d/input-part%d.txt", day, part))
	if err != nil {
		log.Fatal(err)
	}
	Days[day][part](input)
}
