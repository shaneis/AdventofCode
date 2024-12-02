package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isReportSafe(report []int) bool {
	isIncreasing := true
	if report[0] > report[1] {
		isIncreasing = false
	} else if report[0] < report[1] {
		isIncreasing = true
	} else {
		return false
	}
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		if diff < -3 || diff > 3 {
			return false
		}
		switch isIncreasing {
		case true:
			if report[i] >= report[i+1] {
				return false
			}
		case false:
			if report[i] <= report[i+1] {
				return false
			}
		}
	}
	return true
}

func convertStringReportToInt(str string) []int {
	reports := strings.Fields(str)

	reportInt := []int{}
	for _, r := range reports {
		if t, err := strconv.Atoi(r); err != nil {
			panic(err)
		} else {
			reportInt = append(reportInt, t)
		}
	}
	return reportInt
}

func main() {
	fileName := "puzzle_input.txt"
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	goodReports, badReports := 0, 0
	for scanner.Scan() {
		thisReport := convertStringReportToInt(scanner.Text())
		if isSafe := isReportSafe(thisReport); isSafe {
			goodReports++
		} else {
			badReports++
		}
	}

	fmt.Println("Part 01: Number of safe reports =", goodReports)
}
