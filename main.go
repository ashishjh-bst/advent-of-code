package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ashishjh-bst/aoc/common"
	"github.com/ashishjh-bst/aoc/y2015"
	"github.com/ashishjh-bst/aoc/y2022"
	"github.com/ashishjh-bst/aoc/y2023"
)

func main() {
	args := os.Args[1:]
	type Year func(day int, part int, input *string) (string, error)
	day, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("Invalid Day %s, should be an integer", args[1])
	}

	part, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatalf("Invalid Part %s, should be an integer", args[2])
	}

	Years := map[int]Year{
		2015: y2015.Calculate,
		2022: y2022.Calculate,
		2023: y2023.Calculate,
	}

	year, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Invalid Year %s, should be an integer", args[0])
	}
	if Years[year] == nil {
		log.Fatalf("Invalid Year %s, Either not created or doesn't exist", args[0])
	}

	log.Printf("\nExecuting Year %d Day %d Part %d", year, day, part)
	inputPath := fmt.Sprintf("./y%d/day%d/input.txt", year, day)
	input, err := common.ReadFileInput(inputPath)
	if err != nil {
		log.Fatalf("Failed Reading Input from %s", inputPath)
	}

	start := time.Now()
	answer, err := Years[year](day, part, &input)
	if err != nil {
		log.Fatal(err.Error())
	}
	elapsed := time.Since(start)

	log.Printf("\nExecution took %s", elapsed)
	log.Printf("\nThe answer is %s", answer)
}
