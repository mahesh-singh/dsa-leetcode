package main

import (
	"container/list"
	"fmt"
)

func TraversalAllNode_BFS(graph [][]int) []int {
	queue := []int{}
	num_vertices := len(graph)
	visted := make([]bool, num_vertices)
	ret := []int{}

	for i := 0; i < num_vertices; i++ {
		if !visted[i] {
			queue = append(queue, i)
			visted[i] = true
			ret = append(ret, i)

			for len(queue) > 0 {
				curren_vertex := queue[0]
				queue = queue[1:]
				for neighbor := 0; neighbor < num_vertices; neighbor++ {
					if graph[curren_vertex][neighbor] == 1 && !visted[neighbor] {
						queue = append(queue, neighbor)
						visted[neighbor] = true

						ret = append(ret, neighbor)
					}
				}

			}
		}

	}

	return ret
}

func TraversalAllNode_BFS_AdjacencyList(graph [][]int) {
	numVertices := len(graph)
	visited := make([]bool, numVertices)
	queue := list.New()

	queue.PushBack(0)
	visited[0] = true

	for queue.Len() > 0 {
		currentVertex := queue.Remove(queue.Front()).(int)
		fmt.Printf("%d ", currentVertex)

		for _, neighbor := range graph[currentVertex] {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
			}
		}
	}
}

func main() {

	/*
		graph := [][]int{
			{1, 1, 1, 0, 0, 0},
			{1, 1, 0, 0, 1, 0},
			{1, 0, 0, 1, 0, 0},
			{0, 0, 1, 0, 1, 0},
			{0, 1, 0, 1, 0, 1},
			{0, 0, 0, 0, 1, 0},
		}

		want := []int{0, 1, 2, 4, 3, 5}

		got := TraversalAllNode_BFS(graph)

		if !reflect.DeepEqual(got, want) {
			fmt.Printf("expected: %v \n", want)
			fmt.Printf("expected: %v \n", got)
		} else {
			fmt.Println("looks good")
		}*/

	graph := [][]int{
		{1, 3, 4},
		{0, 2},
		{1, 3},
		{0, 2, 4},
		{0, 3},
	}

	// Call the BFS function starting from vertex 0
	TraversalAllNode_BFS_AdjacencyList(graph)

}
