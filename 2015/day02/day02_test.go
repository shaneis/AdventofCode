package main

import "testing"

func TestGetWrappingPaperArea(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"2x3x4", 58},
		{"1x1x10", 43},
	}

	for _, test := range tests {
		got := getWrappingPaperArea(test.input)
		if got != test.output {
			t.Errorf("Expected %d, got %d. Parameters: %s\n", test.output, got, test.input)
		}
	}
}

func TestGetRibbonSize(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"2x3x4", 34},
		{"1x1x10", 14},
	}

	for _, test := range tests {
		got := getRibbonSize(test.input)
		if got != test.input {
			t.Errorf("Expected %d, got %d. Parameters: %s\n", test.output, got, test.input)
		}
	}
}
