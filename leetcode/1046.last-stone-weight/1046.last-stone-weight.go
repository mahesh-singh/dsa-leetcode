package main

import "container/heap"

type intHeap []int

func (h intHeap) Len() int { return len(h) }

func (h intHeap) Less(i, j int) bool { return h[i] > h[j] } //max heap

func (h intHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := intHeap{}
	heap.Init(&h)
	for i := 0; i < len(stones); i++ {
		heap.Push(&h, stones[i])
	}

	for len(h) > 1 {
		y := heap.Pop(&h)
		x := heap.Pop(&h)

		if x != y {
			diff := y.(int) - x.(int)
			heap.Push(&h, diff)
		}
	}

	if len(h) == 1 {
		return h[0]
	}
	return 0
}
