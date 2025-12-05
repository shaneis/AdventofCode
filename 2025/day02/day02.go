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

	var total int64 = 0
	for _, line := range lines {
		invalids := findInvalidIds(line)
		for _, inval := range invalids {
			total += int64(inval)
		}
	}
	fmt.Printf("Day 02-i: %d\n", total)
}
