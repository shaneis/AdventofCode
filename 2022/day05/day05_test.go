package main

import (
	"testing"
)

func TestParseHeader(t *testing.T) {
	var testHeader = []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
	}

	header := parseHeader(testHeader)
	if len(header) != 6 {
		t.Errorf("Expected 6 structs, got %d\n", len(header))
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

func TestDelete(t *testing.T) {
	var testheader = []Header{
		{Column: 1, Location: 1, crateName: "D"},
		{Column: 1, Location: 2, crateName: "S"},
	}
	var testColID int = 1

	returnedHeaders, returnedCV := delete(testheader, testColID)

	if len(returnedHeaders) != 1 {
		t.Error("length of returned headers is not equal to 1")
	}
	if returnedCV != "D" {
		t.Errorf("returned crateValue is %q, not D\n", returnedCV)
	}
	if returnedHeaders[0].Column != 1 {
		t.Errorf("Column returned is unexpected\n")
	}
	if returnedHeaders[0].Location != 2 {
		t.Errorf("Returned location is unexpected: %d\n", returnedHeaders[0].Location)
	}
	if returnedHeaders[0].crateName != "S" {
		t.Errorf("Returned location is unexpected: %q\n", returnedHeaders[0].crateName)
	}
}

func TestAdd(t *testing.T) {
	var testheader = []Header{
		{Column: 1, Location: 2, crateName: "D"},
		{Column: 1, Location: 3, crateName: "S"},
	}

	returnedHeader := add(testheader, "F", 1)

	if returnedHeader.crateName != "F" {
		t.Errorf("returned crate name is unexpected, returned:%q\n", returnedHeader.crateName)
	}
	if returnedHeader.Location != 1 {
		t.Errorf("returned location is unexpected, returned:%d\n", returnedHeader.Location)
	}
	if returnedHeader.Column != 1 {
		t.Errorf("returned Column is unexpected, returned:%d\n", returnedHeader.Column)
	}
}
