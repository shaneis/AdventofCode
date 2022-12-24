package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var debug bool = false

type Header struct {
	Column, Location int
	crateName        string
}

func delete(h []Header, colid int) ([]Header, string) {
	var (
		bottom     int
		headers    []Header
		crateValue string
	)

	for _, x := range h {
		if bottom == 0 {
			bottom = x.Location
		}
		if x.Column == colid && x.Location < bottom {
			bottom = x.Location
		}
	}

	for _, x := range h {
		if x.Column == colid && x.Location == bottom {
			crateValue = x.crateName
			continue
		}
		headers = append(headers, x)
	}
	if debug {
		fmt.Printf("Returning from delete: %s\n\n", crateValue)
	}
	return headers, crateValue
}

func add(h []Header, crateVal string, to int) Header {
	var (
		bottom int
	)

	for _, x := range h {
		if x.Column == to {
			if bottom == 0 {
				bottom = x.Location
			}
			if x.Location < bottom {
				bottom = x.Location
			}
		}
	}
	if bottom == 0 {
		for i := 0; i < len(h); i++ {
			if h[i].Location > bottom {
				bottom = h[i].Location + 1
			}
		}
	}
	if debug {
		fmt.Printf("Bottom: %d\n\n", bottom)
	}

	return Header{
		crateName: crateVal,
		Location:  bottom - 1,
		Column:    to,
	}
}

type Instr struct {
	amount, from, to int
}

func readLines(file string) []string {
	var (
		lines []string
	)

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scnr := bufio.NewScanner(f)
	for scnr.Scan() {
		lines = append(lines, scnr.Text())
	}
	return lines
}

func getHeaderLines(lines []string) []string {
	var lineResult []string
	for _, x := range lines {
		if x == "" {
			break
		}
		lineResult = append(lineResult, x)
	}
	return lineResult
}

func parseHeader(lines []string) []Header {
	var (
		headers []Header
		widest  int
	)

	depth := len(lines) - 1

	for _, l := range lines {
		if len(l) > widest {
			widest = len(l)
		}
	}

	for j := 1; j <= widest; j += 4 {
		row := 1
		for i := 0; i <= depth-1; i++ {
			if j > len(lines[i]) {
				continue
			}
			if string(lines[i][j]) == " " {
				continue
			}
			header := Header{
				Location:  i + 1,
				Column:    j/4 + 1,
				crateName: string(lines[i][j]),
			}
			row++
			headers = append(headers, header)
		}
	}
	return headers
}

func getInstructionLines(lines []string) []string {
	var (
		lineResult []string
		start      bool = false
	)

	for _, x := range lines {
		if x == "" {
			start = true
			continue
		}

		if !start {
			continue
		}

		lineResult = append(lineResult, x)
	}
	return lineResult
}

func parseInstr(instr string) Instr {
	tokens := strings.Fields(instr)
	amount, _ := strconv.Atoi(tokens[1])
	from, _ := strconv.Atoi(tokens[3])
	to, _ := strconv.Atoi(tokens[5])

	return Instr{
		amount: amount,
		from:   from,
		to:     to,
	}
}

func moveCrate(crates []Header, actions Instr) []Header {
	var result []Header
	sort.SliceStable(crates, func(i, j int) bool {
		if crates[i].Column > crates[j].Column {
			return true
		}
		if crates[i].Column == crates[j].Column && crates[i].Location > crates[j].Location {
			return true
		}
		return false
	})

	crs, cv := delete(crates, actions.from)

	sort.SliceStable(crs, func(i, j int) bool {
		if crs[i].Column > crs[j].Column {
			return true
		}
		if crs[i].Column == crs[j].Column && crs[i].Location > crs[j].Location {
			return true
		}
		return false
	})

	var redoLoc bool
	for j := 0; j < len(crs); j++ {
		if crs[j].Column == actions.to && crs[j].Location == 1 {
			redoLoc = true
			if debug {
				fmt.Println("Redo location:", redoLoc)
			}
		}
	}
	if redoLoc {
		if debug {
			fmt.Println("Before location increase:")
			for _, x := range crs {
				fmt.Printf("%+v\n", x)
			}
		}
		for i := 0; i < len(crs); i++ {
			crs[i].Location += 1
		}
		if debug {
			fmt.Println("After location increase:")
			for _, x := range crs {
				fmt.Printf("%+v\n", x)
			}
		}
	}

	addH := add(crs, cv, actions.to)
	if debug {
		fmt.Printf("Adding %+v\n\n", addH)
	}
	result = append(result, crs...)
	result = append(result, addH)
	if debug {
		for _, x := range result {
			fmt.Printf("After moving:%+v\n", x)
		}
	}
	return result
}

