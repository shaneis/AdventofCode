package main

import (
	"fmt"
	"testing"
)

func TestFindStartingMarker(t *testing.T) {
	var tests = []struct {
		stream                    string
		groupSize, markerLocation int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4, 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 4, 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 4, 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4, 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4, 11},
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14, 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 14, 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 14, 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14, 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14, 26},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Stream: %d - (%q:%d), returns %d\n", i, test.stream, test.groupSize, test.markerLocation), func(t *testing.T) {
			returnedMarkerLocation := findStartingMarker(test.stream, test.groupSize)

			if returnedMarkerLocation != test.markerLocation {
				t.Errorf("returnedMarkerLocation %d does not match expected markerLocation %d\n", returnedMarkerLocation, test.markerLocation)
			}
		})
	}
}
