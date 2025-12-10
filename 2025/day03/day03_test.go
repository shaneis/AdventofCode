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

func TestFindMaxPossibleJoltage12(t *testing.T) {
	tests := []struct {
		bank string
		want int
	}{
		{"987654321111111", 987654321111},
		{"811111111111119", 811111111119},
		{"234234234234278", 434234234278},
		{"818181911112111", 888911112111},
	}
	for _, tt := range tests {
		got := FindMaxPossibleJoltage12(tt.bank)
		if got != tt.want {
			t.Errorf("FindMaxPossibleJoltage12(%q) = %d; want %d", tt.bank, got, tt.want)
		}
	}
}
