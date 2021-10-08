package main

import (
	"fmt"

	"github.com/silischev/algorithms/linked_list/implementation"
)

func main() {
	list := implementation.NewLinkedList()
	list.AddBack("1")
	list.AddBack("2")
	list.AddBack("3")
	list.Print()

	fmt.Println("\n")
	list.Reverse()
	list.Print()

	fmt.Println("\n")
	list.ReverseRecursive(list.Head(), list.Head().GetNext())
	list.Print()
}
