package day7

import (
	"strconv"
	"strings"
)

type Directory struct {
	size        int
	name        string
	parent      *Directory
	directories []*Directory
}

func Part1(input *string) string {
	lines := strings.Split(*input, "\n")
	dir := &Directory{size: 0, name: "/", directories: make([]*Directory, 0)}
	fileSystemParser(lines, dir)
	total := getDirSumBelowSize(dir, 100000)
	return strconv.Itoa(total)
}

func Part2(input *string) string {
	lines := strings.Split(*input, "\n")
	dir := &Directory{size: 0, name: "/", directories: make([]*Directory, 0)}
	fileSystemParser(lines, dir)
	spaceNeeded := 30000000 - (70000000 - dir.size)
	dirs := getAllEligibleSubDir(dir, spaceNeeded)
	var smallestEligibeDir *Directory
	for _, subDir := range dirs {
		if smallestEligibeDir == nil {
			smallestEligibeDir = subDir
		}
		if subDir.size < smallestEligibeDir.size {
			smallestEligibeDir = subDir
		}
	}
	spaceFreed := strconv.Itoa(smallestEligibeDir.size)
	return spaceFreed
}

func fileSystemParser(lines []string, dir *Directory) *Directory {
	for index, line := range lines {
		command := strings.Split(line, " ")
		if command[0] == "$" {
			if command[1] == "cd" {
				if command[2] == "/" {
					continue
				}
				if command[2] == ".." {
					return fileSystemParser(lines[index+1:], dir.parent)
				} else {
					newDir := &Directory{size: 0, name: command[2], parent: dir, directories: make([]*Directory, 0)}
					dir.directories = append(dir.directories, newDir)
					return fileSystemParser(lines[index+1:], newDir)
				}
			}
		} else if command[0] != "dir" {
			size, _ := strconv.Atoi(command[0])
			updateDirSize(dir, size)
		}
	}
	return dir
}

func updateDirSize(dir *Directory, increment int) {
	if dir == nil {
		return
	}
	dir.size += increment
	updateDirSize(dir.parent, increment)
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

func getAllEligibleSubDir(dir *Directory, spaceNeeded int) []*Directory {
	dirs := make([]*Directory, 0)
	if dir.size > spaceNeeded {
		dirs = append(dirs, dir)
	}

	for _, subDir := range dir.directories {
		if subDir.size < spaceNeeded {
			continue
		}
		dirs = append(dirs, getAllEligibleSubDir(subDir, spaceNeeded)...)
	}

	return dirs
}
