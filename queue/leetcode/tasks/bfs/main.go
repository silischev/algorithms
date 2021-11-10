package main

import "log"

func main() {
	//edges := [][]int{[]int{0, 1}, []int{1, 2}, []int{2, 0}}
	//edges := [][]int{[]int{4, 3}, []int{1, 4}, []int{4, 8}, []int{1, 7}, []int{6, 4}, []int{4, 2}, []int{7, 4}, []int{4, 0}, []int{0, 9}, []int{5, 4}}
	edges := [][]int{[]int{0, 7}, []int{0, 8}, []int{6, 1}, []int{2, 0}, []int{0, 4}, []int{5, 8}, []int{4, 7}, []int{1, 3}, []int{3, 5}, []int{6, 5}}
	//edges := [][]int{[]int{0, 1}, []int{0, 2}, []int{3, 5}, []int{5, 4}, []int{4, 3}} // false
	log.Println(validPath(10, edges, 7, 5))
}
