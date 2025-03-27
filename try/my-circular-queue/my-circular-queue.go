package main

import "fmt"

type MyCircularQueue struct {
	data        []int
	head, tail  int
	count, size int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		data: make([]int, k),
		size: k,
	}
}

func (cq *MyCircularQueue) IsEmpty() bool {
	return cq.count == 0
}

func (cq *MyCircularQueue) IsFull() bool {
	return cq.count == cq.size
}

func (cq *MyCircularQueue) EnQueue(val int) bool {
	if cq.IsFull() {
		return false
	}
	cq.data[cq.tail] = val
	cq.tail = (cq.tail + 1) % cq.size
	cq.count += 1
	return true
}

func (cq *MyCircularQueue) DeQueue() bool {
	if cq.IsEmpty() {
		return false
	}
	cq.head = (cq.head + 1) % cq.size
	cq.count -= 1
	return true
}

func (cq *MyCircularQueue) Front() int {
	if cq.IsEmpty() {
		return -1
	}
	return cq.data[cq.head]
}

func (cq *MyCircularQueue) Rear() int {
	if cq.IsEmpty() {
		return -1
	}
	return cq.data[(cq.tail-1+cq.size)%cq.size]
}

func main() {
	mcq := Constructor(5)
	mcq.EnQueue(2)
	mcq.EnQueue(3)
	mcq.EnQueue(5)
	mcq.EnQueue(4)
	mcq.EnQueue(8)
	fmt.Println(mcq)
	mcq.DeQueue()
	fmt.Println(mcq)
}
