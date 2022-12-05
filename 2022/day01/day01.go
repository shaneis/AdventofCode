package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func main() {
	var (
		filename   = "input.txt" // "sample_input_01.txt"
		elves      []Elf
		fattestElf int
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

	for i := range elves {
		fmt.Printf("Elf %d has %d calories\n", elves[i].elfid, elves[i].calories)
		if elves[i].calories > elves[fattestElf].calories {
			fattestElf = i
		}
	}
	fmt.Printf("The fattest elf is at index %d %+v\n", fattestElf, elves[fattestElf])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
