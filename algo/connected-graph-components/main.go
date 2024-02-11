package main

import "fmt"

// Number of Connected Components in an Undirected Graph:
// Given n nodes labeled from 0 to 1
// n−1 and a list of undirected edges (each edge is a pair of nodes),
// write a function to find the number of
// connected components in an undirected graph.
func main() {
	graph := map[int][]int{
		0: {1, 2},
		1: {0, 3},
		2: {0, 3},
		3: {1, 2},

		4: {5},
		5: {4},

		6: {7},
		7: {8},
		8: {},
	}
	components := countComponents(graph)
	fmt.Println("Number of connected components:", components) // Output: 2
}

func countComponents(graph map[int][]int) int {
	count := 0
	visited := make(map[int]bool)

	for v := range graph {
		if !visited[v] {
			dfs(graph, visited, v)
			count++
		}
	}

	return count
}

func dfs(graph map[int][]int, visited map[int]bool, node int) {
	visited[node] = true

	for _, neighbour := range graph[node] {
		if !visited[neighbour] {
			dfs(graph, visited, neighbour)
		}
	}
}
