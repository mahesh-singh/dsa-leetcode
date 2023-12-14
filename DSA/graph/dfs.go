package main

import "fmt"

/*
DFS with adjacency matrix
DFS with adjacency list via array
DFS with adjacency list via slice of vertex

*/

func DFS_AdjacencyMatrix(graph [][]int) []int {
	numVertices := len(graph)
	visited := make([]bool, numVertices)
	ret := []int{}

	for i := 0; i < numVertices; i++ {
		if !visited[i] {
			visited[i] = true
			//ret = append(ret, i)
			dfs_AdjacencyMatrixHelper(graph, i, visited, &ret)
		}

	}

	return ret
}

func dfs_AdjacencyMatrixHelper(graph [][]int, curr_vertex int, visited []bool, output *[]int) {
	*output = append(*output, curr_vertex)
	visited[curr_vertex] = true
	for neighbour, value := range graph[curr_vertex] {
		if !visited[neighbour] && value == 1 {
			dfs_AdjacencyMatrixHelper(graph, neighbour, visited, output)
		}
	}
}

func DFS_AdjacencyArray(graph [][]int) []int {
	output := []int{}
	numNodes := len(graph)

	visited := make([]bool, numNodes)
	dfs_AdjacencyArrayHelper(graph, 0, visited, &output)
	return output
}

func dfs_AdjacencyArrayHelper(graph [][]int, currNode int, visited []bool, output *[]int) {

	if visited[currNode] {
		return
	}
	visited[currNode] = true
	*output = append(*output, currNode)

	for _, neighbour := range graph[currNode] {
		dfs_AdjacencyArrayHelper(graph, neighbour, visited, output)
	}
}

type Vertex struct {
	Val        int
	Neighbours []*Vertex
}

func DFS_AdjacencyList(vertex *Vertex, visit func(*Vertex)) {
	visited := make(map[*Vertex]bool)
	dfs_AdjacencyListHelper(vertex, visited, visit)
}

func dfs_AdjacencyListHelper(vertex *Vertex, visited map[*Vertex]bool, visit func(*Vertex)) {
	if visited[vertex] {
		return
	}
	visited[vertex] = true
	visit(vertex)

	for _, neighbour := range vertex.Neighbours {
		dfs_AdjacencyListHelper(neighbour, visited, visit)
	}
}

func main() {
	graph := [][]int{
		//0,1, 2, 3, 4, 5
		{0, 1, 1, 0, 0, 0}, //0
		{1, 0, 0, 0, 1, 0}, //1
		{1, 0, 0, 1, 0, 0}, //2
		{0, 0, 1, 0, 1, 0}, //3
		{0, 1, 0, 1, 0, 1}, //4
		{0, 0, 0, 0, 1, 0}, //5
	}

	ret := DFS_AdjacencyMatrix(graph)
	fmt.Println(ret)

	graphAdjacencyArray := [][]int{
		{1, 2},
		{0, 4},
		{0, 3},
		{2, 4},
		{1, 5},
		{4},
	}

	ret = DFS_AdjacencyArray(graphAdjacencyArray)
	fmt.Println(ret)

	vertices := make([]*Vertex, len(graphAdjacencyArray))
	for i := range graphAdjacencyArray {
		vertices[i] = &Vertex{Val: i}
	}

	for i, neighbours := range graphAdjacencyArray {
		for _, neighbour := range neighbours {
			vertices[i].Neighbours = append(vertices[i].Neighbours, vertices[neighbour])
		}
	}

	ret = []int{}
	visitFunc := func(vertex *Vertex) {
		ret = append(ret, vertex.Val)

	}

	DFS_AdjacencyList(vertices[0], visitFunc)
	fmt.Println(ret)
}
