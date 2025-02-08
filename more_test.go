package main

import (
	"testing"
	"reflect"
)


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{	
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		
		// add more cases here
		
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := cleanInput(c.input)
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
}


