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
		fmt.Printf("\n PARSING %s and %s", leftSignal, rightSignal)
		valid = comparator(leftSignal, rightSignal)
		fmt.Printf("\n IS IT VALID? %t", valid)
		if valid {
			validSignals = append(validSignals, signalCounter)
		}
		signalCounter++
	}
	fmt.Printf("\n %v", validSignals)
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
	fmt.Print("\nNOW PARSING LEFT")
	leftQ, _ := parse(left)
	fmt.Printf("\n %v", leftQ)
	fmt.Print("\nNOW PARSING RIGHT")
	rightQ, _ := parse(right)
	fmt.Printf("\n %v", rightQ)
	leftQ = leftQ[0].([]interface{})
	rightQ = rightQ[0].([]interface{})
	fmt.Printf("\n COMPARING %v %v", leftQ, rightQ)
	for i := 0; i < len(leftQ); i++ {
		if i >= len(rightQ) {
			return false
		}
		comparison := compare(leftQ[i], rightQ[i])
		fmt.Printf("\n comparison at index %d with left %v  right %v is %d", i, leftQ[i], rightQ[i], comparison)
		if comparison == 0 {
			continue
		}
		if comparison == -1 {
			return false
		}
	}

	return true
}

func compare(left interface{}, right interface{}) int {
	fmt.Printf("\n COMPARING %v %v", left, right)
	if isInt(left) && isInt(right) {
		fmt.Printf("\n BOTH %v %v are ints", left, right)
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
		fmt.Printf("\n left %v is an int, converting to slice", left)
		leftQ = append(leftQ, left)
	} else {
		fmt.Printf("\n left %v is a slice", left)
		leftQ = left.([]interface{})
	}
	if isInt(right) {
		fmt.Printf("\n right %v is an int, converting to slice", right)
		rightQ = append(rightQ, right)
	} else {
		fmt.Printf("\n right %v is a slice", right)
		rightQ = right.([]interface{})
	}

	fmt.Println(len(leftQ), len(rightQ))

	if len(leftQ) == 0 && len(rightQ) > 0 {
		return 1
	}

	if len(leftQ) > 0 && len(rightQ) == 0 {
		return -1
	}

	for i := 0; i < len(leftQ); i++ {
		if i >= len(rightQ) {
			return 1
		}
		comparison := compare(leftQ[i], rightQ[i])
		fmt.Printf("\n comparison at index %d with left %v  right %v is %d", i, leftQ[i], rightQ[i], comparison)
		if comparison != 0 {
			return comparison
		}
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
