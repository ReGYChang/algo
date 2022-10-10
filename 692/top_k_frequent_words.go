package main

import (
	"container/heap"
)

type word struct {
	word      string
	frequency int
}

type maxHeap []*word

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h maxHeap) Less(i, j int) bool {
	if h[i].frequency == h[j].frequency {
		return h[i].word < h[j].word
	}
	return h[i].frequency > h[j].frequency
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(*word))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	tmp := old[len(*h)-1]
	*h = old[0 : len(*h)-1]
	return tmp
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}

	h := &maxHeap{}
	heap.Init(h)
	for w, freq := range m {
		heap.Push(h, &word{word: w, frequency: freq})
	}

	res := make([]string, k)
	for i := 0; i < k; i++ {
		w := heap.Pop(h).(*word)
		res[i] = w.word
	}

	return res
}
