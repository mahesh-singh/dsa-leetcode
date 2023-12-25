package main

import "fmt"

type Vertex struct {
	Val       int
	Neighbors []*Vertex
}

type Graph struct {
	Vertices []*Vertex
}

func TopologicalSort(graph *Graph) []int {
	visited := make(map[int]bool)
	stack := []int{}

	for _, vertex := range graph.Vertices {

		dfs(vertex, &visited, &stack)
	}

	return stack
}

func dfs(vertex *Vertex, visited *map[int]bool, stack *[]int) {
	if (*visited)[vertex.Val] {
		return
	}
	(*visited)[vertex.Val] = true

	for _, adjVertex := range vertex.Neighbors {
		if (*visited)[adjVertex.Val] == false {
			dfs(adjVertex, visited, stack)
		}
	}

	*stack = append(*stack, vertex.Val)
}

func main() {
	graphAdjacencyArray := [][]int{
		{1, 2},
		{4},
		{3},
		{},
		{5},
		{},
	}
	vertices := make([]*Vertex, len(graphAdjacencyArray))
	for i := range graphAdjacencyArray {
		vertices[i] = &Vertex{Val: i}
	}

	for i, neighbors := range graphAdjacencyArray {
		for _, neighbor := range neighbors {
			vertices[i].Neighbors = append(vertices[i].Neighbors, vertices[neighbor])
		}
	}
	graph := Graph{Vertices: vertices}

	ret := TopologicalSort(&graph)

	fmt.Println(ret)
}
