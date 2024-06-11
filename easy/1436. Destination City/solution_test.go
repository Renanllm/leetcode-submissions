package main

import "testing"

func TestDestCity(t *testing.T) {
	tests := []struct {
		paths    [][]string
		expected string
	}{
		{
			paths:    [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}},
			expected: "Sao Paulo",
		},
		{
			paths:    [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}},
			expected: "A",
		},
		{
			paths:    [][]string{{"A", "Z"}},
			expected: "Z",
		},
	}

	for _, tt := range tests {
		result := DestCity(tt.paths)
		if result != tt.expected {
			t.Errorf("DestCity(%v) = %s; want %s", tt.paths, result, tt.expected)
		}
	}
}
