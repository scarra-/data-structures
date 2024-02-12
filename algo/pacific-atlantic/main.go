package main

import "fmt"

func main() {
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}

	// [0 0 0 0 1]
	// [0 0 0 1 1]
	// [0 0 1 0 0]
	// [1 1 0 0 0]
	// [1 0 0 0 0]
	result := pacificAtlantic(heights)

	for _, v := range result {
		fmt.Println(v)
	}
}

// Given an m×n matrix of non-negative integers representing the height of each unit cell in a continent,
// the "Pacific ocean" touches the left and top edges of the matrix, and the "Atlantic ocean" touches the right and bottom edges.
// Water can only flow in four directions (up, down, left, or right) from a cell to another one with height equal or lower.
// Find the list of grid coordinates where water can flow to both the Pacific and Atlantic ocean.
func pacificAtlantic(grid [][]int) [][]int {
	pacific := [][]int{}
	atlantic := [][]int{}
	result := [][]int{}

	for _, v := range grid {
		pacific = append(pacific, make([]int, len(v)))
		atlantic = append(atlantic, make([]int, len(v)))
		result = append(result, make([]int, len(v)))
	}

	h, w := len(grid), len(grid[0])

	for y := 0; y < h; y++ {
		dfs(grid, pacific, y, 0, grid[y][0])
		dfs(grid, atlantic, y, w-1, grid[y][w-1])
	}

	for x := 0; x < w; x++ {
		dfs(grid, pacific, 0, x, grid[0][x])
		dfs(grid, atlantic, h-1, x, grid[h-1][x])
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if pacific[y][x] == 1 && atlantic[y][x] == 1 {
				result[y][x] = 1
			}
		}
	}

	return result
}

func dfs(grid [][]int, visited [][]int, y, x, prevHeight int) {
	if x < 0 || x > len(grid[0])-1 || y < 0 || y > len(grid)-1 || grid[y][x] < prevHeight || visited[y][x] == 1 {
		return
	}

	visited[y][x] = 1

	dfs(grid, visited, y+1, x, grid[y][x])
	dfs(grid, visited, y-1, x, grid[y][x])
	dfs(grid, visited, y, x+1, grid[y][x])
	dfs(grid, visited, y, x-1, grid[y][x])
}
