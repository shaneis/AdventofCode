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

func getFileInputAsColumns(filePath string) [][]string {
	fileContents := getFileInput(filePath)
	var results []string
	var nr [][]string
	keepGoing := true
	i := 0
	var thisResult strings.Builder
	for keepGoing {
		columnBreak := true
		for _, line := range fileContents[:len(fileContents)-1] {
			if columnBreak {
				thisResult = strings.Builder{}
			}
			keepGoing = false
			if i >= len(line) {
				// deal with some lines longer than others
				continue
			} else {
				if line[i] != byte(' ') {
					thisResult.WriteByte(line[i])
					columnBreak = false
				}
				keepGoing = true
			}
		}
		if !columnBreak {
			results = append(results, thisResult.String())
		} else {
			nr = append(nr, results)
			results = []string{}
		}
		i++
	}
	return nr
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

func solvePart2(numbers [][]string, actions []string) []int {
	numberActions := len(actions)
	results := make([]int, numberActions)

	for i := 0; i < len(numbers); i++ {
		line := numbers[i]

		first := true
		for j := 0; j < len(line); j++ {
			thisNum, _ := strconv.Atoi(line[j])
			switch actions[i] {
			case "+":
				if first {
					results[i] = thisNum
					first = false
				} else {
					results[i] += thisNum
				}
			case "*":
				if first {
					results[i] = thisNum
					first = false
				} else {
					results[i] *= thisNum
				}
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

	fmt.Printf("File contents: %+v\n", fileContents)

	actions := fileContents[totalLines-1]

	results := solvePart1(fileContents[:totalLines-1], strings.Fields(actions))
	day06Part01Results := 0
	for _, res := range results {
		day06Part01Results += res
	}
	fmt.Println("Day 06 Part 01 results:", day06Part01Results)

	secondFileContents := getFileInputAsColumns(*file)
	secondResults := solvePart2(secondFileContents, strings.Fields(actions))
	day06Part02Results := 0
	for _, res := range secondResults {
		day06Part02Results += res
	}
	fmt.Println("Day 06 Part 02 results:", day06Part02Results)
}
