package main

import "testing"

func TestIsNice(t *testing.T) {
	tests := []struct {
		str    string
		isNice bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, test := range tests {
		got := isNice(test.str)
		if got != test.isNice {
			t.Errorf("Expected %t, got %t. Parameter: '%s'\n", test.isNice, got, test.str)
		}
	}
}

func TestIsNiceV2(t *testing.T) {
	tests := []struct {
		str    string
		isNice bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	for _, test := range tests {
		got := isNiceV2(test.str)
		if got != test.isNice {
			t.Errorf("Expected %t, got %t. Parameter: '%s'\n", test.isNice, got, test.str)
		}
	}
}
