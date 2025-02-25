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

// ShortestPath finds the shortest path distance from start to end using BFS.
// It returns the number of edges in the shortest path or -1 if no path exists.
func (g *Graph) ShortestPath(start, end int) int {
	// Check if the start or end vertices exist in the graph.
	// If either vertex doesn't exist, return -1 (no path).
	// Example: If the graph has vertices {1: [2], 2: [1]}, and start=3, return -1.
	if _, exists := g.vertices[start]; !exists {
		return -1
	}
	if _, exists := g.vertices[end]; !exists {
		return -1
	}

	// If the start and end vertices are the same, the distance is 0.
	// Example: If start=2 and end=2, return 0.
	if start == end {
		return 0
	}

	// Initialize BFS:
	// - queue: A slice to store vertices to visit next. Start with the start vertex.
	// - distance: A map to store the shortest distance from start to each vertex.
	// Example: If start=1, queue = [1], distance = {1: 0}.
	queue := []int{start}
	distance := make(map[int]int)
	distance[start] = 0

	// Process the queue until it's empty.
	for len(queue) > 0 {
		// Get the first vertex from the queue (current vertex).
		// Example: If queue = [1, 2], current = 1, and queue becomes [2].
		current := queue[0]
		queue = queue[1:]

		// Explore all neighbors of the current vertex.
		// Example: If current=1 and neighbors=[2, 3], loop through 2 and 3.
		for _, neighbor := range g.vertices[current] {
			// If the neighbor is the end vertex, return the distance.
			// Distance to end = distance to current + 1.
			// Example: If current=2, neighbor=3, and end=3, return distance[2] + 1.
			if neighbor == end {
				return distance[current] + 1
			}

			// Check if the neighbor has been visited.
			// If not visited, add it to the queue and update its distance.
			// Example: If neighbor=3 and it's not in the distance map:
			// - distance[3] = distance[2] + 1
			// - queue = [3]
			if _, visited := distance[neighbor]; !visited {
				distance[neighbor] = distance[current] + 1
				queue = append(queue, neighbor)
			}
		}
	}

	// If the queue is empty and we haven't found the end vertex, it's unreachable.
	// Return -1 to indicate no path exists.
	// Example: If start=1 and end=4, but 4 is not connected to 1, return -1.
	return -1
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
