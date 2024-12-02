package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	reports := strings.Split(*input, "\n")
	var safeReports int
	for _, report := range reports {
		parsedReport := wordToReport(report)
		isSafe := isSafeReport(parsedReport)
		fmt.Printf("Report: %s, isSafe: %t \n", report, isSafe)
		if isSafe {
			safeReports++
		}
	}
	return strconv.Itoa(safeReports)
}

func Part2(input *string) string {
	reports := strings.Split(*input, "\n")
	var safeReports int
	for _, report := range reports {
		parsedReport := wordToReport(report)
		isSafe := isSafeReportWithDampners(parsedReport)
		fmt.Printf("Report: %s, isSafe: %t \n", report, isSafe)
		if isSafe {
			safeReports++
		}
	}
	return strconv.Itoa(safeReports)
}

func isSafeReport(report []int) bool {
	isIncreasing := true
	isDecreasing := true
	validDiff := true

	for i := 0; i < len(report)-1; i++ {
		if isDecreasing && report[i] < report[i+1] {
			isDecreasing = false
		}
		if isIncreasing && report[i] > report[i+1] {
			isIncreasing = false
		}
		valueDiff := report[i] - report[i+1]
		if valueDiff < 0 {
			valueDiff *= -1
		}
		if valueDiff < 1 || valueDiff > 3 {
			validDiff = false
		}
		if !validDiff || (!isDecreasing && !isIncreasing) {
			return false
		}
	}
	return true
}

func wordToReport(word string) []int {
	values := strings.Split(word, " ")
	valuesInt := make([]int, 0)
	for _, value := range values {
		valueInt, _ := strconv.Atoi(value)
		valuesInt = append(valuesInt, valueInt)
	}
	return valuesInt
}

func removeIndex(s []int, index int) []int {
	removed := make([]int, 0)
	removed = append(removed, s[:index]...)
	removed = append(removed, s[index+1:]...)
	return removed
}

func isSafeReportWithDampners(report []int) bool {
	if isSafeReport(report) {
		return true
	}
	for index := range report {
		oneRemovedReport := removeIndex(report, index)
		if isSafeReport(oneRemovedReport) {
			return true
		}
	}
	return false
}
