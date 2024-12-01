package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func splitListsIntoTwo(str string) [][]int {
	firstList, secondList := []int{}, []int{}
	lineByLine := strings.Split(str, "\n")
	for _, line := range lineByLine {
		first := strings.Fields(line)

		if f, err := strconv.Atoi(first[0]); err != nil {
			panic(err)
		} else {
			firstList = append(firstList, f)
		}
		if f, err := strconv.Atoi(first[1]); err != nil {
			panic(err)
		} else {
			secondList = append(secondList, f)
		}
	}
	return [][]int{firstList, secondList}
}

func getDifferenceBetweenLists(list1, list2 []int) int {
	sumTotal := 0
	slices.Sort(list1)
	slices.Sort(list2)
	for i := 0; i < len(list1) && i < len(list2); i++ {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff *= -1
		}
		sumTotal += diff
	}
	return sumTotal
}

func getSimilarityScore(list1, list2 []int) int {
	list2Map := make(map[int]int, len(list2))
	for i := 0; i < len(list2); i++ {
		if _, ok := list2Map[list2[i]]; !ok {
			list2Map[list2[i]] = 1
		} else {
			list2Map[list2[i]]++
		}
	}
	sumTotal := 0
	for i := 0; i < len(list1); i++ {
		if c, ok := list2Map[list1[i]]; ok {
			sumTotal += (list1[i] * c)
		}
	}
	return sumTotal
}

func main() {
	fileName := "puzzle_input.txt"
	f, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lists := splitListsIntoTwo(string(f))
	fmt.Println("Part 01 : total distance between lists", getDifferenceBetweenLists(lists[0], lists[1]))
	fmt.Println("Part 02 : total similarity between lists", getSimilarityScore(lists[0], lists[1]))
}
