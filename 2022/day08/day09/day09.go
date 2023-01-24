package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

type Knot struct {
	Row, Column int
}

type Rope struct {
	Head, Tail Knot
}

type Instruction struct {
	Direction string
	Amount    int
}

func main() {
	fileName := flag.String("filename", "sample_input_01.txt", "name of the input file")
	flag.Parse()

	contents := parseFile(*fileName)
	// for index, instruction := range instructions {
	// log.Printf("[%d] line - Instruction: %q\n", index, instruction)
	// }
	instructions := parseInstructions(contents)

	log.Printf("Part 01: %d\n", part01(instructions))
}

func parseFile(fileName string) []string {
	var lines []string

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content := bufio.NewScanner(file)
	for content.Scan() {
		lines = append(lines, content.Text())
	}

	return lines
}

func parseInstructions(lines []string) []Instruction {
	var instructions []Instruction

	for _, line := range lines {
		tokens := strings.Fields(line)

		amt, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal(err)
		}

		var dir string
		switch tokens[0] {
		case "R":
			dir = "Right"
		case "U":
			dir = "Up"
		case "L":
			dir = "Left"
		case "D":
			dir = "Down"
		default:
			log.Fatal(tokens[0])
		}

		instructions = append(instructions, Instruction{Direction: dir, Amount: amt})
	}
	return instructions

}

func part01(instructions []Instruction) int {
	return 0
}
