package day1

import (
	"sort"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	rows := strings.Split(*input, "\n")
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	for _, row := range rows {
		rowVals := strings.Split(row, "   ")
		item1, _ := strconv.Atoi(rowVals[0])
		item2, _ := strconv.Atoi(rowVals[1])
		list1 = append(list1, item1)
		list2 = append(list2, item2)
	}
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	var totalDistance int
	for i, item1 := range list1 {
		diff := item1 - list2[i]
		if diff < 0 {
			diff = diff * -1
		}
		totalDistance += diff
	}

	return strconv.Itoa(totalDistance)
}

func Part2(input *string) string {
	rows := strings.Split(*input, "\n")
	list1 := make([]int, 0)
	map2 := make(map[int]int, 0)

	for _, row := range rows {
		rowVals := strings.Split(row, "   ")
		item1, _ := strconv.Atoi(rowVals[0])
		item2, _ := strconv.Atoi(rowVals[1])
		map2[item2]++
		list1 = append(list1, item1)
	}

	var totalDistance int
	for _, item := range list1 {
		if _, ok := map2[item]; ok {
			totalDistance += item * map2[item]
		}
	}

	return strconv.Itoa(totalDistance)
}
