package main

import "testing"

func TestCountFloors(t *testing.T) {
	tests := []struct {
		directions string
		output     int
	}{
		/*
					For example:

			(()) and ()() both result in floor 0.
			((( and (()(()( both result in floor 3.
			))((((( also results in floor 3.
			()) and ))( both result in floor -1 (the first basement level).
			))) and )())()) both result in floor -3.
		*/
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, test := range tests {
		got := countFloors(test.directions)
		if got != test.output {
			t.Errorf("countFloors('%s'): Expected %d, got %d\n", test.directions, test.output, got)
		}
	}
}

func TestFirstEnterBasement(t *testing.T) {
	/*
			For example:

		) causes him to enter the basement at character position 1.
		()()) causes him to enter the basement at character position 5
	*/
	tests := []struct {
		direction string
		output    int
	}{
		{")", 1},
		{"()())", 5},
	}

	for _, test := range tests {
		got := firstEnterBasement(test.direction)
		if got != test.output {
			t.Errorf("firstEnterBasement('%s'): Expected %d, got %d\n", test.direction, test.output, got)
		}
	}
}
