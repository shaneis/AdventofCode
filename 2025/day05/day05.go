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

func MergeRanges(ranges []Range, newRange Range) []Range {
	var merged []Range

	for _, r := range ranges {
		// Cause we build this up over time, it works...
		// And we compare the ranges we have with the new range!
		if newRange.End < r.Start || newRange.Start > r.End {
			merged = append(merged, r)
		} else {
			if newRange.Start > r.Start {
				newRange.Start = r.Start
			}
			if newRange.End < r.End {
				newRange.End = r.End
			}
		}
	}
	merged = append(merged, newRange)
	return merged
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
	day05part01Total := 0
	var skinnyRanges []Range
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
			skinnyRanges = MergeRanges(skinnyRanges, r)
			continue
		}

		num, _ := strconv.Atoi(line)
	fresh:
		for _, r := range skinnyRanges {
			if num > r.End {
				continue
			}

			if num >= r.Start && num <= r.End {
				day05part01Total++
				break fresh
			}
		}
	}

	day05part02Total := 0
	for _, r := range skinnyRanges {
		amt := (r.End - r.Start) + 1
		day05part02Total += amt
	}

	fmt.Println("Day 05-i Total:", day05part01Total)
	fmt.Println("Day 05-ii Total:", day05part02Total)
}
