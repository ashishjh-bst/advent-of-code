package day13

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	lines := strings.Split(*input, "\n")
	validSignals := make([]int, 0)
	signalCounter := 1
	for i := 0; i < len(lines); i = i + 3 {
		valid := true
		leftSignal := lines[i]
		rightSignal := lines[i+1]
		valid = comparator(leftSignal, rightSignal)
		if valid {
			validSignals = append(validSignals, signalCounter)
		}
		signalCounter++
	}

	result := 0
	for _, item := range validSignals {
		result += item
	}
	return fmt.Sprintf("%d", result)
}

func isInt(v interface{}) bool {
	switch v.(type) {
	case int:
		return true
	}
	return false
}

func comparator(left string, right string) bool {
	leftQ, _ := parse(left)
	rightQ, _ := parse(right)
	leftQ = leftQ[0].([]interface{})
	rightQ = rightQ[0].([]interface{})
	for i := 0; i < len(leftQ); i++ {
		if i >= len(rightQ) {
			return false
		}
		comparison := compare(leftQ[i], rightQ[i])
		if comparison == 0 {
			continue
		}
		if comparison == -1 {
			return false
		}
		if comparison == 1 {
			return true
		}
	}
	return true
}

func compare(left interface{}, right interface{}) int {
	if isInt(left) && isInt(right) {
		l := left.(int)
		r := right.(int)
		if l < r {
			return 1
		}
		if l > r {
			return -1
		}
		return 0
	}
	leftQ := make([]interface{}, 0)
	rightQ := make([]interface{}, 0)
	if isInt(left) {
		leftQ = append(leftQ, left)
	} else {
		leftQ = left.([]interface{})
	}
	if isInt(right) {
		rightQ = append(rightQ, right)
	} else {
		rightQ = right.([]interface{})
	}
	for i := 0; i < len(leftQ); i++ {
		if i >= len(rightQ) {
			return -1
		}
		comparison := compare(leftQ[i], rightQ[i])
		if comparison != 0 {
			return comparison
		}
	}
	if len(leftQ) < len(rightQ) {
		return 1
	}
	return 0
}

func parse(sub string) ([]interface{}, int) {
	slice := make([]interface{}, 0)
	i := 0
	for i < len(sub) {
		switch sub[i] {
		case ',':
			i++
			continue
		case '[':
			subSlice, inc := parse(sub[i+1:])
			slice = append(slice, subSlice)
			i += inc + 1
		case ']':
			return slice, i + 1
		default:
			start := i
			for i < len(sub)-1 && (sub[i+1] != '[' && sub[i+1] != ']' && sub[i+1] != ',') {
				i++
			}
			end := i
			num := sub[start : end+1]
			if start == end {
				num = string(sub[i])
			}
			p, _ := strconv.Atoi(num)
			slice = append(slice, p)
			i++
		}
	}
	return slice, i
}

func Part2(input *string) string {
	return ""
}
