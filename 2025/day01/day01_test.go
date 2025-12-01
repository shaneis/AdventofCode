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
