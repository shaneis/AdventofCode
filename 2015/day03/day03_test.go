package main

import "testing"

func TestCountHousesDelivered(t *testing.T) {
	/*
		For example:

		> delivers presents to 2 houses: one at the starting location, and one to the east.
		^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
		^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.
	*/
	tests := []struct {
		input  string
		output int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}

	for _, test := range tests {
		got := countHousesDelivered(test.input)
		if got != test.output {
			t.Errorf("Expected %d, got %d. Parameters: '%s'\n", test.output, got, test.input)
		}
	}
}

func TestCountHousesDelivered2Santas(t *testing.T) {

	tests := []struct {
		input  string
		output int
	}{
		{">V", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}

	for _, test := range tests {
		got := countHousesDelivered2Santas(test.input)
		if got != test.output {
			t.Errorf("Expected %d, got %d. Parameters: '%s'\n", test.output, got, test.input)
		}
	}
}
