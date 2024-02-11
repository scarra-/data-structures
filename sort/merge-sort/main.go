package main

import "fmt"

func main() {
	arr := []int{17, 2, 3, 5, 18, 9, 81, 7, 64}

	// Big O complexity: O(n*log(n)).
	arr = mergeSort(arr)

	fmt.Println(arr)
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	left := mergeSort(items[:len(items)/2])
	right := mergeSort(items[len(items)/2:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}
