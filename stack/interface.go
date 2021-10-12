package stack

import ll "github.com/silischev/algorithms/linked_list/implementation"

type Stacker interface {
	Push(data string)
	Pop() ll.Node
	IsEmpty() bool
	Peek() ll.Node
	AsSlice() []string
}
