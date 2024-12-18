package day15

import (
	"fmt"
	"strings"
)

type Grid [][]string

func Part1(input *string) string {
	grid, moves, pos := parseInput(input)
	//grid.print()
	for _, move := range *moves {
		dir := getDir(move)
		hasMoved := grid.move(pos, dir)
		if hasMoved {
			pos = [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		}
		// grid.print()
	}
	return fmt.Sprintf("%d", grid.getBoxSum())
}

func Part2(input *string) string {
	grid, moves, pos := parseInput2(input)
	//grid.print()
	for _, move := range *moves {
		dir := getDir(move)
		//fmt.Printf("\nPos:%v ,Move: %s, dir: %v", pos, string(move), dir)
		hasMoved := grid.move2(pos, move)

		//fmt.Printf(" hasMoved: %t", hasMoved)
		if hasMoved {
			pos = [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		}
		//grid.print()
	}
	return fmt.Sprintf("%d", grid.getBigBoxSum())
}

func parseInput(input *string) (*Grid, *string, [2]int) {
	values := strings.Split(*input, "\n\n")
	grid, start := makeGrid(&values[0])
	moves := &values[1]
	return grid, moves, start
}

func parseInput2(input *string) (*Grid, *string, [2]int) {
	values := strings.Split(*input, "\n\n")
	grid, start := makeBigGrid(&values[0])
	moves := &values[1]
	return grid, moves, start
}

func makeGrid(input *string) (*Grid, [2]int) {
	grid := make(Grid, 0)
	lines := strings.Split(*input, "\n")
	var pos [2]int
	for i, line := range lines {
		cols := strings.Split(line, "")
		grid = append(grid, cols)
		j := strings.Index(line, "@")
		if j != -1 {
			pos = [2]int{i, j}
		}
	}
	return &grid, pos
}

func makeBigGrid(input *string) (*Grid, [2]int) {
	grid := make(Grid, 0)
	lines := strings.Split(*input, "\n")
	var pos [2]int
	for i, line := range lines {
		cols := strings.Split(line, "")
		widthCounter := 0
		row := make([]string, 0)
		for _, col := range cols {
			switch col {
			case "#":
				row = append(row, "#", "#")
			case "O":
				row = append(row, "[", "]")
			case ".":
				row = append(row, ".", ".")
			case "@":
				row = append(row, "@", ".")
				pos = [2]int{i, widthCounter}
			}
			widthCounter += 2
		}
		grid = append(grid, row)
	}
	return &grid, pos
}

func (g *Grid) move(pos [2]int, dir [2]int) bool {
	next := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
	if next == pos {
		return false
	}
	if g.isOutside(next) {
		return false
	}
	if (*g)[next[0]][next[1]] == "#" {
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
	nextBlock := g.getBlock(next)
	if nextBlock == "." {
		return g.move(pos, dir)
	}
	var box = g.getBox(next)
	return g.isBoxMovable(box, move) && g.moveBox(box, move) && g.move(pos, dir)
}

func (g *Grid) isBoxMovableV(box [2][2]int, move rune) bool {
	dir := getDir(move)
	nextStart := getNext(box[0], dir)
	nextEnd := getNext(box[1], dir)
	if g.isBlocked(nextStart) || g.isBlocked(nextEnd) {
		return false
	}
	nextStartBlock := g.getBlock(nextStart)
	nextEndBlock := g.getBlock(nextEnd)
	if nextStartBlock == "." && nextEndBlock == "." {
		return true
	}
	if nextStartBlock+nextEndBlock == "[]" {
		nextBox := g.getBox(nextStart)
		return g.isBoxMovableV(nextBox, move)
	}
	if nextStartBlock+nextEndBlock == "]." {
		nextBox := g.getBox([2]int{nextStart[0], nextStart[1] - 1})
		return g.isBoxMovableV(nextBox, move)
	}
	if nextStartBlock+nextEndBlock == ".[" {
		nextBox := g.getBox([2]int{nextEnd[0], nextEnd[1] + 1})
		return g.isBoxMovableV(nextBox, move)
	}
	if nextStartBlock+nextEndBlock == "][" {
		box1 := g.getBox([2]int{nextStart[0], nextStart[1] - 1})
		box2 := g.getBox([2]int{nextEnd[0], nextEnd[1] + 1})
		return g.isBoxMovableV(box1, move) && g.isBoxMovableV(box2, move)
	}
	return false
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

	if move == '>' {
		next := getNext(box[1], dir)
		nextBlock := g.getBlock(next)
		if nextBlock == "." {
			return g.move(box[1], dir) && g.move(box[0], dir)
		}
		if nextBlock == "[" {
			nextBox := g.getBox([2]int{next[0], next[1] + 1})
			return g.moveBox(nextBox, move) && g.move(box[1], dir) && g.move(box[0], dir)
		}
	} else if move == '<' {
		next := getNext(box[0], dir)
		nextBlock := g.getBlock(next)
		if nextBlock == "." {
			return g.move(box[0], dir) && g.move(box[1], dir)
		}
		if nextBlock == "]" {
			nextBox := g.getBox([2]int{next[0], next[1] - 1})
			return g.moveBox(nextBox, move) && g.move(box[0], dir) && g.move(box[1], dir)
		}
	}
	return false
}

func (g *Grid) isBoxMovableH(box [2][2]int, move rune) bool {
	dir := getDir(move)

	if move == '>' {
		next := getNext(box[1], dir)
		if g.isBlocked(next) {
			return false
		}
		nextBlock := g.getBlock(next)
		if nextBlock == "." {
			return true
		}
		if nextBlock == "[" {
			nextBox := g.getBox([2]int{next[0], next[1] + 1})
			return g.isBoxMovableH(nextBox, move)
		}
	} else if move == '<' {
		next := getNext(box[0], dir)
		if g.isBlocked(next) {
			return false
		}
		nextBlock := g.getBlock(next)
		if nextBlock == "." {
			return true
		}
		if nextBlock == "]" {
			nextBox := g.getBox([2]int{next[0], next[1] - 1})
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
	nextStartBlock := g.getBlock(nextStart)
	nextEndBlock := g.getBlock(nextEnd)
	if nextStartBlock == "." && nextEndBlock == "." {
		return g.move(box[0], dir) && g.move(box[1], dir)
	}
	if nextStartBlock+nextEndBlock == "[]" {
		nextBox := g.getBox(nextStart)
		return g.isBoxMovableV(nextBox, move) && g.moveBoxV(nextBox, move) && g.move(box[0], dir) && g.move(box[1], dir)
	}
	if nextStartBlock+nextEndBlock == "]." {
		nextBox := g.getBox([2]int{nextStart[0], nextStart[1] - 1})
		return g.isBoxMovableV(nextBox, move) && g.moveBoxV(nextBox, move) && g.move(box[0], dir) && g.move(box[1], dir)
	}
	if nextStartBlock+nextEndBlock == ".[" {
		nextBox := g.getBox([2]int{nextEnd[0], nextEnd[1] + 1})
		return g.isBoxMovableV(nextBox, move) && g.moveBoxV(nextBox, move) && g.move(box[0], dir) && g.move(box[1], dir)
	}
	if nextStartBlock+nextEndBlock == "][" {
		box1 := g.getBox([2]int{nextStart[0], nextStart[1] - 1})
		box2 := g.getBox([2]int{nextEnd[0], nextEnd[1] + 1})
		return g.isBoxMovableV(box1, move) && g.isBoxMovableV(box2, move) && g.moveBoxV(box1, move) && g.moveBoxV(box2, move) && g.move(box[0], dir) && g.move(box[1], dir)
	}
	return false
}

func (g *Grid) isBoxMovable(box [2][2]int, move rune) bool {
	if move == '<' || move == '>' {
		return g.isBoxMovableH(box, move)
	}
	return g.isBoxMovableV(box, move)
}

func (g *Grid) getBlock(pos [2]int) string {
	return (*g)[pos[0]][pos[1]]
}

func (g *Grid) isBlocked(pos [2]int) bool {
	if g.isOutside(pos) {
		return true
	}
	if g.getBlock(pos) == "#" {
		return true
	}
	return false
}

func getNext(pos [2]int, dir [2]int) [2]int {
	return [2]int{pos[0] + dir[0], pos[1] + dir[1]}
}

func (g *Grid) isOutside(pos [2]int) bool {
	if pos[0] < 0 || pos[0] >= len(*g) || pos[1] < 0 || pos[1] >= len((*g)[0]) {
		return true
	}
	return false
}

func (g *Grid) getBox(pos [2]int) [2][2]int {
	if (*g)[pos[0]][pos[1]] == "[" {
		return [2][2]int{pos, {pos[0], pos[1] + 1}}
	}
	return [2][2]int{{pos[0], pos[1] - 1}, pos}
}

func (g *Grid) print() {
	for _, row := range *g {
		fmt.Printf("\n %v", row)
	}
}

func (g *Grid) getBoxSum() int {
	var sum int
	for i, row := range *g {
		for j, col := range row {
			if col == "O" {
				sum += i*100 + j
			}
		}
	}
	return sum
}

func (g *Grid) getBigBoxSum() int {
	var sum int
	for i, row := range *g {
		for j, col := range row {
			if col == "[" {
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
