package main

import (
	"fmt"
	"testing"
)

func TestFindStartingMarker(t *testing.T) {
	var tests = []struct {
		stream         string
		markerLocation int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d - %q, returns %d\n", i, test.stream, test.markerLocation), func(t *testing.T) {
			returnedMarkerLocation := findStartingMarker(test.stream)

			if returnedMarkerLocation != test.markerLocation {
				t.Errorf("returnedMarkerLocation %d does not match expected markerLocation %d\n", returnedMarkerLocation, test.markerLocation)
			}
		})
	}
}
