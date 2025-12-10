package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func FindMaxPossibleJoltage(bank string) int {
	max10 := 0
	maxTotal := 0
	for i := 0; i < len(bank)-1; i++ {
		// get the max 10
		digit10, _ := strconv.Atoi(string(bank[i]))
		if digit10 >= max10 {
			max10 = digit10
			maxOne := 0
			// get the max 1
			for j := i + 1; j < len(bank); j++ {
				digitOne, _ := strconv.Atoi(string(bank[j]))
				if digitOne >= maxOne {
					maxOne = digitOne
				}
			}
			// check if the max total is beaten
			digitTotal := max10*10 + maxOne
			if digitTotal > maxTotal {
				maxTotal = digitTotal
			}
		}
	}
	return maxTotal
}

func FindMaxPossibleJoltage12(bank string) int {
	n := 12
	firstDigit, _ := strconv.Atoi(string(bank[0]))
	stack := []int{firstDigit}
	digitsToRemove := len(bank) - n

	for i := 1; i < len(bank); i++ {
		// do we need this? is byte('1') > byte('0')?
		// don't care; it works - test later...
		c, _ := strconv.Atoi(string(bank[i]))

		for len(stack) > 0 && digitsToRemove > 0 && stack[len(stack)-1] < c {
			stack = stack[:len(stack)-1]
			digitsToRemove--
		}

		if len(stack) < n {
			stack = append(stack, c)
		} else {
			digitsToRemove--
		}
	}
	var sb strings.Builder
	for _, v := range stack {
		sb.WriteString(strconv.Itoa(v))
	}
	finalNumber, _ := strconv.Atoi(sb.String())
	return finalNumber
}

func main() {
	f := flag.String("file", "input.txt", "name of the input file")
	flag.Parse()
	file, err := os.Open(*f)
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := 0
	totalPart2 := 0
	for scanner.Scan() {
		input := scanner.Text()
		total += FindMaxPossibleJoltage(input)
		totalPart2 += FindMaxPossibleJoltage12(input)
	}
	fmt.Printf("Day 03-i: %d\n", total)
	fmt.Printf("Day 03-ii: %d\n", totalPart2)
}
