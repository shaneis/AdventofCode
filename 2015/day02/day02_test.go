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

func TestConvertDimensions(t *testing.T) {
	tests := []struct {
		input   string
		l, w, h int
	}{
		{"2x3x4", 2, 3, 4},
		{"1x1x102", 1, 1, 10},
	}

	for _, test := range tests {
		gotL, gotW, gotH := convertDimensions(test.input)
		if gotL != test.l {
			t.Errorf("Expected %d, got %d. Parameters: %s\n", test.l, gotL, test.input)
		}
		if gotW != test.w {
			t.Errorf("Expected %d, got %d. Parameters: %s\n", test.w, gotW, test.input)
		}
		if gotH != test.h {
			t.Errorf("Expected %d, got %d. Parameters: %s\n", test.h, gotH, test.input)
		}
	}
}
