package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() int {
	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return value
}

func main() {
	stack := &Stack{items: []int{1, 2, 3, 4, 5}}

	stack.Push(6)
	stack.Push(7)
	stack.Push(8)

	fmt.Println(stack.items)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	fmt.Println(stack.items)
}
