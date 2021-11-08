package implementation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyQueue(t *testing.T) {
	q := NewQueue()

	n, err := q.Dequeue()
	require.Error(t, ErrEmptyQueue, err)
	require.Nil(t, n)
}

func TestEnqueueDequeueQueue(t *testing.T) {
	q := NewQueue()

	q.Enqueue("test1")
	require.Equal(t, q.head.GetData(), "test1")
	require.Equal(t, q.tail.GetData(), "test1")
	require.Equal(t, 1, q.len)

	q.Enqueue("test2")
	require.Equal(t, q.head.GetData(), "test1")
	require.Equal(t, q.tail.GetData(), "test2")
	require.Equal(t, 2, q.len)

	q.Enqueue("test3")
	require.Equal(t, q.head.GetData(), "test1")
	require.Equal(t, q.tail.GetData(), "test3")
	require.Equal(t, 3, q.len)

	node, err := q.Dequeue()
	require.Nil(t, err)
	require.Equal(t, node.GetData(), "test1")
	require.Equal(t, q.head.GetData(), "test2")
	require.Equal(t, q.tail.GetData(), "test3")

	node, err = q.Dequeue()
	require.Nil(t, err)
	require.Equal(t, node.GetData(), "test2")
	require.Equal(t, q.head.GetData(), "test3")
	require.Equal(t, q.tail.GetData(), "test3")

	node, err = q.Dequeue()
	require.Nil(t, err)
	require.Equal(t, node.GetData(), "test3")
	require.Nil(t, q.head)
	require.Nil(t, q.tail)
	require.Equal(t, 0, q.len)

	n, err := q.Dequeue()
	require.Error(t, ErrEmptyQueue, err)
	require.Nil(t, n)
}
