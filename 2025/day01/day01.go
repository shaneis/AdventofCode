package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetRotationPassword(rotations []string) int {
	hitsZero := 0

	oldStart, newStart := 50, 50
	for _, rotation := range rotations {
		direction := rotation[0]
		turns := rotation[1:]

		switch direction {
		case 'L':
			value, _ := strconv.Atoi(turns)
			if value >= newStart {
				newStart = (100 - (value - oldStart)) % 100
			} else {
				newStart = (oldStart - value) % 100
			}
		case 'R':
			value, _ := strconv.Atoi(turns)
			if value >= (100 - newStart) {
				newStart = (value - (100 - oldStart)) % 100
			} else {
				newStart = (oldStart + value) % 100
			}
		}
		if newStart == 0 {
			hitsZero++
		}
		fmt.Println("The dial is rotated", rotation, "to point at", newStart, ". (", oldStart, ")")
		oldStart = newStart
	}
	return hitsZero
}

func parseInput(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rotations []string
	for scanner.Scan() {
		rotations = append(rotations, scanner.Text())
	}
	return rotations
}

func main() {
	rotations := parseInput("input.txt")
	fmt.Println("Day 01:", GetRotationPassword(rotations))
}
