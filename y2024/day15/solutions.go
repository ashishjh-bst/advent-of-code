package day15

import (
	"fmt"
	"strings"
)

type Grid [][]string

func Part1(input *string) string {
	grid, moves, pos := parseInput(input, false)
	for _, move := range *moves {
		dir := getDir(move)
		if grid.move(pos, dir) {
			pos = getNext(pos, dir)
		}
	}
	return fmt.Sprintf("%d", grid.getBoxSum("O"))
}

func Part2(input *string) string {
	grid, moves, pos := parseInput(input, true)
	for _, move := range *moves {
		dir := getDir(move)
		if grid.move2(pos, move) {
			pos = getNext(pos, dir)
		}
	}
	return fmt.Sprintf("%d", grid.getBoxSum("["))
}

func parseInput(input *string, big bool) (*Grid, *string, [2]int) {
	values := strings.Split(*input, "\n\n")
	if big {
		grid, start := makeBigGrid(&values[0])
		return grid, &values[1], start
	}
	grid, start := makeGrid(&values[0])
	return grid, &values[1], start
}

func makeGrid(input *string) (*Grid, [2]int) {
	lines := strings.Split(*input, "\n")
	grid := make(Grid, len(lines))
	var pos [2]int
	for i, line := range lines {
		cols := strings.Split(line, "")
		grid[i] = cols
		if j := strings.Index(line, "@"); j != -1 {
			pos = [2]int{i, j}
		}
	}
	return &grid, pos
}

func makeBigGrid(input *string) (*Grid, [2]int) {
	lines := strings.Split(*input, "\n")
	grid := make(Grid, len(lines))
	var pos [2]int
	for i, line := range lines {
		row := make([]string, 0, len(line)*2)
		for j, col := range line {
			switch col {
			case '#':
				row = append(row, "#", "#")
			case 'O':
				row = append(row, "[", "]")
			case '.':
				row = append(row, ".", ".")
			case '@':
				row = append(row, "@", ".")
				pos = [2]int{i, j * 2}
			}
		}
		grid[i] = row
	}
	return &grid, pos
}

func (g *Grid) move(pos [2]int, dir [2]int) bool {
	next := getNext(pos, dir)
	if next == pos || g.isOutside(next) || (*g)[next[0]][next[1]] == "#" {
		return false
	}
	if (*g)[next[0]][next[1]] == "." || g.move(next, dir) {
		(*g)[next[0]][next[1]] = (*g)[pos[0]][pos[1]]
		(*g)[pos[0]][pos[1]] = "."
		return true
	}
	return false
}

func (g *Grid) move2(pos [2]int, move rune) bool {
	dir := getDir(move)
	next := getNext(pos, dir)
	if g.isBlocked(next) {
		return false
	}
	if g.getBlock(next) == "." {
		return g.move(pos, dir)
	}
	box := g.getBox(next)
	return g.moveBox(box, move) && g.move(pos, dir)
}

func (g *Grid) moveBox(box [2][2]int, move rune) bool {
	if move == '<' || move == '>' {
		return g.moveBoxH(box, move)
	}
	return g.moveBoxV(box, move)
}

func (g *Grid) moveBoxH(box [2][2]int, move rune) bool {
	if !g.isBoxMovableH(box, move) {
		return false
	}
	dir := getDir(move)
	next := getNext(box[1], dir)
	if move == '>' {
		if g.getBlock(next) == "." {
			return g.move(box[1], dir) && g.move(box[0], dir)
		}
		if g.getBlock(next) == "[" {
			nextBox := g.getBox(getNext(next, dir))
			return g.moveBox(nextBox, move) && g.move(box[1], dir) && g.move(box[0], dir)
		}
	} else if move == '<' {
		next = getNext(box[0], dir)
		if g.getBlock(next) == "." {
			return g.move(box[0], dir) && g.move(box[1], dir)
		}
		if g.getBlock(next) == "]" {
			nextBox := g.getBox(getNext(next, dir))
			return g.moveBox(nextBox, move) && g.move(box[0], dir) && g.move(box[1], dir)
		}
	}
	return false
}

func (g *Grid) isBoxMovableH(box [2][2]int, move rune) bool {
	dir := getDir(move)
	next := getNext(box[1], dir)
	if move == '>' {
		if g.isBlocked(next) {
			return false
		}
		if g.getBlock(next) == "." {
			return true
		}
		if g.getBlock(next) == "[" {
			nextBox := g.getBox(getNext(next, dir))
			return g.isBoxMovableH(nextBox, move)
		}
	} else if move == '<' {
		next = getNext(box[0], dir)
		if g.isBlocked(next) {
			return false
		}
		if g.getBlock(next) == "." {
			return true
		}
		if g.getBlock(next) == "]" {
			nextBox := g.getBox(getNext(next, dir))
			return g.isBoxMovableH(nextBox, move)
		}
	}
	return false
}

