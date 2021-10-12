package implementation

import (
	"errors"

	ll "github.com/silischev/algorithms/linked_list/implementation"
)

var ErrEmptyStack = errors.New("stack is empty")

type Stack struct {
	Len int
	Top *ll.Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(data string) {
	node := ll.NewNode(s.Top, data)
	s.Top = node
	s.Len++
}

func (s *Stack) Pop() (*ll.Node, error) {
	if s.isEmpty() {
		return nil, ErrEmptyStack
	}

	node := s.Top
	s.Top = node.GetNext()
	s.Len--

	return node, nil
}

func (s *Stack) Peek() (*ll.Node, error) {
	if s.isEmpty() {
		return nil, ErrEmptyStack
	}

	return s.Top, nil
}

func (s *Stack) isEmpty() bool {
	return s.Len == 0
}

func (s *Stack) AsSlice() []string {
	out := []string{}
	next := s.Top
	for next != nil {
		out = append(out, next.GetData())
		next = next.GetNext()
	}

	return out
}
