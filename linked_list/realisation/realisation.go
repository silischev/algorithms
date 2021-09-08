package realisation

import (
	"fmt"
)

type linkedList interface {
	Get(i int) *Node
	Next() *Node
	Delete(i int)
	AddBack(data string)
	AddFront(data string)
	Insert(i int, n Node)
	Put(i int, n Node)
}

type Node struct {
	next *Node
	data string
}

type SingleLinkedList struct {
	len  int
	head *Node
}

func NewLinkedList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func (l *SingleLinkedList) Print() {
	currNode := l.head
	for i := 0; i < l.len; i++ {
		if i == l.len-1 {
			fmt.Print(currNode.data)
			break
		}

		fmt.Print(currNode.data + ", ")
		currNode = currNode.next
	}
}

func (l *SingleLinkedList) GetAll() []string {
	currNode := l.head
	data := []string{}
	for i := 0; i < l.len; i++ {
		data = append(data, currNode.data)
		currNode = currNode.next
	}

	return data
}

func (l *SingleLinkedList) Reverse() {
	var prev *Node
	curr := l.head

	for curr != nil {
		next := curr.next
		t1 := curr
		curr.next = prev

		prev = t1
		curr = next

	}

	l.head = prev
}

func (l *SingleLinkedList) ReverseRecursive() {
}

func (l *SingleLinkedList) AddFront(data string) {
	node := &Node{
		next: nil,
		data: data,
	}

	if l.len == 0 {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}

	l.len++
}

func (l *SingleLinkedList) AddBack(data string) {
	node := &Node{
		next: nil,
		data: data,
	}

	if l.len == 0 {
		l.head = node
	} else {
		currNode := l.head
		for currNode.next != nil {
			currNode = currNode.next
		}

		currNode.next = node
	}

	l.len++
}