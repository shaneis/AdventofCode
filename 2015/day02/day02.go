package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertDimensions(dimension string) (int, int, int) {
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
	return dimensionsInInts[0], dimensionsInInts[1], dimensionsInInts[2]
}

func getWrappingPaperArea(dimension string) int {
	l, w, h := convertDimensions(dimension)
	dimensionsInInts := [3]int{l, w, h}

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

	return ((2 * l * w) + (2 * w * h) + (2 * h * l)) + min
}

func getRibbonSize(dimensions string) int {
	l, w, h := convertDimensions(dimensions)
	dimensionsInInts := [3]int{l, w, h}

	min := dimensionsInInts[0] * dimensionsInInts[1]
	minPerimeter := (2 * dimensionsInInts[0]) + (2 * dimensionsInInts[1])
	for i := 0; i < len(dimensionsInInts); i++ {
		for j := 0; j < len(dimensionsInInts); j++ {
			if i == j {
				continue
			}
			smallestSide := dimensionsInInts[i] * dimensionsInInts[j]
			if smallestSide < min {
				min = smallestSide
				minPerimeter = (2 * dimensionsInInts[i]) + (2 * dimensionsInInts[j])
			}
		}
	}
	return (l * w * h) + minPerimeter
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
