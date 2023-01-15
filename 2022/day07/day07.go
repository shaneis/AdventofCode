package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Parent      *Node
	Name        string
	IsFile      bool
	Order, Size int
	Children    []*Node
}

var ()

func main() {
	f := flag.String("filename", "sample_input_01.txt", "file name for input")
	d := flag.Bool("debug", false, "true for extra output, false (default) otherwise")
	flag.Parse()

	lines := readLines(*f, *d)

	fmt.Println("Part 01:", part01(lines, *d))
	fmt.Println("Part 02:", part02(lines, *d))
}

func readLines(file string, debug bool) []string {
	var lines []string

	content, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if debug {
		for _, x := range lines {
			fmt.Printf("%q\n", x)
		}
	}
	return lines
}

func part01(lines []string, debug bool) int {
	var sizes int

	tree := parseLines(lines, debug)
	nodes := tree.iterate()
	for index, node := range nodes {
		fmt.Printf("[%2d]Iterating over node: %q\n", index, node.Name)
		nodeSize := node.getSize(debug)
		fmt.Printf("Size: %d\n", node.getSize(debug))
		if nodeSize <= 100000 {
			sizes += nodeSize
		}
	}

	return sizes
}

func parseLines(lines []string, debug bool) *Node {
	var currentNode *Node
	root := &Node{
		Parent:   nil,
		Name:     "/",
		IsFile:   false,
		Order:    1,
		Size:     0,
		Children: nil,
	}
	currentNode = root
	for index := 0; index < len(lines); {

		tokens := strings.Fields(lines[index])

		if tokens[0] == "$" {

			instr := tokens[1]
			switch instr {
			case "cd":
				loc := tokens[2]
				currentNode = setLocation(currentNode, loc, debug)
			case "ls":
				startIndex, endIndex := index+1, index+1

				for !strings.HasPrefix(lines[index+1], "$") {
					index++
					endIndex++
					if index == len(lines)-1 {
						if debug {
							fmt.Println("Reached end of file, breaking...")
						}
						break
					}
				}
				parsed := currentNode.addItemstoDir(lines[startIndex:endIndex], debug)
				index += parsed
				continue
			}
		}
		index++
	}
	if debug {
		showTree(root, 1, debug)
	}
	return root
}

func setLocation(StartingPoint *Node, location string, debug bool) *Node {
	var returnNode *Node
	switch location {
	case "/":
		returnNode = StartingPoint
		for returnNode.Parent != nil {
			returnNode = StartingPoint.Parent
		}
		return returnNode
	case "..":
		if StartingPoint.Parent == nil {
			log.Fatal()
		}
		return StartingPoint.Parent
	case location:
		for i := 0; i < len(StartingPoint.Children); i++ {
			if StartingPoint.Children[i].Name == location {
				return StartingPoint.Children[i]
			}
		}
	}
	return nil
}

func (n *Node) addItemstoDir(lines []string, debug bool) int {
	var (
		linesParsed, order int
	)

	for i := 0; i < len(lines); i++ {
		tokens := strings.Fields(lines[i])
		order++

		if tokens[0] == "dir" {
			n.addDirectory(tokens[1], order, debug)
			continue
		}

		size, err := strconv.Atoi(tokens[0])
		if err != nil {
			log.Fatal(err)
		}
		n.addFile(tokens[1], order, size, debug)
	}

	linesParsed += 1
	return linesParsed
}

func (n *Node) addDirectory(name string, order int, debug bool) {
	newN := Node{
		Parent:   n,
		Name:     name,
		IsFile:   false,
		Order:    order + 1,
		Size:     0,
		Children: nil,
	}
	n.Children = append(n.Children, &newN)
}

func (n *Node) addFile(name string, order, size int, debug bool) {
	newN := Node{
		Parent:   n,
		Name:     name,
		IsFile:   true,
		Order:    order + 1,
		Size:     size,
		Children: nil,
	}
	n.Children = append(n.Children, &newN)
}

func (n *Node) getSize(debug bool) int {
	var size int
	size = n.Size
	var stack []Node
	//DFS?
	for i := 0; i <= len(n.Children)-1; i++ {
		if n.Children[i].IsFile == true {
			size += n.Children[i].Size
		}

		if n.Children[i].IsFile != true {
			stack = append(stack, *n.Children[i])
		}
	}

	for len(stack) > 0 {
		f := stack[0]

		for i := 0; i <= len(f.Children)-1; i++ {
			if f.Children[i].IsFile == true {
				size += f.Children[i].Size
			}

			if f.Children[i].IsFile != true {
				stack = append(stack, *f.Children[i])
			}
		}

		stack = stack[1:]
	}

	return size
}

func (n *Node) iterate() []Node {
	var stack []Node
	var visited []Node

	stack = append(stack, *n)

	for i := 0; i <= len(n.Children)-1; i++ {
		if n.Children[i].IsFile != true {
			stack = append(stack, *n.Children[i])
		}

	}

	visited = append(visited, *n)
	stack = stack[1:]

	for len(stack) > 0 {
		node := stack[0]
		fmt.Println("Node:", node)
		for i := 0; i <= len(node.Children)-1; i++ {
			if node.Children[i].IsFile != true {
				stack = append(stack, *node.Children[i])
			}
		}
		visited = append(visited, node)
		stack = stack[1:]
	}
	return visited
}

func addToStack(n *Node, debug bool) []Node {
	var results []Node
	var stack []Node
	keepGoing := true

	stack = append(stack, *n)
	current := stack[len(stack)-1]
	for keepGoing {
		for i := 0; i <= len(current.Children)-1; i++ {
			stack = append(stack, *current.Children[i])
		}
		results = append(results, current)
		current = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		if len(stack) == 0 {
			keepGoing = false
		}
	}

	return results
}

func showTree(n *Node, depth int, debug bool) {
	tabs := strings.Repeat(" ", depth-1)

	if n.IsFile != true {
		fmt.Printf("%s- %s (dir)\n", tabs, n.Name)
	} else {
		fmt.Printf("%s- %s (file, size=%d)\n", tabs, n.Name, n.Size)
	}
	if len(n.Children) != 0 {
		depth++
		for i := 0; i < len(n.Children); i++ {
			showTree(n.Children[i], depth, debug)
		}
	}
}

func part02(lines []string, debug bool) int {
	var sizeFreed int

	totalSize := 70000000

	filetree := parseLines(lines, debug)
	nodes := filetree.iterate()
	for index, node := range nodes {
		if node.Name == "/" {
			fmt.Printf("[%2d] Node %q, Size: %d\n", index, node.Name, node.getSize(debug))
		}
	}
	sizeUsed := filetree.getSize(debug)
	sizeRemaining := totalSize - sizeUsed
	fmt.Printf("Total Size: %d. Size Used: %d. Size Remaining: %d\n", totalSize, sizeUsed, sizeRemaining)

	return sizeFreed
}
