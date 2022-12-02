package common

import (
	"io"
	"log"
	"os"
)

func ReadFileInput(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed Reading file %s %s", path, err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed Reading file %s %s", path, err)
	}
	return string(data)
}
