package main

import (
	"fmt"
	"os"

	"../common"
)

func main() {
	args := os.Args[1:]
	inputFilePath := args[1]
	input := common.ReadFileInput(inputFilePath)
	fmt.Printf(input)
}
