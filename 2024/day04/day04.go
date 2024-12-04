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

func main() {
	fileName := "puzzle_input.txt"
	f, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("error reading file")
		panic(err)
	}
	fmt.Printf("Part 01: %d\n", findXmas(string(f)))
}
