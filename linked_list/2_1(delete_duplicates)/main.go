package main

import (
	"fmt"
	"linked_list/realisation"
)

func main() {
	list := realisation.NewLinkedList()
	list.AddBack("1")
	list.AddBack("2")
	list.AddBack("2")
	list.AddBack("2")
	list.AddBack("3")
	list.AddBack("4")
	list.AddBack("4")
	list.AddBack("5")
	list.Print()
	fmt.Println("\n")

	deleteDuplicates(list)
	list.Print()
}

func deleteDuplicates(l *realisation.SingleLinkedList) {
	currNode := l.Head()
	var prevNode *realisation.Node

	for currNode != nil {
		if prevNode != nil && currNode.Data() == prevNode.Data() {
			prevNode.SetNext(currNode.Next())
			l.SetLength(l.Length() - 1)
			currNode = currNode.Next()
			continue
		}

		prevNode = currNode
		currNode = currNode.Next()
	}
}
