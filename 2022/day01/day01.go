package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	elfid    int
	calories int
}

func (e Elf) addCalories(addedCalories int) Elf {
	e.calories += addedCalories
	return e
}

func parseInput(filename string) []Elf {
	var (
		elves []Elf
	)

	// Dealing with the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		nf    bool = true
		e     Elf
		elfid int = 0
	)

	// foreach text
	for scanner.Scan() {
		c := scanner.Text()

		// if there's a line entry...
		if len(c) > 1 {
			// Convert the string to a int
			cInt, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal(err)
			}

			// and the previous line is empty...
			if nf == true {
				// create a new elf
				elfid += 1
				nf = false

				e = Elf{elfid: elfid, calories: 0}
			}

			// otherwise, add calories to current elf
			e = e.addCalories(cInt)

		} else {
			nf = true
			elves = append(elves, e)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return elves
}

func main() {
	var (
		filename      string = "input.txt" // "sample_input_01.txt"
		fattestElf    int
		totalCalories int = 0
	)
	elves := parseInput(filename)

	for i := range elves {
		if elves[i].calories > elves[fattestElf].calories {
			fattestElf = i
		}
	}

	sort.Slice(elves, func(i, j int) bool { return elves[i].calories > elves[j].calories })

	for i := 0; i < 3; i++ {
		totalCalories += elves[i].calories
	}

	fmt.Printf("Part 01:\t%d\n", elves[fattestElf].calories)
	fmt.Printf("Part 02:\t%d\n", totalCalories)
}
