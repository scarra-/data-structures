package main

import "fmt"

func main() {
	arr := []int{17, 2, 3, 5, 18, 9, 81, 7, 64}

	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

	fmt.Println(arr)
}
