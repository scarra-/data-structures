package main

import "fmt"

func main() {
	fmt.Println("Result - ", CalculateFibSequentially(7))
}

func CalculateFibSequentially(n uint) uint {
	// 0 1 1 2 3 5 8 13 21 34 55

	if n < 2 {
		return n
	}

	to := int(n)

	var a, b uint
	b = 1

	for i := 0; i < to; i++ {
		a += b

		a, b = b, a
	}

	return a
}
