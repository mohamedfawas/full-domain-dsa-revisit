package main

import "fmt"

/*
Use map-based (your implementation) when:

Number of vertices is unknown
Vertices can have arbitrary IDs
Graph is sparse
Need dynamic vertex addition/removal
*/

type Edge struct {
	From   int
	To     int
	Weight int
}

type Graph struct {
	// Using a map instead of a slice has some advantages:
	// 1. We don't need to know the number of vertices beforehand
	// 2. We can have non-consecutive vertex numbers
	// 3. Memory efficient for sparse graphs (graphs with few edges)
	Edges map[int][]Edge
}

// NewGraph creates a new weighted directed graph
func NewGraph() *Graph {
	return &Graph{
		Edges: make(map[int][]Edge),
	}
}

// AddEdge adds a new directed edge to the graph
func (g *Graph) AddEdge(from, to, weight int) {
	g.Edges[from] = append(g.Edges[from], Edge{
		From:   from,
		To:     to,
		Weight: weight,
	})
}

// GetNeighbors returns all vertices that can be reached directly from the given vertex
func (g *Graph) GetNeighbors(vertex int) []Edge {
	return g.Edges[vertex]
}

// HasVertex checks if a vertex exists in the graph
func (g *Graph) HasVertex(vertex int) bool {
	_, exists := g.Edges[vertex]
	return exists
}

// GetVertices returns all vertices in the graph
func (g *Graph) GetVertices() []int {
	vertices := make([]int, 0, len(g.Edges))
	for vertex := range g.Edges {
		vertices = append(vertices, vertex)
	}
	return vertices
}

// PrintGraph prints all edges in the graph
func (g *Graph) PrintGraph() {
	for from, edges := range g.Edges {
		fmt.Printf("From City %d:\n", from)
		for _, edge := range edges {
			fmt.Printf("  → To City %d (cost: %d)\n", edge.To, edge.Weight)
		}
	}
}

func main() {
	// Let's create a graph and add some edges
	graph := NewGraph()

	// Create a graph like this:
	//    5
	// 1 ---> 2
	// | \    |
	// |  \   | 2
	// |   \  |
	// 3 ---> 3
	//    1

	graph.AddEdge(1, 2, 5) // Road from city 1 to 2, cost 5
	graph.AddEdge(1, 3, 3) // Road from city 1 to 3, cost 3
	graph.AddEdge(2, 3, 2) // Road from city 2 to 3, cost 2

	fmt.Println("Our Graph:")
	graph.PrintGraph()

	fmt.Println("\nAll vertices in the graph:", graph.GetVertices())

	// Let's check neighbors of vertex 1
	fmt.Println("\nNeighbors of City 1:")
	for _, edge := range graph.GetNeighbors(1) {
		fmt.Printf("Can go to City %d with cost %d\n", edge.To, edge.Weight)
	}

	// Check if vertices exist
	fmt.Printf("\nDoes City 1 exist? %v\n", graph.HasVertex(1))
	fmt.Printf("Does City 4 exist? %v\n", graph.HasVertex(4))
}
