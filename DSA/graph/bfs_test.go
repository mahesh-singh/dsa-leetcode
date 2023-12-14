package main

import (
	"reflect"
	"testing"
)

func TestTraversalAllNode_BFS(t *testing.T) {

	tests := []struct {
		name  string
		graph [][]int
		want  []int
	}{
		{"Adjacency matrix",
			[][]int{
				{1, 1, 1, 0, 0, 0},
				{1, 1, 0, 0, 1, 0},
				{1, 0, 0, 1, 0, 0},
				{0, 0, 1, 0, 1, 0},
				{0, 1, 0, 1, 0, 1},
				{0, 0, 0, 0, 1, 0},
			},
			[]int{0, 1, 2, 4, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TraversalAllNode_BFS(tt.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TraversalAllNode_BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
