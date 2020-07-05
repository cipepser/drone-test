package app

import (
	"testing"
)

type testCase struct {
	name   string
	input  int
	expect int
}

func TestAddTwo(t *testing.T) {
	tests := []testCase{
		{
			"simple",
			3,
			5,
		},
		//{
		//	"fatal",
		//	3,
		//	2,
		//},
	}

	for _, tt := range tests {
		actual := AddTwo(tt.input)
		if actual != tt.expect {
			t.Fatalf("%s failed: \nexpect: %v, actual: %v", tt.name, tt.expect, actual)
		}
	}
}
