package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func FindMaxPossibleJoltage(bank string) int {
	max10 := 0
	maxTotal := 0
	for i := 0; i < len(bank)-1; i++ {
		// get the max 10
		digit10, _ := strconv.Atoi(string(bank[i]))
		if digit10 >= max10 {
			max10 = digit10
			maxOne := 0
			// get the max 1
			for j := i + 1; j < len(bank); j++ {
				digitOne, _ := strconv.Atoi(string(bank[j]))
				if digitOne >= maxOne {
					maxOne = digitOne
				}
			}
			// check if the max total is beaten
			digitTotal := max10*10 + maxOne
			if digitTotal > maxTotal {
				maxTotal = digitTotal
			}
		}
	}
	return maxTotal
}

func main() {
	f := flag.String("file", "input.txt", "name of the input file")
	flag.Parse()
	file, err := os.Open(*f)
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		input := scanner.Text()
		total += FindMaxPossibleJoltage(input)
	}
	fmt.Printf("Day 03-i: %d\n", total)
}
