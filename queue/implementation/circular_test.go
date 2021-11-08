package implementation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyCircularQueue(t *testing.T) {
	q := NewCircularQueue(1)

	n, err := q.Dequeue()
	require.Error(t, ErrEmptyQueue, err)
	require.Empty(t, n)
}

func TestExceededSizeCircularQueue(t *testing.T) {
	q := NewCircularQueue(1)

	err := q.Enqueue("1")
	require.Nil(t, err)

	err = q.Enqueue("2")
	require.Error(t, ErrExceededCircularQueueSize, err)
}

func TestEnqueueDequeueCircularQueue(t *testing.T) {
	const queueSize = 3
	q := NewCircularQueue(queueSize)

	err := q.Enqueue("test1")
	require.Nil(t, err)
	require.Equal(t, 0, q.head)
	require.Equal(t, 0, q.tail)

	err = q.Enqueue("test2")
	require.Nil(t, err)
	require.Equal(t, 0, q.head)
	require.Equal(t, 1, q.tail)

	err = q.Enqueue("test3")
	require.Nil(t, err)
	require.Equal(t, 0, q.head)
	require.Equal(t, 2, q.tail)

	data, err := q.Dequeue()
	require.Nil(t, err)
	require.Equal(t, "test1", data)
	require.Equal(t, 1, q.head)
	require.Equal(t, 2, q.tail)

	data, err = q.Dequeue()
	require.Nil(t, err)
	require.Equal(t, "test2", data)
	require.Equal(t, 2, q.head)
	require.Equal(t, 2, q.tail)

	data, err = q.Dequeue()
	require.Nil(t, err)
	require.Equal(t, "test3", data)
	require.Equal(t, 0, q.head, "head")
	require.Equal(t, 0, q.tail, "tail")

	n, err := q.Dequeue()
	require.Error(t, ErrEmptyQueue, err)
	require.Empty(t, n)

	err = q.Enqueue("test4")
	require.Nil(t, err)
	require.Equal(t, 0, q.head)
	require.Equal(t, 0, q.tail)

	err = q.Enqueue("test5")
	require.Nil(t, err)
	require.Equal(t, 0, q.head)
	require.Equal(t, 1, q.tail)

	data, err = q.Dequeue()
	require.Nil(t, err)
	require.Equal(t, "test4", data)
	require.Equal(t, 1, q.head)
	require.Equal(t, 1, q.tail)

	data, err = q.Dequeue()
	require.Nil(t, err)
	require.Equal(t, "test5", data)
	require.Equal(t, 0, q.head)
	require.Equal(t, 0, q.tail)

	n, err = q.Dequeue()
	require.Error(t, ErrEmptyQueue, err)
	require.Empty(t, n)
}
