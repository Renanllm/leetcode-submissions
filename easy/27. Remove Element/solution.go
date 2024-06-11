package main

func removeElement(nums []int, val int) int {
	aux := make([]int, 0)
	for _, num := range nums {
		if num != val {
			aux = append(aux, num)
		}
	}
	copy(nums, aux)
	return len(aux)
}
