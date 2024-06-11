package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},   // edge case: empty string
		{"(", false}, // edge case: single open bracket
		{")", false}, // edge case: single close bracket
		{"[{()}]", true},
		{"[{)]}", false},
		{"[({})]", true},
		{"[[[]]]", true},
		{"{{{{}}}}", true},
		{"{[()]}", true},
		{"[{()()[]}]", true},
		{"[", false},           // edge case: single open bracket
		{"]", false},           // edge case: single close bracket
		{"[{()()[]}[]", false}, // missing closing bracket
		{"[{()()}]]", false},   // extra closing bracket
		{"[{(()})", false},     // misnested brackets
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := isValid(test.input) // Replace with the actual function to test
			if result != test.expected {
				t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
			}
		})
	}
}
