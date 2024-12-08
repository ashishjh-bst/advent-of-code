package day8

import (
	"fmt"
	"strings"
)

type Pos [2]int

type AntinaeMap struct {
	Grid        [][]string
	NodeMap     map[string][]Pos
	AntiNodeMap map[Pos]bool
}

func Part1(input *string) string {
	am := ParseInput(input)
	for _, v := range am.NodeMap {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				am.getAntinodes(v[i], v[j])
			}
		}
	}
	return fmt.Sprintf("%d", len(am.AntiNodeMap))
}

func Part2(input *string) string {
	am := ParseInput(input)
	for _, v := range am.NodeMap {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				am.getAllPointsInLine(v[i], v[j])
			}
		}
	}
	return fmt.Sprintf("%d", len(am.AntiNodeMap))
}

func (p Pos) Add(p1 Pos) Pos {
	p2 := Pos{}
	p2[0] = p[0] + p1[0]
	p2[1] = p[1] + p1[1]
	return p2
}

func (p Pos) Sub(p1 Pos) Pos {
	p2 := Pos{}
	p2[0] = p[0] - p1[0]
	p2[1] = p[1] - p1[1]
	return p2
}

func (am *AntinaeMap) getAllPointsInLine(p1, p2 Pos) {
	for i, row := range am.Grid {
		for j := range row {
			n := Pos{i, j}
			if AreColinearPoints(p1, p2, n) {
				am.AntiNodeMap[n] = true
			}
		}
	}
}

func AreColinearPoints(p1, p2, n Pos) bool {
	//slope calculation was complicated, so I just check if they can form a triangle or not, if they can't area of triangle is 0
	area := p1[0]*(p2[1]-n[1]) + p2[0]*(n[1]-p1[1]) + n[0]*(p1[1]-p2[1])
	return area == 0
}

func (am *AntinaeMap) getAntinodes(p1, p2 Pos) {
	n1 := p1.Add(p1.Sub(p2))
	n2 := p2.Add(p2.Sub(p1))
	if !isPosOutside(n1, am.Grid) {
		am.AntiNodeMap[n1] = true
	}
	if !isPosOutside(n2, am.Grid) {
		am.AntiNodeMap[n2] = true
	}
}

func isPosOutside(currentPos Pos, grid [][]string) bool {
	return currentPos[0] < 0 || currentPos[0] >= len(grid) || currentPos[1] < 0 || currentPos[1] >= len((grid)[0])
}

func ParseInput(input *string) *AntinaeMap {
	output := &AntinaeMap{}
	output.Grid = make([][]string, 0)
	output.NodeMap = make(map[string][]Pos)
	output.AntiNodeMap = make(map[Pos]bool)
	rows := strings.Split(*input, "\n")
	for i, row := range rows {
		cols := strings.Split(row, "")
		for j, col := range cols {
			if col != "." {
				_, ok := output.NodeMap[col]
				if !ok {
					output.NodeMap[col] = make([]Pos, 0)
				}
				output.NodeMap[col] = append(output.NodeMap[col], Pos{i, j})
			}
		}
		output.Grid = append(output.Grid, cols)
	}
	return output
}
