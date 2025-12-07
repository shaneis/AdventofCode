package main

import (
	"fmt"
	"testing"
)

func TestFindMaxPossibleJoltage(t *testing.T) {
	tests := []struct {
		bank string
		want int
	}{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
	}
	for _, tt := range tests {
		fmt.Printf("Testing bank: %s\n", tt.bank)
		got := FindMaxPossibleJoltage(tt.bank)
		if got != tt.want {
			t.Errorf("FindMaxPossibleJoltage(%q) = %d; want %d", tt.bank, got, tt.want)
		}
	}
}
