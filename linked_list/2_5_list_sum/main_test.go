package main

import (
	"testing"

	ll "github.com/silischev/algorithms/linked_list/implementation"

	"github.com/stretchr/testify/assert"
)

func TestSum1(t *testing.T) {
	l1 := ll.NewLinkedList()
	l2 := ll.NewLinkedList()

	l1.AddBack("6")
	l1.AddBack("1")
	l1.AddBack("3")

	l2.AddBack("2")

	assert.Equal(t, []string{"3", "1", "8"}, sum(l1, l2).GetAll())
}

func TestSum2(t *testing.T) {
	l1 := ll.NewLinkedList()
	l2 := ll.NewLinkedList()

	l1.AddBack("6")
	l1.AddBack("1")
	l1.AddBack("3")

	l2.AddBack("4")

	assert.Equal(t, []string{"3", "2", "0"}, sum(l1, l2).GetAll())
}

func TestSum3(t *testing.T) {
	l1 := ll.NewLinkedList()
	l2 := ll.NewLinkedList()

	l1.AddBack("6")
	l1.AddBack("1")

	l2.AddBack("2")

	assert.Equal(t, []string{"1", "8"}, sum(l1, l2).GetAll())
}

func TestSum4(t *testing.T) {
	l1 := ll.NewLinkedList()
	l2 := ll.NewLinkedList()

	l1.AddBack("6")
	l1.AddBack("1")
	l1.AddBack("3")

	l2.AddBack("5")

	assert.Equal(t, []string{"3", "2", "1"}, sum(l1, l2).GetAll())
}

func TestSum5(t *testing.T) {
	l1 := ll.NewLinkedList()
	l2 := ll.NewLinkedList()

	l1.AddBack("6")
	l1.AddBack("1")
	l1.AddBack("3")

	l2.AddBack("9")

	assert.Equal(t, []string{"3", "2", "5"}, sum(l1, l2).GetAll())
}

func TestSum6(t *testing.T) {
	l1 := ll.NewLinkedList()
	l2 := ll.NewLinkedList()

	l1.AddBack("6")
	l1.AddBack("1")
	l1.AddBack("3")

	l2.AddBack("0")

	assert.Equal(t, []string{"3", "1", "6"}, sum(l1, l2).GetAll())
}

func TestSum7(t *testing.T) {
	l1 := ll.NewLinkedList()
	l2 := ll.NewLinkedList()

	l1.AddBack("9")

	l2.AddBack("6")
	l2.AddBack("1")
	l2.AddBack("3")

	assert.Equal(t, []string{"3", "2", "5"}, sum(l1, l2).GetAll())
}

func TestSum8(t *testing.T) {
	l1 := ll.NewLinkedList()
	l2 := ll.NewLinkedList()

	l1.AddBack("1")

	l2.AddBack("6")

	assert.Equal(t, []string{"7"}, sum(l1, l2).GetAll())
}

func TestSum9(t *testing.T) {
	l1 := ll.NewLinkedList()
	l2 := ll.NewLinkedList()

	l1.AddBack("7")
	l1.AddBack("1")
	l1.AddBack("6")

	l2.AddBack("5")
	l2.AddBack("9")
	l2.AddBack("2")

	assert.Equal(t, []string{"9", "1", "2"}, sum(l1, l2).GetAll())
}
