package main

import "testing"

func TestGetRotationPassword(t *testing.T) {
	tests := []struct {
		rotations []string
		output    int
	}{
		{[]string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}, 3},
	}

	for id, test := range tests {
		got := GetRotationPassword(test.rotations)
		if got != test.output {
			t.Errorf(
				"Test %d failed! Expected %d, got %d for rotations=%v\n",
				id,
				test.output,
				got,
				test.rotations,
			)
		}
	}
}

func TestGetRotationPassword2(t *testing.T) {
	tests := []struct {
		rotations []string
		output    int
	}{
		{[]string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}, 6},
		{[]string{"R11", "R44"}, 1},
		{[]string{"R2", "L48"}, 0},
		{[]string{"L48", "L574"}, 6},
		{[]string{"L48", "L500"}, 5},
		{[]string{"R500"}, 5},
	}

	for id, test := range tests {
		got := GetRotationPassword2(test.rotations)
		if got != test.output {
			t.Errorf(
				"Test %d failed! Expected %d, got %d for rotations=%v\n",
				id,
				test.output,
				got,
				test.rotations,
			)
		}
	}
}

func TestGetRotationPasswordBruteForce(t *testing.T) {
	tests := []struct {
		rotations []string
		output    int
	}{
		{[]string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}, 6},
		{[]string{"R11", "R44"}, 1},
		{[]string{"R2", "L48"}, 0},
		{[]string{"L48", "L574"}, 6},
		{[]string{"L48", "L500"}, 5},
		{[]string{"R500"}, 5},
	}

	for id, test := range tests {
		got := GetRotationPasswordBruteForce(test.rotations)
		if got != test.output {
			t.Errorf(
				"Test %d failed! Expected %d, got %d for rotations=%v\n",
				id,
				test.output,
				got,
				test.rotations,
			)
		}
	}
}
