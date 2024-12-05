package main

import (
	"fmt"
	"os"
	"strings"
)

type Xmas struct {
	startRow, startCol, endRow, endCol int
}

func isInBounds(row, col, endRow, endCol int) bool {
	return col >= 0 && col < endCol && row >= 0 && row < endRow
}

func findXmas(puzzle string) int {
	puzzleLines := strings.Split(puzzle, "\n")
	rows, columns := len(puzzleLines), len(puzzleLines[0])

	var xmases []Xmas
	for r := range rows {
		for c := range len(puzzleLines[r]) {
			if string(puzzleLines[r][c]) == "X" {
				// right
				if isInBounds(r, c+3, rows, columns) {
					if puzzleLines[r][c:c+4] == "XMAS" {
						xmases = append(xmases, Xmas{r, c, r, c + 4})
					}
				}
				// left
				if isInBounds(r, c-3, rows, columns) {
					if puzzleLines[r][c-3:c+1] == "SAMX" {
						xmases = append(xmases, Xmas{r, c - 3, r, c})
					}
				}
				// up
				if isInBounds(r-3, c, rows, columns) {
					potMat := strings.Join(
						[]string{
							string(puzzleLines[r-3][c]),
							string(puzzleLines[r-2][c]),
							string(puzzleLines[r-1][c]),
							string(puzzleLines[r][c]),
						}, "")
					if potMat == "SAMX" {
						xmases = append(xmases, Xmas{r - 3, c, r, c})
					}
				}
				// down
				if isInBounds(r+3, c, rows, columns) {
					potMat := strings.Join(
						[]string{
							string(puzzleLines[r][c]),
							string(puzzleLines[r+1][c]),
							string(puzzleLines[r+2][c]),
							string(puzzleLines[r+3][c]),
						}, "")
					if potMat == "XMAS" {
						xmases = append(xmases, Xmas{r + 3, c, r, c})
					}
				}
				// up right
				if isInBounds(r-3, c+3, rows, columns) {
					potMat := strings.Join(
						[]string{
							string(puzzleLines[r][c]),
							string(puzzleLines[r-1][c+1]),
							string(puzzleLines[r-2][c+2]),
							string(puzzleLines[r-3][c+3]),
						}, "")
					if potMat == "XMAS" {
						xmases = append(xmases, Xmas{r, c, r - 3, c + 3})
					}
				}
				// up left
				if isInBounds(r-3, c-3, rows, columns) {
					potMat := strings.Join(
						[]string{
							string(puzzleLines[r][c]),
							string(puzzleLines[r-1][c-1]),
							string(puzzleLines[r-2][c-2]),
							string(puzzleLines[r-3][c-3]),
						}, "")
					if potMat == "XMAS" {
						xmases = append(xmases, Xmas{r, c, r - 3, c - 3})
					}
				}
				// down left
				if isInBounds(r+3, c-3, rows, columns) {
					potMat := strings.Join(
						[]string{
							string(puzzleLines[r][c]),
							string(puzzleLines[r+1][c-1]),
							string(puzzleLines[r+2][c-2]),
							string(puzzleLines[r+3][c-3]),
						}, "")
					if potMat == "XMAS" {
						xmases = append(xmases, Xmas{r, c, r + 3, c - 3})
					}
				}
				// down right
				if isInBounds(r+3, c+3, rows, columns) {
					potMat := strings.Join(
						[]string{
							string(puzzleLines[r][c]),
							string(puzzleLines[r+1][c+1]),
							string(puzzleLines[r+2][c+2]),
							string(puzzleLines[r+3][c+3]),
						}, "")
					if potMat == "XMAS" {
						xmases = append(xmases, Xmas{r, c, r + 3, c + 3})
					}
				}
			}
		}
	}
	return len(xmases)
}

type CrossedMases struct {
	sr, sc, er, ec int
}

