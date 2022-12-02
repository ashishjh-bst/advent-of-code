package common

import (
	"fmt"
	"io"
	"os"
)

func ReadFileInput(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s, error: %s", path, err.Error())
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s, error: %s", path, err.Error())
	}
	return string(data), nil
}
