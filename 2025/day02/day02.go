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

func findInvalidIds(numberRange string) []int {
	r := strings.Split(numberRange, "-")
	start, _ := strconv.Atoi(r[0])
	end, _ := strconv.Atoi(r[1])
	var invalids []int
	for i := start; i <= end; i++ {
		s := strconv.Itoa(i)
		if len(s)%2 != 0 {
			continue
		}

		mid := len(s) / 2
		firstHalf := s[:mid]
		secondHalf := s[mid:]
		if firstHalf == secondHalf {
			invalids = append(invalids, i)
		}
	}
	return invalids
}

func findInvalidIdsAtLeastTwice(numberRange string) []int {
	r := strings.Split(numberRange, "-")
	start, _ := strconv.Atoi(r[0])
	end, _ := strconv.Atoi(r[1])
	var invalids []int
	// go through the ranges
	for i := start; i <= end; i++ {
		s := strconv.Itoa(i)
		// split in half then smaller and smaller parts
		for j := 1; j <= len(s); j++ {
			if len(s)%j != 0 {
				continue
			}
			var splitted []string
			for i := 0; i <= len(s)-j; i += j {
				splitted = append(splitted, s[i:i+j])
			}
			if len(splitted) < 2 {
				continue
			}
			// check if all parts are equal
			allEqual := true
			for k := 1; k < len(splitted); k++ {
				if splitted[k] != splitted[0] {
					allEqual = false
					break
				}
			}
			if allEqual {
				invalids = append(invalids, i)
				break
			}
		}
	}
	return invalids
}

func parseInputFile(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ",")...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}

	return lines
}

func main() {
	filepath := flag.String("file", "input.txt", "Input file path")
	flag.Parse()
	lines := parseInputFile(*filepath)

	var totalPart01 int64 = 0
	var totalPart02 int64 = 0
	for _, line := range lines {
		invalids := findInvalidIds(line)
		for _, inval := range invalids {
			totalPart01 += int64(inval)
		}
		invalidsAtLeastTwice := findInvalidIdsAtLeastTwice(line)
		for _, inval := range invalidsAtLeastTwice {
			totalPart02 += int64(inval)
		}
	}
	fmt.Printf("Day 02-i: %d\n", totalPart01)
	fmt.Printf("Day 02-ii: %d\n", totalPart02)
}
