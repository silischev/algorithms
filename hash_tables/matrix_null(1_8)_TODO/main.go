package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	/*data := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}*/

	/*data := [][]int{
		{1, 2},
		{3, 0},
	}*/

	/*data := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{0, 8, 9},
	}*/

	data := [][]int{
		{1, 2, 3},
		{4, 0, 0},
		{0, 8, 9},
	}

	printSlice(data)
	log.Println("***************************")
	printSlice(mark(data))
}

func printSlice(data [][]int) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			fmt.Print(data[i][j], " ")
			if j == len(data)-1 {
				fmt.Println()
			}
		}
	}
}

func mark(data [][]int) [][]int {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {

			curElem := data[i][j]
			/*log.Println("i ", i)
			log.Println("j ", j)
			log.Println("curElem ", curElem)*/
			if curElem == 0 {
				if data[0][j] == 0 {
					continue
				}

				//log.Println("elem is 0 => ", i, j)

				for pos := 0; pos < len(data[i]); pos++ {
					data[i][pos] = 0
				}

				for pos := 0; pos < len(data[j]); pos++ {
					data[pos][j] = 0
				}

				//printSlice(data)

				break
			}
		}
	}

	return data
}
