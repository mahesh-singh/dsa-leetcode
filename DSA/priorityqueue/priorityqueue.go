package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string
	priority int

	//This is maintained by heap.Interface
	index int
}

type PriorityQueue []Item

func (p PriorityQueue) Len() int { return len(p) }

func (p PriorityQueue) Less(i, j int) bool {
	//Because highest priority item need to return so its an maxHeap
	return p[i].priority > p[j].priority
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]

	p[i].index = i //because now p[i] was p[j], hence changing index to i, so that index and value remain in sync
	p[j].index = j
}

func (p *PriorityQueue) Push(x interface{}) {
	n := len(*p) //needed to set the index in x
	item := x.(Item)
	item.index = n
	*p = append(*p, item)
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1] //Less condition is reversed as compare to min-heap
	//old[n-1] = nil   //to avoid memory leak in underline slice, this will help GC to collect memory. It used in DS where slice get modified
	*p = old[0 : n-1]
	return item
}

func (p *PriorityQueue) Update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	fmt.Println(p)
	heap.Fix(p, item.index)
	fmt.Println(p)
}

func main() {
	items := map[string]int{
		"Mango": 2, "banana": 3, "guava": 1,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for val, priority := range items {
		pq[i] = Item{
			value:    val,
			priority: priority,
			index:    i,
		}
		i = i + 1
	}

	heap.Init(&pq)
	item := Item{value: "orange", priority: 4}
	heap.Push(&pq, item)
	pq.Update(&item, item.value, 5)

	for pq.Len() > 0 {
		itm := heap.Pop(&pq).(Item)
		fmt.Printf("%v: %v\n", itm.value, itm.priority)
	}
}
