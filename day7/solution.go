package day7

import (
	"fmt"
	"strconv"
	"strings"
)

type Directory struct {
	size        int
	name        string
	parent      *Directory
	directories map[string]*Directory
}

func Part1(input *string) string {
	lines := strings.Split(*input, "\n")
	dir := &Directory{size: 0, name: "/", directories: make(map[string]*Directory, 0)}
	fileSystemParser(lines, dir)
	total := getDirSumBelowSize(dir, 100000)
	return strconv.Itoa(total)
}

func Part2(input *string) string {
	return ""
}

func fileSystemParser(lines []string, dir *Directory) *Directory {
	for index, line := range lines {
		command := strings.Split(line, " ")
		fmt.Println(command)
		if command[0] == "$" {
			if command[1] == "ls" {
				continue
			} else if command[1] == "cd" {
				if command[2] == "/" {
					continue
				}
				if command[2] == ".." {
					return fileSystemParser(lines[index+1:], dir.parent)
				} else {
					newDir := &Directory{size: 0, name: command[2], parent: dir, directories: make(map[string]*Directory, 0)}
					dir.directories[command[2]] = newDir
					return fileSystemParser(lines[index+1:], newDir)
				}
			}
		} else if command[0] != "dir" {
			size, _ := strconv.Atoi(command[0])
			updateSize(dir, size)
		}
	}
	return dir
}

func updateSize(dir *Directory, increment int) {
	if dir == nil {
		return
	}
	dir.size += increment
	updateSize(dir.parent, increment)
}

func getDirSumBelowSize(dir *Directory, max int) int {
	total := 0
	if dir.size < max {
		total += dir.size
	}
	for _, subDir := range dir.directories {
		total += getDirSumBelowSize(subDir, max)
	}
	return total
}
