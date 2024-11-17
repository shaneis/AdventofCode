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
func countHousesDelivered2Santas(directions string) int {

	type Location struct {
		x, y int
	}
	Santa := new(Location)
	RoboSanta := new(Location)
	houses := make(map[Location]int)
	houses[*Santa]++

	for i := 0; i < len(directions); i++ {
		switch string(directions[i]) {
		case "^":
			if i%2 == 0 {
				Santa.y++
				houses[*Santa]++
			} else {
				RoboSanta.y++
				houses[*RoboSanta]++
			}
		case ">":
			if i%2 == 0 {
				Santa.x++
				houses[*Santa]++
			} else {
				RoboSanta.x++
				houses[*RoboSanta]++
			}
		case "v":
			if i%2 == 0 {
				Santa.y--
				houses[*Santa]++
			} else {
				RoboSanta.y--
				houses[*RoboSanta]++
			}
		case "V":
			if i%2 == 0 {
				Santa.y--
				houses[*Santa]++
			} else {
				RoboSanta.y--
				houses[*RoboSanta]++
			}
		case "<":
			if i%2 == 0 {
				Santa.x--
				houses[*Santa]++
			} else {
				RoboSanta.x--
				houses[*RoboSanta]++
			}
		default:
			panic(fmt.Sprintf("Unreachable! %s\n", string(directions[i])))
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
	fmt.Println("Part 02:", countHousesDelivered2Santas(string(f)), "houses had presents delivered")
}
