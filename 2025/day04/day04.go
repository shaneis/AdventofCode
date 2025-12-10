package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func parsePaperRolls(input string, width int) *strings.Builder {
	output := &strings.Builder{}
	height := len(input) / width

	neighbours := []int{-1, 0, 1}
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			nc := 0
			for _, dh := range neighbours {
				for _, dw := range neighbours {
					if dh == 0 && dw == 0 {
						continue
					}
					nh := h + dh
					nw := w + dw
					nbr := nh*width + nw
					if nh < 0 || nh >= height || nw < 0 || nw >= width {
						continue
					}
					if input[nbr] == '@' {
						nc++
					}
				}
			}
			char := input[h*width+w]
			if nc > 3 && char == '@' {
				fmt.Printf("@")
				output.WriteByte('@')
			} else {
				if char == '@' {
					char = 'X'
				}
				fmt.Printf("%c", char)
				output.WriteByte(char)
			}
		}
		fmt.Println()
	}
	return output
}

func main() {
	fileName := flag.String("file", "input.txt", "Input file name")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input strings.Builder
	width := 0
	height := 0
	for scanner.Scan() {
		if width == 0 {
			width = len(scanner.Text())
		}
		line := scanner.Text()
		input.WriteString(line)
		height++
	}

	output := parsePaperRolls(input.String(), width)
	/*
		fmt.Printf("Input:\n%v\n", input.String())
		for i := 0; i < count; i++ {
			fmt.Printf("%d: ", i)
			for j := 0; j < width; j++ {
				fmt.Printf("%c", input.String()[i*width+j])
			}
			fmt.Println()
		}
	*/

	countX := 0
	for i := 0; i < output.Len(); i++ {
		if output.String()[i] == 'X' {
			countX++
		}
	}
	// 11908, too high
	fmt.Printf("Day 04-i: %d\n", countX)
}
