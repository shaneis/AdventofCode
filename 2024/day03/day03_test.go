package main

import "testing"

func TestParseMulString(t *testing.T) {
	tests := []struct {
		str    string
		output []string
	}{
		{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}},
		{"xmul(2,4)mul(),m()", []string{"mul(2,4)"}},
	}

	for testID, test := range tests {
		got := parseMulString(test.str)

		if len(got) != len(test.output) {
			t.Errorf("FAIL: test %d. Expected output of count %d, got count %d\n", testID+1, len(test.output), len(got))
			break
		}

		for i := 0; i < len(got); i++ {
			if got[i] != test.output[i] {
				t.Errorf("FAIL: test %d. Expected item %d to be '%s', got '%s'\n", testID+1, i+1, test.output[i], got[i])
			}
		}
	}
}

func TestInvokeMulString(t *testing.T) {
	tests := []struct {
		str    string
		output int
	}{
		{"mul(2,4)", 8},
		{"mul(5,5)", 25},
		{"mul(11,8)", 88},
		{"mul(8,5)", 40},
	}

	for testID, test := range tests {
		got := invokeMulString(test.str)

		if got != test.output {
			t.Errorf("FAIL: test %d. Expected %d, got %d\n", testID+1, test.output, got)
		}
	}
}

func TestParseDontDoWhatJohnnyDontDoDoes(t *testing.T) {
	tests := []struct {
		str    string
		output []string
	}{
		{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", []string{"mul(2,4)", "mul(8,5)"}},
		{"mul(2,4)mul(3,7)", []string{"mul(2,4)", "mul(3,7)"}},
		{"[mul(917/;~(:mul(730,865)^]-'when()-;mul(115,72)", []string{"mul(730,865)", "mul(115,72)"}},
	}

	for testID, test := range tests {
		got := parseDontDoWhatJohnnyDontDoDoes(test.str)

		if len(got) != len(test.output) {
			t.Errorf("FAIL: test %d. Expected output of count %d, got count %d\n", testID+1, len(test.output), len(got))
			break
		}

		for i := 0; i < len(got); i++ {
			if got[i] != test.output[i] {
				t.Errorf("FAIL: test %d. Expected item %d to be '%s', got '%s'\n", testID+1, i+1, test.output[i], got[i])
			}
		}
	}
}
