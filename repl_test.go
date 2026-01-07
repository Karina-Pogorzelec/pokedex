package main

import "testing"

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
			input:    "  Hello  World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello  WORLD",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Count of words does not match, got: %v, want: %v", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("the word does not match the expected output, got: %s, want %s", word, expectedWord)
			}
		}
	}
}