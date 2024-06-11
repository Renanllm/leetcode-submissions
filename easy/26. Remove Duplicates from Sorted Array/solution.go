package main

import "slices"

func removeDuplicates(nums []int) int {
	aux := make([]int, 0)
	for _, num := range nums {
		if !slices.Contains(aux, num) {
			aux = append(aux, num)
		}
	}
	copy(nums, aux)
	return len(aux)
}
