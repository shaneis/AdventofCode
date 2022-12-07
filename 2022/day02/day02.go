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

	retorts = map[string]string{
		"X": "lose",
		"Y": "draw",
		"Z": "win",
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
	return sPoints[entry]
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

func getRetort(opp string, want string) string {
	win := map[string]string{
		"Rock":     "Paper",
		"Paper":    "Scissors",
		"Scissors": "Rock",
	}
	lose := map[string]string{
		"Rock":     "Scissors",
		"Paper":    "Rock",
		"Scissors": "Paper",
	}

	if want == "win" {
		return win[opp]
	}
	if want == "lose" {
		return lose[opp]
	}
	return opp
}

func parseFile(filename string, part int) []Outcome {
	var (
		Outcomes []Outcome
	)
	if part != 1 && part != 2 {
		log.Panic("part can only be 1 or 2")
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scnr := bufio.NewScanner(file)

	for scnr.Scan() {
		l := scnr.Text()

		first, second, _ := strings.Cut(l, " ")
		ourC := getRetort(col1[first], retorts[second])
		// _ := getVersusPoints(first)
		var (
			res  string
			s    int
			pnts int
			r    *Outcome
		)
		if part == 1 {
			res = getResultOutcome(col2[second], col1[first])
			s = getSelectedPoints(col2[second])
			pnts = sPoints[col2[second]] + rPoints[res]

			r = NewOutcome(col2[second], col1[first], s, res, pnts)
		} else {
			res = getResultOutcome(ourC, col1[first])
			s = getSelectedPoints(ourC)
			pnts = sPoints[ourC] + rPoints[res]

			r = NewOutcome(ourC, col1[first], s, res, pnts)
		}
		fmt.Printf("Result: %+v\n", r)
		Outcomes = append(Outcomes, *r)
	}
	return Outcomes
}

func main() {
	file := flag.String("filename", "sample_input_01.txt", "the name of the file to parse")
	part := flag.Int("part", 1, "the part you want to solve for")
	flag.Parse()

	fmt.Printf("File passed in: %s\n", *file)

	Outcomes := parseFile(*file, *part)

	var ttl int
	for _, x := range Outcomes {
		ttl += x.total
	}
	fmt.Println("Part 01:", ttl)
}
