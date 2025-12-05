package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetRotationPasswordBruteForce(rotations []string) int {
	hitsZero := 0
	start := 50
	for _, rotation := range rotations {
		direction := rotation[0]
		turns := rotation[1:]
		value, _ := strconv.Atoi(turns)

		for i := 0; i < value; i++ {
			switch direction {
			case 'L':
				start--
			case 'R':
				start++
			}
			if start == 0 || start == 100 {
				hitsZero++
			}
			start = (start + 100) % 100
		}
	}
	return hitsZero
}
func GetRotationPassword2(rotations []string) int {
	hitsZero := 0

	oldStart, newStart := 50, 50
	for _, rotation := range rotations {
		direction := rotation[0]
		turns := rotation[1:]

		value, _ := strconv.Atoi(turns)
		extraIncrement := value / 100
		if extraIncrement > 0 {
		}
		remainder := value % 100
		var newValue int
		if direction == 'L' {
			newValue = oldStart - remainder
		} else {
			newValue = oldStart + remainder
		}

		switch direction {
		case 'L':
			if newStart != 0 && newValue <= 0 && extraIncrement > 0 {
				hitsZero += extraIncrement + 1
			} else if newStart != 0 && newValue <= 0 {
				hitsZero++
			} else if extraIncrement > 0 {
				hitsZero += extraIncrement
			}
			if newValue <= 0 {
				newStart = (100 + newValue) % 100
			} else {
				newStart = newValue
			}
		case 'R':
			if newStart != 0 && newValue >= 100 && extraIncrement > 0 {
				hitsZero += extraIncrement + 1
			} else if newStart != 0 && newValue >= 100 {
				hitsZero++
			} else if extraIncrement > 0 {
				hitsZero += extraIncrement
			}
			if newValue >= 100 {
				newStart = (newValue - 100) % 100
			} else {
				newStart = newValue
			}
		}
		oldStart = newStart
	}
	return hitsZero
}

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
	fmt.Println("Day 01-i:", GetRotationPassword(rotations))
	fmt.Println("Day 01-ii:", GetRotationPassword2(rotations))
	fmt.Println("Day 01-bf:", GetRotationPasswordBruteForce(rotations))
}
