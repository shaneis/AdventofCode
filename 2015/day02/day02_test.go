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
