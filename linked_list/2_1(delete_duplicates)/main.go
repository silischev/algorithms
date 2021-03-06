package main

import (
	"fmt"

	"github.com/silischev/algorithms/linked_list/implementation"
)

func main() {
	list := implementation.NewLinkedList()
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

	list.DeleteDuplicates()
	list.Print()

	list.Delete("1")
	list.Delete("5")
	fmt.Println("\n")
	list.Print()
}
