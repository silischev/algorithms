package implementation

import (
	"errors"
	"log"
)

var (
	ErrEmptyCircularQueue        = errors.New("queue is empty")
	ErrExceededCircularQueueSize = errors.New("queue size exceeded")
)

type CircularQueue struct {
	elements []string
	head     int
	tail     int
	currSize int
}

func NewCircularQueue(size int) *CircularQueue {
	return &CircularQueue{elements: make([]string, size, size)}
}

func (q *CircularQueue) Enqueue(data string) error {
	if q.currSize == cap(q.elements) {
		return ErrExceededCircularQueueSize
	}

	if q.currSize > 0 {
		q.tail++
	}

	q.elements[q.tail] = data
	q.currSize++

	return nil
}

/*func (q *CircularQueue) Enqueue(data string) error {
	if q.currSize == cap(q.elements) {
		return ErrExceededCircularQueueSize
	}

	if q.currSize > 0 {
		q.tail++
	}

	q.elements = append(q.elements, data)
	q.currSize++

	return nil
}*/

func (q *CircularQueue) Dequeue() (string, error) {
	log.Println("q.currSize ", q.currSize)
	if q.currSize == 0 {
		return "", ErrEmptyCircularQueue
	}

	element := q.elements[q.head]
	if q.currSize == 1 {
		q.head = 0
		q.tail = 0
	} else {
		q.head++
	}

	q.currSize--

	return element, nil
}
