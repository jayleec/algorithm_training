package main

import (
	"fmt"
	"github.com/jayleec/algorithm_training/algoutils"
)

func main() {

	slice := algoutils.GenerateSlice(10)
	fmt.Println("Unsorted: ", slice)

	fmt.Println("Sorted: ", InsertionSort(slice))
}

func InsertionSort(slice []int) []int {
	var i, j, temp int
	for i = 1; i < len(slice); i++ {
		temp = slice[i]
		j = i
		for j >= 1 && slice[j-1] > temp {
			slice[j] = slice[j-1]
			j = j - 1
		}
		slice[j] = temp
	}
	return slice
}
