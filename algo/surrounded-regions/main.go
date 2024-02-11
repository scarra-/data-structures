package main

import (
	"fmt"
)

// Given a 2D board containing 'X' and 'O' (the letter O),
// capture all regions surrounded by 'X'.
// A region is captured by flipping all 'O's into 'X's in that surrounded region.
// You can solve this problem using a depth-first search (DFS)
// or breadth-first search (BFS) approach.
func fillSurrounded(grid [][]byte) {
	h, w := len(grid), len(grid[0])

	for y := 0; y < h; y++ {
		fillConnected(grid, '#', y, 0)
		fillConnected(grid, '#', y, w-1)
	}

	for x := 0; x < w; x++ {
		fillConnected(grid, '#', 0, x)
		fillConnected(grid, '#', h-1, x)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			switch grid[i][j] {
			case '#':
				grid[i][j] = 'O'
			case 'O':
				grid[i][j] = '3'
			}
		}
	}
}

func fillConnected(grid [][]byte, symbol byte, y, x int) {
	if y < 0 || y > len(grid)-1 || x < 0 || x > len(grid[0])-1 || grid[y][x] != 'O' {
		return
	}

	grid[y][x] = symbol

	fillConnected(grid, symbol, y+1, x)
	fillConnected(grid, symbol, y-1, x)
	fillConnected(grid, symbol, y, x+1)
	fillConnected(grid, symbol, y, x-1)
}

func main() {
	grid := [][]byte{
		{'O', 'O', 'X', 'X', 'X', 'O', 'X', 'X', 'X', 'X'},
		{'O', 'X', 'X', 'X', 'O', 'O', 'X', 'O', 'O', 'X'},
		{'O', 'X', 'O', 'O', 'X', 'O', 'X', 'O', 'O', 'X'},
		{'O', 'X', 'O', 'O', 'X', 'O', 'O', 'X', 'X', 'O'},
		{'O', 'O', 'X', 'X', 'X', 'O', 'O', 'X', 'O', 'X'},
		{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'X', 'O', 'O'},
		{'O', 'O', 'O', 'O', 'X', 'X', 'X', 'O', 'O', 'O'},
	}

	for _, v := range grid {
		fmt.Println(string(v))
	}

	fmt.Println()

	fillSurrounded(grid)

	for _, v := range grid {
		fmt.Println(string(v))
	}
}