func moveCratesAllTogether(crates []Header, actions Instr) []Header {
	var result []Header

	newCrates, values := getNCrates(crates, actions)

	newCrates = resetLocation(newCrates, actions.to)
	// addNCrates(nCrates)
	result = append(result, newCrates...)
	for i := len(values) - 1; i >= 0; i-- {
		addH := add(result, values[i], actions.to)
		if debug {
			fmt.Printf("[moveCratesAllTogether] Adding %+v\n\n", addH)
		}
		result = append(result, addH)
		if debug {
			for _, x := range result {
				fmt.Printf("[moveCratesAllTogether] After moving:%+v\n", x)
			}
		}
		result = resetLocation(result, actions.to)
	}
	return result
}

func resetLocation(crates []Header, Column int) []Header {
	sort.SliceStable(crates, func(i, j int) bool {
		if crates[i].Column > crates[j].Column {
			return true
		}
		if crates[i].Column == crates[j].Column && crates[i].Location > crates[j].Location {
			return true
		}
		return false
	})

	var redoLoc bool
	for j := 0; j < len(crates); j++ {
		if crates[j].Column == Column && crates[j].Location == 1 {
			redoLoc = true
			if debug {
				fmt.Println("[resetLocation] Redo location:", redoLoc)
			}
		}
	}
	if redoLoc {
		if debug {
			fmt.Println("[resetLocation] Before location increase:")
			for _, x := range crates {
				fmt.Printf("%+v\n", x)
			}
		}
		for i := 0; i < len(crates); i++ {
			crates[i].Location += 1
		}
		if debug {
			fmt.Println("[resetLocation] After location increase:")
			for _, x := range crates {
				fmt.Printf("%+v\n", x)
			}
		}
	}

	return crates
}

func getNCrates(crates []Header, actions Instr) ([]Header, []string) {
	var ret []string
	var cv string
	crs := crates
	for i := 0; i < actions.amount; i++ {
		crs, cv = delete(crs, actions.from)
		ret = append(ret, cv)
	}
	return crs, ret
}

func showCrates(crates []Header) {
	var (
		highest int
		widest  int
	)

	for _, x := range crates {
		if x.Column > widest {
			widest = x.Column
		}
		if x.Location > highest {
			highest = x.Location
		}
	}

	sort.SliceStable(crates, func(i, j int) bool {
		if crates[i].Column < crates[j].Column {
			return true
		}
		if crates[i].Column == crates[j].Column && crates[i].Location < crates[j].Location {
			return true
		}
		return false
	})

	for i := 1; i <= highest; i++ {
		for j := 1; j <= widest; j++ {
			var output string
			var done bool
			for _, c := range crates {
				if c.Column == j && c.Location == i {
					val := c.crateName
					if val == "" {
						val = " "
					}
					output = "[" + val + "]"
					done = true
				} else if !done {
					output = "[ ]"
				}
			}
			fmt.Printf("%s", output)
		}
		fmt.Println()
	}

	fmt.Printf("%s\n", strings.Repeat("-", widest*3))
	for col := 1; col <= widest; col++ {
		fmt.Printf("[%d]", col)
	}
	fmt.Println()
}

