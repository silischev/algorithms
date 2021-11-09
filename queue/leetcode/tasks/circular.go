package tasks

import (
	"fmt"
)

type MyCircularQueue struct {
	elements []int
	head     int
	tail     int
	currSize int
	maxSize  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{elements: make([]int, k, k), maxSize: k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.currSize == this.maxSize {
		return false
	}

	if this.currSize > 0 && this.tail < this.maxSize-1 {
		this.tail++
	} else if this.tail == this.maxSize-1 {
		this.tail = 0
	}

	this.elements[this.tail] = value

	this.currSize++

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.currSize == 0 {
		return false
	}

	if this.currSize == 1 || this.head == this.maxSize-1 {
		this.head = 0
		this.tail = 0
	} else if this.head < this.maxSize-1 {
		this.head++
	}

	this.currSize--

	return true
}

func (this *MyCircularQueue) Front() int {
	if this.currSize == 0 {
		return -1
	}

	return this.elements[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.currSize == 0 {
		return -1
	}

	return this.elements[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.currSize == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.currSize == this.maxSize
}

func (this *MyCircularQueue) Print() {
	fmt.Printf("head: [%d] tail: [%d] size: [%d]\n", this.head, this.tail, this.currSize)

	for i, el := range this.elements {
		fmt.Printf("[%d]%d\n", i, el)
	}
}
