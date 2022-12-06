package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	col1 = map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
	}

	col2 = map[string]string{
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}

	sPoints = map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}

	rPoints = map[string]int{
		"lost": 0,
		"draw": 3,
		"won":  6,
	}
)

type Outcome struct {
	selected string
	versus   string
	points   int
	result   string
	total    int
}

func NewOutcome(
	selected string,
	versus string,
	points int,
	result string,
	total int,
) *Outcome {
	return &Outcome{
		selected: selected,
		versus:   versus,
		points:   points,
		result:   result,
		total:    total,
	}
}

func getSelectedPoints(entry string) int {
	return sPoints[col2[entry]]
}

func getVersusPoints(entry string) int {
	return sPoints[col1[entry]]
}

func getResultOutcome(player1 string, player2 string) string {
	if player1 == "Rock" && player2 == "Scissors" {
		return "won"
	}
	if player1 == "Paper" && player2 == "Rock" {
		return "won"
	}
	if player1 == "Scissors" && player2 == "Paper" {
		return "won"
	}
	if player1 == player2 {
		return "draw"
	}
	return "lost"
}

func parseFile(filename string) []Outcome {
	var (
		Outcomes []Outcome
	)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scnr := bufio.NewScanner(file)

	for scnr.Scan() {
		l := scnr.Text()

		first, second, _ := strings.Cut(l, " ")
		s := getSelectedPoints(second)
		// _ := getVersusPoints(first)
		res := getResultOutcome(col2[second], col1[first])
		pnts := sPoints[col2[second]] + rPoints[res]
		r := NewOutcome(col2[second], col1[first], s, res, pnts)
		fmt.Printf("Result: %+v\n", r)
		Outcomes = append(Outcomes, *r)
	}
	return Outcomes
}

func main() {
	file := flag.String("filename", "sample_input_01.txt", "the name of the file to parse")
	flag.Parse()

	fmt.Printf("File passed in: %s\n", *file)

	Outcomes := parseFile(*file)

	var ttl int
	for _, x := range Outcomes {
		ttl += x.total
	}
	fmt.Println("Part 01:", ttl)
}
