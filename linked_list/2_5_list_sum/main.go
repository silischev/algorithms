package main

import (
	"strconv"

	ll "github.com/silischev/algorithms/linked_list/implementation"
)

func main() {
	list1 := ll.NewLinkedList()
	list1.AddBack("6")
	list1.AddBack("1")
	list1.AddBack("3")
	list1.Print()
	//fmt.Println("\n")

	list2 := ll.NewLinkedList()
	list2.AddBack("2")
	list2.AddBack("4")
	list2.Print()
	//fmt.Println("\n")

	sum(list1, list2).Print()
}

func sum(l1 *ll.SingleLinkedList, l2 *ll.SingleLinkedList) *ll.SingleLinkedList {
	newList := &ll.SingleLinkedList{}

	currRes := 0

	currNodeLarger := l1.Head()
	currNodeSmaller := l2.Head()
	if l2.Length() > l1.Length() {
		currNodeLarger = l2.Head()
		currNodeSmaller = l1.Head()
	}

	for currNodeLarger != nil {
		data1, _ := strconv.Atoi(currNodeLarger.GetData())
		if currNodeSmaller != nil {
			data2, _ := strconv.Atoi(currNodeSmaller.GetData())
			currRes += data1 + data2
			currNodeSmaller = currNodeSmaller.GetNext()
		} else {
			currRes += data1
		}

		currNodeLarger = currNodeLarger.GetNext()

		if currRes > 9 {
			newList.AddFront(strconv.Itoa(currRes % 10))
			currRes = 1
		} else {
			newList.AddFront(strconv.Itoa(currRes))
			currRes = 0
		}
	}

	return newList
}
