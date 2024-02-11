package main

import "fmt"

type MaxHeap struct {
	array []int
}

func NewHeap() *MaxHeap {
	return &MaxHeap{array: []int{}}
}

func (h *MaxHeap) Insert(val int) {
	h.array = append(h.array, val)

	h.heapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) heapifyUp(index int) {
	fmt.Println(index, parent(index))
	for h.array[index] > h.array[parent(index)] {
		h.swap(index, parent(index))

		index = parent(index)
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *MaxHeap) Extract() int {
	val := h.array[0]

	lastIndex := len(h.array) - 1

	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	h.heapifyDown(0)

	return val
}

func (h MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	childToCompare := 0
	l, r := left(index), right(index)

	for l <= lastIndex {
		if lastIndex == l {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func main() {
	heap := NewHeap()

	for _, v := range []int{1, 3, 5, 7, 3, 2, 9, 6} {
		heap.Insert(v)
	}

	fmt.Println(heap)

	for i := 0; i < 3; i++ {
		heap.Extract()
		fmt.Println(heap)
	}
}
