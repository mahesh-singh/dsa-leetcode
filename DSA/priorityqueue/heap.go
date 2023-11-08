package main

import (
	"container/heap"
	"fmt"
)

type intHeap []int

func (h intHeap) Len() int { return len(h) }

func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h intHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {

	//*h is dereference the pointer receiver to get the value of minHeap
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	//to make life easy, otherwise everything can be done via *h as well

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &intHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 4)
	heap.Push(h, 3)

	//fmt.Printf("min value is: %d \n", heap.Pop(h))
	fmt.Printf("Heap : %d \n", (*h))
	fmt.Printf("Heap : %d \n", heap.Pop(h))
	fmt.Printf("Heap : %d \n", (*h))

	//fmt.Printf("3rd value is: %d \n", (*h)[2])
}
