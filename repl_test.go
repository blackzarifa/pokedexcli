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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "single",
			expected: []string{"single"},
		},
		{
			input:    "Multiple   Spaces   Between   Words",
			expected: []string{"multiple", "spaces", "between", "words"},
		},
		{
			input:    "tabs\tand\tnewlines\n",
			expected: []string{"tabs", "and", "newlines"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %d words, expected %d words", c.input, len(actual), len(c.expected))
			continue
		}

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("cleanInput(%q) = %v, expected %v", c.input, actual, c.expected)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("cleanInput(%q)[%d] = %q, expected %q", c.input, i, word, expectedWord)
			}
		}
	}
}
