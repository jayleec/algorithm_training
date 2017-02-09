package main

import (
	"fmt"
	"github.com/jayleec/algorithm_training/algoutils"
)

func main(){
	slice := algoutils.GenerateSlice(10)
	fmt.Println("Unsorted: ", slice)
	fmt.Println("Sorted:", QuickSort(slice, 0, 9))
}

func QuickSort(slice []int, low, high int)[]int {
	var pivot int
	if low < high {
		pivot = Partition(slice, low, high)
		QuickSort(slice, low, pivot)
		QuickSort(slice, pivot + 1 , high)
	}
	return slice
}

func Partition(slice []int, low, high int) int {
	pivot := slice[low]
	wall := low

	for i:= low + 1; i<= high; i++{
		if slice[i] < pivot {
			wall = wall + 1
			algoutils.Swap(slice, i, wall)
		}
	}
	//pivot <-> wall
	algoutils.Swap(slice, low, wall)
	return wall
}