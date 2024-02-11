package main

import "fmt"

// Given a 2d grid map of '1's (land) and '0's (water), count the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
// You may assume all four edges of the grid are all surrounded by water.
func numIslands(grid [][]byte) int {
	result := 0

	h, w := len(grid), len(grid[0])

	for x := 0; x < h; x++ {
		for y := 0; y < w; y++ {

			if grid[x][y] != 0 {
				result++
				clearIsland(grid, x, y)
			}
		}
	}

	return result
}

func clearIsland(grid [][]byte, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 {
		return
	}

	grid[x][y] = 0

	clearIsland(grid, x+1, y)
	clearIsland(grid, x-1, y)
	clearIsland(grid, x, y+1)
	clearIsland(grid, x, y-1)
}

func main() {
	grid := [][]byte{
		{0, 0, 2, 2, 0, 2, 2, 2, 2, 2},
		{0, 2, 2, 2, 0, 0, 0, 0, 0, 2},
		{0, 2, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0, 2, 0, 2},
		{0, 0, 0, 0, 0, 0, 0, 2, 0, 0},
		{2, 0, 0, 0, 2, 2, 2, 2, 0, 0},
	}

	fmt.Println(numIslands(grid))
}
