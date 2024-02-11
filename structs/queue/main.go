package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	item := q.items[0]
	q.items = q.items[1:]

	return item
}

func main() {
	queue := &Queue{items: []int{1, 2, 3, 4, 5}}

	queue.Enqueue(6)
	queue.Enqueue(7)
	queue.Enqueue(8)

	fmt.Println(queue.items)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

	fmt.Println(queue.items)
}
