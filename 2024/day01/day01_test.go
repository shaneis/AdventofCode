package main

import "testing"

func TestSplitListsIntoTwo(t *testing.T) {
	tests := []struct {
		input  string
		output [][]int
	}{
		{`3   4
4   3
2   5
1   3
3   9
3   3`, [][]int{{3, 4, 2, 1, 3, 3}, {4, 3, 5, 3, 9, 3}}},
	}

	for _, test := range tests {
		got := splitListsIntoTwo(test.input)
		if len(got) == 0 {
			t.Errorf("Output should not be empty for '%s'\n", test.input)
		}
		if len(got[0]) != len(got[1]) {
			t.Errorf("Expected same output length (%d), got different (%d)\n", len(got[0]), len(got[1]))
		}
		if len(got[0]) != len(test.output[0]) || len(got[1]) != len(test.output[1]) {
			t.Errorf("Expected same output length (%d), got different (%d)\n", got, test.output)
		}
	}
}

func TestGetDifferenceBetweenLists(t *testing.T) {
	tests := []struct {
		str string
		out int
	}{
		{`3   4
4   3
2   5
1   3
3   9
3   3`, 11},
	}
	for _, test := range tests {
		lists := splitListsIntoTwo(test.str)
		got := getDifferenceBetweenLists(lists[0], lists[1])
		if got != test.out {
			t.Errorf("Expected %d, got %d. Parameters '%s', into %d and %d\n", test.out, got, test.str, lists[0], lists[1])
		}
	}
}

func TestGetSimilarityScore(t *testing.T) {
	tests := []struct {
		str string
		out int
	}{
		{`3   4
4   3
2   5
1   3
3   9
3   3`, 31},
	}

	for _, test := range tests {
		lists := splitListsIntoTwo(test.str)
		got := getSimilarityScore(lists[0], lists[1])
		if got != test.out {
			t.Errorf("Expected %d, got %d\n", test.out, got)
		}
	}
}
