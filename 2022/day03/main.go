package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func readLine(file string, part int) []string {
	var ret []string
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error opening file\n")
		log.Fatal(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)
	for scnr.Scan() {
		ret = append(ret, scnr.Text())
	}
	return ret
}

type RuckSack struct {
	Line    string
	LineLen int
	Comp1   string
	Comp2   string
}

func parseLine(line []string) []RuckSack {
	var (
		RuckSacks []RuckSack
	)
	for _, s := range line {
		ln := len(s)
		half := ln / 2
		comp1 := s[0:half]
		comp2 := s[half:ln]
		rck := RuckSack{
			Line:    s,
			LineLen: ln,
			Comp1:   comp1,
			Comp2:   comp2,
		}
		RuckSacks = append(RuckSacks, rck)
	}
	return RuckSacks
}

func findCommonLetter(first, second string) (rune, error) {
	// create hash, assume gauranteed 1, return that rune
	runeHash := make(map[rune]int)
	for _, s := range first {
		runeHash[s] = 1
	}
	for _, r := range second {
		if runeHash[r] != 0 {
			return r, nil
		}
	}
	return -1, errors.New("No common string found")
}

func part01(file *string, part *int) int {
	// Part 01
	// Read lines
	line := readLine(*file, *part)
	// Parse lines into 2
	rcks := parseLine(line)
	// for _, x := range rcks {
	// fmt.Printf("Parsed: %+v\n", x)
	// }
	// Get common letter
	var (
		dups []int
		ps   int
	)
	for _, r := range rcks {
		var pr int
		cl, e := findCommonLetter(r.Comp1, r.Comp2)
		if e != nil {
			log.Fatal(e)
		}
		// Get priority
		intcl := int(cl) - int('A')
		if intcl < 26 && intcl >= 0 {
			pr = intcl + 27
		} else {
			pr = int(cl) - int('a') + 1
		}
		dups = append(dups, pr)
	}
	// Sum priorities
	for _, p := range dups {
		ps += p
	}
	return ps
}
func main() {

	// fmt.Printf("Hello, world\n")
	var (
		file = flag.String("FileName", "sample_input_01.txt", "The name of the file to parse")
		part = flag.Int("Part", 1, "Solve either part 1 or part 2")
	)
	flag.Parse()
	if *part != 1 && *part != 2 {
		log.Panic("part can only be 1 or 2")
	}
	fmt.Printf("Solving part %d for file %s\n", *part, *file)

	fmt.Println("Part01: ", part01(file, part))
}
