package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getWrappingPaperArea(dimension string) int {
	dimensionsInStrings := strings.FieldsFunc(dimension, func(r rune) bool {
		return r == 'x'
	})

	var dimensionsInInts []int
	for _, dim := range dimensionsInStrings {
		d, err := strconv.Atoi(dim)
		if err != nil {
			panic(err)
		}
		dimensionsInInts = append(dimensionsInInts, d)
	}

	min := dimensionsInInts[0] * dimensionsInInts[1]
	for i := 0; i < len(dimensionsInInts); i++ {
		for j := 0; j < len(dimensionsInInts); j++ {
			if i == j {
				continue
			}
			smallestSide := dimensionsInInts[i] * dimensionsInInts[j]
			if smallestSide < min {
				min = smallestSide
			}
		}
	}
	l := dimensionsInInts[0]
	w := dimensionsInInts[1]
	h := dimensionsInInts[2]
	return ((2 * l * w) + (2 * w * h) + (2 * h * l)) + min
}

func main() {
	fileName := "puzzle_input.txt"
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var dimensions []string
	for scanner.Scan() {
		dimensions = append(dimensions, scanner.Text())
	}

	squareFeet := 0
	for _, dim := range dimensions {
		squareFeet += getWrappingPaperArea(dim)
	}
	fmt.Println("Order", squareFeet, "square feet of wrapping paper.")
}
