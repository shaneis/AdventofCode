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

type Range struct {
	Start int
	End   int
}

func newRange(start, end int) Range {
	return Range{Start: start, End: end}
}

func main() {
	file := flag.String("file", "sample_input.txt", "Path to input file")
	flag.Parse()

	f, err := os.Open(*file)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	isHeader := true
	var freshIDRanges []Range
	day05part01Total := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isHeader = false
			continue
		}

		if isHeader {
			ranges := strings.Split(line, "-")
			start, _ := strconv.Atoi(ranges[0])
			end, _ := strconv.Atoi(ranges[1])
			r := newRange(start, end)
			freshIDRanges = append(freshIDRanges, r)
			continue
		}

		num, _ := strconv.Atoi(line)
	fresh:
		for _, r := range freshIDRanges {
			if num > r.End {
				continue
			}

			if num >= r.Start && num <= r.End {
				day05part01Total++
				break fresh
			}
		}
	}

	fmt.Println("Day 05-i Total:", day05part01Total)
}
