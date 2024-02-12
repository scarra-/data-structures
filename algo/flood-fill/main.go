package main

import (
	"fmt"
	"time"
)

type Point struct {
	x, y int
}

type Queue struct {
	items []Point
}

func NewQueue() *Queue {
	return &Queue{
		items: []Point{},
	}
}

func (q *Queue) Enqueue(item Point) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() Point {
	item := q.items[0]
	q.items = q.items[1:]

	return item
}

func (q *Queue) Empty() bool {
	return len(q.items) == 0
}

func floodFill(grid [][]int, x, y int) [][]int {
	maxH := len(grid)
	maxW := len(grid[0])

	queue := NewQueue()
	queue.Enqueue(Point{x: x, y: y})

	for !queue.Empty() {
		p := queue.Dequeue()
		x, y := p.x, p.y

		if x < 0 || x >= maxH || y < 0 || y >= maxW || grid[x][y] != 0 {
			continue
		}

		grid[x][y] = 1

		fmt.Print("\033[H\033[2J")
		for i := 0; i < 10; i++ {
			fmt.Println()
		}

		fmt.Printf(
			"%v\n%v\n%v\n%v\n%v\n%v\n%v",
			grid[0], grid[1], grid[2], grid[3], grid[4], grid[5], grid[6],
		)

		time.Sleep(100 * time.Millisecond)

		queue.Enqueue(Point{x: x + 1, y: y})
		queue.Enqueue(Point{x: x, y: y - 1})
		queue.Enqueue(Point{x: x - 1, y: y})
		queue.Enqueue(Point{x: x, y: y + 1})
	}

	return grid
}

func printArray(array [][]int, startRow, startCol int) {
	// Move the cursor to the specified position
	fmt.Printf("\033[%d;%dH", startRow, startCol)

	// Print the array
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Printf("%d ", array[i][j])
		}
		// Move to the next line after printing each row
		fmt.Println()
	}
}

func main() {
	grid := [][]int{
		{0, 0, 2, 2, 2, 2, 2, 2, 2, 2},
		{0, 2, 2, 2, 0, 0, 0, 0, 0, 2},
		{0, 2, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 2, 0, 0, 0, 0, 0, 0, 2, 0},
		{0, 2, 0, 0, 0, 0, 0, 2, 0, 2},
		{0, 0, 2, 2, 0, 0, 0, 2, 0, 0},
		{0, 0, 0, 0, 2, 2, 2, 0, 0, 0},
	}

	floodFill(grid, 2, 5)
}
