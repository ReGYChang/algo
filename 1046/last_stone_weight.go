package main

import "container/heap"

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	tmp := old[len(*h)-1]
	*h = old[0 : len(*h)-1]
	return tmp
}

func lastStoneWeight(stones []int) int {
	h := &maxHeap{}
	heap.Init(h)
	for i := 0; i < len(stones); i++ {
		heap.Push(h, stones[i])
	}

	for h.Len() > 1 {
		y := heap.Pop(h).(int)
		x := heap.Pop(h).(int)

		if x != y {
			y -= x
			heap.Push(h, y)
		}
	}

	if h.Len() == 1 {
		return h.Pop().(int)
	}
	return 0
}
