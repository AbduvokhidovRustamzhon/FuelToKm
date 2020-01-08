package main

import "testing"

func Test_distanceCalculation(t *testing.T) {

	tests := []struct {
		name string
		consumption int
		fuel int
		want int
	} {
		{"The distance it travels for the available fuel", 100, 10,9},
	}
	for _, test := range tests {
		got := distanceCalculation(test.consumption, test.fuel)
		if got != test.want{
			t.Error("For fuel", test.fuel, "got", got, "want", test.want )
		}
	}
}