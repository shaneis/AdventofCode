package main

import "testing"

func TestFindXmas(t *testing.T) {
	tests := []struct {
		puzzle string
		output int
	}{
		{`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`, 18},
	}

	for testID, test := range tests {
		got := findXmas(test.puzzle)

		if got != test.output {
			t.Errorf("FAIL: %d - Expected %d, got %d\n", testID+1, test.output, got)
		}
	}
}

func TestOhAnXMas(t *testing.T) {
	tests := []struct {
		puzzle string
		output int
	}{
		{`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`, 9},
	}

	for testID, test := range tests {
		got := ohAnXMas(test.puzzle)

		if got != test.output {
			t.Errorf("FAIL: %d - Expected %d, got %d\n", testID+1, test.output, got)
		}
	}
}
