package main

import (
	"container/heap"
	"math"
)

type point struct {
	pnt   []int
	diss  int
	index int
}

type distHeap []point

func (h distHeap) Len() int { return len(h) }

func (h distHeap) Less(i, j int) bool { return h[i].diss < h[j].diss }

func (h distHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *distHeap) Push(x interface{}) {
	pt := x.(point)
	n := len(*h)
	pt.index = n
	*h = append(*h, pt)
}

func (h *distHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func kClosest(points [][]int, k int) [][]int {
	hp := &distHeap{}
	heap.Init(hp)

	for i := 0; i < len(points); i++ {
		diss := math.Sqrt(float64(points[i][0]*points[i][0] + points[i][1]*points[i][1]))
		pt := point{pnt: points[i], diss: int(diss), index: i}

		heap.Push(hp, pt)
	}

	ret := make([][]int, 0)
	for i := 0; i < k; i++ {
		pt := heap.Pop(hp).(point)
		ret = append(ret, pt.pnt)
	}

	return ret
}
