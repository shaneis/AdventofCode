package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func readLine(file string) []string {
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

func findCommonLetterInLines(line []string) (rune, error) {
	runeMap := make(map[rune]int)
	lineCount := len(line)
	for _, s := range line {
		// make a set cause we have dups per line
		set := make(map[rune]int)
		for _, u := range s {
			_, e := set[u]
			if e == false {
				set[u] = 1
			}
		}
		for r := range set {
			runeMap[r] += 1
		}
	}
	for k, v := range runeMap {
		if v == lineCount {
			return k, nil
		}
	}
	return -1, errors.New("No common element found")
}

func getPriorities(chars []rune) []int {
	var (
		lowerP int = int('a')
		upperP int = int('A')
		res    []int
	)
	for i := 0; i < len(chars); i++ {
		rInt := int(chars[i])
		pr := rInt - upperP
		if pr < 26 && pr >= 0 {
			res = append(res, pr+27)
			continue
		} else {
			res = append(res, rInt-lowerP+1)
		}
	}
	return res
}

func part01(line []string) int {
	// Parse lines into 2
	rcks := parseLine(line)
	// Get common letter
	var (
		dups []int
		ps   int
		// pr   int
		cls []rune
	)
	for _, r := range rcks {
		cl, e := findCommonLetter(r.Comp1, r.Comp2)
		if e != nil {
			log.Fatal(e)
		}
		cls = append(cls, cl)
	}
	// Get priority
	dups = getPriorities(cls)
	// Sum priorities
	for _, p := range dups {
		ps += p
	}
	return ps
}

func part02(line []string) int {
	var (
		init int
		ls   []rune
		bPrs int
	)
	init = 0
	for i := 3; i <= len(line); i += 3 {
		l, e := findCommonLetterInLines(line[init:i])
		if e != nil {
			log.Fatal(e)
		}
		init = i
		ls = append(ls, l)
	}
	// Priorities
	prs := getPriorities(ls)
	for _, p := range prs {
		bPrs += p
	}
	return bPrs
}

func main() {

	// fmt.Printf("Hello, world\n")
	var (
		file = flag.String("FileName", "sample_input_01.txt", "The name of the file to parse")
	)
	flag.Parse()
	line := readLine(*file)

	fmt.Println("Part01: ", part01(line))
	fmt.Println("Part02: ", part02(line))
}
