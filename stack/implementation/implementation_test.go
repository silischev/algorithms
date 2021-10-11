package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	s := NewStack()

	n, err := s.Pop()
	assert.Error(t, ErrEmptyStack, err)
	assert.Nil(t, n)
	assert.True(t, s.isEmpty())

	n, err = s.Peek()
	assert.Error(t, ErrEmptyStack, err)
	assert.Nil(t, n)
	assert.True(t, s.isEmpty())
}

func TestPushPop(t *testing.T) {
	s := NewStack()

	s.Push("test1")
	assert.Equal(t, s.Top.GetData(), "test1")

	s.Push("test2")
	assert.Equal(t, s.Top.GetData(), "test2")

	node, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, node.GetData(), "test2")

	node, err = s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, node.GetData(), "test1")
	assert.Equal(t, 0, s.Len)
}

func TestPeek(t *testing.T) {
	s := NewStack()

	s.Push("test1")
	s.Push("test2")

	n, err := s.Peek()
	assert.Nil(t, err)
	assert.Equal(t, n.GetData(), "test2")
}
