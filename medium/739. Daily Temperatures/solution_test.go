package main

import (
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		temperatures []int
		expected     []int
	}{
		{
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected:     []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			temperatures: []int{30, 40, 50, 60},
			expected:     []int{1, 1, 1, 0},
		},
		{
			temperatures: []int{30, 60, 90},
			expected:     []int{1, 1, 0},
		},
	}

	for _, test := range tests {
		result := dailyTemperatures(test.temperatures)
		if !equal(result, test.expected) {
			t.Errorf("For input %v, expected %v but got %v", test.temperatures, test.expected, result)
		}
	}
}

// Helper function to compare two slices for equality
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
