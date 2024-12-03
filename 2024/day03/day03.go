package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseMulString(str string) []string {
	var answers []string
	splitStr := strings.SplitAfter(str, ")")
	for _, x := range splitStr {
		for j := len(x) - 1; j >= 0; j-- {
			c := string(x[j])
			if c == "m" {
				regMatch, err := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
				if err != nil {
					panic(err)
				}
				if r := regMatch.MatchString(x[j:]); !r {
					continue
				}
				answers = append(answers, x[j:])
				break
			}
		}
	}
	return answers
}

func invokeMulString(str string) int {
	firstSplit := strings.SplitAfter(str, "(")
	secondSplit := strings.Split(firstSplit[1], ")")
	numsStr := strings.Split(secondSplit[0], ",")
	var nums []int
	for _, num := range numsStr {
		if num == "" {
			continue
		}
		if n, err := strconv.Atoi(num); err != nil {
			panic(err)
		} else {
			nums = append(nums, n)
		}
	}
	return nums[0] * nums[1]
}

func main() {
	fileName := "puzzle_input.txt"
	f, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	mulString := parseMulString(string(f))
	var answers int = 0
	for _, mul := range mulString {
		answer := invokeMulString(mul)
		answers += answer
	}
	fmt.Println("Part 01:", answers)
}