func ohAnXMas(puzzle string) int {
	foundYa := []CrossedMases{}
	lines := strings.Split(puzzle, "\n")
	rows, columns := len(lines), len(lines[0])

	type location struct{ r, c int }
	tableOfAngles := make(map[string]location, 4)
	tableOfAngles["up left"] = location{-3, -3}
	tableOfAngles["up right"] = location{-3, 3}
	tableOfAngles["down left"] = location{3, -3}
	tableOfAngles["down right"] = location{3, 3}

	for r := range rows {
		for c := range len(lines[r]) {
			if string(lines[r][c]) != "M" {
				continue
			}
			// we're on an M
			// up left, up right, down left, down right
			for _, firstCheck := range []struct {
				direction   string
				r, c        int
				counterpart struct {
					translateR, translateC int
					check                  location
				}
			}{
				{"up left", -2, -2, struct {
					translateR, translateC int
					check                  location
				}{0, -2, tableOfAngles["up right"]}},
				{"up left", -2, -2, struct {
					translateR, translateC int
					check                  location
				}{-2, 0, tableOfAngles["down left"]}},
				{"up right", -2, 2, struct {
					translateR, translateC int
					check                  location
				}{0, 2, tableOfAngles["up left"]}},
				{"up right", -2, 2, struct {
					translateR, translateC int
					check                  location
				}{-2, 0, tableOfAngles["down right"]}},
				{"down left", 2, -2, struct {
					translateR, translateC int
					check                  location
				}{0, -2, tableOfAngles["down right"]}},
				{"down left", 2, -2, struct {
					translateR, translateC int
					check                  location
				}{2, 0, tableOfAngles["up left"]}},
				{"down right", 2, 2, struct {
					translateR, translateC int
					check                  location
				}{0, 2, tableOfAngles["down left"]}},
				{"down right", 2, 2, struct {
					translateR, translateC int
					check                  location
				}{2, 0, tableOfAngles["up right"]}},
			} {
				if !isInBounds(
					r+firstCheck.r,
					c+firstCheck.c,
					rows,
					columns,
				) {
					continue
				}
				if !isInBounds(
					r+firstCheck.counterpart.translateR,
					c+firstCheck.counterpart.translateC,
					rows,
					columns,
				) {
					continue
				}

				var letters []string
				i, j := 0, 0
				for i != firstCheck.r && j != firstCheck.c {
					thisLetter := string(lines[r+i][c+j])
					letters = append(letters, thisLetter)
					if firstCheck.r < 0 {
						i--
					} else if firstCheck.r > 0 {
						i++
					}
					if firstCheck.c < 0 {
						j--
					} else if firstCheck.c > 0 {
						j++
					}
					if !isInBounds(r+i, c+j, rows, columns) {
						break
					}
				}
				// off by one errors
				letters = append(letters, string(lines[r+i][c+j]))
				if xMas := strings.Join(letters, ""); xMas != "MAS" {
					continue
				}

				var secondLetters []string
				i, j = 0, 0
				for i != firstCheck.counterpart.check.r && j != firstCheck.counterpart.check.c {
					thisLetter := string(lines[r+firstCheck.counterpart.translateR+i][c+firstCheck.counterpart.translateC+j])
					secondLetters = append(secondLetters, thisLetter)
					if firstCheck.counterpart.check.r <= 0 {
						i--
					} else {
						i++
					}
					if firstCheck.counterpart.check.c <= 0 {
						j--
					} else {
						j++
					}
					if !isInBounds(r+firstCheck.counterpart.translateR+i, c+firstCheck.counterpart.translateC+j, rows, columns) {
						break
					}
				}
				if xMas := strings.Join(secondLetters, ""); xMas == "MAS" {
					endR := r + firstCheck.counterpart.translateR
					endC := c + firstCheck.counterpart.translateC
					alreadyThere := false
					for _, x := range foundYa {
						if x.sc == endC && x.sr == endR {
							alreadyThere = true
							break
						}
					}
					if !alreadyThere {
						foundYa = append(foundYa, CrossedMases{r, c, r + firstCheck.counterpart.translateR, c + firstCheck.counterpart.translateC})
					}
				}
			}
		}
	}
	return len(foundYa)
}

func main() {
	fileName := "puzzle_input.txt"
	f, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("error reading file")
		panic(err)
	}
	fmt.Printf("Part 01: %d\n", findXmas(string(f)))
	fmt.Printf("Part 02: %d\n", ohAnXMas(string(f)))
}
