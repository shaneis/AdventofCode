package main

import "testing"

func TestParseHeader(t *testing.T) {
	var testHeader = []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
	}

	header := parseHeader(testHeader)
	if len(header) != 9 {
		t.Errorf("Expected 9 structs, got %d\n", len(header))
	}
}

func TestGetHeaderLines(t *testing.T) {
	var testHeader = []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 3 to 1",
		"move 3 from 2 to 3",
		"move 2 from 3 to 1",
		"move 1 from 2 to 2",
	}

	headerLines := getHeaderLines(testHeader)
	if len(headerLines) != 4 {
		t.Errorf("Expected 4 header lines, got %d header lines\n%q\n", len(headerLines), headerLines)
	}
}

func TestGetInstructionLines(t *testing.T) {
	var testLines = []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 3 to 1",
		"move 3 from 2 to 3",
		"move 2 from 3 to 1",
		"move 1 from 2 to 2",
	}

	instrLines := getInstructionLines(testLines)
	if len(instrLines) != 4 {
		t.Errorf("Expected 4 instruction lines, got %d header lines\n%q\n", len(instrLines), instrLines)
	}
	if instrLines[0] != "move 1 from 3 to 1" {
		t.Errorf("Not expected output, first line %s\n", instrLines[0])
	}
	if instrLines[2] != "move 2 from 3 to 1" {
		t.Errorf("Not expected output, first line %s\n", instrLines[2])
	}
}
