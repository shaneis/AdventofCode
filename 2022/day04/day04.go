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

func part01(file string) int {
	var (
		numTrue int
	)
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Cannot open file: %s", file)
		log.Fatal(err)
	}
	scnr := bufio.NewScanner(f)
	for scnr.Scan() {
		xs := string(scnr.Text())
		// 		fmt.Printf("Parsing: %s\n", xs)
		parsed := parseLine(xs)
		if parsed {
			numTrue++
		}
	}
	return numTrue
}

func parseLine(line string) bool {
	var (
		elves = make(map[int]int)
	)
	// make the hash
	sections := strings.Split(string(line), ",")
	for _, elfSec := range sections {
		eSec := strings.Split(elfSec, "-")
		start, _ := strconv.Atoi(string(eSec[0]))
		finish, _ := strconv.Atoi(string(eSec[1]))
		for i := start; i <= finish; i++ {
			elves[i]++
		}
	}
	// check the hash
	sections = strings.Split(string(line), ",")
	for _, elfSec := range sections {
		eSec := strings.Split(elfSec, "-")
		start, _ := strconv.Atoi(string(eSec[0]))
		finish, _ := strconv.Atoi(string(eSec[1]))
		var hasAll bool = true
		for i := start; i <= finish; i++ {
			k := elves[i]
			if k == 1 {
				hasAll = false
			}
		}
		if hasAll {
			return hasAll
		}
	}
	return false
}

func main() {
	f := flag.String("filename", "sample_input_01.txt", "specify the filename to pass into the function")
	flag.Parse()

	fmt.Printf("Solving part 01 for file: %s\n", *f)
	p1 := part01(*f)

	fmt.Printf("Part 01: %d", p1)
}
