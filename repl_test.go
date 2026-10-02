package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "persistence and  determination",
			expected: []string{"persistence", "and", "determination"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}

	for i, c := range cases {
		actual := cleanInput(c.input)
		// check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			t.Errorf("test case %v expected slice length: %v, received slice length: %v", i, len(c.expected), len(actual))
			continue
		}
		for j := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// check each word in the slice
			if !reflect.DeepEqual(word, expectedWord) {
				// if they don't match, use t.Errorf to print and error message
				t.Errorf("word %v in test case %v expected %v received %v", j, i, expectedWord, word)
			}

			// then fail the test
		}
	}
}
