package main

import "testing"

func TestMergeRanges(t *testing.T) {
	type args struct {
		ranges   []Range
		newRange Range
		output   []Range
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Merge overlapping ranges",
			args: args{
				ranges:   []Range{{Start: 1, End: 3}, {Start: 5, End: 7}, {Start: 8, End: 10}},
				newRange: Range{Start: 6, End: 9},
				output:   []Range{{Start: 1, End: 3}, {Start: 5, End: 10}},
			},
		},
		{
			name: "No overlapping ranges",
			args: args{
				ranges:   []Range{{Start: 1, End: 2}, {Start: 4, End: 5}},
				newRange: Range{Start: 6, End: 7},
				output:   []Range{{Start: 1, End: 2}, {Start: 4, End: 5}, {Start: 6, End: 7}},
			},
		},
		{
			name: "New range encompasses all",
			args: args{
				ranges:   []Range{{Start: 2, End: 3}, {Start: 5, End: 6}},
				newRange: Range{Start: 1, End: 7},
				output:   []Range{{Start: 1, End: 7}},
			},
		},
		{
			name: "Adjacent ranges",
			args: args{
				ranges:   []Range{{Start: 1, End: 3}, {Start: 4, End: 6}},
				newRange: Range{Start: 3, End: 4},
				output:   []Range{{Start: 1, End: 6}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeRanges(tt.args.ranges, tt.args.newRange); !equalRanges(got, tt.args.output) {
				t.Errorf("MergeRanges() = %v, want %v", got, tt.args.output)
			}
		})
	}
}

func equalRanges(a, b []Range) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
