package day12

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ashishjh-bst/aoc/common"
)

var Directions = [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
var VisitedMap = make(map[[2]int]bool)

type Region struct {
	edges     int
	regionMap map[[2]int]bool
	sides     int
}

func Part1(input *string) string {
	garden := makeGrid(input)
	regions := make([]*Region, 0)
	for i, row := range *garden {
		for j := range row {
			pos := [2]int{i, j}
			_, ok := VisitedMap[pos]
			if ok {
				continue
			}
			regions = append(regions, getRegion(pos, garden))
		}
	}
	cost := 0
	for _, region := range regions {
		cost += len(region.regionMap) * region.edges
	}
	return fmt.Sprintf("%d", cost)
}

func Part2(input *string) string {
	garden := makeGrid(input)
	regions := make([]*Region, 0)
	for i, row := range *garden {
		for j := range row {
			pos := [2]int{i, j}
			_, ok := VisitedMap[pos]
			if ok {
				continue
			}
			regions = append(regions, getRegion(pos, garden))
		}
	}
	cost := 0
	for _, region := range regions {
		cost += len(region.regionMap) * region.sides
	}
	return fmt.Sprintf("%d", cost)
}

func makeGrid(input *string) *[][]rune {
	grid := make([][]rune, 0)
	for _, line := range strings.Split(*input, "\n") {
		grid = append(grid, []rune(line))
	}
	return &grid
}

func getRegion(pos [2]int, grid *[][]rune) *Region {
	regionMap := make(map[[2]int]bool)
	sides := make([][2][2]int, 0)
	plant := (*grid)[pos[0]][pos[1]]
	edges := 0
	queue := [][2]int{pos}
	for len(queue) > 0 {
		block := queue[0]
		queue = queue[1:]
		_, ok := regionMap[block]
		if ok {
			continue
		}
		regionMap[block] = true
		VisitedMap[block] = true
		for _, dir := range Directions {
			next := [2]int{block[0] + dir[0], block[1] + dir[1]}
			_, ok := regionMap[next]
			if ok {
				continue
			}
			if !IsEdge(next, grid, plant) {
				queue = append(queue, next)
				continue
			}
			sides = append(sides, [2][2]int{block, dir})
			edges++
		}
	}
	region := &Region{regionMap: regionMap, edges: edges}
	region.sides = countSides(sides)
	return region
}

func IsEdge(pos [2]int, grid *[][]rune, plant rune) bool {
	if common.IsPosOutside(pos, grid) {
		return true
	}
	if (*grid)[pos[0]][pos[1]] != plant {
		return true
	}
	return false
}

func countSides(sides [][2][2]int) int {
	sideMap := make(map[[2][2]int]bool)

	sort.Slice(sides, func(i, j int) bool {
		if sides[i][0][0] == sides[j][0][0] {
			return sides[i][0][1] < sides[j][0][1]
		}
		return sides[i][0][0] < sides[j][0][0]
	})

	totalSides := 0
	for _, side := range sides {
		edges := getEdges(side[0])
		isVisitedEdge := false
		for _, edge := range edges {
			edge[1] = side[1]
			if _, found := sideMap[edge]; found {
				isVisitedEdge = true
			}
		}
		if !isVisitedEdge {
			totalSides++
		}
		sideMap[side] = true
	}

	return totalSides

}

func getEdges(side [2]int) [][2][2]int {
	edges := make([][2][2]int, 0)
	for _, dir := range Directions {
		next := [2]int{side[0] + dir[0], side[1] + dir[1]}
		edges = append(edges, [2][2]int{next, dir})
	}
	return edges
}
