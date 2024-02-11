package main

import "fmt"

func main() {
	arr := []int{17, 2, 3, 5, 18, 9, 81, 7, 64}
	size := len(arr)

	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Println(arr)
}
