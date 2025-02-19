package main

import "fmt"

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

// HasCycle checks if the undirected graph contains any cycle.
// Returns true if a cycle is found, false otherwise.
//
// Example graph with cycle:
//    1 --- 2
//    |     |
//    4 --- 3
// In this case, HasCycle() returns true because 1-2-3-4-1 forms a cycle.
//
// Example graph without cycle:
//    1 --- 2
//    |
//    4     3
// In this case, HasCycle() returns false as there are no cycles.
func (g *Graph) HasCycle() bool {
	// Create a map to keep track of visited vertices
	// For example, if we've visited vertices 1 and 2:
	// visited = {1: true, 2: true}
	visited := make(map[int]bool)

	// Helper function for DFS traversal
	// Parameters:
	// - current: the vertex we're currently examining
	// - parent: the vertex we came from (-1 for starting vertex)
	// Returns true if a cycle is detected, false otherwise
	var dfsHelper func(current, parent int) bool
	dfsHelper = func(current, parent int) bool {
		// Mark current vertex as visited
		// Example: If current = 2, then visited[2] = true
		visited[current] = true

		// Visit all neighbors of current vertex
		// Example: If current = 1 and its neighbors are [2, 4],
		// we'll check each of these neighbors
		for _, neighbor := range g.vertices[current] {
			// If neighbor hasn't been visited, explore it
			// Example: If we're at vertex 1 and neighbor 2 is unvisited,
			// we'll recursively explore vertex 2
			if !visited[neighbor] {
				// Recursively visit the neighbor
				// Example: If we're at vertex 1, we'll call dfsHelper(2, 1)
				if dfsHelper(neighbor, current) {
					return true // Cycle detected in deeper recursion
				}
			} else if neighbor != parent {
				// If we find a visited neighbor that's not our parent,
				// we've found a cycle
				//
				// Example:
				// 1 --- 2
				// |     |
				// 4 --- 3
				// If we're at vertex 3 and see vertex 1 is visited but
				// not our parent (2), we've found a cycle
				return true
			}
		}
		return false
	}

	// Check from each vertex in case graph is not fully connected
	// Example: If we have two separate components:
	//    1 --- 2     3 --- 4
	// We need to check both components
	for vertex := range g.vertices {
		// Only start DFS from unvisited vertices
		// Example: After checking component with vertices 1-2,
		// we'll start new DFS from vertex 3 if it's unvisited
		if !visited[vertex] {
			// Start DFS with no parent (-1)
			if dfsHelper(vertex, -1) {
				return true
			}
		}
	}
	return false
}

/* Usage Example:
func main() {
    g := NewGraph()

    // Create a graph with cycle:
    //    1 --- 2
    //    |     |
    //    4 --- 3
    g.AddEdge(1, 2)
    g.AddEdge(2, 3)
    g.AddEdge(3, 4)
    g.AddEdge(4, 1)

    if g.HasCycle() {
        fmt.Println("Cycle detected!") // This will print
    }

    g2 := NewGraph()

    // Create a graph without cycle:
    //    1 --- 2
    //    |
    //    4     3
    g2.AddEdge(1, 2)
    g2.AddEdge(1, 4)
    g2.AddEdge(3, 3)

    if !g2.HasCycle() {
        fmt.Println("No cycle found!") // This will print
    }
}
*/
