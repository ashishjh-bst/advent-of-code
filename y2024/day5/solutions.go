package day5

import (
	"strconv"
	"strings"
)

type Rules map[string]map[string]bool
type Updates [][]string
type Input struct {
	Rules   Rules
	Updates Updates
}

func Part1(input *string) string {
	in := Input{}
	in.Updates = Updates{}
	in.Rules = Rules{}
	in.ParseInput(input)
	var midSum int
	for _, u := range in.Updates {
		if in.isSafeUpdate(u) {
			mid, _ := strconv.Atoi(u[(len(u) / 2)])
			midSum += mid
		}
	}
	return strconv.Itoa(midSum)
}

func Part2(input *string) string {
	in := Input{}
	in.Updates = Updates{}
	in.Rules = Rules{}
	in.ParseInput(input)
	var midSum int
	for _, u := range in.Updates {
		if !in.isSafeUpdate(u) {
			fixed := in.sortUpdate(u)
			mid, _ := strconv.Atoi(fixed[(len(fixed) / 2)])
			midSum += mid
		}
	}
	return strconv.Itoa(midSum)
}

func (in *Input) ParseInput(input *string) {
	lines := strings.Split(*input, "\n")
	var shouldParseUpdate bool
	for _, line := range lines {
		if len(line) == 0 {
			shouldParseUpdate = true
			continue
		}
		if shouldParseUpdate {
			updates := strings.Split(line, ",")
			in.Updates = append(in.Updates, updates)
		} else {
			rules := strings.Split(line, "|")
			m, ok := in.Rules[rules[0]]
			if !ok {
				m = make(map[string]bool)
			}
			m[rules[1]] = true
			in.Rules[rules[0]] = m
		}
	}
}

func (in *Input) isSafeUpdate(u []string) bool {
	prevMap := make(map[string]int)
	for i, v := range u {
		prevMap[v] = i
		r, ok := in.Rules[v]
		if !ok {
			continue
		}
		for p := range r {
			_, ok := prevMap[p]
			if ok {
				return false
			}
		}
	}
	return true
}

func (input *Input) sortUpdate(u []string) []string {
	for i := 0; i < len(u)-1; i++ {
		min := i
		for j := i + 1; j < len(u); j++ {
			if !input.IsIBeforeJ(u[j], u[min]) {
				min = j
			}
		}
		if min != i {
			temp := u[i]
			u[i] = u[min]
			u[min] = temp
		}
	}
	return u
}

func (input *Input) IsIBeforeJ(i, j string) bool {
	return input.Rules[i][j]
}