func part01(file string) string {
	lines := readLines(file)

	headerLines := getHeaderLines(lines)
	if debug {
		for _, h := range headerLines {
			fmt.Println("Header line:", h)
		}
	}
	header := parseHeader(headerLines)
	if debug {
		fmt.Println("Header:")
		for _, x := range header {
			fmt.Printf("%+v\n", x)
		}
	}

	instrs := getInstructionLines(lines)
	for _, instr := range instrs {
		actions := parseInstr(instr)
		if debug {
			fmt.Printf("Instruction: %+v\n", actions)
		}
		for i := 0; i < actions.amount; i++ {
			header = moveCrate(header, actions)
			if debug {
				showCrates(header)
				fmt.Println()
			}
		}
	}

	var columns = make(map[int]Header, 3)
	var topCrates []string

	for i := 0; i < len(header); i++ {
		currMap, ok := columns[header[i].Column]
		if debug {
			fmt.Println("currMap:", currMap, "ok:", ok)
		}
		if !ok {
			if debug {
				fmt.Println("Setting currMap:", currMap, "to:", header[i])
			}
			currMap = header[i]
			columns[header[i].Column] = header[i]
		}

		if currMap.Location > header[i].Location && currMap.Column == header[i].Column {
			if debug {
				fmt.Printf("From:%v", columns[header[i].Location])
			}
			columns[header[i].Column] = header[i]
			if debug {
				fmt.Printf("To:%v\n", columns[header[i].Location])
			}
		}
		if debug {
			fmt.Println("this iteration currMap is:", currMap)
			fmt.Println("Columns:", columns)
		}
	}

	for i := 1; i <= len(columns); i++ {
		for _, v := range columns {
			if v.Column == i {
				topCrates = append(topCrates, v.crateName)
			}
		}
	}
	answer := strings.Join(topCrates, "")

	return answer
}

func part02(file string) string {
	lines := readLines(file)

	headerLines := getHeaderLines(lines)
	if debug {
		for _, h := range headerLines {
			fmt.Println("Header line:", h)
		}
	}
	header := parseHeader(headerLines)
	if debug {
		fmt.Println("Header:")
		for _, x := range header {
			fmt.Printf("%+v\n", x)
		}
	}

	instrs := getInstructionLines(lines)
	for _, instr := range instrs {
		actions := parseInstr(instr)
		if debug {
			fmt.Printf("Instruction: %+v\n", actions)
		}
		header = moveCratesAllTogether(header, actions)
		if debug {
			showCrates(header)
			fmt.Println()
		}
	}

	var columns = make(map[int]Header, 3)
	var topCrates []string

	for i := 0; i < len(header); i++ {
		currMap, ok := columns[header[i].Column]
		if debug {
			fmt.Println("currMap:", currMap, "ok:", ok)
		}
		if !ok {
			if debug {
				fmt.Println("Setting currMap:", currMap, "to:", header[i])
			}
			currMap = header[i]
			columns[header[i].Column] = header[i]
		}

		if currMap.Location > header[i].Location && currMap.Column == header[i].Column {
			if debug {
				fmt.Printf("From:%v", columns[header[i].Location])
			}
			columns[header[i].Column] = header[i]
			if debug {
				fmt.Printf("To:%v\n", columns[header[i].Location])
			}
		}
		if debug {
			fmt.Println("this iteration currMap is:", currMap)
			fmt.Println("Columns:", columns)
		}
	}

	for i := 1; i <= len(columns); i++ {
		for _, v := range columns {
			if v.Column == i {
				topCrates = append(topCrates, v.crateName)
			}
		}
	}
	answer := strings.Join(topCrates, "")

	return answer
}

func main() {
	f := flag.String("filename", "sample_input_01.txt", "name of the input file")
	flag.Parse()

	fmt.Println("Parsing file:", *f)
	topCrates := part01(*f)
	properTopCrates := part02(*f)

	fmt.Println("Part 01:", topCrates)
	fmt.Println("Part 02:", properTopCrates)
}
