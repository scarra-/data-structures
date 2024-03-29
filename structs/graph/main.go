package main

import "fmt"

// Graph represents adjacency list graph
type Graph struct {
	vertices []*Vertex
}

func NewGraph() *Graph {
	return &Graph{vertices: []*Vertex{}}
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		edges := []int{}
		for _, v := range v.adjacent {
			edges = append(edges, v.Vertex.key)
		}

		fmt.Printf("Vertex [%d], adjacent [%v]\n", v.key, edges)
	}
}

func (g *Graph) AddVertex(k int) {
	for key := range g.vertices {
		if key == k {
			return
		}
	}

	g.vertices = append(g.vertices, NewVertex(k))
}

func (g *Graph) GetVertex(k int) *Vertex {
	for key, v := range g.vertices {
		if key == k {
			return v
		}
	}

	return nil
}

func (g *Graph) AddEdge(from, to int) error {
	fromVertex := g.GetVertex(from)
	if fromVertex == nil {
		return fmt.Errorf("from vertex doesnt exist")
	}

	toVertex := g.GetVertex(to)
	if toVertex == nil {
		return fmt.Errorf("to vertex doesnt exist")
	}

	if fromVertex.EdgeExists(to) {
		return fmt.Errorf("from [%d] edge to [%d] already exists", from, to)
	}

	fromVertex.adjacent = append(fromVertex.adjacent, NewEdge(0, toVertex))

	return nil
}

type Vertex struct {
	key      int
	adjacent []*Edge
}

func NewVertex(k int) *Vertex {
	return &Vertex{key: k, adjacent: []*Edge{}}
}

type Edge struct {
	Weight int
	Vertex *Vertex
}

func NewEdge(w int, v *Vertex) *Edge {
	return &Edge{
		Weight: w,
		Vertex: v,
	}
}

func (v *Vertex) EdgeExists(key int) bool {
	for _, v := range v.adjacent {
		if v.Vertex.key == key {
			return true
		}
	}

	return false
}

// Add vertex
// Add edge

// func main() {
// 	graph := NewGraph()

// 	for i := 0; i < 5; i++ {
// 		graph.AddVertex(i)
// 	}

// 	graph.AddVertex(1)

// 	graph.AddEdge(1, 4)
// 	graph.AddEdge(1, 3)
// 	graph.AddEdge(2, 3)

// 	graph.Print()
// }

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
