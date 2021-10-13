package implementation

import (
	"strconv"

	ll "github.com/silischev/algorithms/linked_list/implementation"
)

type SortedStack struct {
	Stack
}

func NewSortedStack() *SortedStack {
	return &SortedStack{}
}

func (s *SortedStack) Push(data string) {
	node := ll.NewNode(nil, data)
	s.Len++

	if s.Top == nil {
		s.Top = node
		return
	}

	currNode := s.Top
	nextNode := node
	for nextNode != nil {
		nextNodeData, _ := strconv.Atoi(nextNode.GetData())
		currNodeData, _ := strconv.Atoi(currNode.GetData())

		if currNodeData >= nextNodeData {
			s.Top = nextNode
			s.Top.SetNext(currNode)
			break
		} else {
			actualNextNode := currNode.GetNext()
			if actualNextNode != nil {
				actualNextNodeData, _ := strconv.Atoi(actualNextNode.GetData())

				if actualNextNodeData >= nextNodeData {
					currNode.SetNext(nextNode)
					nextNode.SetNext(actualNextNode)
					break
				} else {
					currNode = actualNextNode
					continue
				}
			}

			currNode.SetNext(nextNode)
			nextNode = nextNode.GetNext()
		}
	}
}
