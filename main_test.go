package main

import "testing"

func Test_distanceCalculation(t *testing.T) {

	tests := []struct {
		consumption int
		fuel int
		want int
	} {
		{10, 10, 90},
		{20,20,90},
		{30,30,90},
		{40,30,68},
	}
	for _, test := range tests {
		got := distanceCalculation(test.consumption, test.fuel)
		if got != test.want{
			t.Error("For fuel", test.fuel, "got", got, "want", test.want )
		}
	}
}