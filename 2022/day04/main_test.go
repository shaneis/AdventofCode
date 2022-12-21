package main

import (
	"fmt"
	"testing"
)

// func TestPart01(t *testing.T) {
// 	res := part01("sample_input_01.txt")
// 	if res != 2 {
// 		t.Errorf("Function part01 did not return 2")
// 	}
// }

func TestParseLine(t *testing.T) {
	var tests = []struct {
		line   string
		result bool
	}{
		{"2-4,6-8", false},
		{"2-8,3-7", true},
	}

	for _, tst := range tests {
		testname := fmt.Sprintf("%q, %t", tst.line, tst.result)
		t.Run(testname, func(t *testing.T) {
			x := parseLine(tst.line)
			if x != tst.result {
				t.Errorf("got %t, expected %t", x, tst.result)
			}
		})
	}
}
