package main

import (
	"fmt"
	"os"
)

func countHousesDelivered(directions string) int {

	type Location struct {
		x, y int
	}
	now := new(Location)
	houses := make(map[Location]int)
	houses[*now]++
	for i := 0; i < len(directions); i++ {
		switch string(directions[i]) {
		case "^":
			now.y++
			houses[*now]++
		case ">":
			now.x++
			houses[*now]++
		case "v":
			now.y--
			houses[*now]++
		case "<":
			now.x--
			houses[*now]++
		default:
			panic("Unreachable!")
		}
	}
	return len(houses)
}

func main() {
	fileName := "puzzle_input.txt"
	f, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 01:", countHousesDelivered(string(f)), "houses had presents delivered")
}
