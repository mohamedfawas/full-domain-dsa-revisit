package main

import "fmt"

// Graph represents an adjacency list graph
// Using a map where key is the vertex and value is a slice of its neighbors
type Graph struct {
	vertices map[int][]int
}

// NewGraph creates and initializes a new Graph
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int][]int),
	}
}

// AddEdge adds an edge between vertex v1 and v2
func (g *Graph) AddEdge(v1, v2 int) {
	// Add v2 to v1's neighbor list
	g.vertices[v1] = append(g.vertices[v1], v2)
	// Add v1 to v2's neighbor list (for undirected graph)
	g.vertices[v2] = append(g.vertices[v2], v1)
}

// DFS performs Depth First Search starting from vertex start
func (g *Graph) DFS(start int) {
	// Create a map to keep track of visited vertices
	visited := make(map[int]bool)

	// Define a helper function that will be called recursively
	var dfsHelper func(vertex int)
	dfsHelper = func(vertex int) {
		// Mark current vertex as visited
		visited[vertex] = true

		// Print current vertex (or process it as needed)
		fmt.Printf("%d ", vertex)

		// Visit all neighbors of current vertex
		for _, neighbor := range g.vertices[vertex] {
			// If neighbor hasn't been visited, visit it
			if !visited[neighbor] {
				dfsHelper(neighbor)
			}
		}
	}

	// Start DFS from the start vertex
	dfsHelper(start)
	fmt.Println()
}

// BFS performs Breadth First Search starting from vertex start
func (g *Graph) BFS(start int) {
	// Create a map to keep track of visited vertices
	visited := make(map[int]bool)

	// Create a queue for BFS
	// In Go, we can use a slice as a queue
	queue := []int{start}

	// Mark the start vertex as visited
	visited[start] = true

	// Continue until queue is empty
	for len(queue) > 0 {
		// Remove the first element from queue (dequeue)
		vertex := queue[0]
		queue = queue[1:]

		// Print current vertex (or process it as needed)
		fmt.Printf("%d ", vertex)

		// Visit all neighbors of current vertex
		for _, neighbor := range g.vertices[vertex] {
			// If neighbor hasn't been visited:
			// 1. Mark it as visited
			// 2. Add it to queue
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	fmt.Println()
}

// Example usage
func main() {
	// Create a new graph
	graph := NewGraph()

	// Add edges to create this graph:
	//     1 --- 2
	//     |     |
	//     3 --- 4
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)

	fmt.Println("DFS starting from vertex 1:")
	graph.DFS(1)

	fmt.Println("BFS starting from vertex 1:")
	graph.BFS(1)
}
