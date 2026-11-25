package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	// Test Cases go in testing func

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},

		{
			input:    " the car is driving!",
			expected: []string{"the", "car", "is", "driving!"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		// Check length of slice
		// If they don't match use t.Errorf and continue to the next case
		if len(actual) != (len(c.expected)) {
			fmt.Errorf("Actual doesn't match expected length. Test Failed")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Checks to make sure the words match then print error.
			if word != expectedWord {
				fmt.Errorf("Length of word doesn't match expected. Test Failed. Word %v, Expected %v", word, expectedWord)
			}
		}
	}
}
