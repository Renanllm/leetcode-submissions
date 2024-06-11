package main

import (
	"testing"
)

func TestSumOfMinSubarrays(t *testing.T) {
	mod := int(1e9 + 7)

	tests := []struct {
		arr      []int
		expected int
	}{
		{
			arr:      []int{3, 1, 2, 4},
			expected: 17,
		},
		{
			arr:      []int{11, 81, 94, 43, 3},
			expected: 444,
		},
		{
			arr:      []int{1},
			expected: 1,
		},
		{
			arr:      []int{2, 4, 6, 8, 10},
			expected: 50,
		},
		{
			arr:      []int{10, 9, 8, 7, 6, 5},
			expected: 84,
		},
		{
			arr:      []int{100, 100, 100, 100, 100},
			expected: 1500,
		},
		{
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 220,
		},
		{
			arr:      []int{3 * 104, 3 * 104, 3 * 104},
			expected: (3 * 104) * 6 % mod,
		},
		{
			arr:      []int{3 * 104, 1, 3 * 104},
			expected: (3*104 + 1 + 3*104 + 1 + 1 + 1) % mod,
		},
	}

	for _, test := range tests {
		result := sumOfMinSubarrays(test.arr)
		if result != test.expected {
			t.Errorf("For input %v, expected %d but got %d", test.arr, test.expected, result)
		}
	}
}
