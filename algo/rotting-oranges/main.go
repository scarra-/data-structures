package main

import "fmt"

type Coord struct {
	y, x, i int
}

type Queue struct {
	items []Coord
}

func NewQueue() *Queue {
	return &Queue{
		items: []Coord{},
	}
}

func (q *Queue) Enqueue(item Coord) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() Coord {
	item := q.items[0]
	q.items = q.items[1:]

	return item
}

func (q *Queue) Empty() bool {
	return len(q.items) == 0
}

// Given a grid of size m×n representing a field of cells, each cell can have one of three values:
//
//	0 representing an empty cell,
//	1 representing a fresh orange, or
//	2 representing a rotten orange.
//
// Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.
// Return the minimum number of minutes that must elapse until no cell has a fresh orange.
// If this is impossible, return -1 instead.
func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}

	grid2 := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 0, 1},
	}

	fmt.Println(rottingOranges(grid))
	fmt.Println(rottingOranges(grid2))
}

func rottingOranges(originGrid [][]int) int {
	grid := copyGrid(originGrid)
	h, w := len(grid), len(grid[0])

	rottenLocs := []Coord{}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == 2 {
				rottenLocs = append(rottenLocs, Coord{y, x, 0})
			}
		}
	}

	highestCycle := 0

	for _, rottenLoc := range rottenLocs {
		queue := NewQueue()
		queue.Enqueue(rottenLoc)

		for !queue.Empty() {
			item := queue.Dequeue()
			y, x := item.y, item.x

			if y < 0 || y > h-1 || x < 0 || x > w-1 || grid[y][x] != 1 && item.i != 0 {
				continue
			}

			grid[y][x] = 2

			if highestCycle < item.i {
				highestCycle = item.i
			}

			newCycle := item.i + 1

			queue.Enqueue(Coord{y: y + 1, x: x, i: newCycle})
			queue.Enqueue(Coord{y: y - 1, x: x, i: newCycle})
			queue.Enqueue(Coord{y: y, x: x + 1, i: newCycle})
			queue.Enqueue(Coord{y: y, x: x - 1, i: newCycle})
		}
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == 1 {
				return -1
			}
		}
	}

	return highestCycle
}

func copyGrid(originGrid [][]int) [][]int {
	grid := [][]int{}
	for _, v := range originGrid {
		sl := make([]int, len(v))
		copy(sl, v)
		grid = append(grid, sl)
	}

	return grid
}
