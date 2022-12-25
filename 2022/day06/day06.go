package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	f := flag.String("filename", "sample_input_01.txt", "file name for input")
	flag.Parse()

	p1 := part01(*f)
	p2 := part02(*f)

	fmt.Println("Part 01:", p1)
	fmt.Println("Part 02:", p2)
}

func part01(file string) int {
	var line string
	read, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scnr := bufio.NewScanner(read)
	for scnr.Scan() {
		line = scnr.Text()
	}

	result := findStartingMarker(line, 4)
	return result
}

func part02(file string) int {
	var line string
	read, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scnr := bufio.NewScanner(read)
	for scnr.Scan() {
		line = scnr.Text()
	}

	result := findStartingMarker(line, 14)
	return result
}

func findStartingMarker(line string, groupSize int) int {
	start, end := 0, groupSize
	stopPoint := len(line)

	for end < stopPoint {

		matched := false
		matchedI := start
		matchedJ := 0
	MatchCheck:
		for ; matchedI < end; matchedI++ {
			for matchedJ = matchedI + 1; matchedJ < end; matchedJ++ {
				if line[matchedI] == line[matchedJ] {
					matched = true
					break MatchCheck
				}
			}
		}

		if matched {
			end++
			start++
			continue
		}

		stopPoint = end
	}
	return end
}
