package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
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

func parseDontDoWhatJohnnyDontDoDoes(str string) []string {
	var muls []string
	var isMul bool = false
	var capture bool = true
	for i := 0; i < len(str); i++ {
		// start capturing again...
		if i < len(str)-4 && str[i:i+4] == "do()" {
			capture = true
		}

		// stop capturing...
		if i < len(str)-6 && str[i:i+7] == "don't()" {
			capture = false
		}

		if capture && string(str[i]) == "m" {
			if str[i:i+4] == "mul(" {
				isMul = true
			}
		}

		if isMul && i < len(str)-5 {
			j := i + 4
			for string(str[j-1]) != ")" && j < len(str) {
				r := rune(str[j-1])
				if !unicode.IsNumber(r) && r != ',' && r != '(' {
					isMul = false
					break
				}
				j++
			}
			regMatch, err := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
			if err != nil {
				panic(err)
			}
			if r := regMatch.MatchString(str[i:j]); !r {
				isMul = false
				continue
			}
			muls = append(muls, str[i:j])
			i = j - 1
			isMul = false
		}

		// if not capturing, skip...
		if !capture {
			continue
		}
	}
	return muls
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
	var answers2 int = 0
	for _, mul := range mulString {
		answer := invokeMulString(mul)
		answers += answer
	}
	mulStringPartDeux := parseDontDoWhatJohnnyDontDoDoes(string(f))
	for _, mul := range mulStringPartDeux {
		answer := invokeMulString(mul)
		answers2 += answer
	}
	fmt.Println("Part 01:", answers)
	fmt.Println("Part 02:", answers2)
}
