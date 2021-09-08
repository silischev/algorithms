package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	/*data := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}*/

	/*data := [][]int{
		{1, 2},
		{3, 4},
	}
	*/

	printSlice(data)
	log.Println("***")
	printSlice(turn90(data))
}

func printSlice(data [][]int) {
	for row := 0; row < len(data); row++ {
		for col := 0; col < len(data); col++ {
			fmt.Print(data[row][col], " ")
			if col == len(data)-1 {
				fmt.Println()
			}
		}
	}
}

func turn90(data [][]int) [][]int {
	/*turnData := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}*/

	turnData := make([][]int, len(data))
	for i := range turnData {
		turnData[i] = make([]int, len(data))
	}

	row := 0
	col := 0
	newRowPos := len(data) - 1

	for col < len(data) {
		for row < len(data) {
			turnData[col][newRowPos] = data[row][col]
			row++
			newRowPos--
		}

		newRowPos = len(data) - 1
		row = 0
		col++
	}

	return turnData
}
