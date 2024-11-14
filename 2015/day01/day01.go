package main

import (
	"fmt"
	"os"
)

func countFloors(direction string) int {
	directionMap := make(map[rune]int, 2)

	for _, x := range direction {
		directionMap[x]++
	}
	return directionMap['('] - directionMap[')']
}

func firstEnterBasement(direction string) int {
	floor := 0
	for i := 0; i < len(direction); i++ {
		switch direction[i] {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return -1
}

func main() {
	fileName := "puzzle_input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Floor:", countFloors(string(file)))
	fmt.Println("First enter basement at position:", firstEnterBasement(string(file)))
}
