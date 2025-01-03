package main

import "testing"

func TestCleanInput(t *testing.T) {
	type testCases struct {
		input    string
		expected []string
	}
	tests := []testCases{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more casees here
	}

	for _, c := range tests {
		actual := cleanInput(c.input)
		if len()
		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			
		}
	}
}

