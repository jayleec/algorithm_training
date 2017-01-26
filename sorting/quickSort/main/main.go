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
	if high > low {
		pivot = Partition(slice, low, high)
		QuickSort(slice, low, pivot - 1)
		//fmt.Println("1: ", slice)
		QuickSort(slice, pivot + 1 , high)
		//fmt.Println("2: ", slice)
	}
	return slice
}

func Partition(slice []int, low, high int) int {
	pivot := slice[low]
	wall := low

	for i:= low + 1; i<= high; i++{
		if slice[i] < pivot {
			wall = wall + 1
			Swap(slice, i, wall)
		}
	}
	//pivot <-> wall
	Swap(slice, low, wall)
	return wall
}

func Swap(slice []int, a, b int) {
	temp := slice[a]
	slice[a] = slice[b]
	slice[b] = temp
}

