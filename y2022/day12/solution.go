package day12

import (
	"fmt"
	"strconv"
	"strings"
)

type Pos struct {
	x        int
	y        int
	distance int
}

func Part1(input *string) string {
	grid, visted, start := makeGrid(input)
	distance := traverse(grid, visted, start)
	return strconv.Itoa(distance)
}

func Part2(input *string) string {
	grid, visted, start := makeGrid2(input)
	distance := traverse2(grid, visted, start)
	return strconv.Itoa(distance)
}

func makeGrid(input *string) (*[][]rune, *[][]bool, *Pos) {
	lines := strings.Split(*input, "\n")
	grid := make([][]rune, 0)
	visited := make([][]bool, 0)
	start := &Pos{}
	for x, line := range lines {
		row := make([]rune, 0)
		vr := make([]bool, 0)
		for y, char := range line {
			if char == 'S' {
				start.x = x
				start.y = y
				start.distance = 0
				row = append(row, 'a')
			} else {
				row = append(row, char)
			}
			vr = append(vr, false)
		}
		grid = append(grid, row)
		visited = append(visited, vr)
	}
	return &grid, &visited, start
}

func makeGrid2(input *string) (*[][]rune, *[][]bool, *Pos) {
	lines := strings.Split(*input, "\n")
	grid := make([][]rune, 0)
	visited := make([][]bool, 0)
	start := &Pos{}
	for x, line := range lines {
		row := make([]rune, 0)
		vr := make([]bool, 0)
		for y, char := range line {
			if char == 'E' {
				start.x = x
				start.y = y
				start.distance = 0
				row = append(row, 'z')
			} else {
				row = append(row, char)
			}
			vr = append(vr, false)
		}
		grid = append(grid, row)
		visited = append(visited, vr)
	}
	return &grid, &visited, start
}

func isNextAllowed(currentHeight rune, nextHeight rune) bool {
	if nextHeight == 'E' {
		nextHeight = 'z'
	}
	return nextHeight-currentHeight <= 1
}

func traverse(grid *[][]rune, visted *[][]bool, start *Pos) int {
	queue := &Queue{items: []*Pos{start}}
	for len(queue.items) != 0 {
		front := queue.front()
		if (*visted)[front.x][front.y] {
			continue
		}
		(*visted)[front.x][front.y] = true
		current := (*grid)[front.x][front.y]
		fmt.Printf("\n\nprocessing x:%d y:%d distance:%d  %s ", front.x, front.y, front.distance, string(current))
		if current == 'E' {
			return front.distance
		}

		left := &Pos{x: front.x - 1, y: front.y, distance: front.distance + 1}
		if left.x >= 0 {
			val := (*grid)[left.x][left.y]
			isVisited := (*visted)[left.x][left.y]
			fmt.Print("\nleft:", left, isVisited, isNextAllowed(current, val), string(val))
			if !isVisited && isNextAllowed(current, val) {
				fmt.Print("\ngoing left")
				queue.push(left)
			}
		}

		right := &Pos{x: front.x + 1, y: front.y, distance: front.distance + 1}
		if right.x < len(*grid) {
			val := (*grid)[right.x][right.y]
			isVisited := (*visted)[right.x][right.y]
			fmt.Print("\nright:", right, isVisited, isNextAllowed(current, val), string(val))
			if !isVisited && isNextAllowed(current, val) {
				fmt.Print("\ngoing right")
				queue.push(right)
			}
		}

		up := &Pos{x: front.x, y: front.y - 1, distance: front.distance + 1}
		if up.y >= 0 {
			val := (*grid)[up.x][up.y]
			isVisited := (*visted)[up.x][up.y]
			fmt.Print("\nup:", up, isVisited, isNextAllowed(current, val), string(val))
			if !isVisited && isNextAllowed(current, val) {
				fmt.Print("\ngoing up")
				queue.push(up)
			}
		}

		down := &Pos{x: front.x, y: front.y + 1, distance: front.distance + 1}
		if down.y < len((*grid)[down.x]) {
			val := (*grid)[down.x][down.y]
			isVisited := (*visted)[down.x][down.y]
			fmt.Print("\ndown:", down, isVisited, isNextAllowed(current, val), string(val))
			if !isVisited && isNextAllowed(current, val) {
				fmt.Print("\ngoing down")
				queue.push(down)
			}
		}
	}
	return -1
}

func isNextAllowed2(currentHeight rune, nextHeight rune) bool {
	if nextHeight == 'S' {
		nextHeight = 'a'
	}
	return currentHeight-nextHeight <= 1
}

func traverse2(grid *[][]rune, visted *[][]bool, start *Pos) int {
	queue := &Queue{items: []*Pos{start}}
	for len(queue.items) != 0 {
		front := queue.front()
		if (*visted)[front.x][front.y] {
			continue
		}
		(*visted)[front.x][front.y] = true
		current := (*grid)[front.x][front.y]
		fmt.Printf("\n\nprocessing x:%d y:%d distance:%d  %s ", front.x, front.y, front.distance, string(current))
		if current == 'a' {
			return front.distance
		}

		left := &Pos{x: front.x - 1, y: front.y, distance: front.distance + 1}
		if left.x >= 0 {
			val := (*grid)[left.x][left.y]
			isVisited := (*visted)[left.x][left.y]
			fmt.Print("\nleft:", left, isVisited, isNextAllowed2(current, val), string(val))
			if !isVisited && isNextAllowed2(current, val) {
				fmt.Print("\ngoing left")
				queue.push(left)
			}
		}

		right := &Pos{x: front.x + 1, y: front.y, distance: front.distance + 1}
		if right.x < len(*grid) {
			val := (*grid)[right.x][right.y]
			isVisited := (*visted)[right.x][right.y]
			fmt.Print("\nright:", right, isVisited, isNextAllowed2(current, val), string(val))
			if !isVisited && isNextAllowed2(current, val) {
				fmt.Print("\ngoing right")
				queue.push(right)
			}
		}

		up := &Pos{x: front.x, y: front.y - 1, distance: front.distance + 1}
		if up.y >= 0 {
			val := (*grid)[up.x][up.y]
			isVisited := (*visted)[up.x][up.y]
			fmt.Print("\nup:", up, isVisited, isNextAllowed2(current, val), string(val))
			if !isVisited && isNextAllowed2(current, val) {
				fmt.Print("\ngoing up")
				queue.push(up)
			}
		}

		down := &Pos{x: front.x, y: front.y + 1, distance: front.distance + 1}
		if down.y < len((*grid)[down.x]) {
			val := (*grid)[down.x][down.y]
			isVisited := (*visted)[down.x][down.y]
			fmt.Print("\ndown:", down, isVisited, isNextAllowed2(current, val), string(val))
			if !isVisited && isNextAllowed2(current, val) {
				fmt.Print("\ngoing down")
				queue.push(down)
			}
		}
	}
	return -1
}

type Queue struct {
	items []*Pos
}

func (q *Queue) push(item *Pos) {
	q.items = append(q.items, item)
}

func (q *Queue) front() *Pos {
	front := q.items[0]
	q.items = q.items[1:]
	return front
}
