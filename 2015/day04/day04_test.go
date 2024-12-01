package main

import "testing"

func TestGetMD5Hash(t *testing.T) {
	tests := []struct {
		key             string
		zeroLen, output int
	}{
		{"abcdef", 5, 609043},
		{"pqrstuv", 5, 1048970},
	}

	for _, test := range tests {
		got := getMD5Hash(test.key, test.zeroLen)
		if got != test.output {
			t.Errorf(
				"Expected %d, got %d. Parameter: '%s', %d\n",
				test.output,
				got,
				test.key,
				test.zeroLen,
			)
		}
	}
}
