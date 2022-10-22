package main

type MyQueue struct {
	queue []int
}

func Constructor() MyQueue {
	return MyQueue{queue: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyQueue) Pop() int {
	head := this.queue[0]
	this.queue = this.queue[1:]
	return head
}

func (this *MyQueue) Peek() int {
	return this.queue[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
