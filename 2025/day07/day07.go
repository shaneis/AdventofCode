package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func getFileInput(file string) ([]byte, int) {
	var fc []byte
	var width int

	fileContents, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		return nil, 0
	}

	widthCaptured := false
	for i := 0; i < len(fileContents); i++ {
		if !widthCaptured && fileContents[i] == byte('\r') {
			width = i
			widthCaptured = true
		}
		if fileContents[i] == byte('\r') || fileContents[i] == byte('\n') {
			continue
		}
		fc = append(fc, fileContents[i])
	}

	return fc, width
}

func displayMap(fileContents []byte, width int) {
	for i := 0; i < len(fileContents); i += width {
		line := fileContents[i : i+width]
		fmt.Println(string(line))
	}
}

func updateMap(fc []byte, width int) ([]byte, int) {
	hits := 0
	for i := 0; i < len(fc); i++ {
		if fc[i] == byte('S') {
			if i+width < len(fc) && fc[i+width] == byte('.') {
				fc[i+width] = byte('|')
			}
		}
		if fc[i] == byte('|') {
			if i+width < len(fc) && fc[i+width] == byte('.') {
				fc[i+width] = byte('|')
			}
			if i+width+1 < len(fc) && fc[i+width] == byte('^') {
				hits++
				fc[i+width-1] = byte('|')
				fc[i+width+1] = byte('|')
			}
		}
	}
	return fc, hits
}

func main() {
	file := flag.String("file", "sample_input_01.txt", "Relative path to input file")
	flag.Parse()

	fileContents, width := getFileInput(*file)
	displayMap(fileContents, width)

	newMap := []byte{}
	hits := 0
	prevMap := fileContents
	for _ = range 2 {
		newMap, hits = updateMap(prevMap, width)
		displayMap(newMap, width)
		fmt.Println()
		prevMap = newMap
	}
	fmt.Printf("Day 07, Part 1: %d\n", hits)

}
