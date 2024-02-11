package main

import "fmt"

func main() {
	arr := []int{17, 2, 3, 5, 18, 9, 81, 7, 64}
	size := len(arr)

	for i := 0; i < size; i++ {
		min := i
		for j := i + 1; j < size; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}

		arr[min], arr[i] = arr[i], arr[min]
	}

	fmt.Println(arr)
}
