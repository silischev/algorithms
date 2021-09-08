package main

import (
	"fmt"
	"linked_list/realisation"
)

func main() {
	list := realisation.NewLinkedList()
	list.AddBack("1")
	list.AddBack("2")
	list.AddBack("3")
	list.Print()

	fmt.Println("\n")
	list.Reverse()
	list.Print()
}
