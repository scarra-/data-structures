package main

import "fmt"

func main() {
	grid := [][]byte{
		{'k', 'a', 'r', 'o'},
		{'e', 't', 'k', 'a'},
		{'y', 'e', 'a', 'r'},
		{'h', 'h', 'd', 'd'},
		{'h', 'd', 'e', 'd'},
		{'h', 'y', 'a', 'r'},
	}

	fmt.Println(search(grid, "oak"))   // true
	fmt.Println(search(grid, "atr"))   // false
	fmt.Println(search(grid, "heard")) // true
	fmt.Println(search(grid, "test"))  // false
	fmt.Println(search(grid, "dear"))  // true
}

// Given a 2D board and a word, find if the word exists in the grid.
// The word can be constructed from letters of sequentially adjacent cells,
// where "adjacent" cells are horizontally or vertically neighboring.
// The same letter cell may not be used more than once.
func search(g [][]byte, word string) bool {
	grid := copySlice(g)
	h, w := len(grid), len(grid[0])

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if bfs(grid, word, y, x) {
				return true
			}
		}
	}

	return false
}

func bfs(grid [][]byte, word string, y, x int) bool {
	if len(word) == 0 {
		return true
	}

	if y < 0 || y > len(grid)-1 || x < 0 || x > len(grid[0])-1 || grid[y][x] != word[0] {
		return false
	}

	grid[y][x] = '#'

	return bfs(grid, word[1:], y, x+1) ||
		bfs(grid, word[1:], y, x-1) ||
		bfs(grid, word[1:], y+1, x) ||
		bfs(grid, word[1:], y-1, x)
}

func copySlice(g [][]byte) [][]byte {
	res := make([][]byte, len(g))
	for i := range g {
		res[i] = make([]byte, len(g[i]))
		copy(res[i], g[i])
	}

	return res
}
