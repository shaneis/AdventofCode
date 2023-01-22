package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
)

type Tree struct {
	x, y              int
	isEdge, isVisible bool
	val               string
}

type Trees []Tree

func (t Trees) Len() int {
	return 0
}

func (t Trees) Less(i, j int) bool {
	if t[i].x < t[j].x {
		return true
	}
	if t[i].x == t[j].x && t[i].y < t[j].y {
		return true
	}
	return false
}

func (t Trees) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func main() {
	f := flag.String("filename", "sample_input_01.txt", "name of file to parse")
	flag.Parse()

	input := parseFile(f)
	// for index, line := range input {
	// log.Printf("[%2d] line: %q\n", index, line)
	// }

	log.Printf("Part 01: %d\n", part01(input))
}

func parseFile(filename *string) []string {
	var lines []string

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func part01(input []string) int {
	allTrees := createTrees(input)

	sort.Stable(allTrees)
	for i := 0; i < len(allTrees); i++ {
		allTrees[i].isVisible = visibleFromTop(input, allTrees[i])
		allTrees[i].isVisible = visibleFromBottom(input, allTrees[i])
		allTrees[i].isVisible = visibleFromLeft(input, allTrees[i])
		allTrees[i].isVisible = visibleFromRight(input, allTrees[i])

		//log.Printf("Tree: %+v\n", allTrees[i])
	}

	var seenTrees int = 0
	for _, tree := range allTrees {
		if tree.isVisible == true {
			seenTrees += 1
		}
	}

	return seenTrees
}

func createTrees(input []string) Trees {
	var t Trees
	for rows := 0; rows < len(input); rows++ {
		for cols := 0; cols < len(input[rows]); cols++ {
			isEdge := rows == 0 || cols == 0 || rows == len(input)-1 || cols == len(input[rows])-1
			cT := Tree{x: rows, y: cols, isEdge: isEdge, isVisible: isEdge, val: string(input[rows][cols])}
			t = append(t, cT)
		}
	}
	return t
}

func visibleFromTop(input []string, treeToCheck Tree) bool {
	if treeToCheck.isVisible == true {
		return treeToCheck.isVisible
	}

	col := treeToCheck.y
	val, err := strconv.Atoi(treeToCheck.val)
	if err != nil {
		log.Fatal(err)
	}

	for rows := 0; rows < treeToCheck.x; rows++ {
		upper, err := strconv.Atoi(string(input[rows][col]))
		if err != nil {
			log.Fatal(err)
		}
		if upper >= val {
			return false
		}
	}
	return true
}

func visibleFromBottom(input []string, treeToCheck Tree) bool {
	if treeToCheck.isVisible == true {
		return treeToCheck.isVisible
	}

	col := treeToCheck.y
	val, err := strconv.Atoi(treeToCheck.val)
	if err != nil {
		log.Fatal(err)
	}

	for rows := len(input) - 1; rows > treeToCheck.x; rows-- {
		bottom, err := strconv.Atoi(string(input[rows][col]))
		if err != nil {
			log.Fatal(err)
		}
		if bottom >= val {
			return false
		}
	}
	return true
}

func visibleFromLeft(input []string, treeToCheck Tree) bool {
	if treeToCheck.isVisible == true {
		return treeToCheck.isVisible
	}

	row := treeToCheck.x
	val, err := strconv.Atoi(treeToCheck.val)
	if err != nil {
		log.Fatal(err)
	}

	for cols := 0; cols < treeToCheck.y; cols++ {
		upper, err := strconv.Atoi(string(input[row][cols]))
		if err != nil {
			log.Fatal(err)
		}
		if upper >= val {
			return false
		}
	}
	return true
}

func visibleFromRight(input []string, treeToCheck Tree) bool {
	if treeToCheck.isVisible == true {
		return treeToCheck.isVisible
	}

	row := treeToCheck.x
	val, err := strconv.Atoi(treeToCheck.val)
	if err != nil {
		log.Fatal(err)
	}

	for cols := len(input) - 1; cols > treeToCheck.y; cols-- {
		upper, err := strconv.Atoi(string(input[row][cols]))
		if err != nil {
			log.Fatal(err)
		}
		if upper >= val {
			return false
		}
	}
	return true
}
