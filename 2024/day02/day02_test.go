package main

import "testing"

func TestIsReportSafe(t *testing.T) {
	tests := []struct {
		input  []int
		isSafe bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, test := range tests {
		got := isReportSafe(test.input)
		if got != test.isSafe {
			t.Errorf("Expected %t, got %t. Parameters : %d\n", test.isSafe, got, test.input)
		}
	}
}
