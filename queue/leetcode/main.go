package main

import (
	"queue/leetcode/tasks"
)

func main() {
	obj := tasks.Constructor(2)
	println(obj.EnQueue(8), " EnQueue 8")
	obj.Print()
	println(obj.EnQueue(2), " EnQueue 2")
	obj.Print()

	println(obj.DeQueue(), " DeQueue 8")
	obj.Print()

	println(obj.EnQueue(5), " EnQueue 5")
	obj.Print()

	println(obj.DeQueue(), " DeQueue 2")
	obj.Print()

	println(obj.EnQueue(1), " EnQueue 1")
	obj.Print()

	println(obj.DeQueue(), " DeQueue 5")
	obj.Print()
}