func (g *Grid) moveBoxV(box [2][2]int, move rune) bool {
	if !g.isBoxMovableV(box, move) {
		return false
	}
	dir := getDir(move)
	nextStart := getNext(box[0], dir)
	nextEnd := getNext(box[1], dir)
	if g.getBlock(nextStart) == "." && g.getBlock(nextEnd) == "." {
		return g.move(box[0], dir) && g.move(box[1], dir)
	}
	if g.getBlock(nextStart)+g.getBlock(nextEnd) == "[]" {
		nextBox := g.getBox(nextStart)
		return g.isBoxMovableV(nextBox, move) && g.moveBoxV(nextBox, move) && g.move(box[0], dir) && g.move(box[1], dir)
	}
	if g.getBlock(nextStart)+g.getBlock(nextEnd) == "]." {
		nextBox := g.getBox(getNext(nextStart, [2]int{0, -1}))
		return g.isBoxMovableV(nextBox, move) && g.moveBoxV(nextBox, move) && g.move(box[0], dir) && g.move(box[1], dir)
	}
	if g.getBlock(nextStart)+g.getBlock(nextEnd) == ".[" {
		nextBox := g.getBox(getNext(nextEnd, [2]int{0, 1}))
		return g.isBoxMovableV(nextBox, move) && g.moveBoxV(nextBox, move) && g.move(box[0], dir) && g.move(box[1], dir)
	}
	if g.getBlock(nextStart)+g.getBlock(nextEnd) == "][" {
		box1 := g.getBox(getNext(nextStart, [2]int{0, -1}))
		box2 := g.getBox(getNext(nextEnd, [2]int{0, 1}))
		return g.isBoxMovableV(box1, move) && g.isBoxMovableV(box2, move) && g.moveBoxV(box1, move) && g.moveBoxV(box2, move) && g.move(box[0], dir) && g.move(box[1], dir)
	}
	return false
}

func (g *Grid) isBoxMovableV(box [2][2]int, move rune) bool {
	dir := getDir(move)
	nextStart := getNext(box[0], dir)
	nextEnd := getNext(box[1], dir)
	if g.isBlocked(nextStart) || g.isBlocked(nextEnd) {
		return false
	}
	if g.getBlock(nextStart) == "." && g.getBlock(nextEnd) == "." {
		return true
	}
	if g.getBlock(nextStart)+g.getBlock(nextEnd) == "[]" {
		nextBox := g.getBox(nextStart)
		return g.isBoxMovableV(nextBox, move)
	}
	if g.getBlock(nextStart)+g.getBlock(nextEnd) == "]." {
		nextBox := g.getBox(getNext(nextStart, [2]int{0, -1}))
		return g.isBoxMovableV(nextBox, move)
	}
	if g.getBlock(nextStart)+g.getBlock(nextEnd) == ".[" {
		nextBox := g.getBox(getNext(nextEnd, [2]int{0, 1}))
		return g.isBoxMovableV(nextBox, move)
	}
	if g.getBlock(nextStart)+g.getBlock(nextEnd) == "][" {
		box1 := g.getBox(getNext(nextStart, [2]int{0, -1}))
		box2 := g.getBox(getNext(nextEnd, [2]int{0, 1}))
		return g.isBoxMovableV(box1, move) && g.isBoxMovableV(box2, move)
	}
	return false
}

func (g *Grid) getBlock(pos [2]int) string {
	return (*g)[pos[0]][pos[1]]
}

func (g *Grid) isBlocked(pos [2]int) bool {
	return g.isOutside(pos) || g.getBlock(pos) == "#"
}

func getNext(pos [2]int, dir [2]int) [2]int {
	return [2]int{pos[0] + dir[0], pos[1] + dir[1]}
}

func (g *Grid) isOutside(pos [2]int) bool {
	return pos[0] < 0 || pos[0] >= len(*g) || pos[1] < 0 || pos[1] >= len((*g)[0])
}

func (g *Grid) getBox(pos [2]int) [2][2]int {
	if (*g)[pos[0]][pos[1]] == "[" {
		return [2][2]int{pos, {pos[0], pos[1] + 1}}
	}
	return [2][2]int{{pos[0], pos[1] - 1}, pos}
}

func (g *Grid) getBoxSum(boxType string) int {
	var sum int
	for i, row := range *g {
		for j, col := range row {
			if col == boxType {
				sum += i*100 + j
			}
		}
	}
	return sum
}

func getDir(c rune) [2]int {
	switch c {
	case '^':
		return [2]int{-1, 0}
	case '>':
		return [2]int{0, 1}
	case 'v':
		return [2]int{1, 0}
	case '<':
		return [2]int{0, -1}
	case '\n':
		return [2]int{0, 0}
	}
	fmt.Printf("THIS SHOULD NEVER HAPPEN")
	return [2]int{}
}
