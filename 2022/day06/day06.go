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

	fmt.Println("Part 01:", p1)
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

	result := findStartingMarker(line)
	return result
}

func findStartingMarker(line string) int {
	start, end := 0, 4
	stopPoint := len(line)

	for end < stopPoint {
		// 		fmt.Println("Checking", line[start:end], ", start", start, "end", end, "stopPoint", stopPoint)

		matched := false
		matchedI := start
		matchedJ := 0
	MatchCheck:
		for ; matchedI < end; matchedI++ {
			for matchedJ = matchedI + 1; matchedJ < end; matchedJ++ {
				// 				fmt.Println("Comparing", matchedI, string(line[matchedI]), "with", matchedJ, string(line[matchedJ]))
				if line[matchedI] == line[matchedJ] {
					matched = true
					break MatchCheck
				}
			}
		}

		if matched {
			// 			fs := string(line[matchedI])
			// 			ss := string(line[matchedJ])
			// 			fmt.Printf("Matched %q against %q, (%d:%d)\n", fs, ss, matchedI, matchedJ)
			end++
			start++
			continue
		}

		stopPoint = end
	}
	return end
}
