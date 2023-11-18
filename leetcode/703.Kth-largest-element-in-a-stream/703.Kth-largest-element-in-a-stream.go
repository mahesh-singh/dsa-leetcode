package main

import "container/heap"

type KthLargest struct {
	heap minHeap
	k    int
}

type minHeap []int

func (mh minHeap) Len() int           { return len(mh) }
func (mh minHeap) Less(i, j int) bool { return mh[i] < mh[j] }
func (mh minHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }

func (mh *minHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}

func (mh *minHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}

func Constructor(k int, nums []int) KthLargest {
	minH := minHeap{}
	heap.Init(&minH)

	for i := 0; i < len(nums); i++ {
		if len(minH) < k {
			heap.Push(&minH, nums[i])
		} else if minH[0] < nums[i] {
			heap.Pop(&minH)
			heap.Push(&minH, nums[i])
		} else {
			continue
		}
	}

	return KthLargest{heap: minH, k: k}

}

func (this *KthLargest) Add(val int) int {
	if len(this.heap) < this.k {
		heap.Push(&this.heap, val)
	} else if this.heap[0] < val {
		heap.Pop(&this.heap)
		heap.Push(&this.heap, val)
	}
	return this.heap[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
