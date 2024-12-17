package day14

import (
	"fmt"
	"strconv"
	"strings"
)

type Robot struct {
	p [2]int
	v [2]int
	q int
}

var Bounds = [2]int{101, 103}
var Quadrants = [5]int{0, 0, 0, 0, 0}

func Part1(input *string) string {
	robots := initRobots(input)
	for i := 0; i < 100; i++ {
		for _, r := range *robots {
			r.move()
		}
	}

	return fmt.Sprintf("%d", calcScore())
}

func Part2(input *string) string {
	robots := initRobots(input)
	minScore := 9223372036854775807
	minSecond := Bounds[0]*Bounds[1] + 1
	robotmoves := 0
	for i := 1; i <= Bounds[0]*Bounds[1]; i++ {
		for j, r := range *robots {
			r.move()
			score := calcScore()
			// uncomment to see tree and change values.
			// if i == INSERT_MIN_SECOND_RESULT && j ==INSERT_ROBOT_MOVES_VALUE  {
			// 	print(*robots, i, j)
			// }
			if score < minScore {
				minScore = score
				minSecond = i
				robotmoves = j
			}
		}
	}
	fmt.Printf("%d minSeconds, %d robotMoves ", minSecond, robotmoves)
	return fmt.Sprintf("%d", minSecond)
}

func (r *Robot) move() {
	r.p = [2]int{moveInBound(r.p[0], r.v[0], Bounds[0]), moveInBound(r.p[1], r.v[1], Bounds[1])}
	Quadrants[r.q]--

	r.q = r.findQuadrant()
	Quadrants[r.q]++
}

func calcScore() int {
	score := 1
	for i := 1; i < len(Quadrants); i++ {
		score *= Quadrants[i]
	}
	return score
}

func moveInBound(p, v, b int) int {
	n := p + v
	if n < 0 {
		return b + n
	}
	if n > b-1 {
		return n % b
	}
	return n
}

func initRobots(input *string) *[]*Robot {
	lines := strings.Split(*input, "\n")
	robots := make([]*Robot, 0)
	for _, line := range lines {
		config := strings.Split(line, " ")
		pos := parsePos(config[0], "p=")
		vel := parsePos(config[1], "v=")
		robot := &Robot{
			p: pos,
			v: vel,
		}
		robot.q = robot.findQuadrant()
		Quadrants[robot.q]++
		robots = append(robots, robot)
	}
	return &robots
}

func parsePos(posStr string, prefix string) [2]int {
	posStr, _ = strings.CutPrefix(posStr, prefix)
	pos := strings.Split(posStr, ",")
	x, _ := strconv.Atoi(pos[0])
	y, _ := strconv.Atoi(pos[1])
	return [2]int{x, y}
}

func (r *Robot) findQuadrant() int {
	xMid, yMid := Bounds[0]/2, Bounds[1]/2
	if r.p[0] == xMid || r.p[1] == yMid {
		return 0
	}
	if r.p[0] < xMid {
		if r.p[1] < yMid {
			return 1
		}
		return 2
	}
	if r.p[1] < yMid {
		return 3
	}
	return 4
}

func makeGrid() [][]string {
	grid := make([][]string, 0)
	for i := 0; i < Bounds[0]; i++ {
		grid = append(grid, make([]string, 0))
		for j := 0; j < Bounds[1]; j++ {
			grid[i] = append(grid[i], " ")
		}
	}
	return grid
}

func print(robots []*Robot, i, j int) {
	grid := makeGrid()
	fmt.Printf("\n%d seconds, %d robots", i, j)
	for _, r := range robots {
		grid[r.p[0]][r.p[1]] = "█"
	}
	for _, row := range grid {
		fmt.Printf("\n %v", row)
	}
}
