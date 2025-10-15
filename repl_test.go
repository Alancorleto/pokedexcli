package main

import (
	"fmt"
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
			input:    "foo bar baz",
			expected: []string{"foo", "bar", "baz"},
		},
		{
			input:    "   singleword   ",
			expected: []string{"singleword"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "    ",
			expected: []string{},
		},
		{
			input:    "multiple   spaces   between   words",
			expected: []string{"multiple", "spaces", "between", "words"},
		},
		{
			input:    " leading and trailing spaces ",
			expected: []string{"leading", "and", "trailing", "spaces"},
		},
		{
			input:    "Pikachu   Charizard  Bulbasaur ",
			expected: []string{"pikachu", "charizard", "bulbasaur"},
		},
		{
			input:    "  MixedCASE Input  ",
			expected: []string{"mixedcase", "input"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) == %q, expected %q", c.input, actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			fmt.Println("word: ", word)
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%q) == %q, expected %q", c.input, actual, c.expected)
				break
			}
		}
	}
}
