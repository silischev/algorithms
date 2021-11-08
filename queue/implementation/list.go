package implementation

import (
	"errors"

	ll "github.com/silischev/algorithms/linked_list/implementation"
)

var ErrEmptyQueue = errors.New("queue is empty")

type Queue struct {
	len  int
	head *ll.Node
	tail *ll.Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(data string) {
	node := ll.NewNode(nil, data)
	if q.len == 0 {
		q.head = node
	} else {
		q.tail.SetNext(node)
	}

	q.tail = node
	q.len++
}

func (q *Queue) Dequeue() (*ll.Node, error) {
	if q.len == 0 {
		return nil, ErrEmptyQueue
	} else if q.len == 1 {
		q.tail = nil
	}

	node := q.head
	q.head = q.head.GetNext()
	q.len--

	return node, nil
}
