package main

import "testing"

func TestFindInvalidIds(t *testing.T) {
	tests := []struct {
		numberRange string
		count       int
		invalids    []int
	}{
		{"11-22", 2, []int{11, 22}},
		{"95-115", 1, []int{99}},
		{"998-1012", 1, []int{1010}},
		{"1188511880-1188511890", 1, []int{1188511885}},
		{"222220-222224", 1, []int{222222}},
		{"1698522-1698528", 0, []int{}},
		{"446443-446449", 1, []int{446446}},
		{"38593856-38593862", 1, []int{38593859}},
	}
	for id, test := range tests {
		invalids := findInvalidIds(test.numberRange)
		if len(invalids) != test.count {
			t.Errorf("Test %d: Expected %d invalid IDs, got %d", id, test.count, len(invalids))
		}
		for i, inval := range invalids {
			if inval != test.invalids[i] {
				t.Errorf("Test %d: Expected invalid ID %d at index %d, got %d", id, test.invalids[i], i, inval)
			}
		}
	}
}

func TestFindInvalidIdsAtLeastTwice(t *testing.T) {
	tests := []struct {
		numberRange string
		count       int
		invalids    []int
	}{
		{"11-22", 2, []int{11, 22}},
		{"95-115", 2, []int{99, 111}},
		{"998-1012", 2, []int{999, 1010}},
		{"1188511880-1188511890", 1, []int{1188511885}},
		{"222220-222224", 1, []int{222222}},
		{"1698522-1698528", 0, []int{}},
		{"446443-446449", 1, []int{446446}},
		{"38593856-38593862", 1, []int{38593859}},
		{"565653-565659", 1, []int{565656}},
		{"824824821-824824827", 1, []int{824824824}},
		{"2121212118-2121212124", 1, []int{2121212121}},
	}
	for id, test := range tests {
		invalids := findInvalidIdsAtLeastTwice(test.numberRange)
		if len(invalids) != test.count {
			t.Errorf("Test %d: Expected %d invalid IDs, got %d", id, test.count, len(invalids))
		}
		for i, inval := range invalids {
			if inval != test.invalids[i] {
				t.Errorf("Test %d: Expected invalid ID %d at index %d, got %d", id, test.invalids[i], i, inval)
			}
		}
	}
}
