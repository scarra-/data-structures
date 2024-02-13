package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

type Queue []Point

func (q *Queue) Push(p Point) {
	*q = append(*q, p)
}

func (q *Queue) Pop() Point {
	p := (*q)[0]
	*q = (*q)[1:]
	return p
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

func shortestPath(grid [][]int, start, end Point) []Point {
	directions := []Point{
		{x: 1, y: 0},  // Up
		{x: 0, y: 1},  // Right
		{x: -1, y: 0}, // Down
		{x: 0, y: -1}, // Left
	}

	queue := Queue{}
	queue.Push(start)
	current := start

	visited := make(map[Point]Point)
	visited[start] = Point{x: -1, y: -1}

	for !queue.Empty() {
		current = queue.Pop()

		if current == end {
			break
		}

		for _, dir := range directions {
			next := Point{x: current.x + dir.x, y: current.y + dir.y}

			_, ok := visited[next]

			if !inGrid(grid, next) || ok || grid[next.x][next.y] == 1 {
				continue
			}

			visited[next] = current

			queue.Push(next)
		}
	}

	result := []Point{}

	for current != start {
		result = append(result, current)
		current = visited[current]
	}

	result = append(result, start)

	return result
}

func inGrid(grid [][]int, point Point) bool {
	return point.x >= 0 && point.x < len(grid) && point.y >= 0 && point.y < len(grid[0])
}

func main() {
	maze := [][]int{
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
	}

	start := Point{0, 0}
	end := Point{len(maze) - 1, len(maze[0]) - 1}

	shortestPath := shortestPath(maze, start, end)

	for _, point := range shortestPath {
		maze[point.x][point.y] = 3
	}

	printMaze(maze, shortestPath)
}

func printMaze(maze [][]int, shortestPath []Point) {
	mazeCopy := make([][]int, len(maze))
	for i := range maze {
		mazeCopy[i] = make([]int, len(maze[i]))
		copy(mazeCopy[i], maze[i])
	}

	for _, p := range shortestPath {
		mazeCopy[p.x][p.y] = -1
	}

	for i := range mazeCopy {
		for j := range mazeCopy[i] {
			switch mazeCopy[i][j] {
			case -1:
				fmt.Print("\033[32m*\033[0m ")
			case 0:
				fmt.Print(". ")
			default:
				fmt.Print("# ")
			}
		}
		fmt.Println()
	}
}
