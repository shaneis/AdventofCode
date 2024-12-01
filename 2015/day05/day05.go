package main

import (
	"bufio"
	"fmt"
	"os"
)

func isNice(str string) bool {
	// twice in a row, e.g. aa, bb
	twice := false
	runeHash := make(map[rune]int, len(str))
	for i := 0; i < len(str)-1; i++ {
		r := rune(str[i])
		runeHash[r]++

		if !twice {
			if str[i] == str[i+1] {
				twice = true
			}
		}

		// does not contain 'ab', 'cd', 'pq', 'xy'
		duogram := str[i : i+2]
		for _, bad := range []string{"ab", "cd", "pq", "xy"} {
			if duogram == bad {
				return false
			}
		}
		// fmt.Println("Duogram:", duogram)
	}
	// get the final rune
	runeHash[rune(str[len(str)-1])]++

	// 3 vowels
	numVowels := runeHash['a'] + runeHash['e'] + runeHash['i'] + runeHash['o'] + runeHash['u']

	if !twice {
		return false
	}
	if numVowels < 3 {
		return false
	}
	return true
}

func isNiceV2(str string) bool {
	has2LettersTwice := false
checkLetterLoop:
	for i := 0; i < len(str)-1; i++ {
		for j := i + 2; j < len(str)-1; j++ {
			pre, post := str[i:i+2], str[j:j+2]
			if pre == post {
				has2LettersTwice = true
				break checkLetterLoop
			}
		}
	}
	hasLetterInBetween := false
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+2] {
			hasLetterInBetween = true
			break
		}
	}
	if !has2LettersTwice {
		return false
	}
	if !hasLetterInBetween {
		return false
	}
	return true
}

func main() {
	fileName := "puzzle_input.txt"
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	niceStrings := 0
	reallyNiceStrings := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		isNaughtyOrNice := scanner.Text()
		if isInFactNice := isNice(isNaughtyOrNice); isInFactNice {
			niceStrings++
		}
		if IsNicelyNice := isNiceV2(isNaughtyOrNice); IsNicelyNice {
			reallyNiceStrings++
		}
	}
	fmt.Println("Part 01 : Number of nice strings =", niceStrings)
	fmt.Println("Part 02 : Number of nice strings =", reallyNiceStrings)
}
