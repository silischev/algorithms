package main

import (
	"log"
	stack "stack/implementation"
)

func main() {
	s := stack.NewSortedStack()

	s.Push("3")
	s.Push("8")

	stackLen := s.Len
	for i := 1; i <= stackLen; i++ {
		n, err := s.Pop()
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(n.GetData())
	}
}
