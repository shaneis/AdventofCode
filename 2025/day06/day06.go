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

func getFileInput(filePath string) []string {
	fileContents := []string{}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContents = append(fileContents, scanner.Text())
	}
	return fileContents
}

func solvePart1(numbers []string, actions []string) []int {
	numAnswers := len(actions)
	results := make([]int, numAnswers)

	firstLine := numbers[0]
	firstLineNumbers := strings.Fields(firstLine)
	for i := range numAnswers {
		results[i], _ = strconv.Atoi(firstLineNumbers[i])
	}

	for i := 1; i < len(numbers); i++ {
		line := numbers[i]
		lineNumbers := strings.Fields(line)

		for j := 0; j < numAnswers; j++ {
			thisNum, _ := strconv.Atoi(lineNumbers[j])
			switch actions[j] {
			case "+":
				results[j] += thisNum
			case "*":
				results[j] *= thisNum
			}
		}
	}

	return results
}

func main() {
	file := flag.String("file", "sample_input_01.txt", "Input file path")
	flag.Parse()

	fileContents := getFileInput(*file)
	totalLines := len(fileContents)

	actions := fileContents[totalLines-1]

	results := solvePart1(fileContents[:totalLines-1], strings.Fields(actions))
	day06Part01Results := 0
	for _, res := range results {
		day06Part01Results += res
	}
	fmt.Println("Day 06 Part 01 results:", day06Part01Results)
}
